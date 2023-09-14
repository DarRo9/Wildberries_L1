//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main
import ("fmt"
		"sync")

func main() {
	//Инициализация массива
	list := [5]int{2,4,6,8,10}
	c := 0
	//Мьютекс для безопасного доступа к переменной c
	m := sync.Mutex{}
	//Чтобы main не завершилась раньше цикла можно использовать sync.WaitGroup
	var wg sync.WaitGroup
	for _, el1 := range list {
		wg.Add(1)
		go func(el1 int) {
		//Временно блокируем доступ для остальных горутин
		m.Lock()
		c+= el1*el1
		//После изменения переменной возвращаем лоступ для остальных горутин
		m.Unlock() 
			wg.Done()
		}(el1) //Важно передать переменную цикла функции в качестве аргумета, иначе при каждой интерации будет выводиться информация про один и тот же аргумент
	}

	wg.Wait()
	fmt.Println(c)
}