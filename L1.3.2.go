//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.

package main
import ("fmt"
		"sync")

func main() {
	//Инициализация массива
	list := [5]int{2,4,6,8,10}
	c := 0
	//Мьютекс для безопасного доступа к переменной c
	m := sync.Mutex{}
	//Создание буферизированного канала
	ch := make(chan int, 5)
	//Итерация по массиву, отправка в горутинах данных о результатах вычислений в канал
	for _, el := range list {
		
		go func(el int) {
			ch <- el*el
		}(el)//Важно передать переменную цикла функции в качестве аргумета, иначе при каждой интерации будет передаваться информация про один и тот же аргумент
	}
	// Считываем значения из канала и записываем в результат
	for i := 0; i < len(list); i++ {
		//Временно блокируем доступ для остальных горутин
		m.Lock()
		c += <-ch
		//После изменения переменной возвращаем лоступ для остальных горутин
		m.Unlock() 
	}
	close(ch)
	fmt.Println(c)	

}