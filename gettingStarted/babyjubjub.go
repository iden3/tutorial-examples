package main

import (
	"fmt"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

// BabyJubJub
func main() {

	// generate babyJubjub private key
	privKey := babyjub.NewRandPrivKey()
	// get public Key from private Key
	pubKey := privKey.Public()

	fmt.Println(pubKey)

}
