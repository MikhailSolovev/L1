package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("-------With float64------------")
	start := time.Now()
	a := 2e+25
	b := 10e+25
	fmt.Printf("Mult: %.30e\nDiv: %.30e\nSum: %.30e\nSub: %.30e\n", a*b, a/b, a+b, a-b)
	end := time.Now()
	fmt.Printf("Time: %v\n", end.Sub(start))

	fmt.Println("-------With pkg math/big-------")
	start = time.Now()
	c := new(big.Float)
	d := new(big.Float)
	c.SetString("2e+25")
	d.SetString("10e+25")
	fmt.Printf("Mult: %.30e\nDiv: %.30e\nSum: %.30e\nSub: %.30e\n", new(big.Float).Mul(c, d),
		new(big.Float).Quo(c, d), new(big.Float).Add(c, d), new(big.Float).Sub(c, d))
	end = time.Now()
	fmt.Printf("Time: %v\n", end.Sub(start))

	// Точность		Mult	Div		Sum		Sub
	// float64		15		15		16		15
	// math/big		19		19		abs		abs
	//
	// Номер запуска		1		2		3		4		5		6		7		8		9		10
	// Время float64 (µs)	29.9  	116.2	132.3	80.3	61.9	63.5	265.3	191.8	45.9	49.7
	// Время math/big (µs)	53.4	85.2	212.0	57.3	112.5	74.0	1276.6	410.8	75.3	74.5
	//
	//
	//				float64		math/big
	// Средняя			103.7		243.2
	// Стандартное отклонение	71.0		109.7
	// Относительная ошибка		69%		45%
	//
	// Итог: math/big точнее, но в среднем медленее
}
