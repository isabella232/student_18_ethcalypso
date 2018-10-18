package calypso

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/byzcoin"
	"github.com/dedis/cothority/darc"
	"github.com/dedis/onet"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestCreateLTS_ETH(t *testing.T) {
	fmt.Println("Bjorn")
	l := onet.NewTCPTest(cothority.Suite)
	servers, roster, _ := l.GenTree(3, true)
	l.GetServices(servers, calypsoID)
	defer l.CloseAll()

	signer := darc.NewSignerEd25519(nil, nil)
	msg, err := byzcoin.DefaultGenesisMsg(byzcoin.CurrentVersion, roster, []string{"spawn:dummy"}, signer.Identity())
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
	require.Nil(t, err)
	require.NotNil(t, ltsReply.LTSID)
	require.NotNil(t, ltsReply.X)

	write := NewWrite(cothority.Suite, ltsReply.LTSID, msg.GenesisDarc.BaseID, ltsReply.X, []byte("I am secret message"))
	fmt.Println("bjornbjornbjorn", write.LTSID)
	client, e := ethclient.Dial("http://127.0.0.1:7545")
	defer client.Close()
	require.Nil(t, e)
	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
	if err != nil {
		log.Fatal(err)
	}
	address1 := common.BytesToAddress([]byte("8A749D6B91C35b8a5Ce812278C73C988a97790aA"))
	address2 := common.BytesToAddress([]byte("F28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e"))
	address3 := common.BytesToAddress([]byte("09E360FeD8580641CE129252417a8646709196f5"))
	policy := []common.Address{address1, address2, address3}
	write.ETHAdresses = policy

	addr, _, _, e := gocontracts.ServiceDeployWriteRequest(privateKey, client, write.Data, write.ExtraData, write.LTSID, write.ETHAdresses)

	fmt.Println("Address of deployed contract", addr.Hex())
	require.Nil(t, e)
}
