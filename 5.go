package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	N := 5

	// Создание контекста с таймером
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N))
	defer cancel()

	wg := new(sync.WaitGroup)

	// Writer
	go func(c chan int) {
		for i := 0; ; i++ {
			c <- i
			time.Sleep(time.Second * 1)
		}
	}(c)

	wg.Add(1)
	// Reader
	go func(ctx context.Context, c chan int, wg *sync.WaitGroup) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Time out, %v sec paseed\n", N)
				wg.Done()
				return
			case msg := <-c:
				fmt.Println(msg)
			}
		}
	}(ctx, c, wg)

	wg.Wait()
}
