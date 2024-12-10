package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var age int = 90
	fmt.Println("Age", age)
	var gpa float32 = 9.02
	fmt.Println("GPA", gpa)

	// cant do arithmetic ops accross mulitple datatypes

	var num1 int = 9

	var num2 float32 = 2.5

	fmt.Println(float32(num1) * num2)

	// int division is math floor

	// STRINGS IN GO

	var name string = "Parth\nKalia"

	var name2 string = `Parth
Kalia`

	// double quotes and backticks for string, double quotes cant have formatting, in backtick, will be printed as it was written

	fmt.Println(name)
	fmt.Println(name2)

	// u can use the inbuilt len fuction to find string's length but it counts the bytes and apart from normal ascii characters which are 1 byte in size, others are 2

	// thus u need to count the rune -> rune is a character in GO

	// for that import utf8

	var word string = "Hello!"
	fmt.Println(len(word))
	var word2 string = "©"

	// number of bytes
	fmt.Println(len(word2))

	// rune count in string
	fmt.Println(utf8.RuneCountInString(word2))

	var letter rune = 'a'

	fmt.Println(string(letter))

	// BOOLEANS

	var isVirgin bool = false

	fmt.Println(isVirgin)

	// u can directly define variables without datatypes using :=

	name3 := "Parth"

	fmt.Println(name3)

	var name4 = "L"
	fmt.Println(name4)

	// specify type explicitly when requires

}
