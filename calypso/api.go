package calypso

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/byzcoin"
	"github.com/dedis/cothority/darc"
	"github.com/dedis/kyber"
	"github.com/dedis/onet"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Client is a class to communicate to the calypso service.
type Client struct {
	roster     onet.Roster
	c          *onet.Client
	ltsReply   *CreateLTSReply
	PrivateKey *ecdsa.PrivateKey
}

// WriteReply is returned upon successfully spawning a Write instance.
type WriteReply struct {
	*byzcoin.AddTxResponse
	byzcoin.InstanceID
}

// ReadReply is is returned upon successfully spawning a Read instance.
type ReadReply struct {
	*byzcoin.AddTxResponse
	byzcoin.InstanceID
}

// NewClient instantiates a new Client.
// It takes as input an "initialized" byzcoin client
// with an already created ledger
func NewClient(r onet.Roster, pk *ecdsa.PrivateKey) *Client {
	return &Client{roster: r, c: onet.NewClient(
		cothority.Suite, ServiceName),
		PrivateKey: pk}
}

// CreateLTS creates a random LTSID that can be used to reference
// the LTS group created.
func (c *Client) CreateLTS(cal common.Address) (reply *CreateLTSReply, err error) {
	reply = &CreateLTSReply{}
	err = c.c.SendProtobuf(c.roster.List[0], &CreateLTS{
		Roster:         c.roster,
		CalypsoAddress: cal,
	}, reply)
	if err != nil {
		fmt.Println("protobuf error")
		return nil, err
	}
	return reply, nil
}

// DecryptKey takes as input Read- and Write- Proofs. It verifies that
// the read/write requests match and then re-encrypts the secret
// given the public key information of the reader.
func (c *Client) DecryptKey(dkr *DecryptKey) (reply *DecryptKeyReply, err error) {
	reply = &DecryptKeyReply{}
	err = c.c.SendProtobuf(c.roster.List[0], dkr, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (c *Client) AddRead(wr common.Address, cal common.Address) (kyber.Scalar, *common.Address, error) {
	privateKey := c.PrivateKey
	client, e := ethereum.GetClient()
	if e != nil {
		return nil, nil, e
	}
	reader := darc.NewSignerEd25519(nil, nil)
	Xc := reader.Ed25519.Point
	data, e := Xc.MarshalBinary()
	if e != nil {
		return nil, nil, e
	}
	rAddr, _, _, e := gocontracts.ServiceDeployReadRequest(privateKey, client, wr, data)
	if e != nil {
		return nil, nil, e
	}
	_, e = gocontracts.ServiceAddReadRequest(privateKey, cal, rAddr, client)
	if e != nil {
		return nil, nil, e
	}
	return reader.Ed25519.Secret, &rAddr, nil
}

func (c *Client) AddWrite(LTSID []byte, X kyber.Point, data []byte, cal common.Address, p []common.Address) (*common.Address, error) {
	privateKey := c.PrivateKey
	client, e := ethereum.GetClient()
	signer := darc.NewSignerEd25519(nil, nil)
	darc1 := darc.NewDarc(darc.InitRules([]darc.Identity{signer.Identity()},
		[]darc.Identity{signer.Identity()}), []byte("Signer"))
	if darc1 == nil {
		return nil, errors.New("Error making a darc")
	}
	write := NewWrite(cothority.Suite, LTSID, darc1.GetBaseID(), X, data)
	temp := make([][]byte, 0)
	for _, point := range write.Cs {
		marshalled, e := point.MarshalBinary()
		if e != nil {
			return nil, e
		}
		temp = append(temp, marshalled)
	}
	wrData, e := write.U.MarshalBinary()
	if e != nil {
		return nil, e
	}
	Ubar, e := write.Ubar.MarshalBinary()
	if e != nil {
		return nil, e
	}
	F, e := write.F.MarshalBinary()
	if e != nil {
		return nil, e
	}
	E, e := write.E.MarshalBinary()
	if e != nil {
		return nil, e
	}
	addr, _, _, e := gocontracts.ServiceDeployWriteRequest(privateKey, client, write.Data, write.ExtraData, write.LTSID, p, wrData, temp, Ubar, F, E)
	if e != nil {
		return nil, e
	}
	_, e = gocontracts.ServiceAddWriteRequest(privateKey, cal, addr, client)
	if e != nil {
		return nil, e
	}
	return &addr, nil
}
