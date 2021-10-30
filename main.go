package main

import (
	"fmt"

	"github.com/chrisgardner402/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("chris")
	account.Deposit(1000)
	fmt.Println(account.Balance())
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
