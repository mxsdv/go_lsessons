package main

import (
	"errors"
	"fmt"
	"time"
)

func zero(int, int) {
	fmt.Println("Zero")
}

func named(x int, y int) {
	fmt.Println(x + y)
}

func namedWithOneType(x, y int) {
	fmt.Println(x + y)
}

func namedReturn(x, y int) int {
	return x + y
}

func namedReturnMultiple(x, y int) (int, int) {
	return x, y
}

func namedReturnNamed(x, y int) (z int, h error) {
	z = y
	if x == 6 {
		h = errors.New("6 6 6 error!")
	}
	return
}

// Нет именованных параметров!

func spreadParams(nums ...int) int {
	var sum int = 0
	fmt.Println(nums)
	for _, num := range nums {
		sum = sum + num
		fmt.Printf("num is: %v, total is: %v\n", num, sum)
	}
	return sum
}

func inc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum(x, y int) int {
	return x + y
}

func Estimator(fn func(int, int) int) func(int, int) int {
	return func(x int, y int) int {
		start := time.Now()
		res := fn(x, y)
		elapsed := time.Since(start)
		fmt.Println("Функция выполнялась: ", elapsed)
		return res
	}
}

func Hello() string {
	fmt.Println("Hello func")
	return "Hello()"
}

func main() {
	defer fmt.Println("test")
	est := Estimator(sum)
	fmt.Println(est(4, 5))

}
