package main

import (
	"fmt"
)

func main() {
	e1 := make(chan string)
	values := []interface{}{1, "hello", true, e1}
	for _, val := range values {
		switch v := val.(type) {
		case int:
			fmt.Printf("Type: %T Value: %v\n", v, v)
		case string:
			fmt.Printf("Type: %T Value: %v\n", v, v)
		case bool:
			fmt.Printf("Type: %T Value: %v\n", v, v)
		default:
			fmt.Printf("Type: %T Value: %v\n", v, v)
		}
	}
}

// Создается slice interface{}, в котором могут хрнаиться только int, string, bool, chan различных видов
// Используется type assertion для определения типа в рантайме
