// An encapsulated map[string]string.
// Thread-safe.
//
// Generated from immutable/map.tpl with Key=string Type=string
// options: Comparable:true Stringer:true KeyList:<no value> Mutable:disabled

package immutable


import (

	"bytes"
	"fmt"
)

// TX1StringStringMap is the primary type that represents a thread-safe map
type TX1StringStringMap struct {
	m map[string]string
}

// TX1StringStringTuple represents a key/value pair.
type TX1StringStringTuple struct {
	Key string
	Val string
}

// TX1StringStringTuples can be used as a builder for unmodifiable maps.
type TX1StringStringTuples []TX1StringStringTuple

func (ts TX1StringStringTuples) Append1(k string, v string) TX1StringStringTuples {
	return append(ts, TX1StringStringTuple{k, v})
}

func (ts TX1StringStringTuples) Append2(k1 string, v1 string, k2 string, v2 string) TX1StringStringTuples {
	return append(ts, TX1StringStringTuple{k1, v1}, TX1StringStringTuple{k2, v2})
}

func (ts TX1StringStringTuples) Append3(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string) TX1StringStringTuples {
	return append(ts, TX1StringStringTuple{k1, v1}, TX1StringStringTuple{k2, v2}, TX1StringStringTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1StringStringMap() TX1StringStringMap {
	return TX1StringStringMap{
		m: make(map[string]string),
	}
}

// NewTX1StringStringMap creates and returns a reference to a map containing one item.
func NewTX1StringStringMap1(k string, v string) TX1StringStringMap {
	mm := newTX1StringStringMap()
	mm.m[k] = v
	return mm
}

// NewTX1StringStringMap creates and returns a reference to a map, optionally containing some items.
func NewTX1StringStringMap(kv ...TX1StringStringTuple) TX1StringStringMap {
	mm := newTX1StringStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1StringStringMap) Keys() []string {
	var s []string
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TX1StringStringMap) Values() []string {

	var s []string
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1StringStringMap) ToSlice() []TX1StringStringTuple {
	var s []TX1StringStringTuple
	for k, v := range mm.m {
		s = append(s, TX1StringStringTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1StringStringMap) Get(k string) (string, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1StringStringMap) ContainsKey(k string) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1StringStringMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1StringStringMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX1StringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1StringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1StringStringMap) Forall(fn func(string, string) bool) bool {
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
func (mm TX1StringStringMap) Exists(fn func(string, string) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (mm TX1StringStringMap) Find(fn func(string, string) bool) (TX1StringStringTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return TX1StringStringTuple{k, v}, true
		}
	}

	return TX1StringStringTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1StringStringMap) Filter(fn func(string, string) bool) TX1StringStringMap {
	result := NewTX1StringStringMap()

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
func (mm TX1StringStringMap) Partition(fn func(string, string) bool) (matching TX1StringStringMap, others TX1StringStringMap) {
	matching = NewTX1StringStringMap()
	others = NewTX1StringStringMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Transform returns a new TX1StringMap by transforming every element with a function fn.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TX1StringStringMap) Transform(fn func(string, string) (string, string)) TX1StringStringMap {
	result := NewTX1StringStringMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TX1StringStringMap) Equals(other TX1StringStringMap) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns the same map, which is immutable.
func (mm TX1StringStringMap) Clone() TX1StringStringMap {
	return mm
}


//-------------------------------------------------------------------------------------------------

func (mm TX1StringStringMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1StringStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1StringStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm TX1StringStringMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TX1StringStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

