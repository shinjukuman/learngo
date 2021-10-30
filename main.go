package main

import (
	"fmt"

	"github.com/chrisgardner402/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("chris")
	account.Deposit(1000)
	fmt.Println(account)
}
