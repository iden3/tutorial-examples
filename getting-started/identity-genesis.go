package main

import (
	"context"
	"fmt"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

// Identity
func main() {
	ctx := context.Background()

	// babyJJ key
	pk := babyjub.NewRandPrivKey()
	pubKey := pk.Public()

	// create BBJJ claim
	schemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	revNonce := uint64(1)
	authClaim, _ := core.NewClaim(schemaHash,
		core.WithIndexDataInts(pubKey.X, pubKey.Y),
		core.WithRevocationNonce(revNonce))

	// Generate the identity trees
	clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) // Empty Claims tree
	ret, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) // Empty Revocation tree
	rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) // Empty Roots tree

	// Get the Index of the claim and the Value of the claim
	hIndex, hValue, _ := authClaim.HiHv()
	// add auth claim to claims tree with value hValue at index hIndex
	clt.Add(ctx, hIndex, hValue)

	// calculate Genesis State as hash of the 3 roots
	state, _ := core.IdenState(clt.Root().BigInt(), ret.Root().BigInt(), rot.Root().BigInt())

	fmt.Println("Genesis State:", state)

	// calculate ID starting from the genesis state
	id, _ := core.IdGenesisFromIdenState(core.TypeDefault, state)
	fmt.Println("ID:", id)
}
