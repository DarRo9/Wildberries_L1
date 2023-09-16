//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	var num1, num2 int

	num1 = 1
	num2 = 2

	num1, num2 = num2, num1

	fmt.Printf("num1: %d\nnum2: %d\n", num1, num2)
}