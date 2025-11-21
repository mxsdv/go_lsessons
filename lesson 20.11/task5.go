package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	sum := a + b
	diff := a - b
	prod := a * b

	var div float64
	if b != 0 {
		div = float64(a) / float64(b)
	} else {
		fmt.Println("Ошибка: деление на ноль!")
		return
	}

	intDiv := a / b
	mod := a % b

	power := math.Pow(float64(a), float64(b))

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Разность: %d\n", diff)
	fmt.Printf("Произведение: %d\n", prod)
	fmt.Printf("Частное: %.2f\n", div)
	fmt.Printf("Целочисленное деление: %d\n", intDiv)
	fmt.Printf("Остаток от деления: %d\n", mod)
	fmt.Printf("Возведение в степень: %.2f\n", power)
}
