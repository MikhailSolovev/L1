package main

import "fmt"

// BinarySearch - работает на отсортированном по возврастанию массиве, если элемент найден,
// то возвращает его индекс, в противном случае -1
func BinarySearch(slice *[]int, item int) int {
	var low, mid, guess int
	high := len(*slice) - 1
	for low <= high {
		mid = (low + high) / 2
		guess = (*slice)[mid]
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(BinarySearch(&[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))  // ok
	fmt.Println(BinarySearch(&[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11)) // ok
	fmt.Println(BinarySearch(&[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))  // ok
	fmt.Println(BinarySearch(&[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 11)) // fail
	// ...
}
