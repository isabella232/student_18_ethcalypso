package calypso

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/onet"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Workflow", elapsed)
}

func Test_workflow(t *testing.T) {
	defer timeTrack(time.Now())

	cal, e := gocontracts.GetStaticCalypso()
	require.Nil(t, e)
	privateKey, err := crypto.HexToECDSA("1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820")
	require.Nil(t, err)
	roster := GetRoster()
	writeClient := NewClient(*roster, privateKey)
	ltsReply, err := writeClient.CreateLTS(*cal)
	writeClient.ltsReply = ltsReply
	policy := ethereum.GetTestListOfAddresses()
	wr, err := writeClient.AddWrite(ltsReply.LTSID, ltsReply.X, []byte("Sabrina"), *cal, policy)
	require.Nil(t, err)
	readPK, err := crypto.HexToECDSA("8c98cfdae73e83956a9a5e1b9a37a9c48a8949a0a94aa970c6efacecf4d47789")
	require.Nil(t, err)
	readClient := NewClient(*roster, readPK)
	secret, rr, err := readClient.AddRead(*wr, *cal)
	fmt.Println("RR: ", (*rr).Hex())
	fmt.Println("Secret: ", secret.String())
	dk := &DecryptKey{
		Read:  *rr,
		Write: *wr,
	}
	dkr, err := readClient.DecryptKey(dk)
	require.Nil(t, err)
	fmt.Println("decrypt key reply ", dkr)
}

func GetRoster() *onet.Roster {
	l := onet.NewTCPTest(cothority.Suite)
	servers, roster, _ := l.GenTree(3, true)
	l.GetServices(servers, calypsoID)
	return roster
}
