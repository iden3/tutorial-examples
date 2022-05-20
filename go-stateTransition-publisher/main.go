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
	client, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/<token>")
	if err != nil {
		log.Fatal(err)
	}

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

	address := common.HexToAddress("0xa36786C3E18225da7cc8FC69c6443ecD41827FF5")
	instance, err := sate.NewSate(address, client)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := new(big.Int).SetString("212111378763298689002912895423936774667763498482833067487304921288930492416", 10)
	oldState, _ := new(big.Int).SetString(
		"139039236343966087193440793366889850012233031210552441993012191883787325388", 10)
	newState, _ := new(big.Int).SetString(
		"19279431795911575844572480519225458998491451269153657263205822009113404350164", 10)
	isOldStateGenesis, _ := new(big.Int).SetString("1", 10)

	//{
	//	"pi_a": [
	//"4541249619660857829687485867016105455896181561606422924188923469010365687288",
	//"9284446482755979148706022167193244578612598728029148199089157936700322983826",
	//"1"
	//],
	//"pi_b": [
	//[
	//"20907420752803144437243925552401838181195045774956547021002016312332495825766",
	//"16601533590672542652860109242650013673071338130165323977211541370843731550875"
	//],
	//[
	//"20239518017592861620905954949643304172206749928585247673904298304288034827401",
	//"8687219602545301746881478083618978229259576089323579932683422323510866755499"
	//],
	//[
	//"1",
	//"0"
	//]
	//],
	//"pi_c": [
	//"13187244300989365639399661797965257043648852418899922424970485425148007277552",
	//"1771076586200216620591802306415313397906433476744831179448080251400149245557",
	//"1"
	//],
	//"protocol": "groth16",
	//"curve": "bn128"
	//}

	proofA_0 := stringToInt("20970269039732656914916493235285561025951820969112827040052692952223487539310")
	proofA_1 := stringToInt("10149215757706072068269850248501248866573487622775031978509549607900801297043")

	proofB_0_0 := stringToInt("19596423134274638073730402820752293635804955924315168831515672836286837084757")
	proofB_0_1 := stringToInt("2575068209277623247010697722481422706256769935344127087409997414011557402452")

	proofB_1_0 := stringToInt("5857706721571373261728436022024818566771790305771089074940079367382167043395")
	proofB_1_1 := stringToInt("19273288656430907073158894848704470484100206678044910327894389161975858615073")

	proofC_0 := stringToInt("16324024070991631863253851809731467084739306192519499556149011142627728552678")
	proofC_1 := stringToInt("9804276003107083237032945835113863050750039176185538513173269443221570133307")

	proofA := [2]*big.Int{proofA_0, proofA_1}

	proofB := [2][2]*big.Int{
		{proofB_0_1, proofB_0_0},
		{proofB_1_1, proofB_1_0},
	}
	proofC := [2]*big.Int{proofC_0, proofC_1}

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
