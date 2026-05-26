package clickhouse

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/ClickHouse/ch-go/compress"
	chproto "github.com/ClickHouse/ch-go/proto"

	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

const (
	quotaKeyParamName = "quota_key"
	queryIDParamName  = "query_id"
)

type Pool[T any] struct {
	pool *sync.Pool
}

func NewPool[T any](fn func() T) Pool[T] { _ = "STUB: not implemented"; return nil }

func (p *Pool[T]) Get() T { _ = "STUB: not implemented"; return *new(T) }

func (p *Pool[T]) Put(x T) { _ = "STUB: not implemented"; return }

type HTTPReaderWriter struct {
	reader io.Reader
	writer io.WriteCloser
	err    error
	method CompressionMethod
}

// NewReader will return a reader that will decompress data if needed.
func (rw *HTTPReaderWriter) NewReader(res *http.Response) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (rw *HTTPReaderWriter) reset(pw *io.PipeWriter) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// applyOptionsToRequest applies the client Options (such as auth, headers, client info) to the given http.Request
func applyOptionsToRequest(ctx context.Context, req *http.Request, opt *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func dialHttp(ctx context.Context, addr string, num int, opt *Options) (*httpConnect, error) {
	_ = "STUB: not implemented"
	// Get base logger and enrich with connection-specific context
	return nil, nil
}

// Preflight uses hardcoded revision, may break older versions.
// Encoding data over HTTP must use 0. client_protocol_version does not apply to inserts.

func createHTTPRoundTripper(opt *Options) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

type httpConnect struct {
	id              int
	connectedAt     time.Time
	released        bool
	logger          *slog.Logger
	opt             *Options
	revision        uint64
	encodeRevision  uint64
	url             *url.URL
	client          *http.Client
	buffer          *chproto.Buffer
	compression     CompressionMethod
	blockCompressor *compress.Writer
	compressionPool Pool[HTTPReaderWriter]
	blockBufferSize uint8
	handshake       proto.ServerHandshake
}

func (h *httpConnect) serverVersion() (*ServerVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) connID() int { _ = "STUB: not implemented"; return 0 }

func (h *httpConnect) connectedAtTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (h *httpConnect) getLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

func (h *httpConnect) isReleased() bool { _ = "STUB: not implemented"; return false }

func (h *httpConnect) setReleased(released bool) { _ = "STUB: not implemented"; return }

func (h *httpConnect) freeBuffer() { _ = "STUB: not implemented"; return }

func (h *httpConnect) isBad() bool { _ = "STUB: not implemented"; return false }

func (h *httpConnect) queryHello(ctx context.Context, release nativeTransportRelease) (proto.ServerHandshake, error) {
	_ = "STUB: not implemented"
	return *new(proto.ServerHandshake), nil
}

func createCompressionPool(compression *Compression) (Pool[HTTPReaderWriter], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// trick so we can init the reader to something to Reset when we reuse

func (h *httpConnect) writeData(block *proto.Block) error {
	_ = "STUB: not implemented"
	// Saving offset of compressible data
	return nil
}

// Performing compression. Supported and requires

func (h *httpConnect) readData(reader *chproto.Reader, timezone *time.Location, captureBuffer *bytes.Buffer) (*proto.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to decode the block

// Decode failed - check if captured data contains exception marker
// The decode error typically happens because it tries to read the
// "__exception__" marker as binary data

// Exception block size can be up to 16KiB max.
// https://clickhouse.com/docs/interfaces/http#http_response_codes_caveats
// NOTE: When exception happens, a dedicated block is allocated for exception and only
// that exception will be present in the block. So safe to assume whole block size is same
// exception block max size 16KiB.

// Try to read any remaining data
// allocating 2 * maxSize for just in case

// Check if the captured data contains the exception marker

// This is an exception block, parse it

// Not an exception, return the original decode error

// limitedReader is a helper to read from chproto.Reader up to a limit
type limitedReader struct {
	reader *chproto.Reader
	limit  int64
	read   int64
}

func (lr *limitedReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// parseExceptionFromBytes parses ClickHouse exception block
// Format from ClickHouse server (WriteBufferFromHTTPServerResponse.cpp):
//
//	\r\n
//	__exception__
//	\r\n
//	<TAG>  (16 bytes)
//	\r\n
//	<error message>
//	\n
//	<message_length> <TAG>
//	\r\n
//	__exception__
//	\r\n
func parseExceptionFromBytes(data []byte) error { _ = "STUB: not implemented"; return nil }

// bytes

// Find the first __exception__ marker

// Skip past first __exception__\r\n

// Skip \r\n after first marker

// Skip the exception tag (16 bytes) + \r\n

// tag + \r\n

// Now we're at the start of the error message
// Find the second __exception__ marker

// If we can't find second marker, just extract what we can

// Extract error message between tag and second marker
// The error message ends with: \n<message_length> <TAG>\r\n__exception__

// Find the last line which contains "<message_length> <TAG>"

// The last line is "<message_length> <TAG>", error message is everything before it

func (h *httpConnect) sendStreamQuery(ctx context.Context, r io.Reader, options *QueryOptions, headers map[string]string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) sendQuery(ctx context.Context, query string, options *QueryOptions, headers map[string]string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) readRawResponse(response *http.Response) (body []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) createRequest(ctx context.Context, requestUrl string, reader io.Reader, options *QueryOptions, headers map[string]string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check that query doesn't change format

func (h *httpConnect) prepareRequest(ctx context.Context, query string, options *QueryOptions, headers map[string]string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) createRequestWithExternalTables(ctx context.Context, query string, options *QueryOptions, headers map[string]string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) executeRequest(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpConnect) ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// release func is called by connection pool

// check that we got column 1

func (h *httpConnect) close() error { _ = "STUB: not implemented"; return nil }

// discardAndClose discards remaining data and closes the reader.
// Intended for freeing HTTP connections for re-use.
func discardAndClose(rc io.ReadCloser) { _ = "STUB: not implemented"; return }
