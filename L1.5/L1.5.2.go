/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
Необходима возможность выбора количества воркеров при старте.*/

package main

import (
	"fmt"
	"log"
	"context"
	"math/rand"
	"time"
	"sync"
)

func main(){
	fmt.Println("Введите количество секунд")
	//wg чтобы главная горутина не завершилась раньше остальных
	var wg sync.WaitGroup
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil{
		log.Fatal(err)
	}
	ch := make(chan int)
	wg.Add(1)
	time := time.Duration(num) * time.Second
	// Контекст для завершения по истечении времени
	ctx, _ := context.WithTimeout(context.Background(), time) 
	
		go func(сh chan int) {
			defer wg.Done()
			//Извлечение значений из канала
			for i := range ch{	
				fmt.Println(i)
			}
		}(ch)

		
			for {
			select { 
			//Закрытие канала по истечении времени, иначе - отправление псевдорандомного числа в канал
			case <-ctx.Done(): 
				fmt.Println("завершился")
				close(ch)
				wg.Wait()
				return
			default:	//range по каналу берёт значения из канала пока канал не будет закрыт
				ch <- rand.Intn(10)
			}
		}
}


	
	
