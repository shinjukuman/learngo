package main

import (
	"fmt"
)

func main() {
	chris := map[string]string{"name": "chris", "age": "20"}
	for key, value := range chris {
		fmt.Println(key, value)
	}
}
