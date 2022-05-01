package main

import (
	"fmt"
	"math"
)

type Point struct {
	// Все что начинается с маленькой буквы будет приватным полем
	// Оно не будет доступно из других пакетов при импорте, но из текущего оно доступно
	x float64
	y float64
}

// Конструктор для Point
func New(x, y float64) *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

func (p Point) GetData() {
	fmt.Printf("X: %.4f Y: %.4f", p.x, p.y)
}

func Distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func main() {
	A := New(0, 0)
	B := New(3.0, 4.0)
	fmt.Println("Distance:", Distance(A, B))
}
