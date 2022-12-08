/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 2

	Q			: Given an integer n, return an array ans of length n + 1 such that
				  for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

	Link		: https://leetcode.com/problems/counting-bits/
*/

package leetcode

import (
	"fmt"
	"time"
)

func countBits(n int) []int {
	ans := []int{}
	for i := 0; i <= n; i++ {
		bin_rep_ones := 0
		num := i //tracks the number through repeated divisons to convert to binary
		for num > 0 {
			bin_rep_ones += num % 2
			num /= 2
		}
		ans = append(ans, bin_rep_ones)
	}
	return ans
}

func Run_5() {
	num := 5
	start := time.Now()
	bits := countBits(num)
	fmt.Printf("\nNumber: %d, bits counts from 0 to %d: %v %s\n", num, num, bits, time.Since(start))
}
