// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package bimap provides a bidirectional map implementation.
package bimap

// BiMap is a bidirectional map that allows lookups in both directions.
type BiMap[K comparable, V comparable] struct {
	forward  map[K]V
	backward map[V]K
}

// New creates a new empty BiMap.
func New[K comparable, V comparable]() *BiMap[K, V] {
	return &BiMap[K, V]{
		forward:  make(map[K]V),
		backward: make(map[V]K),
	}
}

// Put inserts a key-value pair into the BiMap.
// Returns true if the pair was inserted, false if either key or value already exists.
func (b *BiMap[K, V]) Put(key K, value V) bool {
	if _, exists := b.forward[key]; exists {
		return false
	}
	if _, exists := b.backward[value]; exists {
		return false
	}
	b.forward[key] = value
	b.backward[value] = key
	return true
}

// Get returns the value associated with the key.
func (b *BiMap[K, V]) Get(key K) (V, bool) {
	value, ok := b.forward[key]
	return value, ok
}

// GetValue is an alias for Get.
func (b *BiMap[K, V]) GetValue(key K) (V, bool) {
	return b.Get(key)
}

// GetKey returns the key associated with the value.
func (b *BiMap[K, V]) GetKey(value V) (K, bool) {
	key, ok := b.backward[value]
	return key, ok
}

// Delete removes the key-value pair by key.
func (b *BiMap[K, V]) Delete(key K) bool {
	if value, ok := b.forward[key]; ok {
		delete(b.forward, key)
		delete(b.backward, value)
		return true
	}
	return false
}

// DeleteValue removes the key-value pair by value.
func (b *BiMap[K, V]) DeleteValue(value V) bool {
	if key, ok := b.backward[value]; ok {
		delete(b.forward, key)
		delete(b.backward, value)
		return true
	}
	return false
}

// Len returns the number of entries in the BiMap.
func (b *BiMap[K, V]) Len() int {
	return len(b.forward)
}

// HasKey returns true if the key exists.
func (b *BiMap[K, V]) HasKey(key K) bool {
	_, ok := b.forward[key]
	return ok
}

// HasValue returns true if the value exists.
func (b *BiMap[K, V]) HasValue(value V) bool {
	_, ok := b.backward[value]
	return ok
}

// Keys returns all keys in the BiMap.
func (b *BiMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(b.forward))
	for k := range b.forward {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values in the BiMap.
func (b *BiMap[K, V]) Values() []V {
	values := make([]V, 0, len(b.backward))
	for v := range b.backward {
		values = append(values, v)
	}
	return values
}
