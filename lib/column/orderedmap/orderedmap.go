package orderedmap

import (
	"cmp"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
)

// Map is a simple implementation of [column.IterableOrderedMap] interface.
// It is intended to be used as a serdes wrapper for map[K]V and not as a general purpose container.
type Map[K comparable, V any] []entry[K, V]

type entry[K comparable, V any] struct {
	key   K
	value V
}

type iterator[K comparable, V any] struct {
	om Map[K, V]
	i  int
}

func FromMap[M ~map[K]V, K cmp.Ordered, V any](m M) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func FromMapFunc[M ~map[K]V, K comparable, V any](m M, compare func(K, K) int) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Collect creates a Map from an iter.Seq2[K,V] iterator.
func Collect[K cmp.Ordered, V any](seq func(yield func(K, V) bool)) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// CollectN creates a Map, pre-sized for n entries, from an iter.Seq2[K,V] iterator.
func CollectN[K cmp.Ordered, V any](seq func(yield func(K, V) bool), n int) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// CollectFunc creates a Map from an iter.Seq2[K,V] iterator with a custom compare function.
func CollectFunc[K comparable, V any](seq func(yield func(K, V) bool), compare func(K, K) int) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// CollectNFunc creates a Map, pre-sized for n entries, from an iter.Seq2[K,V] iterator with a custom compare function.
func CollectNFunc[K comparable, V any](seq func(yield func(K, V) bool), n int, compare func(K, K) int) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (om *Map[K, V]) ToMap() map[K]V { _ = "STUB: not implemented"; return nil }

// Put is part of [column.IterableOrderedMap] interface, it expects to be called by the driver itself,
// provides no type safety and expects the keys to be given in order.
// It is recommended to use [FromMap] and [Collect] to initialize [Map].
func (om *Map[K, V]) Put(key any, value any) { _ = "STUB: not implemented"; return }

// All is an iter.Seq[K,V] iterator that yields all key-value pairs in order.
func (om *Map[K, V]) All(yield func(k K, v V) bool) { _ = "STUB: not implemented"; return }

// Keys is an iter.Seq[K] iterator that yields all keys in order.
func (om *Map[K, V]) Keys(yield func(k K) bool) { _ = "STUB: not implemented"; return }

// Values is an iter.Seq[V] iterator that yields all values in key order.
func (om *Map[K, V]) Values(yield func(v V) bool) { _ = "STUB: not implemented"; return }

// Iterator is part of [column.IterableOrderedMap] interface, it expects to be called by the driver itself.
func (om *Map[K, V]) Iterator() column.MapIterator {
	_ = "STUB: not implemented"
	return *new(column.MapIterator)
}

func (i *iterator[K, V]) Next() bool { _ = "STUB: not implemented"; return false }

func (i *iterator[K, V]) Key() any { _ = "STUB: not implemented"; return *new(any) }

func (i *iterator[K, V]) Value() any { _ = "STUB: not implemented"; return *new(any) }
