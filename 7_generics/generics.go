package main

import "fmt"

func main() {
	// suppose we have to make a function to find sum of an array but that array can be a int32 float64 or a uint

	// we shouldnt  make mutliple functions for this thing

	// so for that we use generics

	var int32Array []int32 = []int32{1, 2, 3, 4, 5, 6, 7}

	var float32Array []float32 = []float32{1, 2, 3, 4, 5, 6, 7}

	var float64Array []float64 = []float64{1, 2, 3, 4, 5, 6, 7}

	calcArrSum[int32](int32Array) // optionally u can add the type urself
	calcArrSum(float32Array)
	calcArrSum(float64Array)
}

func calcArrSum[T int32 | float32 | float64](arr []T) {
	var sum T
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(isEmpty([]int32{1, 2, 3, 4, 5}))
	fmt.Println(isEmpty([]uint{}))
}

// using the any dt

// any is used for any lol it can be anything

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
