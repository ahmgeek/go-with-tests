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

	asserError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
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
	

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		asserBalance(t, wallet, Bitcoin(20))
		asserError(t, err)


	})
}
