package main

import (
	"fmt"

	"github.com/World-of-Cryptopups/eosrpc.go"
)

func main() {
	eos := eosrpc.New("https://wax.greymass.com")

	chain := eos.NewChainAPI()

	info, err := chain.GetInfo()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(info)

	acc, err := chain.GetAccount(eosrpc.AccountProps{AccountName: "5g2vm.wam"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(acc)
}
