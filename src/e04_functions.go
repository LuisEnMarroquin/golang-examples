package main

import "fmt"

func main() {
	fmt.Println(first(5))
}

func first(number int) int { // Will get an int and return an int
	return number * 2
}
