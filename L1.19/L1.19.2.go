//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {
	var str1 string

	if _, err := fmt.Scan(&str1); err != nil {
		fmt.Println(err)
		return
	}

	reversed := make([]byte, len(str1))

	//Запись в массив байт строку в обратном порядке
	for i := len(str1) - 1; i >= 0; i-- {
		reversed = append(reversed, str1[i])
	}

	str2 := string(reversed)
	fmt.Println(str2)
}
