package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func benchmarkRead(conn clickhouse.Conn) error { _ = "STUB: not implemented"; return nil }

func benchmarkString(conn clickhouse.Conn) error { _ = "STUB: not implemented"; return nil }

func main() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	})
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
