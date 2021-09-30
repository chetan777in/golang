package main

import "fmt"

func foo() (result string) {
	defer func() {
		result = "Change World" // change value at the very last moment
	}()
	return "Hello World"
}

func main() {
	fmt.Println(foo())
}
