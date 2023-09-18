//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func calc(num1, num2 *big.Int, operator string) *big.Int { 
	result := new(big.Int) // создаем результирующую переменную
	switch operator {   // определяем переданынй опертор и в кейсах манипулируем числами с методоми пакета big
	case "+":
		return result.Add(num1, num2)
	case "-":
		return result.Sub(num1, num2)
	case "*":
		return result.Mul(num1, num2)
	case "/":
		return result.Div(num1, num2)
	}
	return result
}

func main() {
	//Поскольку число слишком большое, оно задаётся строкой
	num1, _ := new(big.Int).SetString("453457568567967898765654656477667854389654238", 10) 
	num2, _ := new(big.Int).SetString("64622819744678293846578292987172783890817667849", 10) 

	fmt.Println(calc(num1, num2, "+"))
	fmt.Println(calc(num1, num2, "-"))
	fmt.Println(calc(num1, num2, "*"))
	fmt.Println(calc(num1, num2, "/"))
}
