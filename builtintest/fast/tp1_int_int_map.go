// An encapsulated map[int]int.
// Thread-safe.
//
// Generated from fast/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always

package fast

import (

	"bytes"
	"fmt"
)

// TP1IntIntMap is the primary type that represents a thread-safe map
type TP1IntIntMap struct {
	m map[*int]*int
}

// TP1IntIntTuple represents a key/value pair.
type TP1IntIntTuple struct {
	Key *int
	Val *int
}

// TP1IntIntTuples can be used as a builder for unmodifiable maps.
type TP1IntIntTuples []TP1IntIntTuple

func (ts TP1IntIntTuples) Append1(k *int, v *int) TP1IntIntTuples {
	return append(ts, TP1IntIntTuple{k, v})
}

func (ts TP1IntIntTuples) Append2(k1 *int, v1 *int, k2 *int, v2 *int) TP1IntIntTuples {
	return append(ts, TP1IntIntTuple{k1, v1}, TP1IntIntTuple{k2, v2})
}

func (ts TP1IntIntTuples) Append3(k1 *int, v1 *int, k2 *int, v2 *int, k3 *int, v3 *int) TP1IntIntTuples {
	return append(ts, TP1IntIntTuple{k1, v1}, TP1IntIntTuple{k2, v2}, TP1IntIntTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTP1IntIntMap() TP1IntIntMap {
	return TP1IntIntMap{
		m: make(map[*int]*int),
	}
}

// NewTP1IntIntMap creates and returns a reference to a map containing one item.
func NewTP1IntIntMap1(k *int, v *int) TP1IntIntMap {
	mm := newTP1IntIntMap()
	mm.m[k] = v
	return mm
}

// NewTP1IntIntMap creates and returns a reference to a map, optionally containing some items.
func NewTP1IntIntMap(kv ...TP1IntIntTuple) TP1IntIntMap {
	mm := newTP1IntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TP1IntIntMap) Keys() []*int {

	var s []*int
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm TP1IntIntMap) Values() []*int {

	var s []*int
	for _, v := range mm.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TP1IntIntMap) ToSlice() []TP1IntIntTuple {

	var s []TP1IntIntTuple
	for k, v := range mm.m {
		s = append(s, TP1IntIntTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TP1IntIntMap) Get(k *int) (*int, bool) {

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm TP1IntIntMap) Put(k *int, v *int) bool {

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm TP1IntIntMap) ContainsKey(k *int) bool {

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TP1IntIntMap) ContainsAllKeys(kk ...*int) bool {

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *TP1IntIntMap) Clear() {

	mm.m = make(map[*int]*int)
}

// Remove a single item from the map.
func (mm TP1IntIntMap) Remove(k *int) {

	delete(mm.m, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm TP1IntIntMap) Pop(k *int) (*int, bool) {

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TP1IntIntMap) Size() int {

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TP1IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TP1IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm TP1IntIntMap) DropWhere(fn func(*int, *int) bool) TP1IntIntTuples {

	removed := make(TP1IntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, TP1IntIntTuple{k, v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm TP1IntIntMap) Foreach(fn func(*int, *int)) {

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
func (mm TP1IntIntMap) Forall(fn func(*int, *int) bool) bool {

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
func (mm TP1IntIntMap) Exists(fn func(*int, *int) bool) bool {

	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first int that returns true for some function.
// False is returned if none match.
func (mm TP1IntIntMap) Find(fn func(*int, *int) bool) (TP1IntIntTuple, bool) {

	for k, v := range mm.m {
		if fn(k, v) {
			return TP1IntIntTuple{k, v}, true
		}
	}

	return TP1IntIntTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TP1IntIntMap) Filter(fn func(*int, *int) bool) TP1IntIntMap {
	result := NewTP1IntIntMap()

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
func (mm TP1IntIntMap) Partition(fn func(*int, *int) bool) (matching TP1IntIntMap, others TP1IntIntMap) {
	matching = NewTP1IntIntMap()
	others = NewTP1IntIntMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new TP1IntMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1IntIntMap) Map(fn func(*int, *int) (*int, *int)) TP1IntIntMap {
	result := NewTP1IntIntMap()

	for k1, v1 := range mm.m {
	    k2, v2 := fn(k1, v1)
	    result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new TP1IntMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm TP1IntIntMap) FlatMap(fn func(*int, *int) []TP1IntIntTuple) TP1IntIntMap {
	result := NewTP1IntIntMap()

	for k1, v1 := range mm.m {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result.m[t.Key] = t.Val
	    }
	}

	return result
}


// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm TP1IntIntMap) Equals(other TP1IntIntMap) bool {

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || *v1 != *v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm TP1IntIntMap) Clone() TP1IntIntMap {
	result := NewTP1IntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}


//-------------------------------------------------------------------------------------------------

func (mm TP1IntIntMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TP1IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TP1IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm TP1IntIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm TP1IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
