/*
	Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three.
	Otherwise, return false.

	An integer y is a power of three if there exists an integer x such that y == 3x.
*/

package misc

import (
	"math"
)

// TODO This is a quick recursive divide and conquer. Slow, better surely exists
func powers(n int, terms int) int {
	term := math.Pow(float64(3), float64(terms))
	if int(term) == n {
		return 1
	} else if int(term) > n {
		return 0
	} else {
		return powers(n, terms+1) + powers(n-int(term), terms+1)
	}
}

func checkPowersOfThree(n int) bool {
	total_terms := powers(n, 0)
	return total_terms != 0
}
