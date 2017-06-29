// An encapsulated []int.
// Thread-safe.
//
// Generated from threadsafe/list.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always

package threadsafe

import (

	"bytes"
	"fmt"
	"sync"
	"math/rand"
)

// X1IntList contains a slice of type int. Use it where you would use []int.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type X1IntList struct {
	s *sync.RWMutex
	m []int
}


//-------------------------------------------------------------------------------------------------

func newX1IntList(len, cap int) *X1IntList {
	return &X1IntList {
		s: &sync.RWMutex{},
		m: make([]int, len, cap),
	}
}

// NewX1IntList constructs a new list containing the supplied values, if any.
func NewX1IntList(values ...int) *X1IntList {
	result := newX1IntList(len(values), len(values))
	copy(result.m, values)
	return result
}

// BuildX1IntListFromChan constructs a new X1IntList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildX1IntListFromChan(source <-chan int) *X1IntList {
	result := newX1IntList(0, 0)
	for v := range source {
		result.m = append(result.m, v)
	}
	return result
}

// ToSlice returns the elements of the current set as a slice
func (list *X1IntList) ToSlice() []int {
	list.s.RLock()
	defer list.s.RUnlock()

	var s []int
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list *X1IntList) Clone() *X1IntList {
	return NewX1IntList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
func (list *X1IntList) Get(i int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list *X1IntList) Head() int {
	return list.Get(0)
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list *X1IntList) Last() int {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list *X1IntList) Tail() *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list *X1IntList) Init() *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether X1IntList is empty.
func (list *X1IntList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether X1IntList is empty.
func (list *X1IntList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list *X1IntList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list *X1IntList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *X1IntList) Size() int {
	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This implements one of the methods needed by sort.Interface (along with Less and Swap).
func (list *X1IntList) Len() int {
	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This implements one of the methods needed by sort.Interface (along with Len and Less).
func (list *X1IntList) Swap(i, j int) {
	list.s.Lock()
	defer list.s.Unlock()

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list *X1IntList) Contains(v int) bool {
	return list.Exists(func (x int) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list *X1IntList) ContainsAll(i ...int) bool {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of X1IntList return true for the passed func.
func (list *X1IntList) Exists(fn func(int) bool) bool {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of X1IntList return true for the passed func.
func (list *X1IntList) Forall(fn func(int) bool) bool {
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over X1IntList and executes the passed func against each element.
// The function can safely alter the values via side-effects.
func (list *X1IntList) Foreach(fn func(int)) {
	list.s.Lock()
	defer list.s.Unlock()

	for _, v := range list.m {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list *X1IntList) Send() <-chan int {
	ch := make(chan int)
	go func() {
		list.s.RLock()
		defer list.s.RUnlock()

		for _, v := range list.m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of X1IntList with all elements in the reverse order.
func (list *X1IntList) Reverse() *X1IntList {
	list.s.Lock()
	defer list.s.Unlock()

	numItems := len(list.m)
	result := newX1IntList(numItems, numItems)
	last := numItems - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of X1IntList, using a version of the Fisher-Yates shuffle.
func (list *X1IntList) Shuffle() *X1IntList {
	result := list.Clone()
	numItems := len(result.m)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Add adds items to the current list. This is a synonym for Append.
func (list *X1IntList) Add(more ...int) {
	list.Append(more...)
}

// Append adds items to the current list, returning the modified list.
func (list *X1IntList) Append(more ...int) *X1IntList {
	list.s.Lock()
	defer list.s.Unlock()

	list.doAppend(more...)
	return list
}

func (list *X1IntList) doAppend(more ...int) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of X1IntList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1IntList) Take(n int) *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	if n > list.Len() {
		return list
	}
	result := newX1IntList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of X1IntList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1IntList) Drop(n int) *X1IntList {
	if n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, 0)
	l := list.Len()
	if n < l {
		result.m = list.m[n:]
	}
	return result
}

// TakeLast returns a slice of X1IntList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list *X1IntList) TakeLast(n int) *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	l := list.Len()
	if n > l {
		return list
	}
	result := newX1IntList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of X1IntList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *X1IntList) DropLast(n int) *X1IntList {
	if n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := list.Len()
	if n > l {
		list.m = list.m[l:]
	} else {
		list.m = list.m[0 : l-n]
	}
	return list
}

// TakeWhile returns a new X1IntList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list *X1IntList) TakeWhile(p func(int) bool) *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new X1IntList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list *X1IntList) DropWhile(p func(int) bool) *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, 0)
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

// Filter returns a new X1IntList whose elements return true for func.
func (list *X1IntList) Filter(fn func(int) bool) *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, list.Len()/2)

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
func (list *X1IntList) Partition(p func(int) bool) (*X1IntList, *X1IntList) {
	list.s.RLock()
	defer list.s.RUnlock()

	matching := newX1IntList(0, list.Len()/2)
	others := newX1IntList(0, list.Len()/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// CountBy gives the number elements of X1IntList that return true for the passed predicate.
func (list *X1IntList) CountBy(predicate func(int) bool) (result int) {
	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *X1IntList) Less(i, j int) bool {
	return list.m[i] < list.m[j]
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *X1IntList) Min() int {
	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *X1IntList) Max() (result int) {
	list.s.RLock()
	defer list.s.RUnlock()

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

// DistinctBy returns a new X1IntList whose elements are unique, where equality is defined by a passed func.
func (list *X1IntList) DistinctBy(equal func(int, int) bool) *X1IntList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := newX1IntList(0, list.Len())
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
func (list *X1IntList) IndexWhere(p func(int) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list *X1IntList) IndexWhere2(p func(int) bool, from int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list *X1IntList) LastIndexWhere(p func(int) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list *X1IntList) LastIndexWhere2(p func(int) bool, before int) int {
	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *X1IntList) Sum() int {
	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *X1IntList) Equals(other *X1IntList) bool {
	list.s.RLock()
	other.s.RLock()
	defer list.s.RUnlock()
	defer other.s.RUnlock()

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
func (list X1IntList) StringList() []string {
	list.s.RLock()
	defer list.s.RUnlock()

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *X1IntList) String() string {
	return list.MkString3("[", ", ", "]")
}

// implements json.Marshaler interface {
func (list X1IntList) MarshalJSON() ([]byte, error) {
	return list.mkString3Bytes("[\"", "\", \"", "\"]").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *X1IntList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *X1IntList) MkString3(pfx, mid, sfx string) string {
	return list.mkString3Bytes(pfx, mid, sfx).String()
}

func (list X1IntList) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}
