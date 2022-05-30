package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/iden3/go-circuits"
	it "github.com/iden3/go-circuits/testing"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

func main() {
	userPrivKHex := "28156abe7fe2fd433dc9df969286b96666489bac508612d0e16593e944c4f69f"
	issuerPrivKHex := "21a5e7321d0e2f3ca1cc6504396e6594a2211544b08c206847cdee96f832421a"
	challenge := new(big.Int).SetInt64(1)
	ctx := context.Background()

	userIdentity, uClaimsTree, _, _, _, userAuthCoreClaim, userPrivateKey := it.Generate(ctx,
		userPrivKHex)

	state, _ := merkletree.HashElems(
		uClaimsTree.Root().BigInt(),
		merkletree.HashZero.BigInt(),
		merkletree.HashZero.BigInt())

	userAuthTreeState := circuits.TreeState{
		State:          state,
		ClaimsRoot:     uClaimsTree.Root(),
		RevocationRoot: &merkletree.HashZero,
		RootOfRoots:    &merkletree.HashZero,
	}

	hIndexAuthEntryUser, _, _ := claimsIndexValueHashes(*userAuthCoreClaim)

	mtpProofUser, _, _ := uClaimsTree.GenerateProof(ctx,
		hIndexAuthEntryUser, uClaimsTree.Root())

	message := big.NewInt(0).SetBytes(challenge.Bytes())

	challengeSignature := userPrivateKey.SignPoseidon(message)

	// Issuer
	issuerIdentity, iClaimsTree, iRevTree, _, _, issuerAuthClaim, issuerKey := it.Generate(ctx,
		issuerPrivKHex)

	// issuer state
	issuerGenesisState, _ := merkletree.HashElems(
		iClaimsTree.Root().BigInt(),
		merkletree.HashZero.BigInt(),
		merkletree.HashZero.BigInt())

	issuerAuthTreeState := circuits.TreeState{
		State:          issuerGenesisState,
		ClaimsRoot:     iClaimsTree.Root(),
		RevocationRoot: &merkletree.HashZero,
		RootOfRoots:    &merkletree.HashZero,
	}

	hIndexAuthEntryIssuer, _, _ :=
		claimsIndexValueHashes(*issuerAuthClaim)

	mtpProofIssuer, _, _ := iClaimsTree.GenerateProof(ctx,
		hIndexAuthEntryIssuer, iClaimsTree.Root())

	issuerAuthClaimRevNonce := new(big.Int).SetUint64(issuerAuthClaim.GetRevocationNonce())
	issuerAuthNonRevProof, _, _ := iRevTree.GenerateProof(ctx,
		issuerAuthClaimRevNonce, iRevTree.Root())

	// issue issuerClaim for user
	dataSlotA, _ := core.NewElemBytesFromInt(big.NewInt(25))

	nonce := 1
	var schemaHash core.SchemaHash

	schemaBytes, _ := hex.DecodeString("ce6bb12c96bfd1544c02c289c6b4b987")

	copy(schemaHash[:], schemaBytes)

	issuerCoreClaim, _ := core.NewClaim(
		schemaHash,
		core.WithIndexID(*userIdentity),
		core.WithIndexData(dataSlotA, core.ElemBytes{}),
		core.WithExpirationDate(time.Unix(1669884010,
			0)), //Thu Dec 01 2022 08:40:10 GMT+0000
		core.WithRevocationNonce(uint64(nonce)))

	hashIndex, hashValue, _ := claimsIndexValueHashes(*issuerCoreClaim)

	commonHash, _ := merkletree.HashElems(hashIndex, hashValue)

	claimSignature := issuerKey.SignPoseidon(commonHash.BigInt())

	iClaimsTree.Add(ctx, hashIndex, hashValue)

	proof, _, _ := iClaimsTree.GenerateProof(ctx, hashIndex,
		iClaimsTree.Root())

	stateAfterClaimAdd, _ := merkletree.HashElems(
		iClaimsTree.Root().BigInt(),
		merkletree.HashZero.BigInt(),
		merkletree.HashZero.BigInt())

	issuerStateAfterClaimAdd := circuits.TreeState{
		State:          stateAfterClaimAdd,
		ClaimsRoot:     iClaimsTree.Root(),
		RevocationRoot: &merkletree.HashZero,
		RootOfRoots:    &merkletree.HashZero,
	}

	issuerRevTreeStorage := memory.NewMemoryStorage()
	issuerRevTree, _ := merkletree.NewMerkleTree(ctx, issuerRevTreeStorage,
		40)

	proofNotRevoke, _, _ := issuerRevTree.GenerateProof(ctx,
		big.NewInt(int64(nonce)), issuerRevTree.Root())

	inputsAuthClaim := circuits.Claim{
		//Schema:    authClaim.Schema,
		Claim:     userAuthCoreClaim,
		Proof:     mtpProofUser,
		TreeState: userAuthTreeState,
		NonRevProof: &circuits.ClaimNonRevStatus{
			TreeState: userAuthTreeState,
			Proof:     mtpProofUser,
		},
	}

	claimIssuerSignature := circuits.BJJSignatureProof{
		IssuerID:           issuerIdentity,
		IssuerTreeState:    issuerAuthTreeState,
		IssuerAuthClaimMTP: mtpProofIssuer,
		Signature:          claimSignature,
		IssuerAuthClaim:    issuerAuthClaim,
		IssuerAuthNonRevProof: circuits.ClaimNonRevStatus{
			TreeState: issuerAuthTreeState,
			Proof:     issuerAuthNonRevProof,
		},
	}

	inputsUserClaim := circuits.Claim{
		Claim:     issuerCoreClaim,
		Proof:     proof,
		TreeState: issuerStateAfterClaimAdd,
		NonRevProof: &circuits.ClaimNonRevStatus{
			TreeState: issuerStateAfterClaimAdd,
			Proof:     proofNotRevoke,
		},
		IssuerID:       issuerIdentity,
		SignatureProof: claimIssuerSignature,
	}

	query := circuits.Query{
		SlotIndex: 2,
		Values:    []*big.Int{new(big.Int).SetInt64(18)},
		Operator:  2,
	}

	atomicInputs := circuits.AtomicQuerySigInputs{
		ID:        userIdentity,
		AuthClaim: inputsAuthClaim,
		Challenge: challenge,
		Signature: challengeSignature,

		CurrentTimeStamp: time.Unix(1642074362, 0).Unix(),

		Claim: inputsUserClaim,

		Query: query,
	}

	bytesInputs, _ := atomicInputs.InputsMarshal()

	fmt.Println(string(bytesInputs))

}

func claimsIndexValueHashes(c core.Claim) (*big.Int, *big.Int, error) {
	index, value := c.RawSlots()
	indexHash, err := poseidon.Hash(core.ElemBytesToInts(index[:]))
	if err != nil {
		return nil, nil, err
	}
	valueHash, err := poseidon.Hash(core.ElemBytesToInts(value[:]))
	return indexHash, valueHash, err
}
