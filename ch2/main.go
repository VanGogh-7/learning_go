package main

import "fmt"

const z int64 = 10

func main() {
	var flag bool
	flag = true
	var isAwesome = true
	x0 := 10
	x0 *= 2
	var complexNum1 = complex(20, 10)
	var complexNum2 = complex(10, 10)
	var x1 int = 10
	var y float64 = 20.2
	var sum1 float64 = float64(x1) + y
	fmt.Println(sum1)
	fmt.Println(complexNum1 * complexNum2)
	fmt.Println(flag)
	fmt.Println(isAwesome)
	fmt.Println("Hello World")
	fmt.Println(0.123 == 0.123)

}
