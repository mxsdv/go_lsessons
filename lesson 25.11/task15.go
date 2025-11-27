package main

import "fmt"

func main() {
	// Исходный массив из 10 элементов
	arr := [10]int{1, 5, 3, 7, 9, 52, 2, 8, 4, 6}

	// Увеличиваем каждый элемент на 1
	for i := 0; i < len(arr); i++ {
		arr[i] += 1
	}

	fmt.Println(arr)
}
