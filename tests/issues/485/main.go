package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func example() error { _ = "STUB: not implemented"; return nil }

//Debug:           true,

func examplePrep(ctx context.Context, conn clickhouse.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	start := time.Now()
	if err := example(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
}
