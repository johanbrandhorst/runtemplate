// An encapsulated []string.
// Thread-safe.
//
// Generated from fast/list.tpl with Type=*string
// options: Comparable:true Numeric:false Ordered:false Stringer:true Mutable:always

package fast

import (

	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

// P1StringList contains a slice of type *string. Use it where you would use []*string.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type P1StringList struct {
	m []*string
}


//-------------------------------------------------------------------------------------------------

func newP1StringList(len, cap int) *P1StringList {
	return &P1StringList {
		m: make([]*string, len, cap),
	}
}

// NewP1StringList constructs a new list containing the supplied values, if any.
func NewP1StringList(values ...*string) *P1StringList {
	result := newP1StringList(len(values), len(values))
	copy(result.m, values)
	return result
}

// ConvertP1StringList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertP1StringList(values ...interface{}) (*P1StringList, bool) {
	result := newP1StringList(0, len(values))

	for _, i := range values {
		v, ok := i.(*string)
		if ok {
			result.m = append(result.m, v)
		}
	}

	return result, len(result.m) == len(values)
}

// BuildP1StringListFromChan constructs a new P1StringList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildP1StringListFromChan(source <-chan *string) *P1StringList {
	result := newP1StringList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current list as a slice.
func (list *P1StringList) ToSlice() []*string {

	s := make([]*string, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *P1StringList) ToInterfaceSlice() []interface{} {

	var s []interface{}
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *P1StringList) Clone() *P1StringList {

	return NewP1StringList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *P1StringList) Get(i int) *string {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *P1StringList) Head() *string {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *P1StringList) Last() *string {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *P1StringList) Tail() *P1StringList {

	result := newP1StringList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *P1StringList) Init() *P1StringList {

	result := newP1StringList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether P1StringList is empty.
func (list *P1StringList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether P1StringList is empty.
func (list *P1StringList) NonEmpty() bool {
	return list.Size() > 0
}

// IsSequence returns true for lists.
func (list *P1StringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *P1StringList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list.
func (list *P1StringList) Size() int {

	return len(list.m)
}

// Swap exchanges two elements.
func (list *P1StringList) Swap(i, j int) {

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *P1StringList) Contains(v string) bool {
	return list.Exists(func (x *string) bool {
		return *x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *P1StringList) ContainsAll(i ...string) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of P1StringList return true for the predicate p.
func (list *P1StringList) Exists(p func(*string) bool) bool {

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of P1StringList return true for the predicate p.
func (list *P1StringList) Forall(p func(*string) bool) bool {

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over P1StringList and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (list *P1StringList) Foreach(fn func(*string)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *P1StringList) Send() <-chan *string {
	ch := make(chan *string)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of P1StringList with all elements in the reverse order.
//
// The original list is not modified.
func (list *P1StringList) Reverse() *P1StringList {
	return list.Clone().doReverse()
}

// DoReverse alters a P1StringList with all elements in the reverse order.
//
// The modified list is returned.
func (list *P1StringList) DoReverse() *P1StringList {
	return list.doReverse()
}

func (list *P1StringList) doReverse() *P1StringList {
	mid := (len(list.m) + 1) / 2
	last := len(list.m) - 1
	for i := 0; i < mid; i++ {
	    r := last - i
	    if i != r {
		    list.m[i], list.m[r] = list.m[r], list.m[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of P1StringList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *P1StringList) Shuffle() *P1StringList {
	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled P1StringList, using a version of the Fisher-Yates shuffle.
//
// The modified list is returned.
func (list *P1StringList) DoShuffle() *P1StringList {
	return list.doShuffle()
}

func (list *P1StringList) doShuffle() *P1StringList {
	numItems := len(list.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
        list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current list. This is a synonym for Append.
func (list *P1StringList) Add(more ...*string) {
	list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *P1StringList) Append(more ...*string) *P1StringList {
	return list.doAppend(more...)
}

func (list *P1StringList) doAppend(more ...*string) *P1StringList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a P1StringList by inserting elements at a given index.
// This is a generalised version of Append.
//
// The modified list is returned.
// Panics if the index is out of range.
func (list *P1StringList) DoInsertAt(index int, more ...*string) *P1StringList {
    return list.doInsertAt(index, more...)
}

func (list *P1StringList) doInsertAt(index int, more ...*string) *P1StringList {
    if len(more) == 0 {
        return list
    }

    if index == len(list.m) {
        // appending is an easy special case
    	return list.doAppend(more...)
    }

	newlist := make([]*string, 0, len(list.m) + len(more))

    if index != 0 {
        newlist = append(newlist, list.m[:index]...)
    }

    newlist = append(newlist, more...)

    newlist = append(newlist, list.m[index:]...)

    list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a P1StringList by deleting n elements from the start of
// the list.
//
// The modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *P1StringList) DoDeleteFirst(n int) *P1StringList {
    return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a P1StringList by deleting n elements from the end of
// the list.
//
// The modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *P1StringList) DoDeleteLast(n int) *P1StringList {
    return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a P1StringList by deleting n elements from a given index.
//
// The modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *P1StringList) DoDeleteAt(index, n int) *P1StringList {
    return list.doDeleteAt(index, n)
}

func (list *P1StringList) doDeleteAt(index, n int) *P1StringList {
    if n == 0 {
        return list
    }

	newlist := make([]*string, 0, len(list.m) - n)

    if index != 0 {
        newlist = append(newlist, list.m[:index]...)
    }

    index += n

    if index != len(list.m) {
        newlist = append(newlist, list.m[index:]...)
    }

    list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoKeepWhere modifies a P1StringList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The modified list is returned.
func (list *P1StringList) DoKeepWhere(p func(*string) bool) *P1StringList {
    return list.doKeepWhere(p)
}

func (list *P1StringList) doKeepWhere(p func(*string) bool) *P1StringList {
	result := make([]*string, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

    list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of P1StringList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *P1StringList) Take(n int) *P1StringList {

	if n > len(list.m) {
		return list
	}
	result := newP1StringList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of P1StringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *P1StringList) Drop(n int) *P1StringList {
	if n == 0 {
		return list
	}


	result := newP1StringList(0, 0)
	l := len(list.m)
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of P1StringList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *P1StringList) TakeLast(n int) *P1StringList {

	l := len(list.m)
	if n > l {
		return list
	}
	result := newP1StringList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of P1StringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *P1StringList) DropLast(n int) *P1StringList {
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

// TakeWhile returns a new P1StringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *P1StringList) TakeWhile(p func(*string) bool) *P1StringList {

	result := newP1StringList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new P1StringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *P1StringList) DropWhile(p func(*string) bool) *P1StringList {

	result := newP1StringList(0, 0)
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

// Find returns the first string that returns true for predicate p.
// False is returned if none match.
func (list P1StringList) Find(p func(*string) bool) (*string, bool) {

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}


	return nil, false

}

// Filter returns a new P1StringList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *P1StringList) Filter(p func(*string) bool) *P1StringList {

	result := newP1StringList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *P1StringList) Partition(p func(*string) bool) (*P1StringList, *P1StringList) {

	matching := newP1StringList(0, len(list.m)/2)
	others := newP1StringList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new P1StringList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *P1StringList) Map(fn func(*string) *string) *P1StringList {
	result := newP1StringList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = fn(v)
	}

	return result
}

// FlatMap returns a new P1StringList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *P1StringList) FlatMap(fn func(*string) []*string) *P1StringList {
	result := newP1StringList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of P1StringList that return true for the passed predicate.
func (list *P1StringList) CountBy(predicate func(*string) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of P1StringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *P1StringList) MinBy(less func(*string, *string) bool) *string {

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

// MaxBy returns an element of P1StringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *P1StringList) MaxBy(less func(*string, *string) bool) *string {

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

// DistinctBy returns a new P1StringList whose elements are unique, where equality is defined by a passed func.
func (list *P1StringList) DistinctBy(equal func(*string, *string) bool) *P1StringList {

	result := newP1StringList(0, len(list.m))
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
func (list *P1StringList) IndexWhere(p func(*string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *P1StringList) IndexWhere2(p func(*string) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *P1StringList) LastIndexWhere(p func(*string) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *P1StringList) LastIndexWhere2(p func(*string) bool, before int) int {

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
// These methods are included when string is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *P1StringList) Equals(other *P1StringList) bool {

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

type sortableP1StringList struct {
	less func(i, j *string) bool
	m []*string
}

func (sl sortableP1StringList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableP1StringList) Len() int {
	return len(sl.m)
}

func (sl sortableP1StringList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *P1StringList) SortBy(less func(i, j *string) bool) *P1StringList {

	sort.Sort(sortableP1StringList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *P1StringList) StableSortBy(less func(i, j *string) bool) *P1StringList {

	sort.Stable(sortableP1StringList{less, list.m})
	return list
}


//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list P1StringList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *P1StringList) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list P1StringList) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *P1StringList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *P1StringList) MkString3(before, between, after string) string {
	return list.mkString3Bytes(before, between, after).String()
}

func (list P1StringList) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""


	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}
