package main

import (
	"fmt"

	"github.com/chrisgardner402/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "Hello"
	def := "Greeting"
	err := dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)
	err2 := dictionary.Add(word, def)
	if err2 != nil {
		fmt.Println(err2)
	}
}
