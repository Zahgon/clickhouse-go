package clickhouse

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrInvalidValueInNamedDateValue = errors.New("invalid value in NamedDateValue for query parameter")
	ErrUnsupportedQueryParameter    = errors.New("unsupported query parameter type")

	hasQueryParamsRe = regexp.MustCompile("{.+:.+}")
)

func bindQueryOrAppendParameters(paramsProtocolSupport bool, options *QueryOptions, query string, timezone *time.Location, args ...any) (string, error) {
	_ = "STUB: not implemented"
	// prefer native query parameters over legacy bind if query parameters provided explicit
	return "", nil
}

// validate if query contains a {<name>:<data type>} syntax, so it's intentional use of query parameters
// parameter values will be loaded from `args ...any` for compatibility

// using the same format logic for NamedValue typed value in function bindNamed

func formatTimeWithScale(t time.Time, scale TimeUnit) string { _ = "STUB: not implemented"; return "" }
