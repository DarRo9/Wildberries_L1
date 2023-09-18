//Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "snow dog sun"
	// Разделение строки на массив со строками в качестве элементов
	cut := strings.Fields(str1) 
	for i := 0; i < len(cut)/2; i++ {
		cut[i], cut[len(cut)-1-i] = cut[len(cut)-1-i], cut[i] 
	}
	//Объединение массива обратно в строку
	result := strings.Join(cut, " ") 
	fmt.Println(result)
}
