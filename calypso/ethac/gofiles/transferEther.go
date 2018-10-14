package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
	if err != nil {
		fmt.Println("jon")
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Can't get public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, e := client.SuggestGasPrice(context.Background())
	if e != nil {
		log.Fatal(e)
	}
	value := big.NewInt(5000000000000000000)
	gasLimit := uint64(21000)
	var data []byte
	toAddress := common.HexToAddress("0xF28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, e := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if e != nil {
		fmt.Println("bjorn")
		log.Fatal(err)
	}
	client.SendTransaction(context.Background(), signedTx)
	checkIfTransactionExists("0xc6cfda68fc14ee127bcc16e07de08d70b7dc5d17bb20bcc1f662c66520273094")

}

func checkIfTransactionExists(hash string) bool {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	genesisBlock, _ := client.BlockByHash(context.Background(), common.HexToHash(hash))
	fmt.Println("Genesis block hash", genesisBlock.Hash().String())
	fmt.Println("Genesis block number", genesisBlock.Header().Number.String())
	return true
}
