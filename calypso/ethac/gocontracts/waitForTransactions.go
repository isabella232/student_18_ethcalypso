package gocontracts

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitForTransAction(tx *types.Transaction, client *ethclient.Client, duration int) {
	return
	for {
		_, isPending, e := client.TransactionByHash(context.Background(), tx.Hash())
		if e != nil {
			log.Fatal(e)
		}
		if !isPending {
			return
		}
		time.Sleep(time.Duration(duration) * time.Second)
	}
}
