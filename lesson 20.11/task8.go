// todo: Проверить истинность высказывания:
//"Данное четырехзначное число читается одинаково слева направо и справа налево".

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numStr string
	fmt.Print("Введите четырехзначное число: ")
	fmt.Scan(&numStr)
	if len(numStr) != 4 {
		fmt.Println("Ошибка: введено не четырехзначное число!")
		return
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Ошибка: введено не число!")
		return
	}
	thousands := num / 1000
	hundreds := (num / 100) % 10
	tens := (num / 10) % 10
	units := num % 10
	if thousands == units && hundreds == tens {
		fmt.Println("Число читается одинаково слева направо и справа налево.")
	} else {
		fmt.Println("Число НЕ читается одинаково слева направо и справа налево.")
	}
}
