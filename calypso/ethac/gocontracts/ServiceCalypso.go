package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceCalypso(privateKey *ecdsa.PrivateKey, client *ethclient.Client, owners common.Address, wrHolder common.Address, rrHolder common.Address) (common.Address, *types.Transaction, *Calypso, error) {
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		log.Fatal(e)
	}

	address, tx, instance, err := DeployCalypso(auth, client, wrHolder, rrHolder, owners)
	if err != nil {
		log.Fatal(err)
	}
	WaitForTransAction(tx, client, 1)
	return address, tx, instance, err
}

func ServiceAddWriteRequest(privateKey *ecdsa.PrivateKey, cal common.Address, wr common.Address, client *ethclient.Client) (*types.Transaction, error) {
	calTransActor, e := NewCalypsoTransactor(cal, client)
	if e != nil {
		log.Fatal(e)
	}
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		return nil, e
	}
	tx, e := calTransActor.AddWriteRequest(auth, wr)
	return tx, e

}

func ServiceAddReadRequest(privateKey *ecdsa.PrivateKey, cal common.Address, rr common.Address, client *ethclient.Client) (*types.Transaction, error) {
	calTransActor, e := NewCalypsoTransactor(cal, client)
	if e != nil {
		return nil, e
	}
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		return nil, e
	}
	tx, e := calTransActor.AddReadRequest(auth, rr)
	return tx, e
}

func ServiceCheckIfCanRead(privateKey *ecdsa.PrivateKey, cal common.Address, rr common.Address, client *ethclient.Client) bool {
	call, e := CreateCallOpts(context.Background(), privateKey, client)
	if e != nil {
		return false
	}
	calypsoCaller, e := NewCalypsoCaller(cal, client)
	canRead, e := calypsoCaller.IsValidReadRequest(call, rr)
	if e != nil {
		return false
	}
	return canRead
}
