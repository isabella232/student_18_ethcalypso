package gocontracts

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployReadRequestHolder(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (common.Address, *types.Transaction, *ReadRequestHolder, error) {
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		fmt.Println("Bjorn")
		log.Fatal(e)
	}

	address, tx, instance, err := DeployReadRequestHolder(auth, client)
	if err != nil {
		fmt.Println("bjorn")
		log.Fatal(err)
	}
	WaitForTransAction(tx, client, 1)
	return address, tx, instance, err
}
