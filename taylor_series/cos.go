package taylor

import (
	"math"
)

func Cos(x float64) float64 {
	x = math.Mod(x, 2*math.Pi) //restrictng the value from -2*PI to 2*PI
	term := 1.0
	cos := 1.0

	for i := 2; term != 0.0; i += 2 {
		// Cos(x) = 1 - x^2/2! + x^4/4! - x^6/6! + ...
		// OR pow(-1, n) * pow(x, 2n) / fac(2n)
		term *= (-x * x) / float64(i*(i-1))
		cos += term
	}
	return cos
}
