/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
Необходима возможность выбора количества воркеров при старте.*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// worker считывает числа из канала и выводит их
func worker(ctx context.Context, num int, ch <-chan int) {
	for {
		
		select {
		case el := <-ch:
			fmt.Printf("воркер %v получил цифру %v\n",num,el)
		case <-ctx.Done(): // канал отмены
			return
				
		}
}

}

func main() {
	fmt.Println("Введите количество воркеров")
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil{
		panic(err)
	}
	
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())


	// Запуск воркеров
	for i := 1; i <= num; i++ {
		go worker(ctx, i, ch)
	}

	// Запись числа 9 в канал, завершение происходит по сигналу контекста
	go func() {
		
		for {
			select {
			case ch <- 9:
			case <-ctx.Done():
			return
			
			}
		}
		
	}()
	
	// При нажатии ctrl+c отправляется сигнал завершения воркерам
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()
	close(ch)
}