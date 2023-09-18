//Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func NewSleep(d time.Duration) {
	<-time.After(d) 
}

func main() {
	fmt.Println("1")
	NewSleep(3 * time.Second)
	fmt.Println("2")
}