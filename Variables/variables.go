package main

import "fmt"

func main() {
	var num1 = 2
	var num2 = 3

	// or //

	var num0 int // by default int value is zero
	var num3 int = 2
	var num4 int = 3

	//  or //

	var num5, num6 int
	num3, num4 = 6, 9

	var result = num1 + num2
	var result1 = num0 + num3 + num4
	var result2 = num5 + num6

	fmt.Println(result)  // 5
	fmt.Println(result1) //5
	fmt.Println(result2) //15

}
