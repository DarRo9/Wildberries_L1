//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func change(x, y *int) {
	if *x != *y {
		// ^= - побитовое исключающее ИЛИ и оператор присваивания
		*x ^= *y
		*y ^= *x
		*x ^= *y
	}
}

func main() {
	var num1, num2 int

	num1 = 1
	num2 = 2

	change(&num1, &num2)

	fmt.Printf("num1: %d\nnum2: %d\n", num1, num2)
}