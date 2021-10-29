package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	length, uppercase := lenAndUpper("chris")
	fmt.Println(length, uppercase)
}
