package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests/std"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var conn *sql.DB

func main() {
	go func() {
		http.ListenAndServe("127.0.0.1:9876", nil)
	}()

	var err error
	conn, err = clickhouse_tests.GetConnectionFromDSN("tcp://127.0.0.1:9000?debug=false")
	if err != nil {
		log.Fatal(err)
	}
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(1)
	conn.SetConnMaxLifetime(15 * time.Minute)
	if err := conn.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
	conn.Exec("DROP TABLE IF EXISTS example")
	_, err = conn.Exec(`
		CREATE EXISTS example (
			country_code FixedString(2),
			os_id        UInt8,
			browser_id   UInt8,
			categories   Array(Int16),
			action_day   Date,
			action_time  DateTime
		) Engine MergeTree() ORDER BY tuple()
	`)
	defer func() {
		conn.Exec("DROP TABLE example")
	}()
	if err != nil {
		log.Fatal(err)
	}

	for range time.Tick(time.Second) {
		log.Println("time", time.Now())
		//go testInsert()
		//go testQuery()
		testQuery()
	}

	if _, err := conn.Exec("DROP TABLE example"); err != nil {
		log.Fatal(err)
	}
}

func testInsert() { _ = "STUB: not implemented"; return }

func testQuery() { _ = "STUB: not implemented"; return }

//log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
