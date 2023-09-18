//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4,2,7,5,-9,0,-3}
	//Бинарный поиск возможен только в отсортированном массиве
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println("Индекс : ", sort.SearchInts(nums, 0))
}
