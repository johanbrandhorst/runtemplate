// An encapsulated []Apple.
// Thread-safe.
//
// Generated from immutable/list.tpl with Type=Apple
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:false Mutable:disabled

package examples

import (
	"math/rand"
	"sort"
)

// ImmutableAppleList contains a slice of type Apple. Use it where you would use []Apple.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type ImmutableAppleList struct {
	m []Apple
}

//-------------------------------------------------------------------------------------------------

func newImmutableAppleList(len, cap int) *ImmutableAppleList {
	return &ImmutableAppleList{
		m: make([]Apple, len, cap),
	}
}

// NewImmutableAppleList constructs a new list containing the supplied values, if any.
func NewImmutableAppleList(values ...Apple) *ImmutableAppleList {
	result := newImmutableAppleList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertImmutableAppleList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertImmutableAppleList(values ...interface{}) (*ImmutableAppleList, bool) {
	result := newImmutableAppleList(0, len(values))

	for _, i := range values {
		v, ok := i.(Apple)
		if ok {
			result.m = append(result.m, v)
		}
	}

	return result, len(result.m) == len(values)
}

// BuildImmutableAppleListFromChan constructs a new ImmutableAppleList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildImmutableAppleListFromChan(source <-chan Apple) *ImmutableAppleList {
	result := newImmutableAppleList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *ImmutableAppleList) ToSlice() []Apple {

	s := make([]Apple, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *ImmutableAppleList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *ImmutableAppleList) Clone() *ImmutableAppleList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *ImmutableAppleList) Get(i int) Apple {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *ImmutableAppleList) Head() Apple {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *ImmutableAppleList) Last() Apple {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *ImmutableAppleList) Tail() *ImmutableAppleList {

	result := newImmutableAppleList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *ImmutableAppleList) Init() *ImmutableAppleList {

	result := newImmutableAppleList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether ImmutableAppleList is empty.
func (list *ImmutableAppleList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether ImmutableAppleList is empty.
func (list *ImmutableAppleList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *ImmutableAppleList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *ImmutableAppleList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *ImmutableAppleList) Size() int {

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *ImmutableAppleList) Len() int {

	return len(list.m)
}

//-------------------------------------------------------------------------------------------------

// Contains determines if a given item is already in the list.
func (list *ImmutableAppleList) Contains(v Apple) bool {
	return list.Exists(func(x Apple) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *ImmutableAppleList) ContainsAll(i ...Apple) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of ImmutableAppleList return true for the passed func.
func (list *ImmutableAppleList) Exists(fn func(Apple) bool) bool {

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of ImmutableAppleList return true for the passed func.
func (list *ImmutableAppleList) Forall(fn func(Apple) bool) bool {

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over ImmutableAppleList and executes the passed func against each element.
func (list *ImmutableAppleList) Foreach(fn func(Apple)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *ImmutableAppleList) Send() <-chan Apple {
	ch := make(chan Apple)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of ImmutableAppleList with all elements in the reverse order.
func (list *ImmutableAppleList) Reverse() *ImmutableAppleList {

	numItems := len(list.m)
	result := newImmutableAppleList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of ImmutableAppleList, using a version of the Fisher-Yates shuffle.
func (list *ImmutableAppleList) Shuffle() *ImmutableAppleList {
	result := NewImmutableAppleList(list.m...)
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *ImmutableAppleList) Append(more ...Apple) *ImmutableAppleList {
	newList := NewImmutableAppleList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *ImmutableAppleList) doAppend(more ...Apple) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of ImmutableAppleList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *ImmutableAppleList) Take(n int) *ImmutableAppleList {

	if n > len(list.m) {
		return list
	}
	result := newImmutableAppleList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of ImmutableAppleList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *ImmutableAppleList) Drop(n int) *ImmutableAppleList {
	if n == 0 {
		return list
	}

	result := newImmutableAppleList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of ImmutableAppleList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *ImmutableAppleList) TakeLast(n int) *ImmutableAppleList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := newImmutableAppleList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of ImmutableAppleList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *ImmutableAppleList) DropLast(n int) *ImmutableAppleList {
	if n == 0 {
		return list
	}

	l := len(list.m)
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new ImmutableAppleList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *ImmutableAppleList) TakeWhile(p func(Apple) bool) *ImmutableAppleList {

	result := newImmutableAppleList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new ImmutableAppleList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *ImmutableAppleList) DropWhile(p func(Apple) bool) *ImmutableAppleList {

	result := newImmutableAppleList(0, 0)
	adding := false

	for _, v := range list.m {
		if !p(v) || adding {
			adding = true
			result.m = append(result.m, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Apple that returns true for some function.
// False is returned if none match.
func (list ImmutableAppleList) Find(fn func(Apple) bool) (Apple, bool) {

	for _, v := range list.m {
		if fn(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false

}

// Filter returns a new ImmutableAppleList whose elements return true for func.
func (list *ImmutableAppleList) Filter(fn func(Apple) bool) *ImmutableAppleList {

	result := newImmutableAppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new AppleLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *ImmutableAppleList) Partition(p func(Apple) bool) (*ImmutableAppleList, *ImmutableAppleList) {

	matching := newImmutableAppleList(0, len(list.m)/2)
	others := newImmutableAppleList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new ImmutableAppleList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableAppleList) Map(fn func(Apple) Apple) *ImmutableAppleList {
	result := newImmutableAppleList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new ImmutableAppleList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *ImmutableAppleList) FlatMap(fn func(Apple) []Apple) *ImmutableAppleList {
	result := newImmutableAppleList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of ImmutableAppleList that return true for the passed predicate.
func (list *ImmutableAppleList) CountBy(predicate func(Apple) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of ImmutableAppleList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *ImmutableAppleList) MinBy(less func(Apple, Apple) bool) Apple {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[i], list.m[m]) {
			m = i
		}
	}
	return list.m[m]
}

// MaxBy returns an element of ImmutableAppleList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *ImmutableAppleList) MaxBy(less func(Apple, Apple) bool) Apple {

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[m], list.m[i]) {
			m = i
		}
	}

	return list.m[m]
}

// DistinctBy returns a new ImmutableAppleList whose elements are unique, where equality is defined by a passed func.
func (list *ImmutableAppleList) DistinctBy(equal func(Apple, Apple) bool) *ImmutableAppleList {

	result := newImmutableAppleList(0, len(list.m))
Outer:
	for _, v := range list.m {
		for _, r := range result.m {
			if equal(v, r) {
				continue Outer
			}
		}
		result.m = append(result.m, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list *ImmutableAppleList) IndexWhere(p func(Apple) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *ImmutableAppleList) IndexWhere2(p func(Apple) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *ImmutableAppleList) LastIndexWhere(p func(Apple) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *ImmutableAppleList) LastIndexWhere2(p func(Apple) bool, before int) int {

	if before < 0 {
		before = len(list.m)
	}
	for i := len(list.m) - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

//-------------------------------------------------------------------------------------------------
// These methods are included when Apple is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *ImmutableAppleList) Equals(other *ImmutableAppleList) bool {

	if len(list.m) != len(other.m) {
		return false
	}

	for i, v := range list.m {
		if v != other.m[i] {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

type sortableImmutableAppleList struct {
	less func(i, j Apple) bool
	m    []Apple
}

func (sl sortableImmutableAppleList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableImmutableAppleList) Len() int {
	return len(sl.m)
}

func (sl sortableImmutableAppleList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *ImmutableAppleList) SortBy(less func(i, j Apple) bool) *ImmutableAppleList {

	result := NewImmutableAppleList(list.m...)
	sort.Sort(sortableImmutableAppleList{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *ImmutableAppleList) StableSortBy(less func(i, j Apple) bool) *ImmutableAppleList {

	result := NewImmutableAppleList(list.m...)
	sort.Stable(sortableImmutableAppleList{less, result.m})
	return result
}