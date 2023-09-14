//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"log"
	"time"
	"sync"
	"math/rand"
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
		
		go func(сh chan int) {
			defer wg.Done()
			//Извлечение значений из канала
			for i := range ch{	
				fmt.Println(i)
			}
		}(ch)

		time := time.After(time.Duration(num) * time.Second)
			for {
			select { 
			//Закрытие канала по истечении времени, иначе - отправление псевдорандомного числа в канал
			case <-time:
				fmt.Println("завершился")
				close(ch)
				wg.Wait()
				return
			default:	//range по каналу берёт значения из канала пока канал не будет закрыт
				ch <- rand.Intn(10)
			}
		}
}


	
	
