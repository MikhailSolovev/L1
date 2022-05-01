package main

import "fmt"

// Sets the bit at pos in the int64 n.
func setBit(n int64, pos uint8) int64 {
	n |= (1 << pos) // Сдвиг единички на  pos позиций влево и применение побитового ИЛИ
	return n
}

// Clear the bit at pos in the int64 n
func clearBit(n int64, pos uint8) int64 {
	n &= ^(1 << pos) // Сдвиг единички на pos позиций влево и применение побитового НЕ и побитового И
	return n
}

func main() {
	var n int64 = 100
	fmt.Printf("Binary: %b Decimal: %d\n", n, n)
	n = setBit(n, 3)
	n = clearBit(n, 2)
	fmt.Printf("Binary: %b Decimal: %d\n", n, n)
}

// Пример работы, пусть n = 100 или n = b1100100
// Работа setBit(n, 3):
//				- генерация значения b1000
//				- примение побитового ИЛИ
//					b1100100
//					   b1000
//					--------
//					b1101100 или d108
// Работа clearBit(n, 2):
//				- генерация значения b011
//				- применение побитового И
//					b1101100
//					    b011
//					--------
//					b1101000 или d104
