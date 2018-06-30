// Generated from fast/collection.tpl with Type=*Apple
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:false Mutable:always

package fast

// P1AppleSizer defines an interface for sizing methods on *Apple collections.
type P1AppleSizer interface {
	// IsEmpty tests whether P1AppleCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether P1AppleCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// P1AppleCollection defines an interface for common collection methods on *Apple.
type P1AppleCollection interface {
	P1AppleSizer


	// IsSequence returns true for lists.
	IsSequence() bool

	// IsSet returns false for lists.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []*Apple

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of P1AppleCollection return true for the passed func.
	Exists(fn func(*Apple) bool) bool

	// Forall verifies that all elements of P1AppleCollection return true for the passed func.
	Forall(fn func(*Apple) bool) bool

	// Foreach iterates over P1AppleCollection and executes the passed func against each element.
	Foreach(fn func(*Apple))

	// Find returns the first Apple that returns true for some function.
	// False is returned if none match.
	Find(fn func(*Apple) bool) (*Apple, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan *Apple

	// CountBy gives the number elements of P1AppleCollection that return true for the passed predicate.
	CountBy(predicate func(*Apple) bool) int

// Add adds items to the current collection.
	Add(more ...Apple)

// MinBy returns an element of P1AppleCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(*Apple, *Apple) bool) *Apple

	// MaxBy returns an element of P1AppleCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(*Apple, *Apple) bool) *Apple

}