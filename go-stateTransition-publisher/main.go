package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"simple-publisher/sate"
)

func main() {
	// Standard GO configuration to connect to an Ethereum node and run transactions

	// Connect to an ethereum RPC node (Infura or Alchemy) using your own token
	client, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/<token>")
	if err != nil {
		log.Fatal(err)
	}

    // Pass in your private key
    // NOTE: NEVER store your private key in source code if you are gonna publish it or use in production.
	privateKey, err := crypto.HexToECDSA("<private key>")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units, avarage gas
	auth.GasPrice = gasPrice

	// Load an istance of the smart contract
	address := common.HexToAddress("0xa36786C3E18225da7cc8FC69c6443ecD41827FF5")
	instance, err := sate.NewSate(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Hardcode the inputs from the public.json from the Generate Proof section
	id, _ := new(big.Int).SetString("<>", 10)
	oldState, _ := new(big.Int).SetString(
		("<>",, 10)
	newState, _ := new(big.Int).SetString(
		("<>",, 10)
	isOldStateGenesis, _ := new(big.Int).SetString("1", 10)

	// Hardcode the proof.json files from the Generate Proof section
	proofA_0 := stringToInt("<>")
	proofA_1 := stringToInt("<>")

	proofB_0_0 := stringToInt("<>")
	proofB_0_1 := stringToInt("<>")

	proofB_1_0 := stringToInt("<>")
	proofB_1_1 := stringToInt("<>")

	proofC_0 := stringToInt("<>")
	proofC_1 := stringToInt("<>")

	proofA := [2]*big.Int{proofA_0, proofA_1}

	proofB := [2][2]*big.Int{
		{proofB_0_1, proofB_0_0},
		{proofB_1_1, proofB_1_0},
	}
	proofC := [2]*big.Int{proofC_0, proofC_1}

	// Call the smart contract TransitState function passing in the inputs it needs
	tx, err := instance.TransitState(auth, id, oldState, newState, isOldStateGenesis, proofA, proofB, proofC)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}

func stringToInt(s string) *big.Int {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		log.Fatal("can not convert string to big.Int")
	}
	return i
}
