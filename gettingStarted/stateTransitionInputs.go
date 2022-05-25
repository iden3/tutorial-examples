package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/iden3/go-circuits"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

// Identity
func main() {
	ctx := context.Background()

	// generate babyJJ private key
	pk := babyjub.NewRandPrivKey()

	// generate auth claim from babyJJ key
	authClaim, _ := AuthClaim(pk)
	authHi, _, _ := authClaim.HiHv()

	// Create inital (empty) identity trees
	claimsTree, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 40) // Claims tree
	revTree, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 40)    // Revocation tree
	rootsTree, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 40)  // Roots tree

	// GENESIS STATE

	// Retrieve ID after adding the authClaim to the Claims tree
	id, _, _ := NewIdentity(ctx, authClaim, claimsTree)

	// Retrieve Genesis State as Hash of the roots of the 3 identity trees
	genesisState, _ := merkletree.HashElems(
		claimsTree.Root().BigInt(),
		revTree.Root().BigInt(),
		rootsTree.Root().BigInt())

	fmt.Println("ID at t=0:", id)
	fmt.Println("State at t=0:", genesisState.BigInt())

	// Generate Merkle Tree Proof for authClaim
	authMTPProof, _, _ := claimsTree.GenerateProof(ctx, authHi, claimsTree.Root())

	// Generate the Non Revocation Merkle Tree Proof for the authClaim.
	// We are proving that the authClaim hasn't been revoked by checking that the nonce associated to the Auth Claim is not inclued in the Revocaion tree
	nonce := new(big.Int).SetUint64(authClaim.GetRevocationNonce())
	authNonRevMTPProof, _, _ := revTree.GenerateProof(ctx, nonce, revTree.Root())

	genesisTreeState := circuits.TreeState{
		State:          genesisState,
		ClaimsRoot:     claimsTree.Root(),
		RevocationRoot: revTree.Root(),
		RootOfRoots:    rootsTree.Root(),
	}

	// STATE 1

	// add new random claim to Claim tree
	newClaim, _ := CreateTestClaim() // create new claim
	hi, hv, _ := newClaim.HiHv()     // Get hash Index and hash Value of the new claim
	claimsTree.Add(ctx, hi, hv)      // add claim to the Claims tree

	// identity state has changed!

	// 1. The Claims tree has been updated so the root has changed
	claimsTreeRoot := claimsTree.Root() // Claims tree root

	// 2. The Roots tree has to be updated with the new claimsTreeRoot
	rootsTree.Add(ctx, claimsTreeRoot.BigInt(), big.NewInt(0))

	// retrieve the new Identity state
	newState, _ := merkletree.HashElems(
		claimsTree.Root().BigInt(),
		revTree.Root().BigInt(),
		rootsTree.Root().BigInt())

	fmt.Println("ID at t=1:", id)
	fmt.Println("State at t=1:", newState.BigInt())

	// Sign a message (hash of the genesis state + hash of the new state) using your private key
	hashOldAndNewStates, _ := poseidon.Hash(
		[]*big.Int{genesisState.BigInt(), newState.BigInt()})
	signature := pk.SignPoseidon(hashOldAndNewStates)

	// Generate state Transition Inputs
	stateTransitionInputs := circuits.StateTransitionInputs{
		ID:                id,
		OldTreeState:      genesisTreeState,
		NewState:          newState,
		IsOldStateGenesis: true,
		AuthClaim: circuits.Claim{
			Claim: authClaim,
			Proof: authMTPProof,
			NonRevProof: &circuits.ClaimNonRevStatus{
				Proof: authNonRevMTPProof,
			},
		},
		Signature: signature,
	}

	inputBytes, _ := stateTransitionInputs.InputsMarshal()

	fmt.Println(string(inputBytes))
}

func AuthClaim(pk babyjub.PrivateKey) (*core.Claim, error) {

	pubKey := pk.Public()
	// create BBJJ claim
	schemaHash, err := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	if err != nil {
		return nil, err
	}

	revNonce := uint64(1)

	return core.NewClaim(schemaHash,
		core.WithIndexDataInts(pubKey.X, pubKey.Y),
		core.WithRevocationNonce(revNonce))

}

func CreateTestClaim() (*core.Claim, error) {
	schemaHex := hex.EncodeToString([]byte("myAge_test_claim"))
	schema, err := core.NewSchemaHashFromHex(schemaHex)
	if err != nil {
		return nil, err
	}

	// define value for the claim
	code := big.NewInt(51)

	// create claim
	return core.NewClaim(schema, core.WithIndexDataInts(code, nil))

}

func NewIdentity(ctx context.Context, authClaim *core.Claim, clt *merkletree.MerkleTree) (id *core.ID, genesisState *big.Int, err error) {

	// add auth claim to claims tree
	hIndex, hValue, err := authClaim.HiHv()
	clt.Add(ctx, hIndex, hValue)
	if err != nil {
		return nil, nil, err
	}

	// calculate identity initial state (genesis state)
	state, err := core.IdenState(clt.Root().BigInt(), big.NewInt(0), big.NewInt(0))
	if err != nil {
		return nil, nil, err
	}

	// calculate id
	identifier, err := core.IdGenesisFromIdenState(core.TypeDefault, state)
	if err != nil {
		return nil, nil, err
	}

	return identifier, state, nil
}
