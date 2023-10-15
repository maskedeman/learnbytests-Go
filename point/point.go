package point

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("you are broke")

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in deposit: %v", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
