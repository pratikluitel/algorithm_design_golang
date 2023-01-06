package utils

import (
	"fmt"
	"math"
	"strconv"
)

// finds the alphabet wise minimum length 1 substring of str, and its position
// in the string
func StringMin(str string) (int, string) {
	minimum := math.Inf(1)
	min_idx := 0
	for idx, val := range str {
		digit, _ := strconv.Atoi(string(val))
		if float64(digit) < minimum {
			minimum = float64(digit)
			min_idx = idx
		}
	}
	return min_idx, fmt.Sprint(int(minimum))
}

func StringMax(str string) (int, string) {
	maximum := math.Inf(-1)
	min_idx := 0
	for idx, val := range str {
		digit, _ := strconv.Atoi(string(val))
		if float64(digit) > maximum {
			maximum = float64(digit)
			min_idx = idx
		}
	}
	return min_idx, fmt.Sprint(int(maximum))
}
