package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	var counter uint32

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Правильный способ
			atomic.AddUint32(&counter, 1)
			// Неправильный способ, результат недетерменирован
			//counter += 1
		}()
	}
	wg.Wait()

	fmt.Println(counter)

	// Также возможно решение через mutex и через каналы
	// Во внутренней реализации atomic похож на mutex
}
