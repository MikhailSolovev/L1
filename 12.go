package main

import "fmt"

// Допустим у нас множество с повторениями {cat, cat, dog, cat, tree}
// Его собственные подмножества:
// {{cat, cat, dog}, {cat, cat, cat}, {cat, cat, tree}, {dog, cat, tree}, {cat}, {dog}, {tree}
//	{cat, cat}, {cat, dog}, {cat, tree}, {dog, tree}, {cat, cat, cat, dog}, {cat, cat, cat, tree},
//  {cat, cat, tree, dog}}
// Всего: 14 подмножеств

// All возвращает все УНИКАЛЬНЫЕ комбинации из элементов среза string
// Это powerset за исключением пустого сета
func All(set []string) (subsets [][]string) {
	length := uint(len(set))

	// Проход через все возможные комбинации объектов
	// от 1 (один элемент в подмножестве) до 2^length - 1 (все элементы в подмножестве кроме одного)
	for subsetBits := 1; subsetBits < (1<<length)-1; subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {
			// проверка того содержится ли объект в subset
			// смотрим включен ли object в соответсвующий subsetBits
			if (subsetBits>>object)&1 == 1 {
				// добавляем объект в subset
				subset = append(subset, set[object])
			}
		}
		// добавление подмножества в множество всех подмножеств
		subsets = append(subsets, subset)
	}
	return subsets
}

// CountElSubset - подсчет элементов в подмножестве
func CountElSubset(sub []string) map[string]int {
	res := map[string]int{}
	for _, str := range sub {
		if _, ok := res[str]; ok {
			res[str] += 1
		} else {
			res[str] = 1
		}
	}
	return res
}

// CompareTwoSubsets - проверка двух подмножест на равенство
func CompareTwoSubsets(sub1, sub2 []string) bool {
	elements1 := CountElSubset(sub1)
	elements2 := CountElSubset(sub2)
	if len(elements2) != len(elements1) {
		return false
	}
	for k, v := range elements1 {
		if val, ok := elements2[k]; ok {
			if val != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// CheckInSet - проверка того, что  subset находится в множестве
func CheckInSet(subsets [][]string, targetsubset []string) bool {
	for _, subset := range subsets {
		if CompareTwoSubsets(subset, targetsubset) {
			return true
		}
	}
	return false
}

// DeleteDuplicates - Удаление дубликатов в множестве
func DeleteDuplicates(subsets [][]string) [][]string {
	var res [][]string
	for _, subset := range subsets {
		if !CheckInSet(res, subset) {
			res = append(res, subset)
		}
	}
	return res
}

func main() {
	powerset := DeleteDuplicates(All([]string{"cat", "cat", "dog", "cat", "tree"}))

	// Вывод
	for i, el := range powerset {
		fmt.Println(i+1, el)
	}
}

// Для понимания работы необходимо думать в двоичной системе
// Написано очень костыльно, хорошо бы переписать все на интерфесах, чтобы можно было работать не только с множествами
// типа string
