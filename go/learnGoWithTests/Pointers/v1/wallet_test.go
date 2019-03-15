package v2

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T)  {
	wallet := Wallet{0.0}
	wallet.Deposit(10.0)

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	want := 10.0

	if got != want {
		t.Errorf("god %.2f want %.2f, %v", got, want, wallet)
	}
}
