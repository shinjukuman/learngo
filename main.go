package main

import (
	"fmt"
)

func main() {
	names := []string{"chris", "gardner"}
	names = append(names, "dali", "van", "picasso")
	fmt.Println(names)
}
