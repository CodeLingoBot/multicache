package multicache

/**
This file is part of go-multicache, a library for handling caches with multiple
keys and replacement algorithms.

Copyright 2015 Joseph Lewis <joseph@josephlewis.net>
Licensed under the MIT license
**/

/** MultiCacheItem

**/
type MultiCacheItem struct {
	// The tag meaning is defined by the multicache algorithm being used.
	Tag int32
	// The set of keys that reference this item.
	keys []string
	// The actual item stored in this item.
	value interface{}
}

// Resets the cache item to a blank slate
func (m *MultiCacheItem) reset() {
	m.Tag = 0
	m.keys = []string{}
	m.value = nil
}

// Resets the cache item without clearing the Tag
func (m *MultiCacheItem) softReset() {
	m.keys = []string{}
	m.value = nil
}
