package main

import "fmt"

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := NewBlockchain(nodeID)
	us := UTXOSet{bc}
	us.Reindex()
	count := us.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}