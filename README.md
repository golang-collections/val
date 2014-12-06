val
===

Val provides the BigVal element type.

So we have a BigVal, internally a []byte. This is going to be the element stored in the user defined collections.

BigVal currently supports storing of any slice. Storing  of arbitrary interface{} in BigVal as a value is also
planned and expected to be easy, but it's not currently a priority, since arbitrary slices seem to be more flexible.
