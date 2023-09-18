//Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

package main

import "fmt"

func main() {
	str1 := "snow dog sun"
	var str2 string

	for i := len(str1) - 1; i >= 0; i-- {

		slice1 := make([]byte, 0, 20)

		// Записываем в буффер слово до пробела
		for i >= 0 && str1[i] != ' ' {
			slice1 = append(slice1, str1[i])
			i--
		}

		// Переворачиваем слово и записываем в результирующую строку
		if len(slice1) > 0 {
			slice2 := make([]byte, 0, len(slice1))
			for j := len(slice1) - 1; j >= 0; j-- {
				slice2 = append(slice2, slice1[j])
			}

			// Добавляем пробел если нужно
			if i > 0 {
				slice2 = append(slice2, ' ')
			}
			str2 += string(slice2)
		}
	}
	
	fmt.Println(str2)
}
