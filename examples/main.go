package main

import (
	"bitgo"
	"log"
	"time"
	// "github.com/jazzserve/bitgo"
)

func main() {
	b, err := bitgo.New("express_test", time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}

	list, err := b.Token("v2x5a924d9e8186042cef25a971d541e7e90d4499e8fef1630a5f8ced8a60acf852").Coin("tbtc").Debug(true).ListWallets(bitgo.ListParams{
		// list, err := b.Token("{Access token}").Coin("tbtc").Debug(true).ListWallets(bitgo.ListParams{
		AllTokens: true,
	})
	if err != nil {
		log.Fatalf("%#v\n", err.(bitgo.Error))
	}

	for _, w := range list.Wallets {
		log.Println(w.ID, w.Label)
	}
}
