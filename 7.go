package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type mapc struct {
	sync.WaitGroup // Встраивание WaitGroup
	sync.RWMutex   // встраивание RWMutex
	m              map[int]int
}

func main() {
	// Инициализация
	m := mapc{sync.WaitGroup{}, sync.RWMutex{}, make(map[int]int)}

	// Запись в map
	for i := 0; i < 100; i++ {
		m.Add(1)
		go func(i int) {
			defer m.Done()
			m.Lock() // блокировка критической секции на запись и чтение
			m.m[i] = rand.Intn(100)
			m.Unlock()
		}(i)
	}

	m.Wait()

	// Вывод
	for k, v := range m.m {
		fmt.Printf("%v %v\n", k, v)
	}
}
