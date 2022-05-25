package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/iden3/go-circuits"
	it "github.com/iden3/go-circuits/testing"
)

func main() {

	ctx := context.Background()
	privKeyHex := "28156abe7fe2fd433dc9df969286b96666489bac508612d0e16593e944c4f69f"

	// Challenge defined by the verifier to be signed
	challenge := big.NewInt(1)

	// AuthClaimFullInfo method is to generate auth claim, identity, all its trees, state
	// and sign a challenge with the private key.
	identifier, claim, state, claimsTree, revTree, rootsTree, claimEntryMTP, claimNonRevMTP, signature, _ := it.AuthClaimFullInfo(ctx, privKeyHex, challenge)

	treeState := circuits.TreeState{
		State:          state,
		ClaimsRoot:     claimsTree.Root(),
		RevocationRoot: revTree.Root(),
		RootOfRoots:    rootsTree.Root(),
	}

	inputs := circuits.AuthInputs{
		ID: identifier,
		AuthClaim: circuits.Claim{
			Claim:       claim,
			Proof:       claimEntryMTP,
			TreeState:   treeState,
			NonRevProof: &circuits.ClaimNonRevStatus{treeState, claimNonRevMTP},
		},
		Signature: signature,
		Challenge: challenge,
	}

	bytesInputs, _ := inputs.InputsMarshal()

	fmt.Println(string(bytesInputs))

}
