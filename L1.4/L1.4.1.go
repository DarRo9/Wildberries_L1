/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
Необходима возможность выбора количества воркеров при старте.*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//Даст возможность дождаться завершения всех горутин до завершения главной
var wg sync.WaitGroup

func main(){
	fmt.Println("Введите количество воркеров")
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil{
		panic(err)
	}
	ch := make(chan int)

	//Когда канал получит информацию о Ctrl+C или о сигнале TERMINATE программа завершится
	end := make(chan os.Signal,1)

	//Добавление возможности уведомления о нажатии Ctrl+C или о сигнале TERMINATE
	signal.Notify(end, os.Interrupt, syscall.SIGTERM)	
	wg.Add(num)
	for i := 0; i < num; i++{
		go func(i int){
			for v:= range ch{	//range по каналу берёт значения из канала пока канал не будет закрыт
				fmt.Printf("воркер %v получил цифру %v\n",i,v)
			}
			fmt.Printf("воркер %v завершает работу\n",i)
			defer wg.Done()
		}(i)
	}
	for{
		select{
			//Закрытие канала ch, ожидание завершения всех горутин
		case <-end:	
			close(ch)
			fmt.Println("The channel is closed")
			wg.Wait()
			return
			//До получения элемента от канала end в ch будет бесконечно отправляться 9
		default:
			ch <- 9
		}
	}
}


	
	
