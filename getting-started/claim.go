package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	core "github.com/iden3/go-iden3-core"
)

// create basic claim
func main() {

	ageSchema, err := core.NewSchemaHashFromHex("ce38102464833febf36e714922a83050")
	if err != nil {
		log.Fatal(err)
	}

	// define age
	age := big.NewInt(25)

	// create claim based on the ageSchema storing the age in the first index slot
	claim, err := core.NewClaim(ageSchema, core.WithIndexDataInts(age, nil))
	if err != nil {
		log.Fatal(err)
	}

	claimToMarshal, _ := json.Marshal(claim)

	// print the claim in string format
	fmt.Println(string(claimToMarshal))
}
