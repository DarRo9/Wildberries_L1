/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
		justString = v[:100]
 }

func main() {
 	someFunc()
 }

В функции someFunc, создается большая строка через функцию createHugeString и затем сохраняются первые 100 символов данной строки в переменную justString.
При обрезании строки с помощью v[:100], создаётся новая строка, которая ссылается на тот же массив, что и исходная строка. Поэтому исходная строка не будет освобождена и останется в памяти.
Использование глобальной переменной может привести к непредсказуемому результату.
*/

package main

import (
	"bytes"
	"fmt"
)

func createHugeString(size uint) string {
	str2 := make([]byte, 0, size)
	str2 = ([]byte)("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")
	return string(str2)
}

func someFunc(str *string) {
	v := createHugeString(1 << 10)
	buffer := bytes.NewBufferString(v)
	*str = string(buffer.Next(100))
}

func main() {
	var str string
	someFunc(&str)
	fmt.Println(fmt.Sprintf("В строке %s \n %d символов",str, len(str)))
}