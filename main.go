package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// single value return type
	fmt.Println(multiply(2, 2))

	// multiple value return type
	totalLength, upperName := lenAndUpper("chris")
	totalLength2, _ := lenAndUpper("chris")
	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2)

	// multiple arguments
	repeatMe("dali", "van", "picasso")
}
