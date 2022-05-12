package main

import "fmt"

// Сортировка в среднем за O(N log N), затраты на доп память O(1), иногда может работать и за O(N^2)
// + все элементы должны быть различны
func Quicksort(ar []int) {
	// Базовый случай
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	// Рекурсивный случай
	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	// Выбираем опорный элемент
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		// Находим первый элемент слева, который больше опорного, но лежит левее
		for ; ar[left] < pivot; left++ {
		}

		// Находим первый элемент справа, который меньше опорного, но лежит правее
		for ; ar[right] > pivot; right-- {
		}

		// Если получилось разделить массив на две части, что меньше опорного и что больше опорного
		if left >= right {
			return right
		}

		// Меняем два элемента местами
		ar[left], ar[right] = ar[right], ar[left]
	}
}

func main() {
	arr := []int{3, 4, 2, 9, 0, 1, 5, 7, 8, 10}
	fmt.Println("Before:", arr)
	Quicksort(arr)
	fmt.Println("After:", arr)
}
