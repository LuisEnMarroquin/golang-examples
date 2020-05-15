package main

import "fmt"

func main() { // Only "for" bucle exists

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for o := 1; o <= 5; o++ {
		fmt.Println(o)
	}

	input := 0
	for {
		fmt.Println("Enter secret number:")
		fmt.Scanln(&input)
		if input == 12 {
			break
		}
		fmt.Println("The number is not correct")
	}

	var newVar int = 0
	for newVar < 5 {
		fmt.Printf("\n Go value is '%d'", newVar)
		if newVar == 3 {
			fmt.Printf(" - You are at 3")
			newVar++
			continue
		}
		fmt.Printf(" - Hi ")
		newVar++
	}

	var otherVar int = 0
routine:
	for otherVar < 10 {
		otherVar = otherVar + 2
		if otherVar == 4 {
			fmt.Print("\n Going back to routine")
			otherVar++
			goto routine
		}
		fmt.Printf("\n - %d", otherVar)
	}

}
