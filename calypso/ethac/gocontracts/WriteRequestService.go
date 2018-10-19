package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/dedis/cothority"
	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployWriteRequest(privateKey *ecdsa.PrivateKey, client *ethclient.Client, d []byte, ed []byte, ltsid []byte, p []common.Address, U []byte) (common.Address, *types.Transaction, *WriteRequest, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4712388) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployWriteRequest(auth, client, d, ed, ltsid, p, U)
	if err != nil {
		log.Fatal(err)
	}
	return address, tx, instance, err
}

type Write struct {
	Data      []byte
	ExtraData []byte
	U         kyber.Point
	LTSID     []byte
}

func ServiceGetWriteRequest(privateKey *ecdsa.PrivateKey, client *ethclient.Client, a common.Address) (*Write, error) {
	call, e := CreateCallOpts(context.Background(), privateKey, client)
	if e != nil {
		return nil, e
	}
	wrCaller, e := NewWriteRequestCaller(a, client)
	if e != nil {
		return nil, e
	}
	ltsid, e := wrCaller.LTSID(call)
	if e != nil {
		return nil, e
	}
	data, e := wrCaller.Data(call)
	if e != nil {
		return nil, e
	}
	ed, e := wrCaller.ExtraData(call)
	if e != nil {
		return nil, e
	}
	U, e := wrCaller.U(call)
	if e != nil {
		fmt.Println(U)
		return nil, e
	}
	point := cothority.Suite.Point()
	e = point.UnmarshalBinary(U)
	if e != nil {
		return nil, e
	}
	write := &Write{
		Data:      data,
		ExtraData: ed,
		U:         point,
		LTSID:     ltsid,
	}
	return write, nil
}
