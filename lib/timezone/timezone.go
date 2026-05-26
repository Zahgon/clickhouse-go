package timezone

import (
	"sync"
	"time"
)

var cache = struct {
	mutex sync.Mutex
	items map[string]*time.Location
}{
	items: make(map[string]*time.Location),
}

func Load(name string) (*time.Location, error) { _ = "STUB: not implemented"; return nil, nil }
