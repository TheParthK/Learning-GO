package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r, err := intDivision(10, 3)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case r == 0:
		fmt.Printf("Quotient = %v", q)
	default:
		fmt.Printf("Quotient = %v \nRemainder = %v", q, r)
	}

	switch r {
	case 0:

	}
}

func intDivision(num int, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("error: divide by zero")
	} else {
		return num / den, num % den, nil
	}
}
