package maybe

type Value interface{}

// Maybe represents a computation that might fail.
type Maybe interface {
	// Bind connects two operations together. If the first computation
	// (the receiver) succeeds, Maybe applies the function argument
	// to the successful computation.
	Bind(func(v Value) Maybe) Maybe

	// Otherwise returns the receiver's Value if the receiver computation
	// succeeded, or otherwise v.
	// Otherwise(v Value) Value
}

// Nothing represents a failed computation
type Nothing struct{}

func (n Nothing) Bind(f func(v Value) Maybe) Maybe {
	return n
}

// Just represents a computation that has succeeded, or at least
// not yet failed.
type Just struct {
	Value Value
}

func (j Just) Bind(f func(v Value) Maybe) Maybe {
	return f(j.Value)
}
