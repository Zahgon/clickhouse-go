package clickhouse

import (
	"reflect"
	"sync"
)

type structMap struct {
	cache sync.Map
}

func (m *structMap) Map(op string, columns []string, s any, ptr bool) ([]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func structIdx(t reflect.Type) map[string][]int { _ = "STUB: not implemented"; return nil }
