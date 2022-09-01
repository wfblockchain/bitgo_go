package bitgo

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// List Wallets

func TestListWallets(t *testing.T) {
	coin, _ := getTestCoin(t)
	list, err := coin.ListWallets(ListParams{
		Limit:     50,
		AllTokens: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	// sort by label alphabetically (not just lexicographically)
	sort.SliceStable(list.Wallets, func(i, j int) bool {
		si := list.Wallets[i].Label
		sj := list.Wallets[j].Label
		siLower := strings.ToLower(si)
		sjLower := strings.ToLower(sj)
		if siLower == sjLower {
			return si < sj
		}
		return siLower < sjLower
	})
	for _, w := range list.Wallets {
		t.Logf("ID: %s; Label: %s; Balance: %s", w.ID, w.Label, w.BalanceString)
	}
}

// Generate Wallet

func TestGenerateWallet(t *testing.T) {
	coin, params := getTestCoin(t)

	newLabel := randStringRunes(5)

	w, err := coin.GenerateWallet(GenerateWalletParams{
		Label:      fmt.Sprintf("%s-go-%s", params.Coin, newLabel),
		Passphrase: params.Passphrase,
		Enterprise: params.Enterprise,
		GasPrice:   params.GasPrice,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("Newly generated wallet - ", "ID: ", w.ID, "; Label: ", w.Label, "; Balance: ", w.Balance)
}

// Get Wallet

func TestGetWallet(t *testing.T) {
	coin, params := getTestCoin(t)

	w, err := coin.GetWallet(params.WalletId, GetWalletParams{
		AllTokens: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ID: ", w.ID, "; Label: ", w.Label, "; Balance: ", w.Balance)
}

// Update Wallet

func TestUpdateWallet(t *testing.T) {
	coin, params := getTestCoin(t)

	newLabel := "TBTC Sean's Wallet - " + randStringRunes(5)

	wallet, err := coin.UpdateWallet(params.WalletId, UpdateWalletParams{
		Label: newLabel,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if wallet.Label != newLabel {
		t.Errorf("update label %s", newLabel)
	}
}

// Get Wallet By Address

func TestGetWalletByAddress(t *testing.T) {
	coin, params := getTestCoin(t)

	w, err := coin.GetWalletByAddress(params.Address)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ID: ", w.ID, "; Label: ", w.Label, "; Balance: ", w.BalanceString)
}

/***
Send gteth
*/
func TestSendCoins_GTETH(t *testing.T) {
	coin, params := getTestCoin(t, true)

	res, err := coin.SendCoins(params.WalletId, SendCoinsReq{
		Address:          params.Address,
		Amount:           "50000000000",
		WalletPassphrase: params.Passphrase,
		GasPrice:         params.GasPrice,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("txid: ", res.TxID, "; status: ", res.Status, "; value: ", res.Transfer.ValueString)
}

/***
Send tbtc
*/
func TestSendCoins_TBTC(t *testing.T) {
	coin, params := getTestCoin(t)

	res, err := coin.SendCoins(params.WalletId, SendCoinsReq{
		Address:          params.Address,
		Amount:           "5000",
		WalletPassphrase: params.Passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("txid: ", res.TxID, "; status: ", res.Status, "; value: ", res.Transfer.ValueString)
}

func TestStringCompare(t *testing.T) {
	str1, str2 := "WIN", "SYRyQ"
	if str1 < str2 {
		t.Errorf("got %s < %s", str1, str2)
	}

}
