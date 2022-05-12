package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, ch <-chan int, wg *sync.WaitGroup, i int) {
	for {
		select {
		case <-ch:
			fmt.Printf("#%v: %v\n", i, <-ch)
		case <-ctx.Done():
			fmt.Printf("CTRL+C signal\n")
			wg.Done()
			return
		}
	}
}

func main() {
	// Создание контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создание сигнального канала
	signalChan := make(chan os.Signal, 1)
	// Отправка сигнала по CTRL+C
	signal.Notify(signalChan, os.Interrupt)
	// Горутина, которая следит за нажатием CTRL+C, если было нажато отеняет контекст
	go func() {
		select {
		case <-signalChan:
			cancel()
		}
	}()

	// Создание канала
	ch := make(chan int)
	wg := sync.WaitGroup{}

	// Кол-во воркеров
	var N int
	_, err := fmt.Scanln(&N)
	if err != nil {
		panic(err)
	}
	// Запуск воркеров
	wg.Add(N)
	for i := 0; i < N; i++ {
		go worker(ctx, ch, &wg, i)
	}

	// Запись в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()
}
