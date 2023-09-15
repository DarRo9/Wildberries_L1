/*Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. 
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, v := range list {
		i := int(v - math.Mod(v, 10))
		_, ok := groups[i]
		if !ok {
			groups[i] = make([]float64, 0)
		}
		groups[i] = append(groups[i], v)
	}

	for i2, v2 := range groups {
		fmt.Printf("%d: %v\n", i2, v2)
	}
}