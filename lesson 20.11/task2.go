package main

import (
	"fmt"
	"strconv"
)

func main() {

	age := "23"
	foo := "23abc"
	ageNum, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Ошибка при преобразовании age:", err)
	} else {
		fmt.Println("age как число:", ageNum)
	}

	// Преобразование foo в int, ожидается ошибка
	fooNum, err := strconv.Atoi(foo)
	if err != nil {
		fmt.Println("Ошибка при преобразовании foo:", err)
	} else {
		fmt.Println("foo как число:", fooNum)
	}

	// Преобразуйте переменную age в Boolean
	// age := "123abc"
	age = "123abc"
	boolVal, err := strconv.ParseBool(age)
	if err != nil {
		fmt.Println("Ошибка при преобразовании age в Boolean:", err)
	} else {
		fmt.Println("age как Boolean:", boolVal)
	}
	// Преобразуйте переменную flag в Boolean
	// flag := 1
	flag := 1
	boolVal2 := flag != 0
	fmt.Println("flag как Boolean:", boolVal2)

	// Преобразуйте значение в Boolean
	// str_one := "Privet"
	// str_two := ""

	str_one := "Privet"
	str_two := ""
	boolVal, err = strconv.ParseBool(str_one)
	if err != nil {
		fmt.Println("Ошибка при преобразовании str_one в Boolean:")
	} else {
		fmt.Println("str_one как Boolean:", boolVal)
	}

	boolVal, err = strconv.ParseBool(str_two)
	if err != nil {
		fmt.Println("Ошибка при преобразовании str_one в Boolean:", err)
	} else {
		fmt.Println("str_one как Boolean:", boolVal)
	}

	//Преобразуйте false в строку
	var b bool = false
	strBool := strconv.FormatBool(b)
	fmt.Println("false как строка:", strBool)
}
