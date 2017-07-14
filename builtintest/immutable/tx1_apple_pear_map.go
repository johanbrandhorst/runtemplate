// An encapsulated map[Apple]Pear.
// Thread-safe.
//
// Generated from immutable/map.tpl with Key=Apple Type=Pear
// options: Comparable:<no value> Stringer:true Mutable:disabled

package immutable


import (

	"bytes"
	"fmt"

)

// TX1ApplePearMap is the primary type that represents a thread-safe map
type TX1ApplePearMap struct {
	m map[Apple]Pear
}

// TX1ApplePearTuple represents a key/value pair.
type TX1ApplePearTuple struct {
	Key Apple
	Val Pear
}

// TX1ApplePearTuples can be used as a builder for unmodifiable maps.
type TX1ApplePearTuples []TX1ApplePearTuple

func (ts TX1ApplePearTuples) Append1(k Apple, v Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k, v})
}

func (ts TX1ApplePearTuples) Append2(k1 Apple, v1 Pear, k2 Apple, v2 Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k1, v1}, TX1ApplePearTuple{k2, v2})
}

func (ts TX1ApplePearTuples) Append3(k1 Apple, v1 Pear, k2 Apple, v2 Pear, k3 Apple, v3 Pear) TX1ApplePearTuples {
	return append(ts, TX1ApplePearTuple{k1, v1}, TX1ApplePearTuple{k2, v2}, TX1ApplePearTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newTX1ApplePearMap() TX1ApplePearMap {
	return TX1ApplePearMap{
		m: make(map[Apple]Pear),
	}
}

// NewTX1ApplePearMap creates and returns a reference to a map containing one item.
func NewTX1ApplePearMap1(k Apple, v Pear) TX1ApplePearMap {
	mm := newTX1ApplePearMap()
	mm.m[k] = v
	return mm
}

// NewTX1ApplePearMap creates and returns a reference to a map, optionally containing some items.
func NewTX1ApplePearMap(kv ...TX1ApplePearTuple) TX1ApplePearMap {
	mm := newTX1ApplePearMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm TX1ApplePearMap) Keys() []Apple {
	var s []Apple
	for k, _ := range mm.m {
		s = append(s, k)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm TX1ApplePearMap) ToSlice() []TX1ApplePearTuple {
	var s []TX1ApplePearTuple
	for k, v := range mm.m {
		s = append(s, TX1ApplePearTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm TX1ApplePearMap) Get(k Apple) (Pear, bool) {
	v, found := mm.m[k]
	return v, found
}

// ContainsKey determines if a given item is already in the map.
func (mm TX1ApplePearMap) ContainsKey(k Apple) bool {
	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm TX1ApplePearMap) ContainsAllKeys(kk ...Apple) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm TX1ApplePearMap) Size() int {
	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm TX1ApplePearMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm TX1ApplePearMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm TX1ApplePearMap) Forall(fn func(Apple, Pear) bool) bool {
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
func (mm TX1ApplePearMap) Exists(fn func(Apple, Pear) bool) bool {
	for k, v := range mm.m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm TX1ApplePearMap) Filter(fn func(Apple, Pear) bool) TX1ApplePearMap {
	result := NewTX1ApplePearMap()

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
func (mm TX1ApplePearMap) Partition(fn func(Apple, Pear) bool) (matching TX1ApplePearMap, others TX1ApplePearMap) {
	matching = NewTX1ApplePearMap()
	others = NewTX1ApplePearMap()

	for k, v := range mm.m {
		if fn(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Clone returns the same map, which is immutable.
func (mm TX1ApplePearMap) Clone() TX1ApplePearMap {
	return mm
}


//-------------------------------------------------------------------------------------------------

func (mm TX1ApplePearMap) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm TX1ApplePearMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm TX1ApplePearMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.

func (mm TX1ApplePearMap) MkString3(pfx, mid, sfx string) string {
	return mm.mkString3Bytes(pfx, mid, sfx).String()
}

func (mm TX1ApplePearMap) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""


	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = mid
	}

	b.WriteString(sfx)
	return b
}

