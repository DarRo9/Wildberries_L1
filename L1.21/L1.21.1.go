//Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

//Интерфейс, отправляющий файл только в JSON формате
type Sender interface { 
	SendJson()
}

type XMLDoc struct { 

}

//Метод, который конвертирует XML в JSON
func (j *XMLDoc) ConvertToJSON() { 
	fmt.Println("Convert xml to json")
}

//Адаптер для интерфейса, отправляющего только JSON файл
type AdapterXML struct { 
	xmldoc *XMLDoc
}

//Метод, который удовлетворяет интерфейсу Sender и конверирует XML в JSON
func (a *AdapterXML) SendJSON() { 
	a.xmldoc.ConvertToJSON()
	fmt.Println("send JSON document")
}

//Воркер, принимающий интерфейс Sender и работающий с ним
func Worker(send Sender) { 
	send.SendJSON()
}

func main() {
	xmldoc := &XMLDoc{}
	adapter := &AdapterXML{xmldoc: xmldoc}
	Worker(adapter)
}
