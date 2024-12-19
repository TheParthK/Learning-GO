package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	fmt.Println("Go Routines")
	/*
		Go Routines is used for concurrent computing

		// this is similar to parallel computing but not exactly

		in // computing or multi threading, tasks are processed at the same time

		in concurrency, for example, a task has many sub tasks so, so the cpu processes these subtasks from both the tasks one by one


	*/

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// dbCall(i) // use the go keyword to run these task concurrently
		go dbCall(i)

		// program to wait untill everythign is completed import sync for this
	}
	fmt.Printf("\nTotal Execution Time : %v", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is : ", dbData[i])
}
