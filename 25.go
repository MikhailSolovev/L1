package main

import (
	"fmt"
	"time"
)

func OwnSleep(x int) time.Time {
	// time.After создает канал, затем через заданное время выкидывает в канал объект time.Time,
	// который мы ждем, чтобы принять. На время ожидания работа main goroutine прекращается.
	return <-time.After(time.Second * time.Duration(x))
}

func main() {
	duration := OwnSleep(5)
	fmt.Println(duration)
}
