# BitGo Golang SDK V2

The BitGo Platform and SDK makes it easy to build multi-signature Bitcoin applications today.

# install packages and build
go install

# Documentation

View [API Documentation](https://www.bitgo.com/api/v2).

# Example Usage
see examples/main.go

### List Wallets
```go
b, err := bitgo.New("express_test", time.Minute)
if err != nil {
	log.Fatal(err.Error())
}

list, err := b.Token("{Access token}").Coin("tbtc").Debug(true).ListWallets(bitgo.ListParams{
	AllTokens: true,
})
if err != nil {
	log.Fatalf("%#v\n", err.(bitgo.Error))
}

for _, w := range list.Wallets {
	log.Println(w.ID, w.Label)
}
```

# Tests
A good way to learn the usage is by looking at the tests.  
To run the tests, need local bitgo express running. Use the BitGoJS repo. Run express either via docker or in debugger via run configuration.