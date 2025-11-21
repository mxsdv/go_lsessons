//todo: Дана сторона квадрата a. Найти его площадь S = a²
// Примечание: сторону квадрата получаем через консольный ввод.

package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
	fmt.Print("Введите сторону квадрата a: ")
	fmt.Scan(&a)
	S := math.Pow(a, 2)
	fmt.Printf("Площадь квадрата S = %.2f\n", S)

}
