package main

import "fmt"

// Объявляем пустую структуру, чтобы она не занимала места в map
var Empty struct{}

func intersection(a, b map[int]struct{}) map[int]struct{} {
	res := map[int]struct{}{}
	for el := range a {
		if _, ok := b[el]; ok {
			res[el] = Empty
		}
	}
	return res
}

func main() {
	a := map[int]struct{}{1: Empty, 2: Empty, 7: Empty, 4: Empty, 5: Empty}
	b := map[int]struct{}{4: Empty, 5: Empty, 6: Empty, 7: Empty, 8: Empty}

	// Вывод
	for k := range intersection(a, b) {
		fmt.Print(k, " ")
	}
}
