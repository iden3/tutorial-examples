package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
)

// Sparse MT
func main() {
	ctx := context.Background()

	// Tree storage
	store := memory.NewMemoryStorage()

	// Generate a new MerkleTree
	mt, _ := merkletree.NewMerkleTree(ctx, store, 32)

	// Add a leaf to the tree with key 1 and value 10
	key1 := big.NewInt(1)
	value1 := big.NewInt(10)
	mt.Add(ctx, key1, value1)

	fmt.Printf("Adding a leaf to the tree with key %d and value %d\n", key1, value1)

	// Update leaf at the key = key1
	newValue1 := big.NewInt(20)
	mt.Update(ctx, key1, newValue1)

	// Add another leaf to the tree
	key2 := big.NewInt(2)
	value2 := big.NewInt(15)
	mt.Add(ctx, key2, value2)

	fmt.Printf("Adding a leaf to the tree with key %d and value %d\n", key2, value2)

	// Proof the membership of a leaf with key 1
	proofExist, value, _ := mt.GenerateProof(ctx, key1, mt.Root())
	fmt.Printf("Proof of membership of a leaf with key %d is %t\n", key1, proofExist.Existence)
	fmt.Printf("Value corresponding to the queried key: %d\n", value)

	// Proof of non-membership of a leaf with key 4
	proofNotExist, _, _ := mt.GenerateProof(ctx, big.NewInt(4), mt.Root())
	fmt.Printf("Proof of membership of a leaf with key %d is %t\n", big.NewInt(4), proofNotExist.Existence)
}
