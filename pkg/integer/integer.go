package integer

import (
	"strconv"

	. "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"golang.org/x/exp/constraints"
)

func Int(s string) int {
	return Must(strconv.Atoi(s))
}

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
