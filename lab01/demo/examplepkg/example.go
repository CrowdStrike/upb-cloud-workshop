package examplepkg

import "fmt"

const(
	Pi = 3.14
)

const (
	a = iota
	b
	c
)


func testSecret() {
	fmt.Println("can't touch this!")
}


func Increment(x int) int{
	testSecret()
	fmt.Println(a,b,c)
	return x + 1
}