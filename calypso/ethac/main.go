package main

import (
	"fmt"
	"log"

	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
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

	addrW, _, _, _ := gocontracts.ServiceDeployWriteRequest(privateKey, client, data, extraData, LTSID, policy)
	addrR, _, _, _ := gocontracts.ServiceDeployReadRequest(privateKey, client, addrW, address1)
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
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Added a writerequest: ", tx.Hash().Hex())
	_, e = gocontracts.ServiceAddReadRequest(privateKey, cal, addrR, client)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Added read request")

}
