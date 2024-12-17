package main

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Gcd[T constraints.Integer](a T, b T) T {
	for a > 0 {
		tmp := b % a
		b = a
		a = tmp
	}
	return b
}
