package taylor

import (
	"fmt"
	"math"
)

// TODO: NOT WORKING
func Tan(x float64) float64 {
	x = math.Mod(x, 2*math.Pi) //restrictng the value from -2*PI to 2*PI
	/* setting the values the n=1 */
	// term := x / 12.0
	term := x / 2.0
	t4 := 4.0 //4 power term
	tan := 0.0

	i := 1
	for term != 0.0 {
		// Tan(x) = x + x^3/3 + 2^5/15 + 17x^7/315 + ...
		// OR pow(4, n) * (pow(4, n) -1) * pow(x, 2n -1) / fac(2n)
		tan += t4 * (t4 - 1) * term
		i++
		t4 *= 4
		term *= (x * x) / float64(2*i*(2*i-1))

	}
	fmt.Println(Sin(x) / Cos(x))
	return tan
}
