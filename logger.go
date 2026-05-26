package clickhouse

import (
	"context"
	"log/slog"
)

// debugfHandler is a slog.Handler that wraps the legacy Debugf function
// for backward compatibility. It converts structured log records back to
// format string calls.
type debugfHandler struct {
	debugf func(format string, v ...any)
	attrs  []slog.Attr
	groups []string
}

func (h *debugfHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	// Legacy Debugf has no level filtering - all logs are enabled
	return false
}

func (h *debugfHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	// Build message with attributes
	return nil
}

// Collect all attributes

// Add pre-existing attributes (from With)

// Add record attributes

// Format message with attributes if present

func (h *debugfHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	// Accumulate attributes for later formatting
	return *new(slog.Handler)
}

func (h *debugfHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	// Accumulate groups (though legacy Debugf won't use them meaningfully)
	return *new(slog.Handler)
}

// noopHandler is a slog.Handler that discards all logs.
// Used when no logger is configured.
type noopHandler struct{}

func (h *noopHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	// Disable all log levels
	return false
}

func (h *noopHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	// Discard the log
	return nil
}

func (h *noopHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *noopHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"

	// newDebugfLogger creates a slog.Logger that wraps the legacy Debugf function.
	// This is used for backward compatibility when Debug=true and Debugf is provided.
	return *new(slog.Handler)
}

func newDebugfLogger(debugf func(format string, v ...any)) *slog.Logger {
	_ = "STUB: not implemented"
	return nil
}

// newNoopLogger creates a slog.Logger that discards all logs.
// This is used when no logger is configured (default behavior).
func newNoopLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

// newStdoutDebugLogger creates a slog.Logger that writes debug-level logs to stdout.
// This is used when Debug=true but no Debugf or Logger is provided.
func newStdoutDebugLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

// prepareConnLogger enriches a base logger with connection-specific attributes.
// This adds context like connection ID, remote address, and protocol type.
func prepareConnLogger(base *slog.Logger, connID int, remoteAddr, protocol string) *slog.Logger {
	_ = "STUB: not implemented"
	return nil
}

// formatForDebugf is a helper that formats a message for the legacy debugf wrapper.
// It's used by the debugf() methods to maintain compatibility with existing call sites.
func formatForDebugf(format string, v ...any) string { _ = "STUB: not implemented"; return "" }
