package gocontracts

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployWriteRequestHolder(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (common.Address, *types.Transaction, *WriteRequestHolder, error) {
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		log.Fatal(e)
	}
	address, tx, instance, err := DeployWriteRequestHolder(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The nonce is ", tx.Nonce())
	WaitForTransAction(tx, client, 1)
	return address, tx, instance, err
}
