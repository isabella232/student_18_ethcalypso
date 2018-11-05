package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/dedis/cothority"
	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployWriteRequest(privateKey *ecdsa.PrivateKey, client *ethclient.Client, d []byte, ed []byte, ltsid []byte, p []common.Address, U []byte, cs [][]byte) (common.Address, *types.Transaction, *WriteRequest, error) {
	auth, e := GetAuth(privateKey, client)
	if e != nil {
		log.Fatal(e)
	}
	temp := make([]byte, 0)
	for j := 0; j < len(cs); j++ {
		temp = append(temp, cs[j]...)
	}
	address, tx, instance, err := DeployWriteRequest(auth, client, d, ed, ltsid, p, U, temp, int64(len(cs)))
	if err != nil {
		log.Fatal(err)
	}
	WaitForTransAction(tx, client, 1)
	return address, tx, instance, err
}

//Write is a struct that keeps track
//of all the data in a WriteRequest
//store on an ethereum blockchain
type Write struct {
	Data      []byte
	ExtraData []byte
	U         kyber.Point
	LTSID     []byte
	Cs        []kyber.Point
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
		return nil, e
	}
	Cs, e := wrCaller.Cs(call)
	if e != nil {
		return nil, e
	}
	split, e := wrCaller.Split(call)
	if e != nil {
		return nil, e
	}
	fmt.Println("Split ", split)
	temp := make([][]byte, 0)
	n := int64(len(Cs)) / split
	for i := int64(0); i < split-1; i++ {
		temp = append(temp, Cs[i*n:i*(n+1)])
	}
	if int64(len(Cs)) == n {
		temp = append(temp, Cs)
	}
	points := make([]kyber.Point, 0)
	for i := 0; i < len(temp); i++ {
		p := cothority.Suite.Point()
		e := p.UnmarshalBinary(temp[i])
		if e != nil {
			return nil, e
		}
		points = append(points, p)
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
		Cs:        points,
	}
	return write, nil
}
