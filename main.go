package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a
	*b = 2000
	fmt.Println(a)
}
