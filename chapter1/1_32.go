/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Write a function to perform integer division without using either the / or *
				  operators. Find a fast way to do it.

	Note		: In integer division, the remainder part is discarded, the result is an integer
				  quotient, but we'll find a remainder anyway
*/

package chapter1

import (
	"fmt"
	"math"
	"time"
)

func divide_recursive(numerator int, denominator int) (int, int) {
	sign := 1
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		sign = -1
	}
	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))
	if denominator > numerator {
		if numerator > 0 && denominator < 0 {
			numerator = -numerator
		}
		return 0, numerator
	} else {
		var quotient int
		temp_quotient, remainder := divide_recursive((numerator - denominator), denominator)
		if sign == -1 {
			quotient = -(1 + temp_quotient)
		}
		return quotient, remainder //basically, on simplifying: 1+(n-d)/d = n/d
	}
}

func divide_non_recursive(numerator int, denominator int) (int, int) {
	sign := 1
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		sign = -1
	}
	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))

	quotient := 0
	for numerator > denominator {
		numerator -= denominator
		quotient += 1
	}
	if sign == -1 {
		quotient = -quotient
	}
	if numerator > 0 && denominator < 0 {
		numerator = -numerator
	}

	return quotient, numerator
}

func Run_32() {
	for i := -5; i < 6; i++ {
		for j := -5; j < 6; j++ {
			if j == 0 {
				break
			}
			start1 := time.Now()
			ans, rem := divide_recursive(i, j)
			fmt.Printf("Recursive method -> %d / %d => quotient : %d remainder : %d, executed in %s\n", i, j, ans, rem, time.Since(start1))
			start2 := time.Now()
			ans2, rem2 := divide_non_recursive(i, j)
			fmt.Printf("Non recursive method -> %d / %d => quotient : %d remainder : %d, executed in %s\n", i, j, ans2, rem2, time.Since(start2))
		}
	}
}
