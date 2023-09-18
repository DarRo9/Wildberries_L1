//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

//Структура с полем, которое представляет значение счетчика
type counter struct {
	mu sync.Mutex
	count int
}

func main() {
	var counter1 counter
	//wg чтобы главная горутина не завершилась раньше остальных
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			//.Lock()/.Unlock() для потокобезопасного изменения значения поля структуры
			counter1.mu.Lock()
			counter1.count++
			counter1.mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			counter1.mu.Lock()
			counter1.count++
			counter1.mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(counter1.count)
}