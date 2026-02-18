package helpers

import "math/rand/v2"

func RandomInt(min, max int) int {
	return rand.IntN(max-min) + min
}

func RandomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
