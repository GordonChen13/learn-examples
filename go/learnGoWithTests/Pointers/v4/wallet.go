package v3

import "fmt"

type BitCoin float64
type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if w.balance < amount {
		return NewError(ErrorOverDraw)
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() BitCoin{
	return w.balance
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
