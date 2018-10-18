package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceCalypso(privateKey *ecdsa.PrivateKey, client *ethclient.Client, owners common.Address, wrHolder common.Address, rrHolder common.Address) (common.Address, *types.Transaction, *Calypso, error) {
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4712388) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployCalypso(auth, client, owners, wrHolder, rrHolder)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	return address, tx, instance, err
}

func ServiceAddWriteRequest(privateKey *ecdsa.PrivateKey, cal common.Address, wr common.Address, client *ethclient.Client) (*types.Transaction, error) {
	calTransActor, e := NewCalypsoTransactor(cal, client)
	if e != nil {
		log.Fatal(e)
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = fromAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4712388) // in units
	auth.GasPrice = gasPrice
	tx, e := calTransActor.AddWriteRequest(auth, wr)
	return tx, e

}

func ServiceAddReadRequest(privateKey *ecdsa.PrivateKey, cal common.Address, rr common.Address, client *ethclient.Client) (*types.Transaction, error) {
	calTransActor, e := NewCalypsoTransactor(cal, client)
	if e != nil {
		log.Fatal(e)
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = fromAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4712388) // in units
	auth.GasPrice = gasPrice
	tx, e := calTransActor.AddReadRequest(auth, rr)
	return tx, e

}
