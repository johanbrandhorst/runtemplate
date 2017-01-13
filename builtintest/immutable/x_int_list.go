// An encapsulated []int.
// Thread-safe.
//
// Generated from immutable/list.tpl with Type=int
// options: Comparable=true Numeric=true Ordered=true Stringer=true Mutable=disabled

package immutable

import (

	"bytes"
	"fmt"
"math/rand"
)

// XIntList contains a slice of type int. Use it where you would use []int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type XIntList struct {
	m []int
}


//-------------------------------------------------------------------------------------------------

func newXIntList(len, cap int) *XIntList {
	return &XIntList {
		m: make([]int, len, cap),
	}
}

// NewXIntList constructs a new list containing the supplied values, if any.
func NewXIntList(values ...int) *XIntList {
	result := newXIntList(len(values), len(values))
	copy(result.m, values)
	return result
}

// BuildXIntListFromChan constructs a new XIntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildXIntListFromChan(source <-chan int) *XIntList {
	result := newXIntList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current set as a slice
func (list *XIntList) ToSlice() []int {
	s := make([]int, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *XIntList) Get(i int) int {

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *XIntList) Head() int {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *XIntList) Last() int {

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *XIntList) Tail() *XIntList {

	result := newXIntList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *XIntList) Init() *XIntList {

	result := newXIntList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether XIntList is empty.
func (list *XIntList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether XIntList is empty.
func (list *XIntList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list *XIntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *XIntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *XIntList) Size() int {

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This implements one of the methods needed by sort.Interface (along with Less and Swap).
func (list *XIntList) Len() int {

	return len(list.m)
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *XIntList) Contains(v int) bool {
	return list.Exists(func (x int) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *XIntList) ContainsAll(i ...int) bool {

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of XIntList return true for the passed func.
func (list *XIntList) Exists(fn func(int) bool) bool {

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of XIntList return true for the passed func.
func (list *XIntList) Forall(fn func(int) bool) bool {

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over XIntList and executes the passed func against each element.
func (list *XIntList) Foreach(fn func(int)) {

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *XIntList) Send() <-chan int {
	ch := make(chan int)
	go func() {

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of XIntList with all elements in the reverse order.
func (list *XIntList) Reverse() *XIntList {

	numItems := len(list.m)
	result := newXIntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of XIntList, using a version of the Fisher-Yates shuffle.
func (list *XIntList) Shuffle() *XIntList {
	result := NewXIntList(list.m...)
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *XIntList) Append(more ...int) *XIntList {
	newList := NewXIntList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *XIntList) doAppend(more ...int) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of XIntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *XIntList) Take(n int) *XIntList {

	if n > list.Len() {
		return list
	}
	result := newXIntList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of XIntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *XIntList) Drop(n int) *XIntList {
	if n == 0 {
		return list
	}


	result := newXIntList(0, 0)
	l := list.Len()
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of XIntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *XIntList) TakeLast(n int) *XIntList {

	l := list.Len()
	if n > l {
		return list
	}
	result := newXIntList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of XIntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *XIntList) DropLast(n int) *XIntList {
	if n == 0 {
		return list
	}


	l := list.Len()
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new XIntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *XIntList) TakeWhile(p func(int) bool) *XIntList {

	result := newXIntList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new XIntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *XIntList) DropWhile(p func(int) bool) *XIntList {

	result := newXIntList(0, 0)
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

// Filter returns a new XIntList whose elements return true for func.
func (list *XIntList) Filter(fn func(int) bool) *XIntList {

	result := newXIntList(0, list.Len()/2)

	for _, v := range list.m {
		if fn(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new intLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *XIntList) Partition(p func(int) bool) (*XIntList, *XIntList) {

	matching := newXIntList(0, list.Len()/2)
	others := newXIntList(0, list.Len()/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of XIntList that return true for the passed predicate.
func (list *XIntList) CountBy(predicate func(int) bool) (result int) {

	for _, v := range list.m {
		if predicate(v) {
			result++
		}
	}
	return
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Less returns true if the element at index i is less than the element at index j.
// This implements one of the methods needed by sort.Interface (along with Len and Swap).
// Panics if i or j is out of range.
func (list *XIntList) Less(i, j int) bool {
	return list.m[i] < list.m[j]
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *XIntList) Min() int {

	l := list.Len()
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	v := list.m[0]
	m := v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list *XIntList) Max() (result int) {

	l := list.Len()
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	v := list.m[0]
	m := v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if v > m {
			m = v
		}
	}
	return m
}

// DistinctBy returns a new XIntList whose elements are unique, where equality is defined by a passed func.
func (list *XIntList) DistinctBy(equal func(int, int) bool) *XIntList {

	result := newXIntList(0, list.Len())
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
func (list *XIntList) IndexWhere(p func(int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *XIntList) IndexWhere2(p func(int) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *XIntList) LastIndexWhere(p func(int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *XIntList) LastIndexWhere2(p func(int) bool, before int) int {

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
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the list.
func (list *XIntList) Sum() int {

	sum := int(0)
	for _, v := range list.m {
		sum = sum + v
	}
	return sum
}


//-------------------------------------------------------------------------------------------------
// These methods are included when int is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list *XIntList) Equals(other *XIntList) bool {

	if list.Size() != other.Size() {
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

// StringList gets a list of strings that depicts all the elements.
func (list XIntList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *XIntList) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list XIntList) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *XIntList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *XIntList) MkString3(pfx, mid, sfx string) string {
	return list.mkString3Bytes(pfx, mid, sfx).String()
}

func (list XIntList) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""


	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
