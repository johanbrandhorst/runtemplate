// A queue or fifo that holds Apple, implemented via a ring buffer. Unlike the list collections, these
// have a fixed size (although this can be changed when needed). For mutable collection that need frequent
// appending, the fixed size is a benefit because the memory footprint is constrained. However, this is
// not usable unless the rate of removing items from the queue is, over time, the same as the rate of addition.
// For similar reasons, there is no immutable variant of a queue.
//
// The queue provides a method to sort its elements.
//
// Thread-safe.
//
// Generated from threadsafe/queue.tpl with Type=Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Sorted:<no value> Stringer:<no value>
// ToList:<no value> ToSet:<no value>
// by runtemplate v2.3.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package examples

import (
	//
	"sort"
	"sync"
)

// AppleQueue is a ring buffer containing a slice of type Apple. It is optimised
// for FIFO operations.
type AppleQueue struct {
	m         []Apple
	read      int
	write     int
	length    int
	capacity  int
	overwrite bool
	less      func(i, j Apple) bool
	s         *sync.RWMutex
}

// NewAppleQueue returns a new queue of Apple. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
func NewAppleQueue(capacity int, overwrite bool) *AppleQueue {
	return NewAppleSortedQueue(capacity, overwrite, nil)
}

// NewAppleSortedQueue returns a new queue of Apple. The behaviour when adding
// to the queue depends on overwrite. If true, the push operation overwrites oldest values up to
// the space available, when the queue is full. Otherwise, it refuses to overfill the queue.
// If the 'less' comparison function is not nil, elements can be easily sorted.
func NewAppleSortedQueue(capacity int, overwrite bool, less func(i, j Apple) bool) *AppleQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}
	return &AppleQueue{
		m:         make([]Apple, capacity),
		read:      0,
		write:     0,
		length:    0,
		capacity:  capacity,
		overwrite: overwrite,
		less:      less,
		s:         &sync.RWMutex{},
	}
}

// IsSequence returns true for ordered lists and queues.
func (queue *AppleQueue) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (queue *AppleQueue) IsSet() bool {
	return false
}

// IsOverwriting returns true if the queue is overwriting, false if refusing.
func (queue AppleQueue) IsOverwriting() bool {
	return queue.overwrite
}

// IsEmpty returns true if the queue is empty.
func (queue *AppleQueue) IsEmpty() bool {
	if queue == nil {
		return true
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == 0
}

// NonEmpty returns true if the queue is not empty.
func (queue *AppleQueue) NonEmpty() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length > 0
}

// IsFull returns true if the queue is full.
func (queue *AppleQueue) IsFull() bool {
	if queue == nil {
		return false
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length == queue.capacity
}

// Space returns the space available in the queue.
func (queue *AppleQueue) Space() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.capacity - queue.length
}

// Size gets the number of elements currently in this queue. This is an alias for Len.
func (queue *AppleQueue) Size() int {
	if queue == nil {
		return 0
	}
	queue.s.RLock()
	defer queue.s.RUnlock()
	return queue.length
}

// Len gets the current length of this queue. This is an alias for Size.
func (queue *AppleQueue) Len() int {
	return queue.Size()
}

// Cap gets the capacity of this queue.
func (queue *AppleQueue) Cap() int {
	if queue == nil {
		return 0
	}
	return queue.capacity
}

// Less reports whether the element with index i should sort before the element with index j.
// The queue must have been created with a non-nil 'less' comparison function and it must not
// be empty.
func (queue *AppleQueue) Less(i, j int) bool {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	return queue.less(queue.m[ri], queue.m[rj])
}

// Swap swaps the elements with indexes i and j.
// The queue must not be empty.
func (queue *AppleQueue) Swap(i, j int) {
	ri := (queue.read + i) % queue.capacity
	rj := (queue.read + j) % queue.capacity
	queue.m[ri], queue.m[rj] = queue.m[rj], queue.m[ri]
}

// Sort sorts the queue using the 'less' comparison function, which must not be nil.
func (queue *AppleQueue) Sort() {
	sort.Sort(queue)
}

// StableSort sorts the queue using the 'less' comparison function, which must not be nil.
// The result is stable so that repeated calls will not arbitrarily swap equal items.
func (queue *AppleQueue) StableSort() {
	sort.Stable(queue)
}

// frontAndBack gets the front and back portions of the queue. The front portion starts
// from the read index. The back portion ends at the write index.
func (queue *AppleQueue) frontAndBack() ([]Apple, []Apple) {
	if queue == nil || queue.length == 0 {
		return nil, nil
	}
	if queue.write > queue.read {
		return queue.m[queue.read:queue.write], nil
	}
	return queue.m[queue.read:], queue.m[:queue.write]
}

// ToSlice returns the elements of the queue as a slice. The queue is not altered.
func (queue *AppleQueue) ToSlice() []Apple {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.toSlice(make([]Apple, queue.length))
}

func (queue *AppleQueue) toSlice(s []Apple) []Apple {
	front, back := queue.frontAndBack()
	copy(s, front)
	if len(back) > 0 && len(s) >= len(front) {
		copy(s[len(front):], back)
	}
	return s
}

