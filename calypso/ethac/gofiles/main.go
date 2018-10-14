package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
)

func mains() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
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

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	calypsoAddress := common.HexToAddress("0xb269bc939c3B8E4E995EEEC7068D5916D2Db0FA3")

	transferFnSignature := []byte("addWriteRequest(bytes,bytes,bytes,address[])")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))

	dataVar := []byte("Hello, my name is Bjorn")
	extraData := []byte("Hello, my name is Sabrina")
	ltsid := []byte("Hello, my name is Manuel")
	dataVarB := common.LeftPadBytes(dataVar, 32)
	extraDataB := common.LeftPadBytes(extraData, 32)
	ltsidB := common.LeftPadBytes(ltsid, 32)

	account1 := common.HexToAddress("0x09E360FeD8580641CE129252417a8646709196f5")
	account2 := common.HexToAddress("0xF28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e")
	account1Bytes := common.LeftPadBytes(account1.Bytes(), 32)
	account2Bytes := common.LeftPadBytes(account2.Bytes(), 32)
	var accounts []byte
	accounts = append(accounts, account1Bytes...)
	accounts = append(accounts, account2Bytes...)
	var data []byte
	data = append(data, methodID...)
	data = append(data, dataVarB...)
	data = append(data, extraDataB...)
	data = append(data, ltsidB...)
	data = append(data, accounts...)
	/*gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To: &calypsoAddress,
	})
	if err != nil {
		fmt.Println("Could not estimate")
		log.Fatal(err)
	}*/

	//Just set the gas limit to really high, Will have to check on it later
	gasLimit := uint64(7000000)
	fmt.Printf("Gas limit: %d", gasLimit)
	tx := types.NewTransaction(nonce, calypsoAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tokens sent at TX: %s", signedTx.Hash().Hex())
}
