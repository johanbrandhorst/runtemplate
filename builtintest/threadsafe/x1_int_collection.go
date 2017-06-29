// Generated from threadsafe/collection.tpl with Type=int
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always

package threadsafe

// X1IntSizer defines an interface for sizing methods on int collections.
type X1IntSizer interface {
	// IsEmpty tests whether X1IntCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether X1IntCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}


// X1IntMkStringer defines an interface for stringer methods on int collections.
type X1IntMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(pfx, mid, sfx string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// X1IntCollection defines an interface for common collection methods on int.
type X1IntCollection interface {
	X1IntSizer

	X1IntMkStringer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []int

	// Exists verifies that one or more elements of X1IntCollection return true for the passed func.
	Exists(fn func(int) bool) bool

	// Forall verifies that all elements of X1IntCollection return true for the passed func.
	Forall(fn func(int) bool) bool

	// Foreach iterates over X1IntCollection and executes the passed func against each element.
	Foreach(fn func(int))

    // Find returns the first int that returns true for some function.
    // False is returned if none match.
    Find(fn func(int) bool) (int, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan int

	// CountBy gives the number elements of X1IntCollection that return true for the passed predicate.
	CountBy(predicate func(int) bool) int


	// Contains determines if a given item is already in the collection.
	Contains(v int) bool

	// ContainsAll determines if the given items are all in the collection.
	ContainsAll(v ...int) bool

// Add adds items to the current collection.
	Add(more ...int)


	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() int

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() int


	// Sum returns the sum of all the elements in the collection.
	Sum() int
}
