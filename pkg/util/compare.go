package util

import "golang.org/x/exp/constraints"

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
