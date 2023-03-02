package main

import (
	"fmt"
	"github.com/naxmefy/00004_interfaces/pkg"
	"github.com/naxmefy/00004_interfaces/pkg/bank"
)

func serrc(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}

func p(a pkg.Account) {
	fmt.Printf("SBA: %f (%v)\n", a.GetBalance(), a)
}

func main() {
	sba := bank.NewSimpleBankAccount()
	p(sba)

	serrc(sba.Deposit(95.4))
	p(sba)

	serrc(sba.Deposit(12.8))
	p(sba)

	serrc(sba.Withdraw(102.1))
	p(sba)

	serrc(sba.Withdraw(12.0))
	p(sba)

	serrc(sba.Deposit(12.8))
	p(sba)

	serrc(sba.Withdraw(12.0))
	p(sba)
}
