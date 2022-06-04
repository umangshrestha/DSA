package taylor

import (
	"math"
)

func Sin(x float64) float64 {
	x = math.Mod(x, 2*math.Pi) //restrictng the value from -2*PI to 2*PI
	term := x
	sin := x

	for i := 3.0; term != 0.0; i += 2.0 {
		// Sin(x) = x - x^3/3! + x^5/5! - x^7/7! + ...
		// OR  pow(-1, n) * pow(x, 2n+1) / fac(2n+1)
		term *= (-x * x) / (i * (i - 1))
		sin += term
	}
	return sin
}
