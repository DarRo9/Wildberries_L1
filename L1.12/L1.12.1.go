//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

/*Создание map, итерация по последовательности строк, добавление элементов в map. 
Поверка наличия элемента цикла в карте. При отсутствии элемент добавляется в список*/
func set(a []string) []string {
	counter := make(map[string]int)
	var result []string

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
			result = append(result, elem)
		} 
	}
	
	return result
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("%v\n", set(a))
}
