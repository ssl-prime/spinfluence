package main

import (
	"github.com/rightjoin/aqua"
	spin "spinfluence/api/service"
	blockchain "spinfluence/blockchain/service"
)

// main function
func main() {
	server := aqua.NewRestServer()
	server.AddService(&spin.User{})
	server.AddService(&spin.Admin{})
	server.AddService(&blockchain.BlockChain{})
	server.Run()
}
