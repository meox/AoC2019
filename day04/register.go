package main

import "github.com/pkg/errors"

const (
	lenPass = 6
)

type Register struct {
	data [lenPass]int
}

func NewRegister(n int) (Register, error) {
	var r Register
	i := 0
	for n > 0 {
		r.data[lenPass-i-1] = n % 10
		n /= 10
		i++
	}
	if i != lenPass {
		return r, errors.New("not of the len expected")
	}
	return r, nil
}

func (r Register) Next() Register {
	for i := lenPass - 1; i >= 0; i-- {
		n := (r.data[i] + 1) % 10
		if n > r.data[i] {
			// ok I get the next value
			r.data[i] = n
			// fill all the other digits
			for j := i + 1; j < lenPass; j++ {
				r.data[j] = n
			}
			return r
		}
	}
	return r
}

func (r Register) IsValidPass() bool {
	max := -1
	prev := -1
	double := false

	for _, e := range r.data {
		if e == prev {
			double = true
		} else {
			prev = e
		}

		if max == -1 || e >= max {
			max = e
		} else {
			return false
		}
	}

	return double
}
func (r Register) ToDec() int {
	k := 1
	n := 0
	for i := lenPass - 1; i >= 0; i-- {
		n += r.data[i] * k
		k *= 10
	}

	return n
}

func (r Register) IsLess(t Register) bool {
	return r.ToDec() < t.ToDec()
}
