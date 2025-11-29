package main

import (
	"fmt"
	"pattern/pkg/singlton"
)

func main() {

	fmt.Println(singlton.GetInstance())

}
