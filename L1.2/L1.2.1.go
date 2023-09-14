//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main
import ("fmt"
		"sync")

func main() {
	//Инициализация массива
	list := [5]int{2,4,6,8,10}
	//Чтобы main не завершилась раньше цикла можно использовать sync.WaitGroup
	var wg sync.WaitGroup
	for _, el1 := range list {
		wg.Add(1)
		go func(el1 int) {
			fmt.Printf("%v*%v = %v\n", el1, el1, el1*el1)
			wg.Done()
		}(el1) //Важно передать переменную цикла функции в качестве аргумета, иначе при каждой интерации будет выводиться информация про один и тот же аргумент
	}

	wg.Wait()

}
