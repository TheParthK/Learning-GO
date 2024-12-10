package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Parth")

	fmt.Println(intDivision(5, 2))

	q, r, err := intDivision(7, 2)

	if err != nil {
		fmt.Println(err.Error())
	} else if r == 0 {
		fmt.Printf("The result of division is %v", q)
	} else {
		fmt.Println("Quotient =", q, "Remainder =", r)
		fmt.Printf("The result of division is %v and remainder is %v", q, r)
	}

	if 1 == 1/1 && 2 == 4/2 {
		println("Passed && check")
	}

	if 1 == 1/1 || 2 == 3 {
		println("Passed || check")
	}

	a := 10

	if a == 10 {
		println("WOW, it works without fmt as well")
	}

}

func printMe(name string) {
	fmt.Println("Hello " + name)
}

// adding check for errors

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("error: divide by zero")
	}
	var quotient int = numerator / denominator
	var remainder int = numerator % denominator

	return quotient, remainder, err

}
