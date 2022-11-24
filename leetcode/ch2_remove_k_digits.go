/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1
	Title		: Given string num representing a non-negative integer num, and an integer k,
				  return the smallest possible integer after removing k digits from num.
	Link		: https://leetcode.com/problems/remove-k-digits/
*/

package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func removeKdigits(num string, k int) string {
	new_length := len(num) - k
	if new_length <= 0 {
		return "0"
	}
	k_digits_removed := ""
	min_num := math.Inf(1)
	for idx := 0; idx < len(num) && len(k_digits_removed) != new_length; idx++ {
		digit, _ := strconv.Atoi(string(num[idx]))

		fmt.Printf("%d %f\n", digit, min_num)
		if float64(digit) <= min_num {
			k_digits_removed += fmt.Sprint(digit)
			min_num = float64(digit)
		}
	}

	// removing the leading zeros
	for i := 0; len(k_digits_removed) > 1 && string(k_digits_removed[i]) == "0"; i++ {
		k_digits_removed = k_digits_removed[i+1:]
	}

	return k_digits_removed
}

func Run_4() {
	num := "1432219"
	k := 3
	str_removed := removeKdigits(num, k)
	start := time.Now()
	fmt.Printf("Start number: %s, string removed %s, execution time: %s\n", num, str_removed, time.Since(start))
}
