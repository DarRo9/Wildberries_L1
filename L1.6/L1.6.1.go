//Реализовать все возможные способы остановки выполнения горутины. 

//Остановка горутины с помощью отправления сообщения через канал
package main

import (
	"sync"
	"fmt"
	"time"
)

func main(){
	var wg sync.WaitGroup
	stop := make(chan string)
	//wg чтобы главная горутина не завершилась раньше остальных
	wg.Add(1)
	//Горутина
	go func(){
		defer wg.Done()
		for{
		select{
			case <-stop:
				fmt.Println("Завершение горутины")
				return
			default:
				fmt.Println("Горутина работает")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	//Канал для остановки горутины
	stop <- "stop"
	wg.Wait()
}
