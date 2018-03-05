// A simple type derived from map[string]Apple.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=Apple
// options: Comparable:<no value> Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package simple


import (

	"bytes"
	"fmt"
)

// TX1StringAppleMap is the primary type that represents a map
type TX1StringAppleMap map[string]Apple

// TX1StringAppleTuple represents a key/value pair.
type TX1StringAppleTuple struct {
	Key string
	Val Apple
}

// TX1StringAppleTuples can be used as a builder for unmodifiable maps.
type TX1StringAppleTuples []TX1StringAppleTuple

func (ts TX1StringAppleTuples) Append1(k string, v Apple) TX1StringAppleTuples {
	return append(ts, TX1StringAppleTuple{k, v})
}

func (ts TX1StringAppleTuples) Append2(k1 string, v1 Apple, k2 string, v2 Apple) TX1StringAppleTuples {
	return append(ts, TX1StringAppleTuple{k1, v1}, TX1StringAppleTuple{k2, v2})
}

func (ts TX1StringAppleTuples) Append3(k1 string, v1 Apple, k2 string, v2 Apple, k3 string, v3 Apple) TX1StringAppleTuples {
	return append(ts, TX1StringAppleTuple{k1, v1}, TX1StringAppleTuple{k2, v2}, TX1StringAppleTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1StringAppleMap() TX1StringAppleMap {
	return TX1StringAppleMap(make(map[string]Apple))
}

// NewTX1StringAppleMap creates and returns a reference to a map containing one item.
func NewTX1StringAppleMap1(k string, v Apple) TX1StringAppleMap {
	mm := newTX1StringAppleMap()
	mm[k] = v
	return mm
}

// NewTX1StringAppleMap creates and returns a reference to a map, optionally containing some items.
func NewTX1StringAppleMap(kv ...TX1StringAppleTuple) TX1StringAppleMap {
	mm := newTX1StringAppleMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1StringAppleMap) Keys() []string {
	var s []string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1StringAppleMap) Values() []Apple {
	var s []Apple
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1StringAppleMap) ToSlice() []TX1StringAppleTuple {
	var s []TX1StringAppleTuple
	for k, v := range mm {
		s = append(s, TX1StringAppleTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1StringAppleMap) Get(k string) (Apple, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TX1StringAppleMap) Put(k string, v Apple) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1StringAppleMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1StringAppleMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TX1StringAppleMap) Clear() {
	*mm = make(map[string]Apple)
}

// Remove allows the removal of a single item from the map.
func (mm TX1StringAppleMap) Remove(k string) {
	delete(mm, k)
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1StringAppleMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm TX1StringAppleMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1StringAppleMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1StringAppleMap) Forall(fn func(string, Apple) bool) bool {
	for k, v := range mm {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm TX1StringAppleMap) Exists(fn func(string, Apple) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified
func (mm TX1StringAppleMap) Filter(fn func(string, Apple) bool) TX1StringAppleMap {
	result := NewTX1StringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified
func (mm TX1StringAppleMap) Partition(fn func(string, Apple) bool) (matching TX1StringAppleMap, others TX1StringAppleMap) {
	matching = NewTX1StringAppleMap()
	others = NewTX1StringAppleMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new TX1AppleMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1StringAppleMap) Map(fn func(string, Apple) (string, Apple)) TX1StringAppleMap {
	result := NewTX1StringAppleMap()

	for k1, v1 := range mm {
	    k2, v2 := fn(k1, v1)
	    result[k2] = v2
	}

	return result
}

// FlatMap returns a new TX1AppleMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1StringAppleMap) FlatMap(fn func(string, Apple) []TX1StringAppleTuple) TX1StringAppleMap {
	result := NewTX1StringAppleMap()

	for k1, v1 := range mm {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TX1StringAppleMap) Clone() TX1StringAppleMap {
	result := NewTX1StringAppleMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm TX1StringAppleMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1StringAppleMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1StringAppleMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm TX1StringAppleMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TX1StringAppleMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

