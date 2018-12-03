// An encapsulated map[{{.Type}}]struct{} used as a set.
// Thread-safe.
//
// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Comparable:always Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} ToList:{{.ToList}}
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

import (
{{- if or .Stringer .GobEncode}}
	"bytes"
{{- end}}
{{- if .GobEncode}}
	"encoding/gob"
{{- end}}
{{- if .Stringer}}
	"encoding/json"
	"fmt"
{{- end}}
	"sync"
{{- if .HasImport}}
	{{.Import}}
{{end}}
)

// {{.UPrefix}}{{.UType}}Set is the primary type that represents a set.
type {{.UPrefix}}{{.UType}}Set struct {
	s *sync.RWMutex
	m map[{{.Type}}]struct{}
}

// New{{.UPrefix}}{{.UType}}Set creates and returns a reference to an empty set.
func New{{.UPrefix}}{{.UType}}Set(values ...{{.Type}}) *{{.UPrefix}}{{.UType}}Set {
	set := &{{.UPrefix}}{{.UType}}Set{
		s: &sync.RWMutex{},
		m: make(map[{{.Type}}]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// Convert{{.UPrefix}}{{.UType}}Set constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func Convert{{.UPrefix}}{{.UType}}Set(values ...interface{}) (*{{.UPrefix}}{{.UType}}Set, bool) {
	set := New{{.UPrefix}}{{.UType}}Set()
{{if and .Numeric (not .TypeIsPtr)}}
	for _, i := range values {
		switch i.(type) {
		case int:
			set.m[{{.PType}}(i.(int))] = struct{}{}
		case int8:
			set.m[{{.PType}}(i.(int8))] = struct{}{}
		case int16:
			set.m[{{.PType}}(i.(int16))] = struct{}{}
		case int32:
			set.m[{{.PType}}(i.(int32))] = struct{}{}
		case int64:
			set.m[{{.PType}}(i.(int64))] = struct{}{}
		case uint:
			set.m[{{.PType}}(i.(uint))] = struct{}{}
		case uint8:
			set.m[{{.PType}}(i.(uint8))] = struct{}{}
		case uint16:
			set.m[{{.PType}}(i.(uint16))] = struct{}{}
		case uint32:
			set.m[{{.PType}}(i.(uint32))] = struct{}{}
		case uint64:
			set.m[{{.PType}}(i.(uint64))] = struct{}{}
		case float32:
			set.m[{{.PType}}(i.(float32))] = struct{}{}
		case float64:
			set.m[{{.PType}}(i.(float64))] = struct{}{}
		}
	}
{{else}}
	for _, i := range values {
		v, ok := i.({{.PType}})
		if ok {
			set.m[v] = struct{}{}
		}
	}
{{end}}
	return set, len(set.m) == len(values)
}

// Build{{.UPrefix}}{{.UType}}SetFromChan constructs a new {{.UPrefix}}{{.UType}}Set from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func Build{{.UPrefix}}{{.UType}}SetFromChan(source <-chan {{.PType}}) *{{.UPrefix}}{{.UType}}Set {
	set := New{{.UPrefix}}{{.UType}}Set()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}
{{- if .ToList}}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *{{.UPrefix}}{{.UType}}Set) ToList() *{{.UPrefix}}{{.UType}}List {
	if set == nil {
		return nil
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return &{{.UPrefix}}{{.UType}}List{
		s: &sync.RWMutex{},
		m: set.slice(),
	}
}
{{- end}}

// ToSet returns the set; this is an identity operation in this case.
func (set *{{.UPrefix}}{{.UType}}Set) ToSet() *{{.UPrefix}}{{.UType}}Set {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *{{.UPrefix}}{{.UType}}Set) slice() []{{.Type}} {
	if set == nil {
		return nil
	}

	s := make([]{{.Type}}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *{{.UPrefix}}{{.UType}}Set) ToSlice() []{{.Type}} {
	set.s.RLock()
	defer set.s.RUnlock()

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *{{.UPrefix}}{{.UType}}Set) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *{{.UPrefix}}{{.UType}}Set) Clone() *{{.UPrefix}}{{.UType}}Set {
	if set == nil {
		return nil
	}

	clonedSet := New{{.UPrefix}}{{.UType}}Set()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *{{.UPrefix}}{{.UType}}Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *{{.UPrefix}}{{.UType}}Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists and queues.
func (set *{{.UPrefix}}{{.UType}}Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *{{.UPrefix}}{{.UType}}Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *{{.UPrefix}}{{.UType}}Set) Size() int {
	if set == nil {
		return 0
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *{{.UPrefix}}{{.UType}}Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *{{.UPrefix}}{{.UType}}Set) Add(more ...{{.Type}}) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set *{{.UPrefix}}{{.UType}}Set) doAdd(i {{.Type}}) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *{{.UPrefix}}{{.UType}}Set) Contains(i {{.Type}}) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *{{.UPrefix}}{{.UType}}Set) ContainsAll(i ...{{.Type}}) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set *{{.UPrefix}}{{.UType}}Set) IsSubset(other *{{.UPrefix}}{{.UType}}Set) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *{{.UPrefix}}{{.UType}}Set) IsSuperset(other *{{.UPrefix}}{{.UType}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *{{.UPrefix}}{{.UType}}Set) Union(other *{{.UPrefix}}{{.UType}}Set) *{{.UPrefix}}{{.UType}}Set {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *{{.UPrefix}}{{.UType}}Set) Intersect(other *{{.UPrefix}}{{.UType}}Set) *{{.UPrefix}}{{.UType}}Set {
	if set == nil || other == nil {
		return nil
	}

	intersection := New{{.UPrefix}}{{.UType}}Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set *{{.UPrefix}}{{.UType}}Set) Difference(other *{{.UPrefix}}{{.UType}}Set) *{{.UPrefix}}{{.UType}}Set {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := New{{.UPrefix}}{{.UType}}Set()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *{{.UPrefix}}{{.UType}}Set) SymmetricDifference(other *{{.UPrefix}}{{.UType}}Set) *{{.UPrefix}}{{.UType}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *{{.UPrefix}}{{.UType}}Set) Clear() {
	if set != nil {
		set.s.Lock()
		defer set.s.Unlock()

		set.m = make(map[{{.Type}}]struct{})
	}
}

// Remove a single item from the set.
func (set *{{.UPrefix}}{{.UType}}Set) Remove(i {{.Type}}) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *{{.UPrefix}}{{.UType}}Set) Send() <-chan {{.Type}} {
	ch := make(chan {{.Type}})
	go func() {
		if set != nil {
			set.s.RLock()
			defer set.s.RUnlock()

			for v := range set.m {
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function p to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set *{{.UPrefix}}{{.UType}}Set) Forall(p func({{.Type}}) bool) bool {
	if set == nil {
		return true
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *{{.UPrefix}}{{.UType}}Set) Exists(p func({{.Type}}) bool) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over {{.Type}}Set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *{{.UPrefix}}{{.UType}}Set) Foreach(f func({{.Type}})) {
	if set == nil {
		return
	}

	set.s.Lock()
	defer set.s.Unlock()

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first {{.Type}} that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *{{.UPrefix}}{{.UType}}Set) Find(p func({{.PType}}) bool) ({{.PType}}, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}
{{- if eq .TypeStar "*"}}

	return nil, false
{{- else}}

	var empty {{.Type}}
	return empty, false
{{- end}}
}

// Filter returns a new {{.UPrefix}}{{.UType}}Set whose elements return true for the predicate p.
//
// The original set is not modified
func (set *{{.UPrefix}}{{.UType}}Set) Filter(p func({{.Type}}) bool) *{{.UPrefix}}{{.UType}}Set {
	if set == nil {
		return nil
	}

	result := New{{.UPrefix}}{{.UType}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new {{.Type}}Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *{{.UPrefix}}{{.UType}}Set) Partition(p func({{.Type}}) bool) (*{{.UPrefix}}{{.UType}}Set, *{{.UPrefix}}{{.UType}}Set) {
	if set == nil {
		return nil, nil
	}

	matching := New{{.UPrefix}}{{.UType}}Set()
	others := New{{.UPrefix}}{{.UType}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new {{.UPrefix}}{{.UType}}Set by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *{{.UPrefix}}{{.UType}}Set) Map(f func({{.PType}}) {{.PType}}) *{{.UPrefix}}{{.UType}}Set {
	if set == nil {
		return nil
	}

	result := New{{.UPrefix}}{{.UType}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		result.m[f(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new {{.UPrefix}}{{.UType}}Set by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *{{.UPrefix}}{{.UType}}Set) FlatMap(f func({{.PType}}) []{{.PType}}) *{{.UPrefix}}{{.UType}}Set {
	if set == nil {
		return nil
	}

	result := New{{.UPrefix}}{{.UType}}Set()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of {{.UPrefix}}{{.UType}}Set that return true for the predicate p.
func (set *{{.UPrefix}}{{.UType}}Set) CountBy(p func({{.Type}}) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}
{{- if .Ordered}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *{{.UPrefix}}{{.UType}}Set) Min() {{.Type}} {
	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set *{{.UPrefix}}{{.UType}}Set) Max() (result {{.Type}}) {
	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}
{{- end}}

// MinBy returns an element of {{.UPrefix}}{{.UType}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *{{.UPrefix}}{{.UType}}Set) MinBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of {{.UPrefix}}{{.UType}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *{{.UPrefix}}{{.UType}}Set) MaxBy(less func({{.Type}}, {{.Type}}) bool) {{.Type}} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m {{.Type}}
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
}
{{- if .Numeric}}

//-------------------------------------------------------------------------------------------------
// These methods are included when {{.Type}} is numeric.

// Sum returns the sum of all the elements in the set.
func (set *{{.UPrefix}}{{.UType}}Set) Sum() {{.Type}} {
	set.s.RLock()
	defer set.s.RUnlock()

	sum := {{.Type}}(0)
	for v := range set.m {
		sum = sum + {{.TypeStar}}v
	}
	return sum
}
{{- end}}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *{{.UPrefix}}{{.UType}}Set) Equals(other *{{.UPrefix}}{{.UType}}Set) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	if set.Size() != other.Size() {
		return false
	}

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}
{{- if .Stringer}}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set *{{.UPrefix}}{{.UType}}Set) StringList() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *{{.UPrefix}}{{.UType}}Set) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *{{.UPrefix}}{{.UType}}Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *{{.UPrefix}}{{.UType}}Set) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *{{.UPrefix}}{{.UType}}Set) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set *{{.UPrefix}}{{.UType}}Set) UnmarshalJSON(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	values := make([]{{.PType}}, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := New{{.UPrefix}}{{.UType}}Set(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *{{.UPrefix}}{{.UType}}Set) MarshalJSON() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *{{.UPrefix}}{{.UType}}Set) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
{{- end}}
{{- if .GobEncode}}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this set type.
// You must register {{.Type}} with the 'gob' package before this method is used.
func (set *{{.UPrefix}}{{.UType}}Set) GobDecode(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register {{.Type}} with the 'gob' package before this method is used.
func (set {{.UPrefix}}{{.UType}}Set) GobEncode() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
{{- end}}
