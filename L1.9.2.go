//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"
		
func main() {
	var array = [5]int{1,2,3,4,5}
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	//Ожидание остальных операций произойдёт путём отправки 'true' в/из 'finish' 
	finish := make(chan bool)
	
	go func(){ 
		for _, el1 := range array {
		ch1 <- el1
		}  
		close(ch1)
		
	}()

	go func() {
		for el2 := range ch1 {
		ch2 <- el2*2
		}
		close(ch2)
	}()

	go func() {
		for el3 := range ch2 {
			fmt.Println(el3)
		}
		finish <- true
	}()
		<- finish	
}