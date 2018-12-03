// A simple type derived from map[{{.Key}}]{{.Type}}.
// Not thread-safe.
//
// Generated from {{.TemplateFile}} with Key={{.Key}} Type={{.Type}}
// options: Comparable:{{.Comparable}} Stringer:{{.Stringer}} KeyList:{{.KeyList}} ValueList:{{.ValueList}} Mutable:always
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
{{- if .Stringer}}
	"bytes"
{{- end}}
	"fmt"
{{- if .HasImport}}
	{{.Import}}
{{- end}}
)

// {{.UPrefix}}{{.UKey}}{{.UType}}Map is the primary type that represents a map
type {{.UPrefix}}{{.UKey}}{{.UType}}Map map[{{.PKey}}]{{.PType}}

// {{.UPrefix}}{{.UKey}}{{.UType}}Tuple represents a key/value pair.
type {{.UPrefix}}{{.UKey}}{{.UType}}Tuple struct {
	Key {{.PKey}}
	Val {{.PType}}
}

// {{.UPrefix}}{{.UKey}}{{.UType}}Tuples can be used as a builder for unmodifiable maps.
type {{.UPrefix}}{{.UKey}}{{.UType}}Tuples []{{.UPrefix}}{{.UKey}}{{.UType}}Tuple

// Append1 adds one item.
func (ts {{.UPrefix}}{{.UKey}}{{.UType}}Tuples) Append1(k {{.PKey}}, v {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	return append(ts, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k, v})
}

// Append2 adds two items.
func (ts {{.UPrefix}}{{.UKey}}{{.UType}}Tuples) Append2(k1 {{.PKey}}, v1 {{.PType}}, k2 {{.PKey}}, v2 {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	return append(ts, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k1, v1}, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k2, v2})
}

// Append3 adds three items.
func (ts {{.UPrefix}}{{.UKey}}{{.UType}}Tuples) Append3(k1 {{.PKey}}, v1 {{.PType}}, k2 {{.PKey}}, v2 {{.PType}}, k3 {{.PKey}}, v3 {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	return append(ts, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k1, v1}, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k2, v2}, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k3, v3})
}

// {{.UPrefix}}{{.UKey}}{{.UType}}Zip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the New{{.UPrefix}}{{.UKey}}{{.UType}}Map
// constructor function.
func {{.UPrefix}}{{.UKey}}{{.UType}}Zip(keys ...{{.PKey}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	ts := make({{.UPrefix}}{{.UKey}}{{.UType}}Tuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with {{.UPrefix}}{{.UKey}}{{.UType}}Zip.
func (ts {{.UPrefix}}{{.UKey}}{{.UType}}Tuples) Values(values ...{{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func new{{.UPrefix}}{{.UKey}}{{.UType}}Map() {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	return {{.UPrefix}}{{.UKey}}{{.UType}}Map(make(map[{{.PKey}}]{{.PType}}))
}

// New{{.UPrefix}}{{.UKey}}{{.UType}}Map1 creates and returns a reference to a map containing one item.
func New{{.UPrefix}}{{.UKey}}{{.UType}}Map1(k {{.PKey}}, v {{.PType}}) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	mm := new{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	mm[k] = v
	return mm
}

// New{{.UPrefix}}{{.UKey}}{{.UType}}Map creates and returns a reference to a map, optionally containing some items.
func New{{.UPrefix}}{{.UKey}}{{.UType}}Map(kv ...{{.UPrefix}}{{.UKey}}{{.UType}}Tuple) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	mm := new{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Keys() {{if .KeyList}}{{.KeyList}}{{else}}[]{{.PKey}}{{end}} {
	s := make({{if .KeyList}}{{.KeyList}}{{else}}[]{{.PKey}}{{end}}, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Values() {{if .ValueList}}{{.ValueList}}{{else}}[]{{.PType}}{{end}} {
	s := make({{if .ValueList}}{{.ValueList}}{{else}}[]{{.PType}}{{end}}, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) ToSlice() []{{.UPrefix}}{{.UKey}}{{.UType}}Tuple {
	s := make([]{{.UPrefix}}{{.UKey}}{{.UType}}Tuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Get(k {{.PKey}}) ({{.PType}}, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Put(k {{.PKey}}, v {{.PType}}) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) ContainsKey(k {{.PKey}}) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) ContainsAllKeys(kk ...{{.PKey}}) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *{{.UPrefix}}{{.UKey}}{{.UType}}Map) Clear() {
	*mm = make(map[{{.PKey}}]{{.PType}})
}

// Remove a single item from the map.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Remove(k {{.PKey}}) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Pop(k {{.PKey}}) ({{.PType}}, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) DropWhere(fn func({{.PKey}}, {{.PType}}) bool) {{.UPrefix}}{{.UKey}}{{.UType}}Tuples {
	removed := make({{.UPrefix}}{{.UKey}}{{.UType}}Tuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k, v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Foreach(f func({{.PKey}}, {{.PType}})) {
	for k, v := range mm {
		f(k, v)
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Forall(p func({{.PKey}}, {{.PType}}) bool) bool {
	for k, v := range mm {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Exists(p func({{.PKey}}, {{.PType}}) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first {{.Type}} that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Find(p func({{.PKey}}, {{.PType}}) bool) ({{.UPrefix}}{{.UKey}}{{.UType}}Tuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{k, v}, true
		}
	}

	return {{.UPrefix}}{{.UKey}}{{.UType}}Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Filter(p func({{.PKey}}, {{.PType}}) bool) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	result := New{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	for k, v := range mm {
		if p(k, v) {
			result[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Partition(p func({{.PKey}}, {{.PType}}) bool) (matching {{.UPrefix}}{{.UKey}}{{.UType}}Map, others {{.UPrefix}}{{.UKey}}{{.UType}}Map) {
	matching = New{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	others = New{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new {{.UPrefix}}{{.UType}}Map by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Map(f func({{.PKey}}, {{.PType}}) ({{.PKey}}, {{.PType}})) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	result := New{{.UPrefix}}{{.UKey}}{{.UType}}Map()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new {{.UPrefix}}{{.UType}}Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) FlatMap(f func({{.PKey}}, {{.PType}}) []{{.UPrefix}}{{.UKey}}{{.UType}}Tuple) {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	result := New{{.UPrefix}}{{.UKey}}{{.UType}}Map()

	for k1, v1 := range mm {
		ts := f(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}
{{- if .Comparable}}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Equals(other {{.UPrefix}}{{.UKey}}{{.UType}}Map) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm {
		v2, found := other[k]
		if !found || {{.TypeStar}}v1 != {{.TypeStar}}v2 {
			return false
		}
	}
	return true
}
{{- end}}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) Clone() {{.UPrefix}}{{.UKey}}{{.UType}}Map {
	result := New{{.UPrefix}}{{.UKey}}{{.UType}}Map()
	for k, v := range mm {
		result[k] = v
	}
	return result
}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) String() string {
	return mm.MkString3("map[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
{{- if .HasKeySlice}}
// The map entries are sorted by their keys.{{- end}}
func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm {{.UPrefix}}{{.UKey}}{{.UType}}Map) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
{{if .HasKeyList}}
	keys := make({{.KeyList}}, 0, len(mm))
	for k, _ := range mm {
		keys  = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm[k]
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}
{{else}}
	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}
{{end}}
	b.WriteString(after)
	return b
}
{{- end}}
