package numset

import (
	"errors"
	"runtime"
)

// #include "numset.h"
import "C"

var (
	OutOfRange = errors.New("Index out of range.")
	MaxTooBig = errors.New("max is too big to hold in a C unsigned int.")
)

// The largest value that can be stored in a C unsigned int
const MaxCUint = uint(^C.uint(0))

// A Set holds uints below a fixed maximum value.
// This behaves like [max]bool; it takes up about twice as much space
// but can be zeroed in constant time.
type Set struct {
	data *C.struct_numset
	max  uint
}

// Creates a new *Set that will hold uints in [0, max).
// The value of max must be small enough to fit in a C unsigned int (that is, no bigger than MaxCUint).
// If not, the error MaxTooBig is returned.
//
// O(1) plus allocation costs
func New(max uint) (*Set, error) {
	if max > MaxCUint {
		return nil, MaxTooBig
	}
	ret := &Set{max: max}
	ret.data = C.make_ns(C.uint(max))
	runtime.SetFinalizer(ret, finalizer)
	return ret, nil
}

func finalizer(s *Set) {
	C.free_ns(s.data)
}

// Returns the maximum passed to New().
//
// O(1)
func (s *Set) Max() uint {
	return s.max
}

// Returns whether the set contains the given index;
// returns false, OutOfRange if index >= max.
// Equivalent to set[index].
//
// O(1)
func (s *Set) Get(index uint) (bool, error) {
	if index >= s.max {
		return false, OutOfRange
	}
	return (C.get_ns(s.data, C.uint(index)) != 0), nil
}

// Makes the set contain the given index; idempotent.
// Equivalent to set[index] = true.
//
// O(1)
func (s *Set) Set(index uint) error {
	if index >= s.max {
		return OutOfRange
	}
	if C.get_ns(s.data, C.uint(index)) != 1 {
		C.set_ns(s.data, C.uint(index))
	}
	return nil
}

// Makes the set not contain the given index; idempotent.
// Equivalent to set[index] = false.
//
// O(1)
func (s *Set) Unset(index uint) error {
	if index >= s.max {
		return OutOfRange
	}
	if C.get_ns(s.data, C.uint(index)) != 0 {
		C.unset_ns(s.data, C.uint(index))
	}
	return nil
}

// Makes the set empty; idempotent.
// Equivalent to a for-loop that sets set[index]=false for each index.
//
// O(1)
func (s *Set) Clear() {
	C.clear_ns(s.data)
}

// The number of elements in a set. Equivalent to len() applied to a map.
//
// O(1)
func (s *Set) Size() uint {
	return uint(C.size_ns(s.data))
}
