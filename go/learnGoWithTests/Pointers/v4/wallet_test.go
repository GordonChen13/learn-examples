package v3

import (
	"testing"
)

func TestWallet(t *testing.T)  {
	assertBalance := func(t *testing.T, wallet *Wallet, want BitCoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("god %s want %s, %v", got, want, wallet)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := &Wallet{BitCoin(0.0)}
		wallet.Deposit(BitCoin(10.0))

		want := BitCoin(10.0)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := &Wallet{BitCoin(20.0)}
		wallet.Withdraw(BitCoin(10.0))

		want := BitCoin(10.0)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := &Wallet{BitCoin(10.0)}
		err := wallet.Withdraw(BitCoin(20.0))

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		assertBalance(t, wallet, BitCoin(10.0))

		if err.(*OverDrawError).code != ErrorOverDraw {
			t.Errorf("wanted over draw error but didn't get one")
		}
	})
}
