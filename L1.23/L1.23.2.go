//Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	str1 := []string{"go", "is", "awesome"}
	var result []string
	//Если i равняемся индексу i2, который нужно удалить, начнётся следующая итерация и элемент не будет добавлен
	i := 2
	for i2, e := range str1 { 
		if i2 == i { 
			continue
		}
		result = append(result, e) 
	}
	fmt.Println(result)
}
