//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

func quicksort(arr []int) []int {
	//Проверка для выхода из рекурсии
	if len(arr) < 2 { 
		return arr
	}
	//Индексы первого и последнего элементов массива
	first, last := 0, len(arr)-1   
	//Опорный индекс                 
	pivot := (first + last) / 2    
	//Перемена индексов опорного элемента с последним элементом                 
	arr[pivot], arr[last] = arr[last], arr[pivot] 
	
	for i := range arr {
		if arr[i] < arr[last] {
			//В цикле меньший элемент помещается на левую сторону опорного элемента
			arr[first], arr[i] = arr[i], arr[first] 
			first++
		}
	}
	//Перемещение опорного элемента, на его окончательное место
	arr[first], arr[last] = arr[last], arr[first] 

	//Вызов первой части
	quicksort(arr[:first]) 
	//Вызов второй части  
	quicksort(arr[first+1:])

	return arr
}

func main() {
	arr := []int{4, 1, 8, 3, 6}
	fmt.Println(quicksort(arr))
}