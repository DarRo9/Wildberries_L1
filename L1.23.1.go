//Удалить i-ый элемент из слайса.

package main

import "fmt"

func DeleteIdx(str []string, idx int) []string { // принимаем срез и индекс который нужно удалить,
	// возвращаем новый срез
	str = append(str[:idx], str[idx+1:]...) // нарезаем срез append ом
	return str
}

func main() {
	str1 := []string{"go", "is", "awesome"}
	i:= 2
	str2 := append(str1[:i], str1[i+1:]...)
	fmt.Println(str2)
}