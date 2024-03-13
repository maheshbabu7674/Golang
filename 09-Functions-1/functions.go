package main

import "fmt"

// Add function

func add(x int, y int) int { //func add(x ,y int){
	var out = x + y
	return out
}

// Sub function

// func sub(x int, y int) int {
// 	var out = x - y
// 	return out
// }

func main() {
	num1 := 5
	num2 := 3

	result := add(num1, num2) //sub(num1, num2)
	fmt.Println(result)

}

// calc function:

// func calc(x int, y int) (int, int) {
// 	var out1 = x + y
// 	var out2 = x - y
// 	return out1, out2
// }

// func main() {
// 	num1 := 5
// 	num2 := 3

// 	result1, result2 := calc(num1, num2)
// 	fmt.Println(result1, result2)
// }
