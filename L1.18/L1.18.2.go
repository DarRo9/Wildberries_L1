//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Структура с полем, которое представляет значение счетчика
type counter struct { 
	count int64
}

//Потокобезопасное увеличение счетчика на 1
func (c *counter) increment() {
	atomic.AddInt64(&c.count, 1) 
}

//Потокобезопасное чтение счетчика на 1
func (c *counter) value() int64 {
	return atomic.LoadInt64(&c.count) 
}

func main() {
	var counter1 counter
	//wg чтобы главная горутина не завершилась раньше остальных
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ { 
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				counter1.increment()
			}
		}()
	}
	wg.Wait()                  
	fmt.Println(counter1.value())
}
