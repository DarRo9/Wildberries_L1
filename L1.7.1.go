//Реализовать конкурентную запись данных в map.

//Реализация через sync.Map
package main

import (
	"sync"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	var map1 sync.Map
	//wg чтобы главная горутина не завершилась раньше остальных
	wg.Add(5)
	//Горутина
	for i:=1; i<=5; i++ {
		go func(i int){
			defer wg.Done()
			map1.Store(i,i)
		}(i)
	} 
	wg.Wait()
	map1.Range(func(k, v interface{}) bool {
		fmt.Println(v.(int))
		return true
	})
}