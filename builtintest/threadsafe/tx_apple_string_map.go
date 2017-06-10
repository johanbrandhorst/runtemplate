// An encapsulated map[Apple]string.
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=Apple Type=string
// options: Comparable:<no value> Stringer:<no value> Mutable:always

package threadsafe

import (

	"sync"
)

// TXAppleStringMap is the primary type that represents a thread-safe map
type TXAppleStringMap struct {
	s *sync.RWMutex
	m map[Apple]string
}

// TXAppleStringTuple represents a key/value pair.
type TXAppleStringTuple struct {
	Key Apple
	Val string
}

// TXAppleStringTuples can be used as a builder for unmodifiable maps.
type TXAppleStringTuples []TXAppleStringTuple

func (ts TXAppleStringTuples) Append1(k Apple, v string) TXAppleStringTuples {
	return append(ts, TXAppleStringTuple{k, v})
}

func (ts TXAppleStringTuples) Append2(k1 Apple, v1 string, k2 Apple, v2 string) TXAppleStringTuples {
	return append(ts, TXAppleStringTuple{k1, v1}, TXAppleStringTuple{k2, v2})
}

func (ts TXAppleStringTuples) Append3(k1 Apple, v1 string, k2 Apple, v2 string, k3 Apple, v3 string) TXAppleStringTuples {
	return append(ts, TXAppleStringTuple{k1, v1}, TXAppleStringTuple{k2, v2}, TXAppleStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTXAppleStringMap() TXAppleStringMap {
	return TXAppleStringMap{
		s: &sync.RWMutex{},
		m: make(map[Apple]string),
	}
}

// NewTXAppleStringMap creates and returns a reference to a map containing one item.
func NewTXAppleStringMap1(k Apple, v string) TXAppleStringMap {
	mm := newTXAppleStringMap()
	mm.m[k] = v
	return mm
}

// NewTXAppleStringMap creates and returns a reference to a map, optionally containing some items.
func NewTXAppleStringMap(kv ...TXAppleStringTuple) TXAppleStringMap {
	mm := newTXAppleStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TXAppleStringMap) Keys() []Apple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TXAppleStringMap) Values() []string {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []string
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TXAppleStringMap) ToSlice() []TXAppleStringTuple {
	mm.s.RLock()
	defer mm.s.RUnlock()

	var s []TXAppleStringTuple
	for k, v := range mm.m {
		s = append(s, TXAppleStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TXAppleStringMap) Get(k Apple) (string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TXAppleStringMap) Put(k Apple, v string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TXAppleStringMap) ContainsKey(k Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TXAppleStringMap) ContainsAllKeys(kk ...Apple) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TXAppleStringMap) Clear() {
	mm.s.Lock()
	defer mm.s.Unlock()

	mm.m = make(map[Apple]string)
}

// Remove allows the removal of a single item from the map.
func (mm TXAppleStringMap) Remove(k Apple) {
	mm.s.Lock()
	defer mm.s.Unlock()

	delete(mm.m, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TXAppleStringMap) Size() int {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TXAppleStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TXAppleStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TXAppleStringMap) DropWhere(fn func(Apple, string) bool) TXAppleStringTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(TXAppleStringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
		    removed = append(removed, TXAppleStringTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TXAppleStringMap) Foreach(fn func(Apple, string)) {
	mm.s.Lock()
	defer mm.s.Unlock()

	for k, v := range mm.m {
		fn(k, v)
	}
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TXAppleStringMap) Forall(fn func(Apple, string) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TXAppleStringMap) Exists(fn func(Apple, string) bool) bool {
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TXAppleStringMap) Filter(fn func(Apple, string) bool) TXAppleStringMap {
	result := NewTXAppleStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			result.m[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
func (mm TXAppleStringMap) Partition(fn func(Apple, string) bool) (matching TXAppleStringMap, others TXAppleStringMap) {
	matching = NewTXAppleStringMap()
	others = NewTXAppleStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TXAppleStringMap) Clone() TXAppleStringMap {
	result := NewTXAppleStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}



//-------------------------------------------------------------------------------------------------
// Lock Accessors

// Lock locks the map for writing. You can use this if the values are themselves datastructures
// that need to be restricted within the same lock.
//
// Do not forget to unlock! Also, do not set this write lock then attempt any read-locked operations (e.g. Get).
func (mm TXAppleStringMap) Lock() {
	mm.s.Lock()
}

// Unlock unlocks the map's write-lock.
func (mm TXAppleStringMap) Unlock() {
	mm.s.Unlock()
}

// RLock locks the map for reading. You can use this if the values are themselves datastructures
// that need to be restricted within the same lock.
//
// Do not forget to unlock! Also, do not set this read lock then attempt any write-locked operations (e.g. Put).
func (mm TXAppleStringMap) RLock() {
	mm.s.RLock()
}

// RUnlock unlocks the map's read-lock.
func (mm TXAppleStringMap) RUnlock() {
	mm.s.RLock()
}
