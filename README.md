# CoinSpot API client
A golang API for the CoinSpot API. 

Currently

## Example Usage

```go
package main

import (
	"fmt"
	cs "github.com/threetoes/coinspot-api"
)

func main() {
	api := cs.NewCoinSpotApi("key", "secret")
	balances, err := api.ListBalances()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", balances)
}
```

## Todo
* Proper string field processing. Parts of the API returns 
  strings where it really should be floats
* Truncate floats in requests where applicable
* ~OpenAPI doc (Coinspot docs are a bit vague)~
  * Tidy up OpenAPI spec
  * Add readonly methods
  * Once OpenAPI is confirmed correct, generate
    client from it