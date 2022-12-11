/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 2
	Q			: Find the number of ways that a given integer, , can be expressed as the sum of the
				  powers of unique, natural numbers.
 				  For example, if and x=13 and n=2, we have to find all combinations of unique squares adding up to 13.
				  The only solution is 2^2 + 3^2.

	Link		: https://www.hackerrank.com/challenges/the-power-sum/problem

*/

package hackerrank

import (
	"fmt"
	"math"
	"time"
)

func powers(X int, N int, terms int) int32 {
	term := math.Pow(float64(terms), float64(N))
	if int(term) == X {
		return 1
	} else if int(term) > X {
		return 0
	} else {
		return powers(X, N, terms+1) + powers(X-int(term), N, terms+1)
	}
}

func powerSum(X int32, N int32) int32 {
	// finding max number of terms possible : when 1**N+2**N+3**N+...k**N >=X -> k=max terms
	return powers(int(X), int(N), 1)
}

func Run_5() {
	X := int32(12)
	N := int32(3)
	start := time.Now()
	powersum := powerSum(X, N)
	fmt.Printf("\nX - %d, N - %d, powersum - %d execution time: %s\n", X, N, powersum, time.Since(start))
}
