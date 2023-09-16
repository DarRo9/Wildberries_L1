/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)

  Есть вероятность возникновения проблем при нарезке слайса, некоторые символы могут занимать несколько байт
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"fmt"
	"strings"
)

func createHugeString(lenght int) string {
	var str1 strings.Builder 
	for i := 0; i < lenght; i++ {
		str1.WriteString("String")
	}
	return str1.String()
}

// Из функции правильно было бы возвращать строку, а не глобальную переменную
func someFunc() string {   
	// Вместо создания большой строки и ее нарезки лучше сразу создать строку нужного размера
	return createHugeString(100) 
}

func main() {
	// Теперь можно изменить глобальную переменную на локальную
	var justString string   
	justString = someFunc()
	fmt.Println(justString)
}
