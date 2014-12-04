package val

type Wrapper interface {
	Wrap() BigVal
	Unwrap(BigVal)
}

func (b *Bytes) Wrap() BigVal {
	return BigVal(*b)
}

func (b *Bytes) Unwrap(o BigVal) {
	*b = Bytes(o)
}
