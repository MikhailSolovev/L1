package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var data = [...]int{2, 4, 6, 8, 10}

	fmt.Println("----------Parallel----------")
	start := time.Now()
	for _, num := range data {
		wg.Add(1)
		go func(num int) {
			fmt.Println(num * num)
			defer wg.Done()
		}(num)
	}

	wg.Wait()
	end := time.Now()
	fmt.Printf("\nIt tooks: %v\n", end.Sub(start))

	fmt.Println("--------Sequentially--------")
	start = time.Now()
	for _, num := range data {
		fmt.Println(num * num)
	}
	end = time.Now()
	fmt.Printf("\nIt tooks: %v\n", end.Sub(start))

	// Из результата видно, что накладные расходы на синхронизацию и на создание горутин достаточно велики
	// Номер запуска 						1		2		3		4		5		6		7		8		9		10
	// Время параллельной (µs)				142.9	145.7	126.3	82.3	45.8	84.5	78.8	96.4	110.0	108.5
	// Время последовательной (µs) 			8.5		10.6	8.9		8.5		10.2	8.1		11.3	8.8		9.1		8.6
	// 											Параллельная			Последовательная
	//				Средняя						102.1					9.3
	//				Стандартное отклонение		29.5					1.0
	//				Относительная ошибка		29%						11%
	//
	//				Последовательная программа работает в среднем в 11 раз быстрее
	//
	// Безусловно это все очень наивные оценки
}
