package taylor

import (
	"math"
)

func Cos(x float64) float64 {
	term := 1.0
	x = math.Mod(x, 2*math.Pi) //restrictng the value from -2*PI to 2*PI
	out := 1.0

	for i := 1; term != 0.0; i++ {
		// Cos(x) = 1 - x^2/2! + x^4/4! - x^6/6! + ...
		// OR pow(-1, n) * pow(x, 2n) / fac(2n)
		term *= x / float64(i)
		if i%4 == 0 {
			out += term
		} else if i%2 == 0 {
			out -= term
		}
	}
	return out
}
