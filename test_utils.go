package bitgo

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type TestParams struct {
	Env        string
	Token      string
	Coin       string
	WalletId   string
	Address    string
	TransferId string
	SequenceId string
	Passphrase string
	Email      string
	Enterprise string
	GasPrice   int
}

func getTestParams() (params *TestParams) {
	return &TestParams{
		Env:   "express_test",                                                        //os.Getenv("ENV"),
		Token: "v2x5a924d9e8186042cef25a971d541e7e90d4499e8fef1630a5f8ced8a60acf852", //os.Getenv("ACCESS_TOKEN"),
		//Token:      "b5917623a6bc3a49fc40f3b7f25d6ddf49f1992d0e46ed3f37c5e029d46601d2", //os.Getenv("ACCESS_TOKEN"),
		//mabel Token:		"v2x2b82d9e33e3945dd58e6daf01da92d1d076c68ee59ff9d0b97b0b10dec2e2f86"
		// sean 2:	v2x93b91018fe97c551c7bb1ecdf57f9e5bd1c604e36e31492a70c226df2e568c7f
		Coin:     "tbtc",                             //os.Getenv("COIN"),
		WalletId: "62f807ec0cc3620007642fe80afae86f", // TBTC Sean's Wallet-go-1 //os.Getenv("WALLET_ID"),
		//WalletId: "62ed99f429259f000768eff5c2a6c23a",    //os.Getenv("WALLET_ID"),
		Address: "2NAjebSdrRWL5G6XCooV9BbMHoNYhZwecvG", //TBTC Sean's Wallet-go-2 //os.Getenv("ADDRESS"),
		//Address:    "2N2KRsoNQ2dZDn5w4u3B5uznSTmNyvB3bAw", //os.Getenv("ADDRESS"),
		TransferId: os.Getenv("TRANSFER_ID"),
		SequenceId: os.Getenv("SEQUENCE_ID"),
		Email:      "xyz@abc.com", //os.Getenv("EMAIL"),
		Passphrase: "changeme",    //"forgetit" // os.Getenv("PASSPHRASE"),
		Enterprise: "62e59e5119dfaf0007e7e36aff6db842",
		GasPrice:   500,
	}
}

func getTestParams_gteth() (params *TestParams) {
	return &TestParams{
		Env:   "express_test",                                                        //os.Getenv("ENV"),
		Token: "v2x5a924d9e8186042cef25a971d541e7e90d4499e8fef1630a5f8ced8a60acf852", //os.Getenv("ACCESS_TOKEN"),
		//Token:      "b5917623a6bc3a49fc40f3b7f25d6ddf49f1992d0e46ed3f37c5e029d46601d2", //os.Getenv("ACCESS_TOKEN"),
		//mabel Token:		"v2x2b82d9e33e3945dd58e6daf01da92d1d076c68ee59ff9d0b97b0b10dec2e2f86"
		// sean 2:	v2x93b91018fe97c551c7bb1ecdf57f9e5bd1c604e36e31492a70c226df2e568c7f
		Coin:     "gteth",                            //os.Getenv("COIN"),
		WalletId: "6305c785e899fe0007ebbbc37a035e4b", // gteth - tINvR //os.Getenv("WALLET_ID"),
		//WalletId: "62ed99f429259f000768eff5c2a6c23a",    //os.Getenv("WALLET_ID"),
		Address: "0xef886999eb6bde17d7e70f13318dcd1e8e2750f6", //GTETH Sean's Wallet - 2 //os.Getenv("ADDRESS"),
		//Address:    "2N2KRsoNQ2dZDn5w4u3B5uznSTmNyvB3bAw", //os.Getenv("ADDRESS"),
		TransferId: os.Getenv("TRANSFER_ID"),
		SequenceId: os.Getenv("SEQUENCE_ID"),
		Email:      "xyz@abc.com", //os.Getenv("EMAIL"),
		Passphrase: "changeme",    //"forgetit" // os.Getenv("PASSPHRASE"),
		Enterprise: "62e59e5119dfaf0007e7e36aff6db842",
		GasPrice:   500000,
	}
}
func getTestBitGo(t *testing.T) (b *BitGo, params *TestParams) {
	params = getTestParams()

	b, err := New(params.Env, time.Minute*5)
	if err != nil {
		t.Fatal(err)
	}

	return b.Token(params.Token).Debug(true), params
}

func getTestCoin(t *testing.T) (coin *BitGo, params *TestParams) {
	b, params := getTestBitGo(t)
	coin = b.Coin(params.Coin)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
