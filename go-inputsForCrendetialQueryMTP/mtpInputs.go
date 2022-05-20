package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/iden3/go-circuits"
	it "github.com/iden3/go-circuits/testing"
	"github.com/iden3/go-iden3-crypto/poseidon"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

func main() {

	// Generate Bob private key
	userPrivKHex := "28156abe7fe2fd433dc9df969286b96666489bac508612d0e16593e944c4f69f"

	// Generate Alice private key
	issuerPrivKHex := "21a5e7321d0e2f3ca1cc6504396e6594a2211544b08c206847cdee96f832421a"

	ctx := context.Background()

	// Create Bob identity adding auth claim to Bob's Claims Tree
	userIdentity, uClaimsTree, uRevsTree, _, _, userAuthCoreClaim, userPrivateKey := it.Generate(ctx,
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

	// Retrieve Index for Bob's auth claim
	hIndexAuthEntryUser, _, _ := claimsIndexValueHashes(*userAuthCoreClaim)

	// Retrieve Merkle Proof for Bob's auth claim
	mtpProofUser, _, _ := uClaimsTree.GenerateProof(ctx,
		hIndexAuthEntryUser, uClaimsTree.Root())

	// Create Alice identity adding auth claim to the Claims Tree
	issuerID, iClaimsTree, _, _, _, _, _ := it.Generate(ctx,
		issuerPrivKHex)

	// Alice issues age claim for Bob saying that he is 25 yo.

	// Add age in the dataslotA
	dataSlotA, _ := core.NewElemBytesFromInt(big.NewInt(25))

	// Add nonce
	nonce := 1
	var schemaHash core.SchemaHash

	// Use the schema age
	schemaBytes, _ := hex.DecodeString("ce6bb12c96bfd1544c02c289c6b4b987")

	copy(schemaHash[:], schemaBytes)

	// Create age claim
	issuerCoreClaim, _ := core.NewClaim(
		schemaHash,
		core.WithIndexID(*userIdentity),
		core.WithIndexData(dataSlotA, core.ElemBytes{}),
		core.WithExpirationDate(time.Unix(1669884010,
			0)), //Thu Dec 01 2022 08:40:10 GMT+0000
		core.WithRevocationNonce(uint64(nonce)))

	hIndexClaimEntry, hValueClaimEntry, _ := claimsIndexValueHashes(*issuerCoreClaim)

	// Add the age claim to Alice's Claims Tree
	_ = iClaimsTree.Add(ctx, hIndexClaimEntry, hValueClaimEntry)

	// Retrieve Merkle Proof for age claim
	proof, _, _ := iClaimsTree.GenerateProof(ctx, hIndexClaimEntry,
		iClaimsTree.Root())

	// Fetch new Alice's state
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

	// Create (empty) Alice's revocation tree
	issuerRevTreeStorage := memory.NewMemoryStorage()
	issuerRevTree, _ := merkletree.NewMerkleTree(ctx, issuerRevTreeStorage,
		40)

	// Retrieve Non revocation Merkle Proof for age claim = Proof of non-membership of nonce 1 to the Alice's revocation tree
	proofNotRevoke, _, _ := issuerRevTree.GenerateProof(ctx,
		big.NewInt(int64(nonce)), issuerRevTree.Root())

	// Retrieve Non-revocation Merkle Proof for Bob's auth claim = Proof of non-membership of Auth Claim's nonce to the Bob's revocation tree
	authClaimRevNonce := new(big.Int).
		SetUint64(userAuthCoreClaim.GetRevocationNonce())
	proofAuthClaimNotRevoked, _, _ :=
		uRevsTree.GenerateProof(ctx, authClaimRevNonce, nil)

		// Design Carl's (Verifier) Query
	query := circuits.Query{
		SlotIndex: 2,
		Values:    []*big.Int{new(big.Int).SetInt64(18)},
		Operator:  2,
	}

	// Create a challenge specific to Bob
	challenge := new(big.Int).SetInt64(1)

	// Bob signs a message containing the challenge presented by Carl
	message := big.NewInt(0).SetBytes(challenge.Bytes())
	challengeSignature := userPrivateKey.SignPoseidon(message)

	// Generate inputs for the circuit

	inputsAuthClaim := circuits.Claim{
		Claim:     userAuthCoreClaim,
		Proof:     mtpProofUser,
		TreeState: userAuthTreeState,
		NonRevProof: &circuits.ClaimNonRevStatus{
			TreeState: userAuthTreeState,
			Proof:     proofAuthClaimNotRevoked,
		},
	}

	inputsUserClaim := circuits.Claim{
		Claim:     issuerCoreClaim,
		Proof:     proof,
		TreeState: issuerStateAfterClaimAdd,
		IssuerID:  issuerID,
		NonRevProof: &circuits.ClaimNonRevStatus{
			TreeState: issuerStateAfterClaimAdd,
			Proof:     proofNotRevoke,
		},
	}

	atomicInputs := circuits.AtomicQueryMTPInputs{
		ID:        userIdentity,
		AuthClaim: inputsAuthClaim,
		Challenge: challenge,
		Signature: challengeSignature,

		Claim: inputsUserClaim,

		CurrentTimeStamp: time.Unix(1642074362, 0).Unix(),

		Query: query,
	}

	bytesInputs, _ := atomicInputs.InputsMarshal()

	// Print out the inputs
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
