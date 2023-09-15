//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import ("fmt"
		"sync"
		)

func main() {
	//Даст возможность дождаться завершения всех горутин до завершения главной
	var wg sync.WaitGroup
	var array = [5]int{1,2,3,4,5}
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	wg.Add(15)
	go func(){ 
		for _, el1 := range array {
		defer wg.Done()
		ch1 <- el1
		}  
		close(ch1)
		
	}()
	go func() {
		for el2 := range ch1 {
		defer wg.Done()
		ch2 <- el2*2
		}
		close(ch2)
	}()
	go func() {
		for el3 := range ch2 {
			defer wg.Done()
			fmt.Println(el3)
		}
	}()
	wg.Wait()
}