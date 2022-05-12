package main

import "fmt"

var justString string

func createHugeString(n int) string {
	m := make([]rune, n, n)
	for i := 0; i < n; i++ {
		m[i] = 'ж'
	}
	return string(m)
}

func someFunc() {
	v := createHugeString(1 << 10) // 1 << 10 = 2 ** 10 = 1024 символов 'ж'
	fmt.Println(v)
	// Это срез byte, видно, что делаем срез на 100 элементов и, кажется, что должно быть 100 символов 'ж', но их в два
	// раза меньше, так как символы кириллицы занимают два байта. Это плохой способ, лучше оперировать срезом rune
	justString = v[:100]
	fmt.Println(justString)
	// Правильный вариант, видно, что теперь все символы на месте
	justString = string([]rune(v)[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}
