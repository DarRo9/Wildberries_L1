//Реализовать все возможные способы остановки выполнения горутины. 

//Завершение горутины через контекст 

package main

import (
	"time"
	"fmt"
	"context"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	//Горутина
	go func(){
		for{
		select{
			case <-ctx.Done():
				fmt.Println("Завершение горутины")
				return
			default:
				fmt.Println("Горутина работает")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
}
