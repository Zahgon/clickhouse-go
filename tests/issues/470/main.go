package main

import (
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

type DatabaseFrame struct {
	name        string
	ColumnNames []string
	rows        *sql.Rows
	columnTypes []*sql.ColumnType
	vars        []any
}

func NewDatabaseFrame(name string, rows *sql.Rows) (DatabaseFrame, error) {
	_ = "STUB: not implemented"
	return *new(DatabaseFrame), nil
}

func (f DatabaseFrame) Next() ([]any, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func NewNativeClient(host string, port uint16, username string, password string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	// debug output ?debug=true
	return nil, nil
}

func main() {
	c, err := NewNativeClient("localhost", 9000, "", "")
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	log.Printf("Reading system.%s", "system.query_thread_log")
	rows, err := c.Query("SELECT * FROM system.query_thread_log")
	if err != nil {
		log.Printf("Query failed")
		log.Fatal(err)
	}
	frame, err := NewDatabaseFrame("db_frame", rows)
	if err != nil {
		log.Println("Cant' construct frame")
		log.Fatal(err)
	}
	//iterate to exhaustion

	for {
		_, ok, err := frame.Next()
		if !ok {
			if err != nil {
				log.Println("Failed on termination")
				log.Println(err)
				break
			}
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	log.Printf("Success with %d rows!!", i)
}
