package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"

	"net/http"
	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type App struct {
	conn   driver.Conn
	signal chan os.Signal
}

func (app *App) invalidPrepare() { _ = "STUB: not implemented"; return }

func (app *App) worker() { _ = "STUB: not implemented"; return }

func (app *App) batch() { _ = "STUB: not implemented"; return }

const ddl = `
CREATE TABLE stress (
	  Col1 UInt8
	, Col2 UUID
	, Col3 DateTime
	, Col4 Array(Array(DateTime))
	, Col5 Map(String, String)
) Engine Null
`

// http://127.0.0.1:8080/debug/pprof/
// http://127.0.0.1:8080/debug/charts/
func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	conn, err := clickhouse_tests.GetConnectionWithOptions(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		MaxOpenConns:    20,
		MaxIdleConns:    15,
		ConnMaxLifetime: 3 * time.Minute,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		// Debug: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.Exec(context.Background(), "DROP TABLE IF EXISTS stress"); err != nil {
		log.Fatal(err)
	}
	if err := conn.Exec(context.Background(), ddl); err != nil {
		log.Fatal(err)
	}
	var (
		app = App{
			conn:   conn,
			signal: make(chan os.Signal),
		}
		signals = []os.Signal{
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGKILL,
		}
	)
	go app.invalidPrepare()
	for i := 0; i < 20; i++ {
		go app.worker()
	}
	signal.Notify(app.signal, signals...)
	{
		signal := <-app.signal
		{
			log.Println("got signal:", signal)
		}
		conn.Exec(context.Background(), "DROP TABLE IF EXISTS stress")
		os.Exit(0)
	}
}
