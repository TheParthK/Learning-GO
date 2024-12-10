package main

import "fmt"

func main() {
	// trying to print something
	fmt.Println("Printing something with a \\n at the end")

	// using Printf
	fmt.Printf("Printing using a f string %v", 10)

	// To add enters into the strign
	// oh variables

	var string1 string = `\nHello
Hi
Bruh
ok`

	fmt.Println(string1)

}
