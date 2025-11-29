package singlton

import "fmt"

var instance *dbConnection

type dbConnection struct {
	host    string
	port    int
	dbName  string
	dblogin string
	dbPass  string
}

func init() {
	fmt.Println("Сработал инит")
}

func GetInstance() *dbConnection {
	if instance == nil {
		instance = &dbConnection{"localhost", 5432, "to-do", "to-do", "123"}
	}
	return instance
}
