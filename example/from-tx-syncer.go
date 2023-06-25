package main

import (
	"fmt"

	"github.com/dstor-team/arsyncer"
)

// syncer from txs
func main() {
	ownerFilterParams := arsyncer.FilterParams{
		OwnerAddress: "cSYOy8-p1QFenktkDBFyRM3cwZSTrQ_J4EsELLho_UE", // arTx from address
	}

	startHeight := int64(804524)
	arNode := "https://arweave.net"
	concurrencyNumber := 100 // runtime concurrency number, default 10
	s := arsyncer.New(startHeight, ownerFilterParams, arNode, concurrencyNumber, 15, "subscribe_tx")

	// run
	s.Run()

	// subscribe tx
	for {
		select {
		case sTx := <-s.SubscribeTxCh():
			// process synced txs
			fmt.Println(sTx)
		}
	}
}
