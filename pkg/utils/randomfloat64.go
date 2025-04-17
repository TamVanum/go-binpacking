package utils

import "math/rand/v2"

func RandomRangeFloat64() (min float64, max float64) {
	a := rand.Float64() * 50
	b := rand.Float64() * 50

	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}

	return
}
