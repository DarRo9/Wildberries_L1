//Реализовать все возможные способы остановки выполнения горутины. 

//Завершение горутины с помощью Ctrl+C или TERMINATE и range по каналу 

package main

import (
	"sync"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan int)
	//wg чтобы главная горутина не завершилась раньше остальных
	wg.Add(1)
	
	//Когда канал получит информацию о Ctrl+C или о сигнале TERMINATE программа завершится
	end := make(chan os.Signal,1)

	//Добавление возможности уведомления о нажатии Ctrl+C или о сигнале TERMINATE
	signal.Notify(end, os.Interrupt, syscall.SIGTERM)	
	//Горутина
		go func(){
			for v:= range ch{	//range по каналу берёт значения из канала пока канал не будет закрыт
				fmt.Print(v)
			}
			fmt.Printf("Завершение горутины")
			defer wg.Done()
		}()
	
	for {
		select{
			//Закрытие канала ch, ожидание завершения всех горутин
		case <-end:	
			close(ch)
			//Канал для остановки горутины
			wg.Wait()
			return
			//До получения элемента от канала end в ch будет бесконечно отправляться 9
		default:
			ch <- 9
		}
	}
}
	
