package main

import "fmt"

/* here is where the cool stuff happens */

func retName(operation int) (errCode int, nextOp int) {
	fmt.Printf("Executing operation %d/n", operation)
	errCode = executeOperation(operation)
	if errCode == -2 {
		fmt.Printf("Converting the errCode/n")
		errCode = -1
	}
	nextOp = operation + 1
	return
}

func executeOperation(operation int) int {
	if operation%2 == 0 {
		return -1
	}
	if operation%3 == 0 {
		return -2
	}
	return 0
}

func deferStackExample() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

