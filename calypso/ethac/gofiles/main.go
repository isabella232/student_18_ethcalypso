package gofiles

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, e := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Something went wrong! ", e)
	}

}
