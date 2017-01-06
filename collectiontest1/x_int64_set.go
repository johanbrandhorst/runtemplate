// Generated from set.tpl with Type=int64
// options: Numeric=true Ordered=true Stringer=true Mutable=false

package collectiontest1


import (
	"bytes"
	"fmt"
)

// Int64Set is the primary type that represents a set
type Int64Set map[int64]struct{}

// NewInt64Set creates and returns a reference to an empty set.
func NewInt64Set(a ...int64) Int64Set {
	set := make(Int64Set)
	for _, i := range a {
		set[i] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set Int64Set) ToSlice() []int64 {
	var s []int64
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set Int64Set) Clone() Int64Set {
	clonedSet := NewInt64Set()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set Int64Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set Int64Set) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set Int64Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set Int64Set) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set Int64Set) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set Int64Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------


func (set Int64Set) doAdd(i int64) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set Int64Set) Contains(i int64) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set Int64Set) ContainsAll(i ...int64) bool {
	for _, v := range i {
		_, found := set[v]
		if !found {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set Int64Set) IsSubset(other Int64Set) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set Int64Set) IsSuperset(other Int64Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set Int64Set) Append(more ...int64) Int64Set {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set Int64Set) Union(other Int64Set) Int64Set {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set Int64Set) Intersect(other Int64Set) Int64Set {
	intersection := NewInt64Set()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set Int64Set) Difference(other Int64Set) Int64Set {
	differencedSet := NewInt64Set()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set Int64Set) SymmetricDifference(other Int64Set) Int64Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}


//-------------------------------------------------------------------------------------------------

// Send returns a channel of type int64 that you can range over.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set Int64Set) Send() <-chan int64 {
	ch := make(chan int64)
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set Int64Set) Forall(fn func(int64) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set Int64Set) Exists(fn func(int64) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over int64Set and executes the passed func against each element.
func (set Int64Set) Foreach(fn func(int64)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new Int64Set whose elements return true for func.
func (set Int64Set) Filter(fn func(int64) bool) Int64Set {
	result := NewInt64Set()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new int64Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set Int64Set) Partition(p func(int64) bool) (Int64Set, Int64Set) {
	matching := NewInt64Set()
	others := NewInt64Set()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of Int64Set that return true for the passed predicate.
func (set Int64Set) CountBy(predicate func(int64) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Int64Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set Int64Set) MinBy(less func(int64, int64) bool) int64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m int64
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of Int64Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set Int64Set) MaxBy(less func(int64, int64) bool) int64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m int64
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is numeric.

// Sum returns the sum of all the elements in the set.
func (set Int64Set) Sum() int64 {
	sum := int64(0)
	for v, _ := range set {
		sum = sum + v
	}
	return sum
}


//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set Int64Set) Equals(other Int64Set) bool {
	if set.Size() != other.Size() {
		return false
	}
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list Int64Set) Min() int64 {
	return list.MinBy(func(a int64, b int64) bool {
		return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list Int64Set) Max() (result int64) {
	return list.MaxBy(func(a int64, b int64) bool {
		return a < b
	})
}



//-------------------------------------------------------------------------------------------------

func (set Int64Set) StringList() []string {
	strings := make([]string, 0)
	for v := range set {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set Int64Set) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set Int64Set) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set Int64Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set Int64Set) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set Int64Set) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

