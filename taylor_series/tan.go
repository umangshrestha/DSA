package taylor

import (
	"math"
)

func Tan(x float64) float64 {
	term := 1.0
	t4 := 1.0 //4 power term
	x = math.Mod(x, 2*math.Pi)
	tan := 0.0 // first term
	for i := 1; term != 0.0 && math.Inf(-1) < tan && tan < math.Inf(1); i++ {
		// Tan(x) = x - x^3/3 + 2^5/15 - 17x^7/315 + ...
		// OR pow(-4, n) * (1 - pow(-4, n)) * pow(x, 2n -1) / fac(2n)
		t4 *= -4.0
		term *= x / float64(i)
		if i%2 == 1 { // all odd numbers
			tan += t4 * (1 - t4) * term
		}
	}
	return tan
}
