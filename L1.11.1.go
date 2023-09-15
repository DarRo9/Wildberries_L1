//Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

/*Создание map, итерация по первому множству, добавление элементов в карту. 
Итерация по второму множеству, проверка наличия элемента цикла в карте. При наличии элемент пересечения добавляется в список и возвращается*/
func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} 
	}
	for _, elem := range b {
		if _, ok := counter[elem]; ok {
			result = append(result, elem)
		}
	}
	return result
}

func main() {

	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}