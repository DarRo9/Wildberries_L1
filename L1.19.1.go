//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 string
	if _, err := fmt.Scan(&str1); err != nil {
		fmt.Println(err)
		return
	}
	//Получение количества символов в строке
	lenght := utf8.RuneCountInString(str1) 
	//Создание среза рун из строки
	reversed := []rune(str1) 
	for i := 0; i < lenght/2; i++ {
		//Перемена местами элементов
		reversed[i], reversed[lenght-1-i] = reversed[lenght-1-i], reversed[i] 
	}
	fmt.Println(string(reversed))
}