// ToInterfaceSlice returns the elements of the queue as a slice of arbitrary type.
// The queue is not altered.
func (queue *AppleQueue) ToInterfaceSlice() []interface{} {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

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
func (queue *AppleQueue) Clone() *AppleQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	buffer := queue.toSlice(make([]Apple, queue.capacity))

	return &AppleQueue{
		m:         buffer,
		read:      0,
		write:     queue.length,
		length:    queue.length,
		capacity:  queue.capacity,
		overwrite: queue.overwrite,
		less:      queue.less,
		s:         &sync.RWMutex{},
	}
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the queue.
// Panics if the index is out of range or the queue is nil.
func (queue *AppleQueue) Get(i int) Apple {
	queue.s.RLock()
	defer queue.s.RUnlock()

	ri := (queue.read + i) % queue.capacity
	return queue.m[ri]
}

// Head gets the first element in the queue. Head is the opposite of Last.
// Panics if queue is empty or nil.
func (queue *AppleQueue) Head() Apple {
	queue.s.RLock()
	defer queue.s.RUnlock()

	return queue.m[queue.read]
}

// HeadOption returns the oldest item in the queue without removing it. If the queue
// is nil or empty, it returns the zero value instead.
func (queue *AppleQueue) HeadOption() Apple {
	if queue == nil {
		return *(new(Apple))
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	if queue.length == 0 {
		return *(new(Apple))
	}

	return queue.m[queue.read]
}

// Last gets the the newest item in the queue (i.e. last element pushed) without removing it.
// Last is the opposite of Head.
// Panics if queue is empty or nil.
func (queue *AppleQueue) Last() Apple {
	queue.s.RLock()
	defer queue.s.RUnlock()

	i := queue.write - 1
	if i < 0 {
		i = queue.capacity - 1
	}

	return queue.m[i]
}

// LastOption returns the newest item in the queue without removing it. If the queue
// is nil empty, it returns the zero value instead.
func (queue *AppleQueue) LastOption() Apple {
	if queue == nil {
		return *(new(Apple))
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

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
func (queue *AppleQueue) Reallocate(capacity int, overwrite bool) *AppleQueue {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}

	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doReallocate(capacity, overwrite)
}

func (queue *AppleQueue) doReallocate(capacity int, overwrite bool) *AppleQueue {
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
//func (queue *AppleQueue) Insert(items ...Apple) []Apple {
//	queue.s.Lock()
//	defer queue.s.Unlock()
//	return queue.doInsert(items...)
//}
//
//func (queue *AppleQueue) doInsert(items ...Apple) []Apple {
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
//func (queue *AppleQueue) doInsertOne(item Apple) {
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

// Push appends items to the end of the queue. If the queue does not have enough space,
// more will be allocated: how this happens depends on the overwriting mode.
//
// When overwriting, the oldest items are overwritten with the new data; it expands the queue
// only if there is still not enough space.
//
// Otherwise, the queue might be reallocated if necessary, ensuring that all the data is pushed
// without any older items being affected.
//
// The modified queue is returned.
func (queue *AppleQueue) Push(items ...Apple) *AppleQueue {
	queue.s.Lock()
	defer queue.s.Unlock()

	n := queue.capacity
	if queue.overwrite && len(items) > queue.capacity {
		n = len(items)
		// no rounding in this case because the old items are expected to be overwritten
	} else if !queue.overwrite && len(items) > (queue.capacity-queue.length) {
		n = len(items) + queue.length
		// rounded up to multiple of 128
		n = ((n + 127) / 128) * 128
	}

	if n > queue.capacity {
		queue = queue.doReallocate(n, queue.overwrite)
	}

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
// The queue capacity is never altered.
func (queue *AppleQueue) Offer(items ...Apple) []Apple {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPush(items...)
}

func (queue *AppleQueue) doPush(items ...Apple) []Apple {
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
func (queue *AppleQueue) Pop1() (Apple, bool) {
	queue.s.Lock()
	defer queue.s.Unlock()

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
func (queue *AppleQueue) Pop(n int) []Apple {
	queue.s.Lock()
	defer queue.s.Unlock()
	return queue.doPop(n)
}

func (queue *AppleQueue) doPop(n int) []Apple {
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

// Exists verifies that one or more elements of AppleQueue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *AppleQueue) Exists(p func(Apple) bool) bool {
	if queue == nil {
		return false
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

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

// Forall verifies that all elements of AppleQueue return true for the predicate p.
// The function should not alter the values via side-effects.
func (queue *AppleQueue) Forall(p func(Apple) bool) bool {
	if queue == nil {
		return true
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

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

// Foreach iterates over AppleQueue and executes function fn against each element.
// The function can safely alter the values via side-effects.
func (queue *AppleQueue) Foreach(fn func(Apple)) {
	if queue == nil {
		return
	}

	queue.s.Lock()
	defer queue.s.Unlock()

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
func (queue *AppleQueue) Find(p func(Apple) bool) (Apple, bool) {
	if queue == nil {
		return *(new(Apple)), false
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

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

// Filter returns a new AppleQueue whose elements return true for predicate p.
//
// The original queue is not modified. See also DoKeepWhere (which does modify the original queue).
func (queue *AppleQueue) Filter(p func(Apple) bool) *AppleQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	result := NewAppleSortedQueue(queue.length, queue.overwrite, queue.less)
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
func (queue *AppleQueue) Partition(p func(Apple) bool) (*AppleQueue, *AppleQueue) {
	if queue == nil {
		return nil, nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	matching := NewAppleSortedQueue(queue.length, queue.overwrite, queue.less)
	others := NewAppleSortedQueue(queue.length, queue.overwrite, queue.less)
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

// Map returns a new AppleQueue by transforming every element with a function fn.
// The resulting queue is the same size as the original queue.
// The original queue is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (queue *AppleQueue) Map(fn func(Apple) Apple) *AppleQueue {
	if queue == nil {
		return nil
	}

	queue.s.RLock()
	defer queue.s.RUnlock()

	result := NewAppleSortedQueue(queue.length, queue.overwrite, queue.less)
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

// CountBy gives the number elements of AppleQueue that return true for the passed predicate.
//func (queue *AppleQueue) CountBy(predicate func(Apple) bool) (result int) {
//	queue.s.RLock()
//	defer queue.s.RUnlock()
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
