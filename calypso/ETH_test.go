package calypso

import (
	"crypto/ecdsa"
	"fmt"
	"testing"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/byzcoin"
	"github.com/dedis/cothority/darc"
	"github.com/dedis/onet"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func getLTS(t *testing.T, cal common.Address, privateKey *ecdsa.PrivateKey) (*CreateLTSReply, darc.ID, *Client, *darc.Signer) {
	l := onet.NewTCPTest(cothority.Suite)
	servers, roster, _ := l.GenTree(3, true)
	l.GetServices(servers, calypsoID)

	signer := darc.NewSignerEd25519(nil, nil)
	provider1 := darc.NewSignerEd25519(nil, nil)
	reader1 := darc.NewSignerEd25519(nil, nil)
	msg, err := byzcoin.DefaultGenesisMsg(byzcoin.CurrentVersion, roster, []string{"spawn:" + byzcoin.ContractDarcID}, signer.Identity())
	msg.BlockInterval = 100 * time.Millisecond
	require.Nil(t, err)
	// The darc inside it should be valid.
	d := msg.GenesisDarc
	require.Nil(t, d.Verify(true))
	//Create Ledger
	//c, _, err := byzcoin.NewLedger(msg, false)
	require.Nil(t, err)
	//Create a Calypso Client (Byzcoin + Onet)
	calypsoClient := NewClient(*roster, privateKey)
	//Invoke CreateLTS
	ltsReply, err := calypsoClient.CreateLTS(cal)
	calypsoClient.ltsReply = ltsReply
	require.Nil(t, err)
	require.NotNil(t, ltsReply.LTSID)
	require.NotNil(t, ltsReply.X)
	darc1 := darc.NewDarc(darc.InitRules([]darc.Identity{provider1.Identity()},
		[]darc.Identity{provider1.Identity()}), []byte("Provider1"))
	// provider1 is the owner, while reader1 is allowed to do read
	require.NotNil(t, darc1)
	//_, err = calypsoClient.SpawnDarc(signer, d, *darc1, 10)
	//require.Nil(t, err)
	return ltsReply, darc1.GetBaseID(), calypsoClient, &reader1

}
func TestCreateLTS_ETH(t *testing.T) {
	cal, e := gocontracts.GetStaticCalypso()
	require.Nil(t, e)
	privateKey, err := ethereum.GetPrivateKey()
	require.Nil(t, err)
	_, gDarc, calypsoClient, signer := getLTS(t, *cal, privateKey)
	write := NewWrite(cothority.Suite, calypsoClient.ltsReply.LTSID, gDarc, calypsoClient.ltsReply.X, []byte("Sabrina"))
	client, e := ethereum.GetClient()
	require.Nil(t, e)
	temp := make([][]byte, 0)
	for _, point := range write.Cs {
		marshalled, e := point.MarshalBinary()
		require.Nil(t, e)
		temp = append(temp, marshalled)
	}
	policy := ethereum.GetTestListOfAddresses()
	write.ETHAdresses = policy
	Xc := signer.Ed25519.Point
	data, e := Xc.MarshalBinary()
	require.Nil(t, e)
	wrData, e := write.U.MarshalBinary()
	require.Nil(t, e)
	addr, _, _, e := gocontracts.ServiceDeployWriteRequest(privateKey, client, write.Data, write.ExtraData, write.LTSID, write.ETHAdresses, wrData, temp)
	require.Nil(t, e)
	rAddr, _, _, e := gocontracts.ServiceDeployReadRequest(privateKey, client, addr, data)
	require.Nil(t, e)
	fmt.Println("Read address ", rAddr)

	calypsoAddr := *cal
	_, e = gocontracts.ServiceAddWriteRequest(privateKey, calypsoAddr, addr, client)
	require.Nil(t, e)
	_, e = gocontracts.ServiceAddReadRequest(privateKey, calypsoAddr, rAddr, client)
	require.Nil(t, e)
	wr, e := gocontracts.ServiceGetWriteRequest(privateKey, client, addr)
	require.Nil(t, e)
	require.True(t, wr.U.Equal(write.U))
	_, e = gocontracts.ServiceGetRead(rAddr, client, privateKey)
	require.Nil(t, e)
	dkr := &DecryptKey{
		Read:  rAddr,
		Write: addr,
	}

	dkReply, e := calypsoClient.DecryptKey(dkr)
	require.Nil(t, e)
	require.Equal(t, len(write.Cs), len(dkReply.Cs))
	for i, c := range write.Cs {
		require.True(t, c.Equal(dkReply.Cs[i]))
	}
	require.True(t, dkReply.X.Equal(calypsoClient.ltsReply.X))
	key, e := DecodeKey(cothority.Suite, calypsoClient.ltsReply.X, dkReply.Cs, dkReply.XhatEnc, signer.Ed25519.Secret)
	require.Nil(t, e)
	fmt.Println(string(key))

}

type structPK struct {
	PrivateKey ecdsa.PrivateKey
}
