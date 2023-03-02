package bank

import (
	"errors"
	"fmt"
)

type Credit float64

type SimpleBankAccount struct {
	credit Credit
}

func (b SimpleBankAccount) GetBalance() float64 {
	return float64(b.credit)
}

func (b *SimpleBankAccount) Deposit(f float64) error {
	if f < 0 {
		return b.Withdraw(f)
	}
	b.credit += Credit(f)
	return nil
}

func (b *SimpleBankAccount) Withdraw(f float64) error {
	tmpCredit := b.credit - Credit(f)
	if tmpCredit < 0 {
		return errors.New("not enough funds")
	}
	b.credit = tmpCredit
	return nil
}

func (b SimpleBankAccount) String() string {
	return fmt.Sprintf("%f Credits", float64(b.credit))
}

func NewSimpleBankAccount() *SimpleBankAccount {
	return &SimpleBankAccount{0}
}
