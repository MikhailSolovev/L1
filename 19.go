package main

import "fmt"

func ReverseString(s string) string {
	// Создаем срез рун из строки и итерируемся с конца
	sint32 := []rune(s)
	res := ""
	for i := len(sint32) - 1; i > -1; i-- {
		res += string(sint32[i])
	}
	return res
}

func main() {
	fmt.Println(ReverseString("Мама мыла раму"))
	fmt.Println(ReverseString("Mom washed the frame"))
	fmt.Println(ReverseString("∀       ∃"))
}
