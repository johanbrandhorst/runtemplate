// A derived string-based type compatible with marshalling and database APIs.
//
// Generated from {{.TemplateFile}} with Type={{.PType}}

package {{.Package}}

import (
	"errors"
	"sort"
	"strings"
	"database/sql/driver"
	"fmt"
)

// {{.Type}} is a specialised kind of string.
type {{.Type}} string

// Ptr returns the address of a {{.Type}}.
func ({{.LType}} {{.Type}}) Ptr() *{{.Type}} {
	return &{{.LType}}
}

// String converts to a string and implements fmt.Stringer.
func ({{.LType}} {{.Type}}) String() string {
	return string({{.LType}})
}

// TrimSpace removes surrounding whitespace.
func ({{.LType}} {{.Type}}) TrimSpace() {{.Type}} {
	return {{.Type}}(strings.TrimSpace({{.LType}}.String()))
}

// ToLower converts the value to lowercase.
func ({{.LType}} {{.Type}}) ToLower() {{.Type}} {
	return {{.Type}}(strings.ToLower(string({{.LType}})))
}

// ToUpper converts the value to uppercase.
func ({{.LType}} {{.Type}}) ToUpper() {{.Type}} {
	return {{.Type}}(strings.ToUpper(string({{.LType}})))
}

//-------------------------------------------------------------------------------------------------

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func ({{.LType}} *{{.Type}}) Scan(value interface{}) (err error) {
	err = nil
	switch value.(type) {
	case string:
		*{{.LType}} = {{.Type}}(value.(string))
	case []byte:
		*{{.LType}} = {{.Type}}(string(value.([]byte)))
	default:
		err = errors.New(fmt.Sprintf("%+v", value))
	}
	return
}

// Value converts the value to a string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func ({{.LType}} {{.Type}}) Value() (driver.Value, error) {
	return string({{.LType}}), nil
}

//-------------------------------------------------------------------------------------------------

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// https://golang.org/pkg/encoding/#TextMarshaler
func ({{.LType}} {{.Type}}) MarshalText() (text []byte, err error) {
	return []byte({{.LType}}.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
// https://golang.org/pkg/encoding/#TextUnmarshaler
func ({{.LType}} *{{.Type}}) UnmarshalText(text []byte) error {
	return {{.LType}}.Scan(text)
}

//-------------------------------------------------------------------------------------------------

// {{.UType}}Slice attaches the methods of sort.Interface to []{{.Type}}, sorting in increasing order.
type {{.UType}}Slice []{{.Type}}

func (p {{.UType}}Slice) Len() int           { return len(p) }
func (p {{.UType}}Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p {{.UType}}Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p {{.UType}}Slice) Sort() { sort.Sort(p) }
