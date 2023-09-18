/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например: 
abcd — true
abCdefAaf — false
	aabcd — false*/


package main

import (
	"fmt"
	"strings"
)
	
func main() {
	str1 := "qweRty"
	//Объявление map для теста на уникальность
	map1 := make(map[uint8]bool)
	
	//Перевод символов в нижный регистр
	lowstr := strings.ToLower(str1)
	
	for i := 0; i < len(lowstr); i++ {
		//Добавление символа, если его еще нет
		if ok, _ := map1[lowstr[i]]; !ok {
			map1[lowstr[i]] = true
		} else {
			//Если элемент уже есть в map - возвращение "false"
			fmt.Println(fmt.Sprintf("false"))
			return
		}
	}
	fmt.Println(fmt.Sprintf("true"))
}
