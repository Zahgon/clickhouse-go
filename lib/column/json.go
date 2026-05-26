package column

import (
	"math"
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/ClickHouse/clickhouse-go/v2/lib/chcol"
)

const JSONDeprecatedObjectSerializationVersion uint64 = 0
const JSONStringSerializationVersion uint64 = 1
const JSONObjectSerializationVersion uint64 = 3
const JSONUnsetSerializationVersion uint64 = math.MaxUint64
const DefaultMaxDynamicPaths = 1024

// JSON implements ClickHouse JSON column on the native protocol layer.
// We choose how to encode the JSON based on serializationVersion.
// It can be either plain string or object type with newer server support.
type JSON struct {
	chType Type
	sc     *ServerContext
	name   string
	rows   int

	serializationVersion uint64
	// pendingNullRows holds null-row counts received before a serialization
	// mode has been latched. A null row carries no mode preference, so we
	// defer the mode decision until a non-null row arrives. Pending nulls
	// are flushed to the latched backing column inside reconcileMode, or —
	// if the whole batch is null — at WriteStatePrefix using String mode.
	pendingNullRows int

	jsonStrings String

	typedPaths      []string
	typedPathsIndex map[string]int
	typedColumns    []Interface

	skipPaths      []string
	skipPathsIndex map[string]int

	totalDynamicPaths int
	dynamicPaths      []string
	dynamicPathsIndex map[string]int
	dynamicColumns    []*Dynamic

	maxDynamicPaths int
	maxDynamicTypes int
}

// jsonMode is the mode preference a value carries when offered to a JSON
// column. It is deliberately separate from the wire-format constants
// (JSONObjectSerializationVersion / JSONStringSerializationVersion) so the
// classifier stays free of wire concerns and so we can express
// "no preference" (jsonModeAny) for null-equivalent inputs.
type jsonMode uint8

const (
	jsonModeAny    jsonMode = iota // nil / typed-nil — does not latch
	jsonModeObject                 // struct, map, *clickhouse.JSON, JSONSerializer
	jsonModeString                 // string, []byte, json.RawMessage, sql.NullString, Stringer, Valuer
)

func (c *JSON) parse(t Type, sc *ServerContext) (_ *JSON, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *JSON) hasTypedPath(path string) bool { _ = "STUB: not implemented"; return false }

func (c *JSON) hasDynamicPath(path string) bool { _ = "STUB: not implemented"; return false }

func (c *JSON) hasSkipPath(path string) bool { _ = "STUB: not implemented"; return false }

// pathHasNestedValues returns true if the provided path has child paths in typed or dynamic paths
func (c *JSON) pathHasNestedValues(path string) bool { _ = "STUB: not implemented"; return false }

