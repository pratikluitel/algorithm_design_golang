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
)

//TODO not giving optimal solution, fix

func removeKdigits(num string, k int) string {
	final_len := len(num) - k

	if final_len <= 0 {
		return "0"
	}
	if k == 0 {
		return num
	}

	k_removed_num := string(num[0]) //initialization

	for idx := 1; idx < len(num); idx++ {
		val := num[idx]
		fmt.Printf("Stack: %s\n", k_removed_num)
		if len(k_removed_num) != 0 && k_removed_num[len(k_removed_num)-1] > val {
			fmt.Printf("curr: %s, %s > %s, Popped %s\n", string(val), string(k_removed_num[len(k_removed_num)-1]), string(val), string(k_removed_num[len(k_removed_num)-1]))
			k_removed_num = string(k_removed_num[:len(k_removed_num)-1])
		} else if len(k_removed_num) != final_len {
			fmt.Printf("curr: %s, %s < %s, Pushed %s\n", string(val), string(k_removed_num), string(val), string(val))
			k_removed_num += string(val)
		}
		fmt.Printf("Stack: %s\n\n", k_removed_num)
	}
	// if len(k_removed_num) != final_len {
	// 	k_removed_num = num[:final_len]
	// }
	// else {
	// 	to_remove := k - len(num) + len(k_removed_num)
	// 	k_removed_num = k_removed_num[to_remove:]
	// }

	// removing leading zeros
	converted_int, _ := strconv.Atoi(k_removed_num)

	return fmt.Sprint(converted_int)
}

func Run_4() {
	num := "2463235"
	k := 4
	str_removed := removeKdigits(num, k)
	start := time.Now()
	fmt.Printf("Start number: %s, %d strings removed: %s, execution time: %s\n", num, k, str_removed, time.Since(start))
}
