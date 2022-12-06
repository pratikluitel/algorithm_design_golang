// An unsorted array of size n contains distinct integers between 1 and n + 1,
// with one element missing. Give an O(n) algorithm to ﬁnd the missing integer,
// without using any extra space.

// Soln :
// Find the sum of first n natural numbers
// subtract the sum of the array from this sum
// Profit

package misc

import (
	"fmt"
	"time"
)

func find_missing(arr []int) int {
	ideal_sum := (len(arr) + 1) * (len(arr) + 2) / 2
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return ideal_sum - sum
}

func Run_1() {
	arr := []int{}
	for i := 1; i <= 9; i++ {
		if i != 5 {
			arr = append(arr, i)
		}
	}
	missing := find_missing(arr)
	start := time.Now()
	fmt.Printf("\nOriginal Array: %v, Missing value: %d execution time: %s\n", arr, missing, time.Since(start))
}
