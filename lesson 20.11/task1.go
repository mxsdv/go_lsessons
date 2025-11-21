package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e interface{} = nil

	var pa *int = &a
	var pb *float64 = &b
	var pc *bool = &c
	var pd *string = &d
	var pe *interface{} = &e

	// Используем указатели, чтобы не возникала ошибка unused
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))

	// Вывод указателей чтобы можно было скомпилировать
	fmt.Println(*pa, *pb, *pc, *pd, *pe)
}
