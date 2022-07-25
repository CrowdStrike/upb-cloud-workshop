package examplepkg

import "fmt"

func testSecret() {
	fmt.Println("can't touch this!")
}

func Increment(a int) int{
	// testSecret()
	return a + 1
}