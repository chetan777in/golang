package main

import (
	"fmt"
	"strings"
)

func foo() (result string) {
	defer func() {
		result = "Change World" // change value at the very last moment
	}()
	return "Hello World"
}

func testMap() map[string]int {
	m := make(map[string]int)
	str := "abc def abc"
	for _, value := range strings.Fields(str) {
		fmt.Println(value)
		m[value] = m[value] + 1
	}

	return m
}

func main() {
	fmt.Println(foo())
	fmt.Println(testMap())
}
