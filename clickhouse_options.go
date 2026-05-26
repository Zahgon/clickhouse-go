package clickhouse

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/ClickHouse/ch-go/compress"
)

type CompressionMethod byte

func (c CompressionMethod) String() string { _ = "STUB: not implemented"; return "" }

const (
	CompressionNone    = CompressionMethod(compress.None)
	CompressionLZ4     = CompressionMethod(compress.LZ4)
	CompressionLZ4HC   = CompressionMethod(compress.LZ4HC)
	CompressionZSTD    = CompressionMethod(compress.ZSTD)
	CompressionGZIP    = CompressionMethod(0x95)
	CompressionDeflate = CompressionMethod(0x96)
	CompressionBrotli  = CompressionMethod(0x97)
)

var compressionMap = map[string]CompressionMethod{
	"none":    CompressionNone,
	"zstd":    CompressionZSTD,
	"lz4":     CompressionLZ4,
	"lz4hc":   CompressionLZ4HC,
	"gzip":    CompressionGZIP,
	"deflate": CompressionDeflate,
	"br":      CompressionBrotli,
}

type Auth struct { // has_control_character
	Database string

	Username string
	Password string
}

type Compression struct {
	Method CompressionMethod
	// this only applies to lz4, lz4hc, zlib, and brotli compression algorithms
	Level int
}

type ConnOpenStrategy uint8

const (
	ConnOpenInOrder ConnOpenStrategy = iota
	ConnOpenRoundRobin
	ConnOpenRandom
)

type Protocol int

const (
	Native Protocol = iota
	HTTP
)

func (p Protocol) String() string { _ = "STUB: not implemented"; return "" }

func ParseDSN(dsn string) (*Options, error) { _ = "STUB: not implemented"; return nil, nil }

type Dial func(ctx context.Context, addr string, opt *Options) (DialResult, error)
type DialResult struct {
	conn nativeTransport
}

type HTTPProxy func(*http.Request) (*url.URL, error)

type Options struct {
	Protocol   Protocol
	ClientInfo ClientInfo

	TLS          *tls.Config
	Addr         []string
	Auth         Auth
	DialContext  func(ctx context.Context, addr string) (net.Conn, error)
	DialStrategy func(ctx context.Context, connID int, options *Options, dial Dial) (DialResult, error)

	// Deprecated: Use Logger instead. Debug enables legacy debug logging to stdout.
	// For structured logging with levels, use the Logger field.
	Debug bool

	// Deprecated: Use Logger instead. Debugf provides a custom debug logging function.
	// For structured logging with levels and custom handlers, use the Logger field with
	// a custom slog.Handler.
	Debugf func(format string, v ...any)

	// Logger provides structured logging using Go's standard log/slog package.
	// If nil, no logging occurs (default). To enable logging, provide a configured
	// slog.Logger:
	//
	//   logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	//       Level: slog.LevelDebug,
	//   }))
	//   opts := &clickhouse.Options{
	//       Logger: logger,
	//   }
	//
	// For backward compatibility, if Debug=true and Debugf is set, those will be used
	// instead of Logger.
	Logger *slog.Logger

	Settings             Settings
	Compression          *Compression
	DialTimeout          time.Duration // default 30 second
	MaxOpenConns         int           // default MaxIdleConns + 5
	MaxIdleConns         int           // default 5
	ConnMaxLifetime      time.Duration // default 1 hour
	ConnOpenStrategy     ConnOpenStrategy
	FreeBufOnConnRelease bool              // drop preserved memory buffer after each query
	HttpHeaders          map[string]string // set additional headers on HTTP requests
	HttpUrlPath          string            // set additional URL path for HTTP requests
	HttpMaxConnsPerHost  int               // MaxConnsPerHost for http.Transport
	BlockBufferSize      uint8             // default 2 - can be overwritten on query
	MaxCompressionBuffer int               // default 10485760 - measured in bytes  i.e.

	// HTTPProxy specifies an HTTP proxy URL to use for requests made by the client.
	HTTPProxyURL *url.URL

	// GetJWT should return a JWT for authentication with ClickHouse Cloud.
	// This is called per connection/request, so you may cache the token in your app if needed.
	// Use this instead of Auth.Username and Auth.Password if you're using JWT auth.
	GetJWT GetJWTFunc

	scheme string

	// ReadTimeout is the maximum duration the client will wait for ClickHouse
	// to respond to a single Read call for bytes over the connection.
	// Can be overridden with context.WithDeadline.
	ReadTimeout time.Duration

	// Set a custom transport for the http client.
	// The default transport configured by the library is passed in as an argument.
	TransportFunc func(*http.Transport) (http.RoundTripper, error)
}

func (o *Options) fromDSN(in string) error { _ = "STUB: not implemented"; return nil }

// default for now same as Clickhouse - https://clickhouse.com/docs/en/operations/settings/settings#settings-http_zlib_compression_level

// a level alone doesn't enable compression

// receive copy of Options, so we don't modify original - so its reusable
func (o Options) setDefaults() *Options { _ = "STUB: not implemented"; return nil }

// logger returns the appropriate logger based on the Options configuration.
// Priority order:
// 1. If Debug=true and Debugf is set, use legacy Debugf (backward compatibility)
// 2. If Logger is set, use the provided logger
// 3. If Debug=true but no Debugf is provided, use a default stdout logger
// 4. Otherwise, use a noop logger (no logging)
func (o *Options) logger() *slog.Logger { _ = "STUB: not implemented"; return nil }
