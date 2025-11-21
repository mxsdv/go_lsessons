// todo: Данные две переменные:
// age := 36.6
// temperature := 25
// Нужно обменять значения переменных местами. В итого age
// должен равнятся 25 а temperature – 36.6

package main

import (
	"fmt"
	"reflect"
)

func main() {
	age := 36.6
	temperature := 25

	var temp interface{}

	temp = age
	age = float64(temperature)
	temperature = int(temp.(float64))

	fmt.Println(reflect.TypeOf(age))         // float64
	fmt.Println(reflect.TypeOf(temperature)) // int

	fmt.Println("Age: ", age)                 // 25.0
	fmt.Println("temperature: ", temperature) // 36
}
