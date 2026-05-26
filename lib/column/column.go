package column

import (
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

// column names which match this must be escaped - see https://clickhouse.com/docs/en/sql-reference/syntax/#identifiers
var escapeColRegex = regexp.MustCompile("^[a-zA-Z_][0-9a-zA-Z_]*$")

// to escape and unescape special chars
var colEscape = strings.NewReplacer("`", "\\`", "\\", "\\\\")
var colUnEscape = strings.NewReplacer("\\`", "`", "\\\\", "\\")

type Type string

func (t Type) params() string { _ = "STUB: not implemented"; return "" }

type Error struct {
	ColumnType string
	Err        error
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

type ColumnConverterError struct {
	Op       string
	Hint     string
	From, To string
}

func (e *ColumnConverterError) Error() string { _ = "STUB: not implemented"; return "" }

type UnsupportedColumnTypeError struct {
	t Type
}

func (e *UnsupportedColumnTypeError) Error() string { _ = "STUB: not implemented"; return "" }

type Interface interface {
	Name() string
	Type() Type
	Rows() int
	Row(i int, ptr bool) any
	ScanRow(dest any, row int) error
	Append(v any) (nulls []uint8, err error)
	AppendRow(v any) error
	Decode(reader *proto.Reader, rows int) error
	Encode(buffer *proto.Buffer)
	ScanType() reflect.Type
	Reset()
}

type CustomSerialization interface {
	ReadStatePrefix(*proto.Reader) error
	WriteStatePrefix(*proto.Buffer) error
}

type ServerContext struct {
	Revision     uint64
	VersionMajor uint64
	VersionMinor uint64
	VersionPatch uint64
	Timezone     *time.Location
}
