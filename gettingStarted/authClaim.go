package main

import (
	"encoding/json"
	"fmt"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

// Create auth claim
func main() {

	// Generate babyjubjub keypair
	privKey := babyjub.NewRandPrivKey()
	pubKey := privKey.Public()

	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")

	// Add revNonce
	revNonce := uint64(1)

	// create claim storing the X coordinate of the pub key in the first index data slot
	// and the Y coordinate in the second index data slot
	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(pubKey.X, pubKey.Y),
		core.WithRevocationNonce(revNonce))

	authClaimToMarshal, _ := json.Marshal(authClaim)

	// print the claim in string format
	fmt.Println(string(authClaimToMarshal))
}
