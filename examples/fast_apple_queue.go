// A queue or fifo that holds Apple, implemented via a ring buffer. Unlike the list collections, these
// have a fixed size (although this can be changed when needed). For mutable collection that need frequent
// appending, the fixed size is a benefit because the memory footprint is constrained. However, this is
// not usable unless the rate of removing items from the queue is, over time, the same as the rate of addition.
// For similar reasons, there is no immutable variant of a queue.
//
// The queue provides a method to sort its elements.
//
// Not thread-safe.
//
// Generated from fast/queue.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Sorted:<no value> Stringer:<no value>
// ToList:<no value> ToSet:<no value>
// by runtemplate v2.3.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	//
	"sort"
)

// FastAppleQueue is a ring buffer containing a slice of type Apple. It is optimised
// for FIFO operations.
type FastAppleQueue struct {
	m         []Apple
	read      int
	write     int
	length    int
	capacity  int
	overwrite bool
	less      func(i, j Apple) bool
}

// NewFastAppleQueue returns a new queue of Apple. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
// If the 'less' comparison function is not nil, elements can be easily sorted.
func NewFastAppleQueue(capacity int, overwrite bool, less func(i, j Apple) bool) *FastAppleQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}
	return &FastAppleQueue{
		m:         make([]Apple, capacity),
		read:      0,
		write:     0,
		length:    0,
		capacity:  capacity,
		overwrite: overwrite,
		less:      less,
	}
}

// IsSequence returns true for ordered lists and queues.
func (queue *FastAppleQueue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *FastAppleQueue) IsSet() bool {
	return false
}

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue FastAppleQueue) IsOverwriting() bool {
	return queue.overwrite
}

// IsEmpty returns true if the queue is empty.
func (queue *FastAppleQueue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *FastAppleQueue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue *FastAppleQueue) IsFull() bool {
	if queue == nil {
		return false
	}
	return queue.length == queue.capacity
}

// Space returns the space available in the queue.
func (queue *FastAppleQueue) Space() int {
	if queue == nil {
		return 0
	}
	return queue.capacity - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *FastAppleQueue) Size() int {
	if queue == nil {
		return 0
	}
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *FastAppleQueue) Len() int {
	if queue == nil {
		return 0
	}
	return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue *FastAppleQueue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.capacity
}

// Less reports whether the element with index i should sort before the element with index j.
// The queue must have been created with a non-nil 'less' comparison function and it must not
// be empty.
func (queue *FastAppleQueue) Less(i, j int) bool {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	return queue.less(queue.m[ri], queue.m[rj])
}

// Swap swaps the elements with indexes i and j.
// The queue must not be empty.
func (queue *FastAppleQueue) Swap(i, j int) {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	queue.m[ri], queue.m[rj] = queue.m[rj], queue.m[ri]
}

// Sort sorts the queue using the 'less' comparison function, which must not be nil.
func (queue *FastAppleQueue) Sort() {
	sort.Sort(queue)
}

// StableSort sorts the queue using the 'less' comparison function, which must not be nil.
// The result is stable so that repeated calls will not arbtrarily swap equal items.
func (queue *FastAppleQueue) StableSort() {
	sort.Stable(queue)
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *FastAppleQueue) frontAndBack() ([]Apple, []Apple) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.m[queue.read:queue.write], nil
	}
	return queue.m[queue.read:], queue.m[:queue.write]
}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *FastAppleQueue) ToSlice() []Apple {
	if queue == nil {
		return nil
	}

	return queue.toSlice(make([]Apple, queue.length))
}

func (queue *FastAppleQueue) toSlice(s []Apple) []Apple {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *FastAppleQueue) ToInterfaceSlice() []interface{} {
	if queue == nil {
		return nil
	}

	front, back := queue.frontAndBack()
	var s []interface{}
	for _, v := range front {
		s = append(s, v)
	}

	for _, v := range back {
		s = append(s, v)
	}

	return s
}

