package main

import (
	"fmt"

	"github.com/dstor-team/arsyncer"
)

// syncer target tx
func main() {
	ownerFilterParams := arsyncer.FilterParams{
		Target: "cSYOy8-p1QFenktkDBFyRM3cwZSTrQ_J4EsELLho_UE", // arTx target address
	}

	startHeight := int64(811484)
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
