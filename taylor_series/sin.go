package taylor

import (
	"math"
)

func Sin(x float64) float64 {
	term := 1.0                // first term
	x = math.Mod(x, 2*math.Pi) //restrictng the value from -2*PI to 2*PI
	sin := 0.0

	for i := 1; term != 0.0; i++ {
		// Sin(x) = x - x^3/3! + x^5/5! - x^7/7! + ...
		// OR  pow(-1, n) * pow(x, 2n+1) / fac(2n+1)
		term *= x / float64(i)
		if i%4 == 3 { //divisible by 3
			sin -= term
		} else if i%4 == 1 { // divisible by 5
			sin += term
		}
	}
	return sin
}
