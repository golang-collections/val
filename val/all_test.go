// Copyright 2014 The b Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package val

import (
	"testing"
//	"fmt"
)

func TestWrapUnwrapBytes(t *testing.T) {


	g := Bytes([]byte("Gopher"))


	e := (&g).Wrap()

	(&g).Unwrap(e)

	if string(g) != "Gopher" {
		t.Fatal("Nope.")
	}
}

func TestWrapUnwrapIntegers(t *testing.T) {


	g := Ints([]int{1,3,3,7})


	e := (&g).Wrap()

	(&g).Unwrap(e)

	print()

	if g[0] != 1 || g[1] != 3 || g[2] != 3 || g[3] != 7 {
		t.Fatal("Nope.")
	}

//	for _, j := range e {
//		fmt.Printf("%d ", j);
//	}
}
