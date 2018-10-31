package gocontracts

import (
	"fmt"

	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var calypsoAddr *common.Address
var calypsoError error

func init() {
	if calypsoAddr != nil {
		fmt.Println("whaddup")
		return
	}
	fmt.Println("Is it always nil")
	policy := ethereum.GetTestListOfAddresses()
	privateKey, e := ethereum.GetPrivateKey()
	if e != nil {
		return
	}
	client, e := ethereum.GetClient()
	if e != nil {
		return
	}
	calypsoAddr, calypsoError = DeployEmptyCalypso(privateKey, client, policy)
	if calypsoAddr == nil && calypsoError == nil {
		fmt.Println("what is going on")
	}
	owners := ethereum.GetTestListOfAddresses()
	for _, owner := range owners {
		_, e = ServiceUpdatOwners(owner, *calypsoAddr, client, privateKey)
		if e != nil {
			calypsoAddr = nil
			calypsoError = e
			return
		}
	}
}

func GetStaticCalypso() (*common.Address, error) {
	return calypsoAddr, calypsoError
}
