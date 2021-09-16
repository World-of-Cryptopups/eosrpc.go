# eosrpc.go

Simple api wrapper for EOS RPC Api (https://developers.eos.io/manuals/eos/latest/nodeos/rpc_apis/index)

## Usage

```go
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
}
```

## Notes

This library is not guaranteed to work perfectly and all functions are not fully tested, validated and confirmed to work functionally. Please be wary of this and report any issues.

PR's are welcome.

##

#### &copy; 2021 | World of Cryptopups | TheBoringDude
