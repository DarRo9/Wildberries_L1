//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"strings"
)


func main() {
	var text interface{}
	text = "text"
	
	switch v1 := text.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		t := fmt.Sprintf("%T", v1)
		if strings.Compare(t[0:4], "chan") == 0 {
			fmt.Println("channel")
		} else {
			fmt.Println("other type")
		}
	}
}
