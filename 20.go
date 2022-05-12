package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	res := ""
	for i := len(words) - 1; i > -1; i-- {
		// Копирование в срез byte быстрее, но так как конкатенируются мало элементов,
		// то можно оставить и так
		res += words[i] + " "
	}
	return res
}

func main() {
	fmt.Println(ReverseWords("snow dog sun"))
}
