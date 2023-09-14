//Реализовать все возможные способы остановки выполнения горутины. 

//Завершение горутины через указатель

package main

import (
	"time"
	"fmt"
	
)

func main(){
	
	//Указатель
	var stop bool
	//Горутина
	go func(){
		for{
			if stop {
				fmt.Println("Завершение горутины")
				return
			}
			fmt.Println("Горутина работает")
			
		}
	}()
	time.Sleep(1 * time.Second)
	stop = true
}
