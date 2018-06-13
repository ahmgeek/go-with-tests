package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	asserBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()

		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		asserBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		asserBalance(t, wallet, Bitcoin(10))
	})
	
}
