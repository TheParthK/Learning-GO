package main

import (
	"fmt"
	"unicode/utf8"
)

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

	// integer
	var num1 int = 10

	fmt.Println(num1)

	num2 := 20

	fmt.Println(num2)

	var gpa float32 = 9.02

	fmt.Println(gpa)

	// int division is math floor

	var num3 int = 7
	var num4 int = 3

	fmt.Println(num3 / num4)

	var n1 int = 9

	var n2 float32 = 2.5

	fmt.Println(n1 * int(n2))

	fmt.Println(float32(n1) * n2)

	fmt.Println(intDivision(9, 2))

	a, b := intDivision(7, 3)

	fmt.Println(a, b)

	arraysInGo()

	mapsInGO()

	loopsInGo()

	stringsRunesBytesInGo()

	structsAndInterfacesInGo()

	pointersInGo()

}

func intDivision(a int, b int) (int, int) {
	return a / b, a % b
}

func arraysInGo() {
	var arr []int32 = []int32{1, 2, 3}

	primes := [5]int32{2, 3, 5, 7, 11} // adding the size is optional

	fmt.Println(arr)
	fmt.Println(primes)

	// fixed length

	var intArr [3]int32 = [3]int32{1, 2, 3}

	// slicing is same as any other language, thank god

	fmt.Println(intArr[0])
	fmt.Println(arr[1:3])

	// print out memory location of an element using the & symbol

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// immediately initiallising array

	arr1 := [...]int32{1, 2, 3} // automatically assigsn the length of the list/array

	fmt.Println(arr1)

	// if we dont have a length thats a slice

	var slice1 []int32 = []int32{4, 3, 2}

	fmt.Printf("Slices in GO \n%v\n", slice1)

	// slices are variable length variable length arrays
	// appending to a slice

	slice1 = append(slice1, 7)

	fmt.Println(slice1)

	slice1 = append(slice1, []int32{8, 9, 10}...)

	fmt.Println(slice1)

	//

	// cap() for capacity of the slice

	// make function can also be used to create a slice

	var slice2 []int32 = make([]int32, 2, 10)

	slice2[0] = 100
	slice2[1] = 99

	slice2 = append(slice2, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	fmt.Println(slice2)

}

func mapsInGO() {
	// map[string]int32

	var details map[string]int32 = map[string]int32{"Parth": 9, "Suraj": 8}

	fmt.Println(details)

	var myMap2 map[string]float32 = map[string]float32{"Parth": 9.02, "John": 7.63}

	fmt.Println(myMap2)

	// a map always return something even if the key doesnt exist

	fmt.Println(myMap2["Parth"])

	var name, exists = myMap2["Rachit"]

	if exists {
		fmt.Printf("The name is %v\n", name)
	} else {
		fmt.Printf("Name Doesn't exist\n")
	}

	fmt.Println(name, exists)

	var myMap3 map[string]int32 = make(map[string]int32)

	fmt.Println(myMap3)

	var peopleDetails map[string]uint8 = map[string]uint8{"Parth": 10, "Rachit": 20}

	fmt.Println(peopleDetails)

	fmt.Println(peopleDetails["Akshit"])

	var _, ok = peopleDetails["Parth"]

	fmt.Println(ok)

	// delete a value from a map, inbuilt delete function

	delete(peopleDetails, "Rachit")

	fmt.Println(peopleDetails)

}

func loopsInGo() {

	var list []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range list {
		fmt.Printf("%v ", i)
	}

	for i := range [5]int32{} {
		fmt.Println(i)
	}

	// thats how u print a range lol

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range [5]uint8{} {
		fmt.Println(i + 1)
	}

	// normal for loop like cpp

	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}

}

func stringsRunesBytesInGo() {
	fmt.Println("STRINGS AND RUNES IN GO")
	var myString string = "résumé"

	fmt.Println(myString)

	fmt.Println(len(myString)) // the len function gives the byte count not the character count so é is counted as 2 bytes bruh as it's not ascii

	// %T returns the type of the variable/constant

	var indexed = myString[0]

	fmt.Printf("%v %T\n", indexed, indexed)

	// if we iterate over the string with 2 varibales, we get the utf-8 encoding and the index (index, ascii_value)

	for i, v := range myString {
		fmt.Println(i, v)
		// utf-8 encoding
	}

	fmt.Println(myString)
	fmt.Printf("Printing rune count in string: %v", utf8.RuneCountInString(myString))

	// string slice

	var stringSlice []string = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	fmt.Println(stringSlice)
	for _, j := range stringSlice {
		fmt.Printf("%v", j)
	}

	// string builder, nhi padd rha maa chudaye

}

// ------------------- STRUCTS IN GO -------------------

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// then u can also have a struct inside a struct

type ownerInfo struct {
	name   string
	engine gasEngine
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg

	// method of a struct, like method of a class

}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func enoughMilesLeft(e engine, miles uint8) {
	if e.milesLeft() >= miles {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You need to fuel up")
	}
}
func structsAndInterfacesInGo() {
	fmt.Println("\n---------- STRUCTS AND INTERFACES ----------")
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var inline6 gasEngine = gasEngine{mpg: 10, gallons: 20}

	fmt.Println(inline6)

	// u can also just use positional arguments

	var inline4 gasEngine = gasEngine{30, 50}

	fmt.Println(inline4)

	var parth ownerInfo = ownerInfo{name: "Parth Kalia", engine: gasEngine{mpg: 20, gallons: 60}}

	var suraj ownerInfo = ownerInfo{name: "Suraj Patil"}

	fmt.Println(parth.engine.mpg)

	fmt.Println(suraj.engine.gallons)

	// declaring an anonymous struct

	var engine2 = struct {
		mpg     uint8
		gallons uint8
	}{10, 30}

	fmt.Println(engine2)

	fmt.Println("----- METHODS IN STRUCTS -----")

	var engine3 gasEngine = gasEngine{mpg: 10, gallons: 2} // line 266

	fmt.Println(engine3.milesLeft())

	var buggatiW16 gasEngine = gasEngine{mpg: 10, gallons: 2}

	var tesla electricEngine = electricEngine{mpkwh: 10, kwh: 1}

	enoughMilesLeft(buggatiW16, 10)
	enoughMilesLeft(tesla, 20)

}

//  ---------- POINTERS IN GO -------------

func pointersInGo() {

	// for delcaring a new pointer

	// var p *int32 // -> this is a int32 pointer

	var p *int32 = new(int32)

	// for it to point to something we need to do var p *int32 = new(int32)

	var i int32

	fmt.Println("The pointer p : ", p)
	fmt.Println("The pointer p points to: ", *p)
	fmt.Println("The value i : ", i)

	// to set value to which a pointer is pointing to

	*p = 10
	fmt.Println("The pointer p points to: ", *p)

	// & to get address loc of a variable

	var marks int32 = 69

	var pointer *int32 = &marks

	fmt.Printf("Marks : %v  Memory Add : %v  Value w/ pointer : %v\n", marks, pointer, *pointer)

	// changing value using the pointer

	*pointer += 10

	fmt.Printf("Updating value using the pointer : %v | %v", *pointer, marks)

	// when u directly do var1 = var2 , it creates a copy of var1 and puts in var 2, so when u update the varibale it doesnt get changed in the other one

	// this is not the same with slices, if u create the copy of a slice the same way, it doesnt really copy

	// slice1  = slice 2 ; slice2[2] = 2 ; this will update slice1 as well

	// pointers in functions

	// old concept -> instead of passing an array in the function param, put in the pointer to the array using &array_name

	// doing this doesnt waste memory anymore

}
