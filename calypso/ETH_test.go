package calypso

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/byzcoin"
	"github.com/dedis/cothority/darc"
	"github.com/dedis/cothority/darc/expression"
	"github.com/dedis/onet"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func getLTS(t *testing.T) (*CreateLTSReply, darc.ID, *Client, *darc.Signer) {
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
	c, _, err := byzcoin.NewLedger(msg, false)
	require.Nil(t, err)
	//Create a Calypso Client (Byzcoin + Onet)
	calypsoClient := NewClient(c)
	//Invoke CreateLTS
	ltsReply, err := calypsoClient.CreateLTS()
	calypsoClient.ltsReply = ltsReply
	require.Nil(t, err)
	require.NotNil(t, ltsReply.LTSID)
	require.NotNil(t, ltsReply.X)
	fmt.Println("Before: ", reader1.Ed25519.Secret)
	darc1 := darc.NewDarc(darc.InitRules([]darc.Identity{provider1.Identity()},
		[]darc.Identity{provider1.Identity()}), []byte("Provider1"))
	// provider1 is the owner, while reader1 is allowed to do read
	darc1.Rules.AddRule(darc.Action("spawn:"+ContractWriteID),
		expression.InitOrExpr(provider1.Identity().String()))
	darc1.Rules.AddRule(darc.Action("spawn:"+ContractReadID),
		expression.InitOrExpr(reader1.Identity().String()))
	require.NotNil(t, darc1)
	_, err = calypsoClient.SpawnDarc(signer, d, *darc1, 10)
	require.Nil(t, err)
	fmt.Println("After ", reader1.Ed25519.Secret)
	return ltsReply, darc1.GetBaseID(), calypsoClient, &reader1

}
func TestCreateLTS_ETH(t *testing.T) {
	_, gDarc, calypsoClient, signer := getLTS(t)

	write := NewWrite(cothority.Suite, calypsoClient.ltsReply.LTSID, gDarc, calypsoClient.ltsReply.X, []byte("Sabrina"))
	client, e := ethclient.Dial("http://127.0.0.1:7545")
	defer client.Close()
	require.Nil(t, e)
	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
	if err != nil {
		log.Fatal(err)
	}
	temp := make([][]byte, 0)
	for _, point := range write.Cs {
		marshalled, e := point.MarshalBinary()
		require.Nil(t, e)
		temp = append(temp, marshalled)
	}
	address1 := common.BytesToAddress([]byte("8A749D6B91C35b8a5Ce812278C73C988a97790aA"))
	address2 := common.BytesToAddress([]byte("F28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e"))
	address3 := common.BytesToAddress([]byte("09E360FeD8580641CE129252417a8646709196f5"))
	policy := []common.Address{address1, address2, address3}
	write.ETHAdresses = policy
	data, e := calypsoClient.ltsReply.X.MarshalBinary()
	require.Nil(t, e)
	wrData, e := write.U.MarshalBinary()
	require.Nil(t, e)
	addr, _, _, e := gocontracts.ServiceDeployWriteRequest(privateKey, client, write.Data, write.ExtraData, write.LTSID, write.ETHAdresses, wrData, temp)
	rAddr, _, _, e := gocontracts.ServiceDeployReadRequest(privateKey, client, addr, data)
	require.Nil(t, e)
	_, e = gocontracts.ServiceGetWriteRequest(privateKey, client, addr)
	require.Nil(t, e)
	_, e = gocontracts.ServiceGetRead(rAddr, client, privateKey)
	require.Nil(t, e)
	dkr := &DecryptKey{
		Read:  rAddr,
		Write: addr,
	}

	dkReply, e := calypsoClient.DecryptKey(dkr)
	require.Nil(t, e)
	fmt.Println("DkReply cs ", dkReply.Cs)
	fmt.Println("Ethereum CS", temp)
	require.Equal(t, len(write.Cs), len(dkReply.Cs))
	for i, c := range write.Cs {
		require.True(t, c.Equal(dkReply.Cs[i]))
	}
	require.True(t, dkReply.X.Equal(calypsoClient.ltsReply.X))
	secret, e := DecodeKey(cothority.Suite, calypsoClient.ltsReply.X, dkReply.Cs, dkReply.XhatEnc, signer.Ed25519.Secret)
	require.Nil(t, e)
	fmt.Println(string(secret))

}
