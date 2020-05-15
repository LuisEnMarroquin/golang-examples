package main // Can be the same name as the file

import ( // If only import one package you can avoid ()
	"fmt"
	"strconv"
)

// Declare global variables
var num1 int    // By default is initialized in 0
var txt1 string // By default is initialized in ""
var boo1 bool   // By default is initialized in false

func main() {
	fmt.Println("Hello world", num1, txt1, boo1)

	// Declare and initialize variables
	num2 := 1000
	txt2 := "Hi"
	boo2 := true
	fmt.Println(num2, txt2, boo2)

	// Declare multiple variables
	var n1, n2 int
	n1, n2 = 1, 2
	fmt.Println(n1, n2)

	// Declare and initialize multiple variables
	x1, x2 := 10, "text"
	fmt.Println(x1, x2)

	// Call function
	newFunction()

	// Parse data
	var coin int32 = 2
	newCoin := int(coin)
	fmt.Println(newCoin)

	// Parse int to text
	parsedText := fmt.Sprintf("%d", newCoin)
	fmt.Println(parsedText)

	// Parse int to text with strconv
	parsedText2 := strconv.Itoa(int(coin)) // 'Itoa' function only accepts a generic int
	fmt.Println(parsedText2)

	// Conditional
	state := true
	if state == true {
		fmt.Println(state)
	} else {
		fmt.Println(state)
	}

	// Move some value before evaluating variable
	state2 := true
	if state2 = false; state2 != true {
		fmt.Println("Is false")
	}

	// Switch case
	switch newNumber := 2; newNumber {
	case 1:
		fmt.Println("Case", 1)
	case 2:
		fmt.Println("Case", 2)
	case 3:
		fmt.Println("Case", 3)
	default:
		fmt.Println("Case default")
	}
}

func newFunction() { // If a variable or function starts with upper case it will be public
	fmt.Println("Call success")
}
