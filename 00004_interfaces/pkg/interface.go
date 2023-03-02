package pkg

type Account interface {
	GetBalance() float64
	Deposit(float64) error
	Withdraw(float64) error
}
