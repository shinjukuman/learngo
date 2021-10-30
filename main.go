package main

import (
	"fmt"

	"github.com/chrisgardner402/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	// update success
	word := "Hello"
	dictionary.Add(word, "First")
	err1 := dictionary.Update(word, "Second")
	if err1 != nil {
		fmt.Println(err1)
	}
	def, _ := dictionary.Search(word)
	if def != "" {
		println(def)
		fmt.Println("[update success]")
	}
	// update fail
	err2 := dictionary.Update("Hi", "Second")
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("[update fail]")
	}
	// delete success
	err3 := dictionary.Delete(word)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("[delete success]")
	// delete fail
	err4 := dictionary.Delete(word)
	if err4 != nil {
		fmt.Println(err4)
		fmt.Println("[delete fail]")
	}
	fmt.Println(dictionary)
}
