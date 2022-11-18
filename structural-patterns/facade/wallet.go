package facade

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("Balance is not sufficient")
	}
	fmt.Println("Wallet valance is Sufficient")
	return nil
}
