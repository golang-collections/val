// Package Val provides a value to be stored inside a collection 
package val

type BigVal []byte

// The types (concrete slices) that can be BigVal
type Ints []int
type Bytes []byte
type Ptrs uintptr	// a pointer to concrete slice
//TODO etc..

