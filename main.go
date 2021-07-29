package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	var x int = 2
	var x32 int32 = 3
	var x64 int64 = 4
	var x16 int16 = 5
	var x8 int8 = 9
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", x32)
	fmt.Printf("%T\n", x64)
	fmt.Printf("%T\n", x8)
	fmt.Printf("%T\n", x16)
	println()
	fmt.Printf("%T, %T, %T, %T, %T", x, x8, x16, x32, x64)

	x1:= string(x32)
	println()
	fmt.Printf("%T", x1)
	println("unsigned integers")

	/*
	The term "unsigned" in computer programming indicates a variable
	that can hold only positive numbers. The term "signed" in computer
	code indicates that a variable can hold negative and positive values.
	The property can be applied to most of the numeric data types
	including int, char, short and long.
	 */

	var unsigned uint = 1
	fmt.Println(unsigned)

	ninjas:= 1e100
	println(ninjas)

	fmt.Println(math.Ceil(9.99999)) //10
	fmt.Println(math.Floor(9.99999)) // 9
	fmt.Println(math.Ceil(3.001)) // 4
	fmt.Println(math.Floor(3.99999)) // 3
	fmt.Println(math.Ceil(math.Max(2.001,5.999)))
	fmt.Println(math.Min(2,5))
	fmt.Println(math.Abs(-4)) //4
	fmt.Println(math.Sqrt2)
	fmt.Println(math.Sqrt(math.Ceil(50)))
	fmt.Println(math.Pow(2, 3))

	fmt.Println("complex numbers")

	fmt.Println(complex(1, 2))


}
