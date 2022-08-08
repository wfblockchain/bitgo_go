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
}

func getTestParams() (params *TestParams) {
	return &TestParams{
		Env:        "test", //os.Getenv("ENV"),
		Token:      "b5917623a6bc3a49fc40f3b7f25d6ddf49f1992d0e46ed3f37c5e029d46601d2", //os.Getenv("ACCESS_TOKEN"),
		Coin:       "tbtc", //os.Getenv("COIN"),
		WalletId:   os.Getenv("WALLET_ID"),
		Address:    os.Getenv("ADDRESS"),
		TransferId: os.Getenv("TRANSFER_ID"),
		SequenceId: os.Getenv("SEQUENCE_ID"),
		Email:      os.Getenv("EMAIL"),
		Passphrase: os.Getenv("PASSPHRASE"),
	}
}

func getTestBitGo(t *testing.T) (b *BitGo, params *TestParams) {
	params = getTestParams()

	b, err := New(params.Env, time.Minute*5)
	if err != nil {
		t.Fatal(err)
	}

	return b.Token(params.Token).Debug(false), params
}

func getTestCoin(t *testing.T) (coin *BitGo, params *TestParams) {
	b, params := getTestBitGo(t)
	coin = b.Coin(params.Coin)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
