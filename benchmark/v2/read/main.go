package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func benchmarkRead(conn *sql.DB) error { _ = "STUB: not implemented"; return nil }

func benchmarkString(conn *sql.DB) error { _ = "STUB: not implemented"; return nil }

func main() {
	conn, err := sql.Open("clickhouse", "clickhouse://127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	if err := benchmarkRead(conn); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("benchmarkRead: %v\n", time.Since(start))
	start = time.Now()
	if err := benchmarkString(conn); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("benchmarkString: %v\n", time.Since(start))
}
