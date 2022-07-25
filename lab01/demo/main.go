package main

import (
	"fmt"
	"math"
	"reflect"
	pkg "example.com/demo/examplepkg"
)

var (
	i int = 2
	f float64 = 34.5436
	s = "abc"
)

func g(i int64) int64{
	return i*i
}

func sum(n int, stop bool) int {
	s := 0
	// var s int
	i := 0

	for j:=0 ; j <2 ;j++ {
		for {
			s += i
			if i % 2 == 0 {
				continue
			}
			i++
		}
	}
	
	return s
} 

func f2()  {
	defer fmt.Println("abc")
	fmt.Println("def")
}

func f1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println(4)
	f2()
}

func f3() {
	var p *int
	var i int
	_ = i

	// p = &i
	fmt.Println(p, *p)
}

func main() {
	var a bool = true 
	fmt.Println("Hello world!", 1, a)
	fmt.Printf("%f\n", math.Sqrt(2))

	fmt.Println(pkg.Increment(int(f)))

	s := fmt.Sprintf("%v %q suqare", g(int64(i)), "a")
	fmt.Println(s)
	var f = "abc"
	// _ = f 

	// _, t := swap("1","2")

	fmt.Println(reflect.TypeOf(f))
	fmt.Println(sum(5,false))

	f1()
	f3()

	b := [5]int{1,2,3,4,5}

	fmt.Printf("%v\n", b)

	slice := b[:2] // [low,high)
	fmt.Println(slice)
	
	slice[1] = 7
	var s1 []string
	fmt.Printf("%v %v %v\n", len(slice), cap(slice), s1)

	if s1 == nil {
		fmt.Println("nil")
	}

	var s2 = make([]string, 0) 

	for i := 0; i < 10; i++ {
		s2 = append(s2, fmt.Sprintf("%d",i))
		// fmt.Println(cap(s2))
	}

	fmt.Println(s2)

	var s3 = make([]string, 5) 

	for i := 0; i < 15; i++ {
		// s3[i] = fmt.Sprintf("%d  ",i)
		s3 = append(s3, fmt.Sprintf("%d",i))
	}

	fmt.Println(s3, cap(s3))

	var s4 = make([]string, 10) 

	for idx, v := range s4 {
		// s3[i] = fmt.Sprintf("%d  ",i)
		fmt.Println(v)
		s4[idx] = fmt.Sprintf("%d",idx)
	}

	fmt.Println(s3, cap(s3))


	v1 := []int{1,2,3,4,5,6,7}

	for i := range v1 {
		v1[i] = 1
	}

	fmt.Println(v1)

	str := "abcdef" 
	var str1 int32

	for _, v := range str {
		str1 += v
		fmt.Println(int(v), reflect.TypeOf(v))
	}

	for _, v := range []byte(str) {
		fmt.Println(int(v), string(v), reflect.TypeOf(v))
	}
}

