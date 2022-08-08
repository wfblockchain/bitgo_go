package main

import (
	"log"
	"time"
	"bitgo"

	// "github.com/jazzserve/bitgo"
)

func main() {
	b, err := bitgo.New("test", time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}

	list, err := b.Token("b5917623a6bc3a49fc40f3b7f25d6ddf49f1992d0e46ed3f37c5e029d46601d2").Coin("tbtc").Debug(true).ListWallets(bitgo.ListParams{
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
