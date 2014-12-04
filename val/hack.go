package val

// This provides fast additional wrappers for a different collection BigVals

import (
//	"github.com/davecgh/go-spew/spew"
	u "unsafe"	// our good friends reflect and unsafe
	r "reflect"
)

// FIXME if gc eats our data, we may want to stretch the
// byte slice size a bit to accomodate the larger size of integers

func (i *Ints) Wrap() (o BigVal) {
	return Get(i).Wrap()
}

func (i *Ints) Unwrap(o BigVal) {
	Get(i).Unwrap(o)
}

func (i Ptrs) Wrap() (o BigVal) {
	p := (*r.SliceHeader)(u.Pointer(i))
	q := (*r.SliceHeader)(u.Pointer(&o))

	*q = *p

	return
}

func (i Ptrs) Unwrap(o BigVal) {
	p := (*r.SliceHeader)(u.Pointer(i))
	q := (*r.SliceHeader)(u.Pointer(&o))

	*p = *q
}

func Get(i interface{}) Ptrs {
	return Ptrs(r.ValueOf(i).Pointer())
}