// valueAtPath returns the row value at the specified path, typed or dynamic
func (c *JSON) valueAtPath(path string, row int, ptr bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// scanTypedPathToValue scans the provided typed path into a `reflect.Value`
func (c *JSON) scanTypedPathToValue(path string, row int, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// scanDynamicPathToValue scans the provided typed path into a `reflect.Value`
func (c *JSON) scanDynamicPathToValue(path string, row int, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) rowAsJSON(row int) *chcol.JSON { _ = "STUB: not implemented"; return nil }

func (c *JSON) Name() string { _ = "STUB: not implemented"; return "" }

func (c *JSON) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (c *JSON) Rows() int { _ = "STUB: not implemented"; return 0 }

func (c *JSON) Row(row int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (c *JSON) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (c *JSON) scanRowObject(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (c *JSON) scanRowString(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (c *JSON) Append(v any) (nulls []uint8, err error) { _ = "STUB: not implemented"; return nil, nil }

// Typed-slice fast paths. Each routes through reconcileMode so any
// pending null rows (from prior AppendRow(nil) calls) get flushed
// into the now-latched backing column before we iterate.

// Slice of JSON-text-like values — string serialization.
// Append is columnar; single values (string, *string, etc.)
// should go through AppendRow instead.

// Named []byte-compatible slices (e.g. user aliases of []byte).

// Fallback: appendObject iterates the slice reflectively and calls
// AppendRow per element. AppendRow's own classify+reconcile decides
// the column's mode based on what the elements actually are — do
// not force a mode here, or all-null slices and []string would be
// miscategorized.
//
// If appendObject errors *after* writing some rows (e.g.
// []any{"str", someStruct{}} latches String on element 0 then
// errors on the mode-conflict at element 1), the appendString
// retry would re-iterate from the start and double-write the
// already-appended rows. Snapshot state and skip the retry if
// appendObject mutated anything.

// appendString iterates element-by-element through appendRowString,
// which never calls reconcileMode. For an empty or all-null slice
// the version stays Unset; latch String here so WriteStatePrefix
// emits the correct serialization for the rows we just wrote.

func (c *JSON) appendObject(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unwrap the reflect.Value to the underlying element before dispatch.
// Passing value.Index(i) directly would wrap a reflect.Value struct
// as the `any` argument, which hits structToJSON's struct path and
// produces an empty {} for every row — silent data loss.

func (c *JSON) appendString(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	// Iterate v element-wise and route each row through appendRowString.
	// Delegating to String.Append directly would write "" for nil pointers
	// and invalid sql.NullString values — both invalid JSON on the wire
	// (server code 117) — and would panic on nil *json.RawMessage.
	// appendRowString runs each element through isNullishForJSON, converting
	// null-equivalent inputs to the JSON literal "null". Mirrors
	// appendObject's reflective pattern.
	//
	// nulls mirrors the per-row null mask: 1 if the element was
	// null-equivalent. Nullable(JSON).Append consumes this to populate its
	// null bitmap; the underlying string still carries the "null" literal
	// so the server's JSON parser does not reject the payload.
	return nil, nil
}

// Two separate concerns share the same predicate: the outer check
// records the element in the per-row null mask returned to the
// Nullable wrapper, while appendRowString rewrites null-equivalent
// inputs to the JSON literal "null" so the wire payload still parses.

// classifyJSONValue decides whether v should take the object or string
// serialization path, or whether it carries no preference (null).
//  1. Untyped nil and typed-nil pointers → jsonModeAny (deferred).
//  2. Known object-mode types (chcol.JSON, chcol.JSONSerializer, struct/map/ptr-to-either).
//  3. Known string-mode types (string/[]byte and their pointers, json.RawMessage,
//     sql.NullString, named byte slices, driver.Valuer, fmt.Stringer).
//  4. Everything else → error; never silently store as {}.
func classifyJSONValue(v any) (jsonMode, error) {
	_ = "STUB: not implemented"
	return *new(jsonMode), nil
}

// Guard: a typed nil pointer carries no mode — same as untyped nil.

// can be *any. Should be treated as jsonModeAny and deferred as well.

// Fast path for common exact types.

// Named []byte types (e.g. user aliases of []byte) → string mode.

// Interface-based fallbacks. driver.Valuer before fmt.Stringer — the SQL
// ecosystem convention is that Value() drives serialization.

// reconcileMode latches the column's serialization version on the first
// non-null row and rejects any subsequent row that disagrees. When
// latching transitions the column from Unset to Object/String, any pending
// null rows are flushed into the newly-chosen backing column.
func (c *JSON) reconcileMode(m jsonMode) error { _ = "STUB: not implemented"; return nil }

// flushPendingNulls writes buffered null rows into the currently-latched
// backing column. c.rows was already bumped when each null was queued, so
// we subtract the pending count first and let the per-row append helpers
// re-increment it.
func (c *JSON) flushPendingNulls() error { _ = "STUB: not implemented"; return nil }

func serializationVersionName(v uint64) string { _ = "STUB: not implemented"; return "" }

func (c *JSON) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

// Null rows defer the mode decision. Flushed later, either when a
// non-null row arrives (via reconcileMode) or at WriteStatePrefix for
// all-null batches.

func (c *JSON) appendRowObject(v any) error { _ = "STUB: not implemented"; return nil }

// A nil pointer, or a pointer whose target is a nil interface,
// represents a null row — normalize to nil so the tail handles it.

// A nil value represents a null row (e.g. from Nullable(JSON)) — the
// null mask at the wrapper layer hides this payload, so emitting an
// empty object keeps typed/dynamic sub-column row counts in sync.
// For a non-nil value that reached this point, no branch above knew
// how to convert it; returning an error is essential to avoid silent
// data loss (caller sees a success and reads back `{}`).

// Match typed paths first

// Even if value is nil, we must append a value for this row.
// nil is a valid value for most column types, with most implementations putting a zero value.
// If the column doesn't support appending nil, then the user must provide a zero value.

// Verify all dynamic paths have an equal number of rows by appending nil for all unspecified dynamic paths

// Match or add dynamic paths

// Path doesn't exist, add new dynamic path + column

// New path must back-fill nils for each row

func (c *JSON) appendRowString(v any) error {
	_ = "STUB: not implemented"
	// In String serialization the server parses every row's payload as JSON,
	// even Nullable(JSON) rows that the null mask will later hide. An empty
	// payload ("") fails server-side with "Cannot parse JSON object here"
	// (code 117). For null-equivalent inputs we emit the JSON literal "null"
	// instead — valid JSON that parses, then the null mask (when present)
	// masks it back to NULL on read.
	return nil
}

// isNullishForJSON returns true if v should be encoded as the JSON literal
// "null" when written through the String column. This catches the common
// null-equivalent spellings that the underlying String column would
// otherwise write as "" (invalid JSON on the wire).
func isNullishForJSON(v any) bool { _ = "STUB: not implemented"; return false }

func (c *JSON) encodeObjectHeader(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) encodeObjectData(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *JSON) encodeStringData(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *JSON) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	// If the batch is entirely null rows (all deferred), latch to String
	// now — it's the cheapest on-wire encoding — and flush the pending
	// nulls as JSON "null" payloads so the server accepts the block.
	return nil
}

// If the column is an array, it can be empty but still require a prefix.
// Use string encoding since it's smaller.

func (c *JSON) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *JSON) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (c *JSON) Reset() { _ = "STUB: not implemented"; return }

// Clear the latched mode so the next batch starts unbiased. An all-null
// batch only latches String at WriteStatePrefix; without this reset, a
// subsequent batch of structs/maps would fail reconcileMode with a
// mode-conflict error even though no real value was ever appended in the
// previous batch.

func (c *JSON) decodeObjectHeader(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) decodeObjectData(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) decodeStringData(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) ReadStatePrefix(reader *proto.Reader) error { _ = "STUB: not implemented"; return nil }

// c.serializationVersion is still in deprecated state until c.Decode()

func (c *JSON) Decode(reader *proto.Reader, rows int) error { _ = "STUB: not implemented"; return nil }

// Ensure the rest of the append/scan logic uses object mode

// splitWithDelimiters splits the string while considering backticks and parentheses
func splitWithDelimiters(s string) []string { _ = "STUB: not implemented"; return nil }
