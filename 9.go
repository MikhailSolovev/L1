package main

import "fmt"

// Читает данные и отправляет в канал ch1
func ping(d *[]int, cout chan<- int) {
	for _, num := range *d {
		cout <- num
	}
	// Закрывает канал ch1
	close(cout)
}

// Вычисляет произведение и отправляет в канал ch2
func pong(cin <-chan int, cout chan<- int) {
	for num := range cin {
		cout <- 2 * num
	}
	// Закрывает канал ch2
	close(cout)
}

// Выводит данные в stdout из канала ch2
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go ping(&[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, ch1)
	go pong(ch1, ch2)
	for num := range ch2 {
		fmt.Println(num)
	}
}