// Clone returns a shallow copy of the queue. It does not clone the underlying elements.
func (queue *FastAppleQueue) Clone() *FastAppleQueue {
	if queue == nil {
		return nil
	}

	buffer := queue.toSlice(make([]Apple, queue.capacity))

	return &FastAppleQueue{
		m:         buffer,
		read:      0,
		write:     queue.length,
		length:    queue.length,
		capacity:  queue.capacity,
		overwrite: queue.overwrite,
	}
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the queue.
// Panics if the index is out of range or the queue is nil.
func (queue *FastAppleQueue) Get(i int) Apple {

	ri := (queue.read + i) % queue.capacity
	return queue.m[ri]
}

// Head gets the first element in the queue. Head is the opposite of Last.
// Panics if queue is empty or nil.
func (queue *FastAppleQueue) Head() Apple {

	return queue.m[queue.read]
}

// HeadOption returns the oldest item in the queue without removing it. If the queue
// is nil or empty, it returns the zero value instead.
func (queue *FastAppleQueue) HeadOption() Apple {
	if queue == nil {
		return *(new(Apple))
	}

	if queue.length == 0 {
		return *(new(Apple))
	}

	return queue.m[queue.read]
}

// Last gets the the newest item in the queue (i.e. last element pushed) without removing it.
// Last is the opposite of Head.
// Panics if queue is empty or nil.
func (queue *FastAppleQueue) Last() Apple {

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is nil empty, it returns the zero value instead.
func (queue *FastAppleQueue) LastOption() Apple {
	if queue == nil {
		return *(new(Apple))
	}

	if queue.length == 0 {
		return *(new(Apple))
	}

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

//-------------------------------------------------------------------------------------------------

// Reallocate adjusts the allocated capacity of the queue and allows the overwriting behaviour to be changed.
// If the new queue capacity is less than the old capacity, the oldest items in the queue are discarded so
// that the remaining data can fit in the space available.
//
// If the new queue capacity is the same as the old capacity, the queue is not altered except for adopting
// the new overwrite flag's value. Therefore this is the means to change the overwriting behaviour.
//
// Reallocate adjusts the storage space but does not clone the underlying elements.
//
// The queue must not be nil.
func (queue *FastAppleQueue) Reallocate(capacity int, overwrite bool) *FastAppleQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}

	queue.overwrite = overwrite

	if capacity < queue.length {
		// existing data is too big and has to be trimmed to fit
		n := queue.length - capacity
		queue.read = (queue.read + n) % queue.capacity
		queue.length -= n
	}

	if capacity != queue.capacity {
		oldLength := queue.length
		queue.m = queue.toSlice(make([]Apple, capacity))
		if oldLength > len(queue.m) {
			oldLength = len(queue.m)
		}
		queue.read = 0
		queue.write = oldLength
		queue.length = oldLength
		queue.capacity = capacity
	}

	return queue
}

//-------------------------------------------------------------------------------------------------

// Insert adds items to the queue in sorted order.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
//func (queue *FastAppleQueue) Insert(items ...Apple) []Apple {
//	return queue.doInsert(items...)
//}
//
//func (queue *FastAppleQueue) doInsert(items ...Apple) []Apple {
//	n := len(items)
//
//	space := queue.capacity - queue.length
//	overwritten := n - space
//
//	if queue.overwrite {
//		space = queue.capacity
//	}
//
//	if space < n {
//		// there is too little space; reject surplus elements
//		surplus := items[space:]
//		queue.doInsert(items[:space]...)
//		return surplus
//	}
//
//  for _, item := range items {
//      queue.doInsertItem(item)
//  }
//	return nil
//}
//
//func (queue *FastAppleQueue) doInsertOne(item Apple) {
//	if queue.write < queue.capacity {
//		// easy case: enough space at end for the item
//		queue.m[queue.write] = item
//		queue.write = (queue.write + n) % queue.capacity
//		queue.length++
//		return
//	}
//
//	end := queue.capacity - queue.write
//	queue.m[queue.write] = items
//	//copy(queue.m, items[end:])
//	queue.write = n - end
//	queue.length++
//	if queue.length > queue.capacity {
//		queue.length = queue.capacity
//	}
//	if overwritten > 0 {
//		queue.read = (queue.read + overwritten) % queue.capacity
//	}
//}

// Push appends items to the end of the queue.
// This panics if the queue does not have enough space.
func (queue *FastAppleQueue) Push(items ...Apple) *FastAppleQueue {

	overflow := queue.doPush(items...)
	if len(overflow) > 0 {
		panic(len(overflow))
	}
	return queue
}

// Offer appends as many items to the end of the queue as it can.
// If the queue is already full, what happens depends on whether the queue is configured
// to overwrite. If it is, the oldest items will be overwritten. Otherwise, it will be
// filled to capacity and any unwritten items are returned.
//
// If the capacity is too small for the number of items, the excess items are returned.
func (queue *FastAppleQueue) Offer(items ...Apple) []Apple {
	return queue.doPush(items...)
}

func (queue *FastAppleQueue) doPush(items ...Apple) []Apple {
	n := len(items)

	space := queue.capacity - queue.length
	overwritten := n - space

	if queue.overwrite {
		space = queue.capacity
	}

	if space < n {
		// there is too little space; reject surplus elements
		surplus := items[space:]
		queue.doPush(items[:space]...)
		return surplus
	}

	if n <= queue.capacity-queue.write {
		// easy case: enough space at end for all items
		copy(queue.m[queue.write:], items)
		queue.write = (queue.write + n) % queue.capacity
		queue.length += n
		return nil
	}

	// not yet full
	end := queue.capacity - queue.write
	copy(queue.m[queue.write:], items[:end])
	copy(queue.m, items[end:])
	queue.write = n - end
	queue.length += n
	if queue.length > queue.capacity {
		queue.length = queue.capacity
	}
	if overwritten > 0 {
		queue.read = (queue.read + overwritten) % queue.capacity
	}
	return nil
}

// Pop1 removes and returns the oldest item from the queue. If the queue is
// empty, it returns the zero value instead.
// The boolean is true only if the element was available.
func (queue *FastAppleQueue) Pop1() (Apple, bool) {

	if queue.length == 0 {
		return *(new(Apple)), false
	}

	v := queue.m[queue.read]
	queue.read = (queue.read + 1) % queue.capacity
	queue.length--

	return v, true
}

// Pop removes and returns the oldest items from the queue. If the queue is
// empty, it returns a nil slice. If n is larger than the current queue length,
// it returns all the available elements, so in this case the returned slice
// will be shorter than n.
func (queue *FastAppleQueue) Pop(n int) []Apple {
	return queue.doPop(n)
}

func (queue *FastAppleQueue) doPop(n int) []Apple {
	if queue.length == 0 {
		return nil
	}

	if n > queue.length {
		n = queue.length
	}

	s := make([]Apple, n)
	front, back := queue.frontAndBack()
	// note the length copied is whichever is shorter
	copy(s, front)
	if n > len(front) {
		copy(s[len(front):], back)
	}

	queue.read = (queue.read + n) % queue.capacity
	queue.length -= n

	return s
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of FastAppleQueue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *FastAppleQueue) Exists(p func(Apple) bool) bool {
	if queue == nil {
		return false
	}

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			return true
		}
	}
	for _, v := range back {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of FastAppleQueue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *FastAppleQueue) Forall(p func(Apple) bool) bool {
	if queue == nil {
		return true
	}

	front, back := queue.frontAndBack()
	for _, v := range front {
		if !p(v) {
			return false
		}
	}
	for _, v := range back {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over FastAppleQueue and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (queue *FastAppleQueue) Foreach(fn func(Apple)) {
	if queue == nil {
		return
	}

	front, back := queue.frontAndBack()
	for _, v := range front {
		fn(v)
	}
	for _, v := range back {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first Apple that returns true for predicate p.
// False is returned if none match.
func (queue *FastAppleQueue) Find(p func(Apple) bool) (Apple, bool) {
	if queue == nil {
		return *(new(Apple)), false
	}

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			return v, true
		}
	}
	for _, v := range back {
		if p(v) {
			return v, true
		}
	}

	var empty Apple
	return empty, false
}

// Filter returns a new FastAppleQueue whose elements return true for predicate p.
//
// The original queue is not modified. See also DoKeepWhere (which does modify the original queue).
func (queue *FastAppleQueue) Filter(p func(Apple) bool) *FastAppleQueue {
	if queue == nil {
		return nil
	}

	result := NewFastAppleQueue(queue.length, queue.overwrite, queue.less)
	i := 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			result.m[i] = v
			i++
		}
	}
	for _, v := range back {
		if p(v) {
			result.m[i] = v
			i++
		}
	}
	result.length = i
	result.write = i

	return result
}

