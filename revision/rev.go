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
		fmt.Printf("Name Doesn't exist")
	}

	fmt.Println(name, exists)

}
