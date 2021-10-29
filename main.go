package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"sushi", "ramen", "tonkatsu"}
	chris := person{name: "chris", age: 20, favFood: favFood}
	fmt.Println(chris)
}
