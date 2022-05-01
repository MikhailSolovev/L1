package main

import "fmt"

func main() {
	// Быстрая версия за O(1), НЕ сохраняет порядок
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	// Копировать последний элемент в индекс i
	a[i] = a[len(a)-1]
	// Удалить последний элемент (записать нулевое значение)
	a[len(a)-1] = ""
	// Усечь срез
	a = a[:len(a)-1]
	fmt.Println(a)

	// Медленная версия за O(n), сохраняет порядок
	b := []string{"A", "B", "C", "D", "E"}
	j := 2
	// Выполнить сдвиг a[j+1:] влево на один индекс, операция выполняется за O(n)
	copy(b[j:], b[j+1:])
	// Удалить последний элемент (записать нулевое значение)
	b[len(b)-1] = ""
	// Усечь срез
	b = b[:len(b)-1]
	fmt.Println(b)
}