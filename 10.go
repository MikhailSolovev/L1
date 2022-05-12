package main

import "fmt"

// GetDozens - получение кол-ва десятков числа float32
func GetDozens(f float32) int {
	return int(f) / 10 * 10
}

// Min - поиск минимума в срезе float32 и вывод кол-ва десятков в числе, сложность O(N)
func Min(d []float32) int {
	var m float32
	for i, el := range d {
		if i == 0 || el < m {
			m = el
		}
	}
	return GetDozens(m)
}

// Max - поиск максимума в срезе float32 и вывод кол-ва десятков в числе, сложность O(N)
func Max(d []float32) int {
	var m float32
	for i, el := range d {
		if i == 0 || el > m {
			m = el
		}
	}
	return GetDozens(m)
}

func main() {
	data := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Объявляем пустую структуру, чтобы она не занимала места в map
	var empty struct{}
	// Объявляам map сетов, если нужно сохранить порядок значений интервалов, то можно использовать
	// OrderedMap (go get github.com/elliotchance/orderedmap), извлечение из такого map будет происходить за O(1),
	// но затраты по памяти увеличатся (https://medium.com/swlh/an-ordered-map-in-go-436634692381)
	res := make(map[int]map[float32]struct{})

	for i := Min(data); i <= Max(data); i += 10 {
		res[i] = map[float32]struct{}{}
		// Инициализация внутреннего map
		for _, val := range data {
			if GetDozens(val) == i {
				res[i][val] = empty
			}
		}
		// Проверка интервала на пустоту
		if len(res[i]) == 0 {
			delete(res, i)
		}
	}

	// Вывод
	for interval, values := range res {
		str := fmt.Sprintf("%v: ", interval)
		for k := range values {
			str += fmt.Sprintf("%.1f ", k)
		}
		fmt.Println(str)
	}
}
