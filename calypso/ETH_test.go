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

func TestDeploy_Calypso(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("Hello my name is Bjorn")
	extraData := []byte("Hello, my name is Sabrina")
	LTSID := []byte("Hello, my name is Manuel")
	address1 := common.BytesToAddress([]byte("8A749D6B91C35b8a5Ce812278C73C988a97790aA"))
	address2 := common.BytesToAddress([]byte("F28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e"))
	address3 := common.BytesToAddress([]byte("09E360FeD8580641CE129252417a8646709196f5"))
	policy := []common.Address{address1, address2, address3}

	addrW, _, _, _ := gocontracts.ServiceDeployWriteRequest(privateKey, client, data, extraData, LTSID, policy, data)
	addrR, _, _, _ := gocontracts.ServiceDeployReadRequest(privateKey, client, addrW, data)
	fmt.Println(addrR)

	addrWR, _, _, _ := gocontracts.ServiceDeployWriteRequestHolder(privateKey, client)
	fmt.Println("WriteRequestHolder deployed")
	addrRR, _, _, _ := gocontracts.ServiceDeployReadRequestHolder(privateKey, client)
	fmt.Println("ReadRequestHolder deployed")

	owners, _, _, _ := gocontracts.ServiceDeployOwners(privateKey, client, policy)
	fmt.Println("Owners deployed")

	cal, _, _, _ := gocontracts.ServiceCalypso(privateKey, client, owners, addrWR, addrRR)
	fmt.Println("Calypso deployed at ", cal.Hex())
	tx, e := gocontracts.ServiceAddWriteRequest(privateKey, cal, addrW, client)
	require.Nil(t, e)
	fmt.Println("Added a writerequest: ", tx.Hash().Hex())
	_, e = gocontracts.ServiceAddReadRequest(privateKey, cal, addrR, client)
	require.Nil(t, e)
	fmt.Println("Added read request")
}
func getLTS(t *testing.T) (*CreateLTSReply, darc.ID) {
	l := onet.NewTCPTest(cothority.Suite)
	servers, roster, _ := l.GenTree(3, true)
	l.GetServices(servers, calypsoID)

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
	return ltsReply, msg.GenesisDarc.BaseID

}
func TestCreateLTS_ETH(t *testing.T) {
	ltsReply, gDarc := getLTS(t)

	write := NewWrite(cothority.Suite, ltsReply.LTSID, gDarc, ltsReply.X, []byte("I am secret message"))
	client, e := ethclient.Dial("http://127.0.0.1:7545")
	defer client.Close()
	require.Nil(t, e)
	fmt.Println(write.Data)
	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey.X)
	address1 := common.BytesToAddress([]byte("8A749D6B91C35b8a5Ce812278C73C988a97790aA"))
	address2 := common.BytesToAddress([]byte("F28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e"))
	address3 := common.BytesToAddress([]byte("09E360FeD8580641CE129252417a8646709196f5"))
	policy := []common.Address{address1, address2, address3}
	write.ETHAdresses = policy
	data, e := write.U.MarshalBinary()
	require.Nil(t, e)
	fmt.Println("We hvae this many bytes of data ", len(write.Data))
	addr, _, _, e := gocontracts.ServiceDeployWriteRequest(privateKey, client, write.Data, write.ExtraData, write.LTSID, write.ETHAdresses, data)
	fmt.Println("Address of deployed contract", addr.Hex())

	rAddr, _, _, e := gocontracts.ServiceDeployReadRequest(privateKey, client, addr, data)
	require.Nil(t, e)
	fmt.Println("Address of the deployed read request is ", rAddr.Hex())

}
