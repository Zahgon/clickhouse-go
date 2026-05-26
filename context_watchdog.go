package clickhouse

import "context"

// contextWatchdog is a helper function to run a callback when the context is done.
// It has a cancellation function to prevent the callback from running.
// Useful for interrupting some logic when the context is done,
// but you want to not bother about context cancellation if your logic is already done.
// Example:
// stopCW := contextWatchdog(ctx, func() { /* do something */ })
// // do something else
// defer stopCW()
func contextWatchdog(ctx context.Context, callback func()) (cancel func()) {
	_ = "STUB: not implemented"
	return nil
}
