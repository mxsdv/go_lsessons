package main

import (
	"fmt"
	"math"
)

func findMinDistance(arr []int) {
	valueIndices := make(map[int][]int)

	// Собираем индексы для каждого значения
	for i, value := range arr {
		valueIndices[value] = append(valueIndices[value], i)
	}

	// Для каждого значения находим минимальное расстояние
	for value, indices := range valueIndices {
		if len(indices) < 2 {
			fmt.Printf("Для числа %d нет минимального расстояния т.к. элемент в массиве один\n", value)
			continue
		}

		minDistance := math.MaxInt32
		var minIndices [2]int

		// Ищем пару индексов с минимальным расстоянием
		for i := 0; i < len(indices)-1; i++ {
			for j := i + 1; j < len(indices); j++ {
				distance := indices[j] - indices[i]
				if distance < minDistance {
					minDistance = distance
					minIndices[0] = indices[i]
					minIndices[1] = indices[j]
				}
			}
		}

		fmt.Printf("Для числа %d минимальное расстояние в массиве по индексам: %d и %d\n",
			value, minIndices[0], minIndices[1])
	}
}

func main() {
	mass := []int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2}
	findMinDistance(mass)
}
