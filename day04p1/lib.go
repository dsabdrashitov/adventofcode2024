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

func check[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
