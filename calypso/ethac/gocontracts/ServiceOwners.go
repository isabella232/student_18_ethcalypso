package gocontracts

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployOwners(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (common.Address, *types.Transaction, *Owners, error) {
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		log.Fatal(e)
	}

	address, tx, instance, err := DeployOwners(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The nonce is ", tx.Nonce())
	WaitForTransAction(tx, client, 1)
	return address, tx, instance, err
}

func ServiceUpdatOwners(a common.Address, deployed common.Address, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	calypsoTransActor, e := NewCalypsoTransactor(deployed, client)
	if e != nil {
		return nil, e
	}
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		return nil, e
	}
	tx, e := calypsoTransActor.AddOwner(auth, a)
	return tx, e
}
