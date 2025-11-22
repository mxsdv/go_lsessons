package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func main() {
	// Настройка подключения c IP, портом, пользователем, паролем и базой
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"CLICKHOUSE_HOST"},
		Auth: clickhouse.Auth{
			Database: "CLICKHOUSE_DB",
			Username: "CLICKHOUSE_USER",
			Password: "CLICKHOUSE_PASSWORD",
		},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping ClickHouse: %v", err)
	}

	var version string
	if err := conn.QueryRow(ctx, "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query error: %v", err)
	}
	fmt.Println("ClickHouse version:", version)
}
