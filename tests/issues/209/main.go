package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func getClickhouseClient() driver.Conn { _ = "STUB: not implemented"; return *new(driver.Conn) }

// Debug: true,

func main() {
	conn := getClickhouseClient()
	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		var result []struct {
			Test string `ch:"test"`
		}
		sql := `SELECT 'test' AS test FROM system.numbers LIMIT 10`
		if response := conn.Select(context.Background(), &result, sql); response != nil {
			fmt.Println(response.Error())
		}
		fmt.Println(result, conn.Stats())
	})
	http.ListenAndServe(":8080", nil)
}
