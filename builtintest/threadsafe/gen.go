// This package contains example collection types using the thread-safe templates.
// Encapsulation of the underlying data is also provided.
package threadsafe

// Code generation with non-pointer values

//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X1 Type=string Stringer:true Comparable:true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X1 Type=int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X1 Type=Apple  Stringer:false
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X1 Type=Pear
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=X2 Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=X1 Type=string Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=X1 Type=int    Stringer:true  Comparable:true Ordered:true  Numeric:true GobEncode:true JsonEncode:true
//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=X1 Type=Apple  Stringer:false Comparable:true
//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=X2 Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl threadsafe/queue.tpl      Prefix=X1 Type=string  Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl threadsafe/queue.tpl      Prefix=X1 Type=int     Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl threadsafe/queue.tpl      Prefix=X1 Type=Apple   Stringer:false Comparable:true
//go:generate runtemplate -tpl threadsafe/queue.tpl      Prefix=X2 Type=big.Int Import:"math/big"

//go:generate runtemplate -tpl threadsafe/set.tpl        Prefix=X1 Type=string Stringer:true  Ordered:false Numeric:false
//go:generate runtemplate -tpl threadsafe/set.tpl        Prefix=X1 Type=int    Stringer:true  Ordered:true  Numeric:true GobEncode:true JsonEncode:true
//go:generate runtemplate -tpl threadsafe/set.tpl        Prefix=X1 Type=Apple  Stringer:false
//go:generate runtemplate -tpl threadsafe/set.tpl        Prefix=X2 Type=url.URL Stringer:true  Comparable:true Import:"net/url"
//go:generate runtemplate -tpl threadsafe/set.tpl        Prefix=X2 Type=testtypes.Email Import:"github.com/rickb777/runtemplate/builtintest/testtypes"

//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX1 Key=int    Type=int     Comparable:true Stringer:true Numeric:true GobEncode:true JsonEncode:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX1 Key=string Type=string  Comparable:true Stringer:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX1 Key=string Type=Apple                   Stringer:true KeySlice:sort.StringSlice
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX1 Key=Email  Type=string                  Stringer:true KeySlice:EmailSlice
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX1 Key=Apple  Type=string
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX1 Key=Apple  Type=Pear                    Stringer:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TX2 Key=Apple  Type=big.Int  Import:"math/big"

//go:generate runtemplate -tpl types/stringy.tpl         Prefix=X1 Type=Email SortableSlice:true
//go:generate runtemplate -tpl plumbing/plumbing.tpl     Prefix=X1 Type=Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl        Prefix=X1 Type=Apple ToPrefix=X1 ToType=Pear

//go:generate runtemplate -tpl ../collection_test.tpl    Type=int Mutable:true Numeric:true Comparable:true
//go:generate runtemplate -tpl ../list_test.tpl          Type=int Mutable:true Numeric:true Comparable:true M:.slice() GobEncode:true JsonEncode:true Append:true
//go:generate runtemplate -tpl ../queue_test.tpl         Type=int Mutable:true M:
//go:generate runtemplate -tpl ../set_test.tpl           Type=int Mutable:true Numeric:true Comparable:true M:.slice() GobEncode:true JsonEncode:true Append:true
//go:generate runtemplate -tpl ../map_test.tpl   Key=int Type=int Mutable:true Numeric:true Comparable:true M:.slice() GobEncode:true JsonEncode:true

// Code generation with pointer values

//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P1 Type=*string Stringer:true Comparable:true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P1 Type=*int    Stringer:true Comparable:true Ordered:true Numeric:true
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P1 Type=*Apple  Stringer:false
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P1 Type=*Pear
//go:generate runtemplate -tpl threadsafe/collection.tpl Prefix=P2 Type=*big.Int Import:"math/big"

//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=P1 Type=*string Stringer:true  Comparable:true Ordered:false Numeric:false
//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=P1 Type=*int    Stringer:true  Comparable:true Ordered:true  Numeric:true
//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=P1 Type=*Apple  Stringer:false Comparable:true
//go:generate runtemplate -tpl threadsafe/list.tpl       Prefix=P2 Type=*big.Int Import:"math/big"

//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TP1 Key=*int    Type=*int     Comparable:true Stringer:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TP1 Key=*string Type=*string  Comparable:true
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TP1 Key=*string Type=*Apple
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TP1 Key=*Apple  Type=*string
//go:generate runtemplate -tpl threadsafe/map.tpl        Prefix=TP1 Key=*Apple  Type=*Pear

//go:generate runtemplate -tpl plumbing/plumbing.tpl     Prefix=P1 Type=*Apple
//go:generate runtemplate -tpl plumbing/mapTo.tpl        Prefix=P1 Type=*Apple ToPrefix=P1 ToType=*Pear

type Apple struct {
	N int
}

type Pear struct {
	K int
}

var _ X1StringCollection = NewX1StringList()
var _ X1IntCollection = NewX1IntList()
var _ X1AppleCollection = NewX1AppleList()

var _ X1StringCollection = NewX1StringSet()
var _ X1IntCollection = NewX1IntSet()
var _ X1AppleCollection = NewX1AppleSet()
