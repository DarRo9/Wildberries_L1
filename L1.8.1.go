//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"log"
)

//Функция установки бита, index = позиция , bit = нужный бит
func bit_operation(num *int64, index uint, bit uint) { 
	if bit == 1 {
		//Установка единицы в нужную позицию путем наложенмя маски и логического ИЛИ(+)
		*num |= int64(1) << index 
	} else if bit == 0 {
		//Установка нуля в нужную позицию путем наложенмя маски и логического И(*)
		*num &= ^(int64(1) << index) 
	} else {
		log.Fatal("bit must be 0 or 1")
	}
}

func main() {
	var num int64 = 1
	//Изменение первого бита на ноль
	bit_operation(&num, 0, 0) 
	fmt.Println(num)
	//Изменение 4-го бита на единцу
	bit_operation(&num, 3, 1) 
	fmt.Println(num)
}