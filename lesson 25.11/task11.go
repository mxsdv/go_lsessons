package main

import (
	"fmt"
)

func main() {
	var mounth int
	fmt.Print("Введите номер месяца: ")
	fmt.Scanln(&mounth)
	//fmt.Println(mounth)
	switch {
	case (mounth >= 1 && mounth <= 2) || mounth == 12:
		fmt.Println("Зима")
	case mounth >= 3 && mounth <= 5:
		fmt.Println("Весна")
	case mounth >= 6 && mounth <= 8:
		fmt.Println("Лето")
	case mounth >= 9 && mounth <= 11:
		fmt.Println("Осень")
	default:
		fmt.Println("Неверный номер месяца")
	}
}
