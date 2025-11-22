// todo: Даны три точки A , B , C на числовой оси. Найти длины отрезков AC и BC и их сумму.
// Примечание: все точки получаем через консольный ввод.
package main

import (
	"fmt"
	"math"
)

func main() {

	var ax, ay, bx, by, cx, cy int32
	fmt.Println("Введите координаты точки A: ")
	fmt.Scan(&ax, &ay)
	fmt.Println("Введите координаты точки B: ")
	fmt.Scan(&bx, &by)
	fmt.Println("Введите координаты точки C: ")
	fmt.Scan(&cx, &cy)

	var AB, BC float64

	AB = float64(math.Sqrt(float64((ax-bx)*(ax-bx) + (ay-by)*(ay-by))))
	BC = float64(math.Sqrt(float64((bx-cx)*(bx-cx) + (by-cy)*(by-cy))))

	fmt.Println("AB =", AB)
	fmt.Println("BC =", BC)
}
