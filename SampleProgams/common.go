package main

import (
	"fmt"
	"strings"
)

func stringCount(input string) {
	m := make(map[string]int)
	for _, c := range input {
		// Used strings count function to get actual char count, used type conversion
		//fmt.Println(string(c), "==>", strings.Count(input, string(c)))
		m[string(c)] = strings.Count(input, string(c))
		fmt.Println(m)
	}
}

func prime(input int) {
	var count, c int
	var i = 3

	if input >= 1 {
		fmt.Printf("First %d prime numbers are :  ", input)
		fmt.Printf("2 ")
	}

	for count = 2; count <= input; i++ {
		// iteration to check c is prime or not
		for c = 2; c < i; c++ {
			if i%c == 0 {
				break
			}
		}
		if c == i { // c is prime
			fmt.Printf("%d ", i)
			count++ // increment the count of prime numbers
		}
	}
	fmt.Println(input)

}

func demo() {
	fmt.Println("HI")
}
