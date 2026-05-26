package clickhouse

import (
	"bytes"
	"context"
	"io"
)

// capturingReader wraps a reader and captures all data that passes through it
type capturingReader struct {
	reader io.Reader
	buffer bytes.Buffer
}

func (r *capturingReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// release is ignored, because http used by std with empty release function
func (h *httpConnect) query(ctx context.Context, release nativeTransportRelease, query string, args ...any) (*rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// request encoding

//nolint:bodyclose // false positive

// The HTTPReaderWriter.NewReader will create a reader that will decompress it if needed,

// Wrap reader with capturing reader to detect exceptions

// allow block buffer size to be overridden per query

// ch-go wraps EOF errors

func (h *httpConnect) queryRow(ctx context.Context, release nativeTransportRelease, query string, args ...any) *row {
	_ = "STUB: not implemented"
	return nil
}
