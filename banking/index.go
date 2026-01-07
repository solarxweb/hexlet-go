package banking

import "fmt";

type Account struct {
	ID int
	Owner string
	Balance float64
}

func Deposit(acc *Account, amount float64) {
	fmt.Println("old balance:", acc.Balance);
	acc.Balance += amount;
	fmt.Println("new balance:", acc.Balance)
}