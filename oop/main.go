package main

import (
	"fmt"
	"oop/pkg/profile"
)

func main() {
	obj := profile.NewProfile("admin", "123", false)
	fmt.Println(obj)
	obj.ChangeStatus(true)
	fmt.Println(obj)
}
