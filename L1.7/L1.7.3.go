//Реализовать конкурентную запись данных в map.

//Реализация через структуру с полем sync.RWMutex
package main

import (
	"sync"
	"fmt"
)
type Counters struct {
    mx sync.RWMutex
    m map[string]int
}
func main(){
	var wg sync.WaitGroup
	var map1 Counters
	map1.m = make(map[string]int)
	map1.m["сумма"] = 0
	//wg чтобы главная горутина не завершилась раньше остальных
	wg.Add(5)
	//Горутина
	for i:=1; i<=5; i++ {
		go func(){
			map1.mx.RLock()
			defer wg.Done()
    		defer map1.mx.RUnlock()
			
			map1.m["сумма"]+=1 
		}()
	} 
	wg.Wait()
		fmt.Println(map1.m["сумма"])
//Результат всегда 5, значит доступ к map потокобезопасный		
}