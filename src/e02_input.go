package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1 int
var number2 int
var legend string

func main() {

	fmt.Println("Enter first number:")
	fmt.Scanf("%d", &number1)

	fmt.Println("Enter second number:")
	fmt.Scanf("%d", &number2)

	fmt.Printf("Result: %d\n", number1+number2)

	fmt.Println("Description: ")
	scanner := bufio.NewScanner(os.Stdin) // Other way to enter text by keyboard

	if scanner.Scan() { // If executing Scan() returns something
		legend = scanner.Text()
		fmt.Println("Text:", legend)
	}

}
