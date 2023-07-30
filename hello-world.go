package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World!!!")
	fmt.Println("go" + "lang")
	fmt.Println("7.0/3.0", 7.0/3.0, "Answer")
	fmt.Println(!false)
	fmt.Println(!!false) // multiple negations.
	fmt.Println(true && false)

	// Variables
	myName := "Danyal"
	fmt.Println("myName ->", myName)

	// Constants
	const PI = 3.14
	fmt.Println("PI + 5 ->", PI+5)
	fmt.Println("3e20 / PI ->", 3e20/PI)

	// Defaults
	var str string
	fmt.Println("str ->", str, "is uninitialized str === ''", str == "")
	var num int
	fmt.Println("num ->", num, "is uninitialized num === ''", num == 0)
	var floating float64
	fmt.Println("floating ->", floating, "is uninitialized floating === ''", floating == 0)
	var boolean bool
	fmt.Println("boolean ->", boolean, "is uninitialized boolean === ''", boolean == false)

	// Arthematic left shift.
	println("1 << 2 ->", 1<<2)
	println("1 << 3 ->", 1<<3)
	println("1 << 4 ->", 1<<4)

	// Iota is used in constant blocks.
	// They auto increment.
	const (
		_ = iota
		a
		b
		c
		d = 1 << iota
		e
		f
	)
	fmt.Println(a, b) // 1 2 (0 is skipped)
	fmt.Println(c, d)
	fmt.Println(e, f)

	const (
		One = iota + 2
		Two
		Three
		Four
	)
	// 2 3 4 5
	println(One, Two, Three, Four)

	// Used similar to enums
	const (
		_ = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	println()

	rand.Seed(time.Now().UnixNano())
	var randomNum = rand.Intn(10) + 1
	switch randomNum {
	case Sunday:
		println("Sunday")
	case Saturday:
		println("Saturday")
	case Monday:
		println("Monday")
	case Tuesday:
		println("Tuesday")
	case Thursday:
		println("Thursday")
	case Wednesday:
		println("Wednesday")
	case Friday:
		println("Friday")
	default:
		println("Not Found")
	}
}
