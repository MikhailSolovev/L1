package main

import (
	"fmt"
	"unicode"
)

// Empty - пустая структура, чтобы val в map не занимал место
var empty struct{}

// Скорость работы O(N), так как итерируемся по всей строке, проверка ключа в hash map занимает O(1)
func UniqSymbols(s string) bool {
	dict := map[rune]struct{}{}
	for _, sym := range []rune(s) {
		// Lowercase символ
		syml := unicode.ToLower(sym)
		if _, ok := dict[syml]; ok {
			return false
		}
		dict[syml] = empty
	}
	return true
}

func main() {
	fmt.Println("abcd", UniqSymbols("abcd"))
	fmt.Println("abCdefAaf", UniqSymbols("abCdefAaf"))
	fmt.Println("aabcd", UniqSymbols("aabcd"))
}
