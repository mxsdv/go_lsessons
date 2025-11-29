package singlton

var instance *dbConnection

type dbConnection struct {
	host    string
	port    int
	dbName  string
	dblogin string
	dbPass  string
}

func GetInstance() *dbConnection {
	if instance == nil {
		instance = &dbConnection{"localhost", 5432, "to-do", "to-do", "123"}
	}
	return instance
}
