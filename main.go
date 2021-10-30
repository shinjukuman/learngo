package main

import (
	"fmt"

	"github.com/chrisgardner402/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("chris")
	fmt.Println(account)
}
