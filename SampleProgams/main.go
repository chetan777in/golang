package main

import "fmt"

func main() {
	stringCount("India")

	var n int
	fmt.Println("Enter the number of prime numbers required ")
	fmt.Scanln(&n)
	prime(n)
}
