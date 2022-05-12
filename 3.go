package main

import (
	"fmt"
	"math"
)

func square(m []int, c chan int) {
	res := 0
	for _, num := range m {
		res += num * num
	}
	c <- res
}

func main() {
	data := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	// Размер одной порции для воркера
	chunk := len(data) / 3
	res := 0
	// Можно написать интерфейс для выбора кол-ва воркеров, но я ограничусь тем, что у нас их всего три
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go square(data[:chunk], chan1)
	go square(data[chunk:2*chunk], chan2)
	go square(data[2*chunk:], chan3)

	res += <-chan1
	res += <-chan2
	res += <-chan3

	fmt.Println("By exp:", res)
	// Faulhaber's formula
	fmt.Println("By formula:", 2*(2*math.Pow(15, 3)+3*math.Pow(15, 2)+15)/3)
}
