package maybe

import (
	"testing"
	"testing/quick"
	"math/rand"
	"reflect"
)

func (m Nothing) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Nothing{})
}

func (m Just) Generate(rand *rand.Rand, size int) reflect.Value {
	// Return a Just{i} where -size/2 <= i < size/2
	return reflect.ValueOf(Just{rand.Intn(size) - size/2})
}

// While Maybe is generic (in the sense that it wraps a Value (an interface{})),
// quick can't generate test cases for a Value. These tests thus use int values,
// purely so we can use quick. Imagine that the tests generated random arbitrary
// values!
func TestJustThenBindOperatesOnJust(t *testing.T) {
	f := func(v int) Maybe {
		return Just{v * 2}
	}
	g := func(v int) Maybe {
		return Just{v}.Bind(func (x Value) Maybe { return Just{x.(int) * 2}})
	}

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func TestNothingThenBindReturnsNothing(t *testing.T) {
	f := func(v int) Maybe {
		return Nothing{}
	}
	g := func(v int) Maybe {
		return Nothing{}.Bind(func (x Value) Maybe { return Just{x.(int) * 2}})
	}

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func TestJustOtherwiseReturnsOtherwise(t *testing.T) {
	answer := 22
	f := func(v int) Value {
		return answer
	}
	g := func(v int) Value {
		return Just{answer}.Otherwise(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}


func TestNothingOtherwiseReturnsOtherwise(t *testing.T) {
	f := func(v int) Value {
		return v
	}
	g := func(v int) Value {
		return Nothing{}.Otherwise(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
