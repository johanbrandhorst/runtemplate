// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Comparable:{{.Comparable}} Numeric:{{.Numeric}} Ordered:{{.Ordered}} Stringer:{{.Stringer}} Mutable:always
// by runtemplate {{.AppVersion}}
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package {{.Package}}

{{if .HasImport}}
import (
	{{.Import}}
)

{{end -}}
// {{.UPrefix}}{{.UType}}Sizer defines an interface for sizing methods on {{.PType}} collections.
type {{.UPrefix}}{{.UType}}Sizer interface {
	// IsEmpty tests whether {{.UPrefix}}{{.UType}}Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether {{.UPrefix}}{{.UType}}Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}
{{- if .Stringer}}

// {{.UPrefix}}{{.UType}}MkStringer defines an interface for stringer methods on {{.PType}} collections.
type {{.UPrefix}}{{.UType}}MkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	//MarshalJSON() ([]byte, error)

	// implements json.Unmarshaler interface {
	//UnmarshalJSON(b []byte) error

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}
{{- end}}

// {{.UPrefix}}{{.UType}}Collection defines an interface for common collection methods on {{.PType}}.
type {{.UPrefix}}{{.UType}}Collection interface {
	{{.UPrefix}}{{.UType}}Sizer
{{- if .Stringer}}
	{{.UPrefix}}{{.UType}}MkStringer
{{- end}}

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for list and queues.
	IsSet() bool
{{- if .ToList}}

    // ToList returns a shallow copy as a list.
    ToList() {{.UPrefix}}{{.UType}}List
{{- end}}
{{- if .ToSet}}

    // ToSet returns a shallow copy as a set.
    ToSet() {{.UPrefix}}{{.UType}}Set
{{- end}}

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []{{.PType}}

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of {{.UPrefix}}{{.UType}}Collection return true for the predicate p.
	Exists(p func({{.PType}}) bool) bool

	// Forall verifies that all elements of {{.UPrefix}}{{.UType}}Collection return true for the predicate p.
	Forall(p func({{.PType}}) bool) bool

	// Foreach iterates over {{.UPrefix}}{{.UType}}Collection and executes the function f against each element.
	Foreach(f func({{.PType}}))

	// Find returns the first {{.Type}} that returns true for the predicate p.
	// False is returned if none match.
	Find(p func({{.PType}}) bool) ({{.PType}}, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan {{.PType}}

	// CountBy gives the number elements of {{.UPrefix}}{{.UType}}Collection that return true for the predicate p.
	CountBy(p func({{.PType}}) bool) int
{{- if .Comparable}}

	// Contains determines if a given item is already in the collection.
	Contains(v {{.Type}}) bool

	// ContainsAll determines if the given items are all in the collection.
	ContainsAll(v ...{{.Type}}) bool
{{- end}}

	// Add adds items to the current collection.
	//Add(more ...{{.Type}})
{{- if .Ordered}}

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() {{.Type}}

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() {{.Type}}
{{- end}}

	// MinBy returns an element of {{.UPrefix}}{{.UType}}Collection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}}

	// MaxBy returns an element of {{.UPrefix}}{{.UType}}Collection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func({{.PType}}, {{.PType}}) bool) {{.PType}}
{{- if .Numeric}}

	// Sum returns the sum of all the elements in the collection.
	Sum() {{.Type}}
{{- end}}
}