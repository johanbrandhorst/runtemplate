// An encapsulated map[testtypes.Email]struct{} used as a set.
// Thread-safe.
//
// Generated from threadsafe/set.tpl with Type=testtypes.Email
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:<no value>

package threadsafe

import (

	"sync"
    "github.com/rickb777/runtemplate/builtintest/testtypes"

)

// X2EmailSet is the primary type that represents a set
type X2EmailSet struct {
	s *sync.RWMutex
	m map[testtypes.Email]struct{}
}

// NewX2EmailSet creates and returns a reference to an empty set.
func NewX2EmailSet(a ...testtypes.Email) X2EmailSet {
	set := X2EmailSet{
		s: &sync.RWMutex{},
		m: make(map[testtypes.Email]struct{}),
	}
	for _, i := range a {
		set.m[i] = struct{}{}
	}
	return set
}

// BuildX2EmailSetFromChan constructs a new X2EmailSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX2EmailSetFromChan(source <-chan testtypes.Email) X2EmailSet {
	set := NewX2EmailSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice
func (set X2EmailSet) ToSlice() []testtypes.Email {
	set.s.RLock()
	defer set.s.RUnlock()

	var s []testtypes.Email
	for v, _ := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set X2EmailSet) Clone() X2EmailSet {
	clonedSet := NewX2EmailSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set X2EmailSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set X2EmailSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set X2EmailSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set X2EmailSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set X2EmailSet) Size() int {
	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set X2EmailSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set X2EmailSet) Add(more ...testtypes.Email) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set X2EmailSet) doAdd(i testtypes.Email) {
	set.m[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set X2EmailSet) Contains(i testtypes.Email) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines if the given items are all in the set.
func (set X2EmailSet) ContainsAll(i ...testtypes.Email) bool {
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

// IsSubset determines if every item in the other set is in this set.
func (set X2EmailSet) IsSubset(other X2EmailSet) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set X2EmailSet) IsSuperset(other X2EmailSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set X2EmailSet) Union(other X2EmailSet) X2EmailSet {
	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v, _ := range other.m {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set X2EmailSet) Intersect(other X2EmailSet) X2EmailSet {
	intersection := NewX2EmailSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v, _ := range set.m {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v, _ := range other.m {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set X2EmailSet) Difference(other X2EmailSet) X2EmailSet {
	differencedSet := NewX2EmailSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v, _ := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set X2EmailSet) SymmetricDifference(other X2EmailSet) X2EmailSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *X2EmailSet) Clear() {
	set.s.Lock()
	defer set.s.Unlock()

	set.m = make(map[testtypes.Email]struct{})
}

// Remove removes a single item from the set.
func (set X2EmailSet) Remove(i testtypes.Email) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set X2EmailSet) Send() <-chan testtypes.Email {
	ch := make(chan testtypes.Email)
	go func() {
		set.s.RLock()
		defer set.s.RUnlock()

		for v, _ := range set.m {
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
func (set X2EmailSet) Forall(fn func(testtypes.Email) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set X2EmailSet) Exists(fn func(testtypes.Email) bool) bool {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over testtypes.EmailSet and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (set X2EmailSet) Foreach(fn func(testtypes.Email)) {
	set.s.Lock()
	defer set.s.Unlock()

	for v, _ := range set.m {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new X2EmailSet whose elements return true for func.
func (set X2EmailSet) Filter(fn func(testtypes.Email) bool) X2EmailSet {
	result := NewX2EmailSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if fn(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new testtypes.EmailLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set X2EmailSet) Partition(p func(testtypes.Email) bool) (X2EmailSet, X2EmailSet) {
	matching := NewX2EmailSet()
	others := NewX2EmailSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of X2EmailSet that return true for the passed predicate.
func (set X2EmailSet) CountBy(predicate func(testtypes.Email) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v, _ := range set.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of X2EmailSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set X2EmailSet) MinBy(less func(testtypes.Email, testtypes.Email) bool) testtypes.Email {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m testtypes.Email
	first := true
	for v, _ := range set.m {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of X2EmailSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set X2EmailSet) MaxBy(less func(testtypes.Email, testtypes.Email) bool) testtypes.Email {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

	var m testtypes.Email
	first := true
	for v, _ := range set.m {
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

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set X2EmailSet) Equals(other X2EmailSet) bool {
	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	if set.Size() != other.Size() {
		return false
	}
	for v, _ := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

