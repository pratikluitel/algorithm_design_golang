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
	"strconv"
	"time"

	"github.com/pratikluitel/o_riley_algorithm_design_manual_excercises/utils"
)

// TODO working for 28/43 need to apply stack for an optimal solution
func pre_zero_removal(num string, k int) string {
	new_length := len(num) - k
	fmt.Printf("Removing %d digits from %s\n", k, num)
	if k == 0 {
		return num
	}
	if new_length <= 0 {
		return "0"
	}
	possible_options := num[:k+1]
	max_idx, maximum_num := utils.StringMax(possible_options)
	fmt.Printf("Checked in %s, minimum number: %s\n", possible_options, maximum_num)

	new_num := num[:max_idx] + num[max_idx+1:] // remove the element

	return pre_zero_removal(new_num, k-1)
}

func removeKdigits(num string, k int) string {
	k_removed_num := pre_zero_removal(num, k)
	converted_int, _ := strconv.Atoi(k_removed_num)

	return fmt.Sprint(converted_int)
}

func Run_4() {
	num := "112"
	k := 1
	str_removed := removeKdigits(num, k)
	start := time.Now()
	fmt.Printf("Start number: %s, %d strings removed: %s, execution time: %s\n", num, k, str_removed, time.Since(start))
}
