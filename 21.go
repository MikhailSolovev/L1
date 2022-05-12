package main

import "fmt"

// DataService работает в формате XML, это внешний интерфейс, который неизменен
type DataService interface {
	SendXMLData()
}

type XMLDocument struct {
}

func (doc XMLDocument) SendXMLData() {
	fmt.Println("Send xml document")
}

// ClientService работает в формате JSON то, чем мы пользуемся
type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXML() string {
	return "<xml></xml>"
}

// JSONDocumentAdapter
type JSONDocumentAdapter struct {
	JsonDocument *JsonDocument
}

func (adapter JSONDocumentAdapter) SendXMLData() {
	adapter.JsonDocument.ConvertToXML()
	fmt.Println("Send xml document")
}

// ClientService
func main() {
	// На клиенте создали JSON файл
	somejsondoc := JsonDocument{}

	// Создали адаптер для JSON файла
	adapter := JSONDocumentAdapter{JsonDocument: &somejsondoc}
	// Конвертация и отправка XML файлов на сервер
	adapter.SendXMLData()
}

// Паттерн Адаптер позволяет не реализовывать связь внутри JSONDocument, а связать их через независимую структуру.
// Например, если у нас поменяется DataService (например, начнет принимать данные в формате CSV)
// или ClientService (начнем оперировать например с TSV файлами), то нам нужно будет изменить только Адаптер
// - Усложнение кода
