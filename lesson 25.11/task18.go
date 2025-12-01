//todo: Заданы массивы

//Даны читатели книг
//var readers_books = [...] string {"id3", "id5", "id9", "id8", "id2", "id1" }

//Даны читатели газет
//var readers_magazines = [...] string { "id8", "id2", "id1", "id4", "id6", "id7", "id10"}

//Найти пользователей кто читает и книги и газеты

package main

import "fmt"

func main() {
	readers_books := []string{"id3", "id5", "id9", "id8", "id2", "id1"}
	readers_magazines := []string{"id8", "id2", "id1", "id4", "id6", "id7", "id10"}

	// карта должна быть map[string]bool
	m := make(map[string]bool)

	for _, v := range readers_books {
		m[v] = true
	}

	var intersection []string

	for _, v := range readers_magazines {
		if m[v] {
			intersection = append(intersection, v)
		}
	}

	fmt.Println("Читают и книги и журналы", intersection)
}
