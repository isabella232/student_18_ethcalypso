package gocontracts

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitForTransAction(tx *types.Transaction, client *ethclient.Client, duration int) {
	for {
		_, isPending, e := client.TransactionByHash(context.Background(), tx.Hash())
		if e != nil {
			fmt.Println("waiting for transaction")
			log.Fatal(e)
		}
		if !isPending {
			fmt.Println("it stopped pending")
			return
		}
		time.Sleep(time.Duration(duration) * time.Second)
	}
}
