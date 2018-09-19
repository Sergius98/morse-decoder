package main

import "fmt"

func main() {
	defer fmt.Println("what?")

	fmt.Println("hi, World!")
}