// Partition returns two new AppleQueues whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original queue.
//
// The original queue is not modified
func (queue *FastAppleQueue) Partition(p func(Apple) bool) (*FastAppleQueue, *FastAppleQueue) {
	if queue == nil {
		return nil, nil
	}

	matching := NewFastAppleQueue(queue.length, queue.overwrite, queue.less)
	others := NewFastAppleQueue(queue.length, queue.overwrite, queue.less)
	m, o := 0, 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		if p(v) {
			matching.m[m] = v
			m++
		} else {
			others.m[o] = v
			o++
		}
	}
	for _, v := range back {
		if p(v) {
			matching.m[m] = v
			m++
		} else {
			others.m[o] = v
			o++
		}
	}
	matching.length = m
	matching.write = m
	others.length = o
	others.write = o

	return matching, others
}

// Map returns a new FastAppleQueue by transforming every element with a function fn.
// The resulting queue is the same size as the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *FastAppleQueue) Map(fn func(Apple) Apple) *FastAppleQueue {
	if queue == nil {
		return nil
	}

	result := NewFastAppleQueue(queue.length, queue.overwrite, queue.less)
	i := 0

	front, back := queue.frontAndBack()
	for _, v := range front {
		result.m[i] = fn(v)
		i++
	}
	for _, v := range back {
		result.m[i] = fn(v)
		i++
	}
	result.length = i
	result.write = i

	return result
}

// CountBy gives the number elements of FastAppleQueue that return true for the passed predicate.
//func (queue *FastAppleQueue) CountBy(predicate func(Apple) bool) (result int) {
//
//	front, back := queue.frontAndBack()
//	for _, v := range front {
//		if predicate(v) {
//			result++
//		}
//	}
//	for _, v := range back {
//		if predicate(v) {
//			result++
//		}
//	}
//	return
//}
//