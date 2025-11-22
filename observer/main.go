package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Чтение параметров из окружения
	host := os.Getenv("CLICKHOUSE_HOST")
	db := os.Getenv("CLICKHOUSE_DB")
	user := os.Getenv("CLICKHOUSE_USER")
	password := os.Getenv("CLICKHOUSE_PASSWORD")

	// Подключение к ClickHouse
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host},
		Auth: clickhouse.Auth{
			Database: db,
			Username: user,
			Password: password,
		},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Выполнение запроса (укажите вашу таблицу и колонки)
	rows, err := conn.Query(ctx, "SELECT date,log_type,Messege FROM GOAPP")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	// Обработка результатов
	for rows.Next() {
		var col1 time.Time
		var col2 string
		var col3 string
		if err := rows.Scan(&col1, &col2, &col3); err != nil {
			log.Fatalf("Scan error: %v", err)
		}
		fmt.Printf("Row: col1=%d, col2=%s\n", col1, col2, col3)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Rows error: %v", err)
	}
}
