/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
				  You may assume the input array always has a valid answer.

	Link		: https://leetcode.com/problems/wiggle-sort-ii/
*/

package leetcode

import (
	"fmt"
	"time"
)

func sortList(nums []int) []int {
	sorted_list := []int{}
	for _, num := range nums {
		sorted_list = append(sorted_list, num)
		for idx := len(sorted_list) - 1; idx > 0; idx-- {
			if sorted_list[idx] < sorted_list[idx-1] {
				temp_num := sorted_list[idx]
				sorted_list[idx] = sorted_list[idx-1]
				sorted_list[idx-1] = temp_num
			}
		}
	}
	return sorted_list
}

// wiggles the slice:
//
// first sorts it, then inserts elements in sorted order
//
// inserts every even index first (0,2,4,...), -> the smallest n/2 numbers in even indices
//
// once the slive length runs out, inserts them in every odd index (1,3,5,...) -> the largest n/2 numbers in odd indices
func wiggleSort_brute(nums []int) []int {
	sorted_list := sortList(nums)
	list_len := len(sorted_list)
	mid_point := int((list_len + 1) / 2) // mid point is n/2+1 for odd length slice and n/2 for even length slice

	for idx, jdx := 0, 0; jdx < mid_point; idx, jdx = idx+2, jdx+1 {
		nums[idx] = sorted_list[jdx]
	}
	for idx, jdx := 1, mid_point; jdx < list_len; idx, jdx = idx+2, jdx+1 {
		nums[idx] = sorted_list[jdx]
	}

	recur := true
	// for adjacent elements equal cases, swap and place the elements as far away as possible from each other
	// until there are no adjacent elements in the slice
	for recur {
		recur = false
		for idx := 1; idx < list_len; idx++ {
			if nums[idx] == nums[idx-1] && idx+1 < list_len {
				recur = true
				temp := nums[idx]
				nums[idx] = nums[idx-2]
				nums[idx-2] = temp
				temp = nums[idx-1]
				nums[idx-1] = nums[idx+1]
				nums[idx+1] = temp
			}
		}
	}
	return nums
}

// Slightly faster than brute force. Just a small optimization with ref:
//
// https://leetcode.com/problems/wiggle-sort-ii/solutions/77684/summary-of-the-various-solutions-to-wiggle-sort-for-your-reference/?orderBy=most_votes
//
// In the brute force algorithm, we place the sorted elements serially, i.e. smallest -> median first (even indices), then median -> largest (odd indices)
// Equal values can occur near the median, which can occur in adjacent indices (since we are placing everything serially)
// so repeated elements may appear as adjacent in brute force, to prevent this we have run an infinite loop in the above function.
//
// i.e. final list [0] = smallest, final list [2] = smallest greater than smallest, and so on for even indices
//
// and final list [1] = median, final list [3] = smallest greater than median. and so on for odd indices
//
// But here, we start placing the sorted elements from the mid point (median) -> smallest first (even indices), then largest -> median (odd indices)
//
// i.e. final list [0] = median, final list[2] = largest smaller than median and so on for even
//
// and final list [1] = largest, final list [3] = largest smaller than largest, and so on for odd indices
//
// similar values have a chance of occuring near the median, as the medians are at either extreme left (0 index) or extreme right (end of array) they don't occur adjacently
func wiggleSort_nlogn(nums []int) []int {
	sorted_list := sortList(nums)
	list_len := len(sorted_list)
	mid_point := int((list_len + 1) / 2) // mid point is n/2+1 for odd length slice and n/2 for even length slice

	for idx, jdx := 0, mid_point-1; jdx >= 0; idx, jdx = idx+2, jdx-1 {
		nums[idx] = sorted_list[jdx]
	}
	for idx, jdx := 1, list_len-1; jdx >= mid_point; idx, jdx = idx+2, jdx-1 {
		nums[idx] = sorted_list[jdx]
	}

	return nums
}

func Run_3() {
	nums := []int{4, 5, 5, 6}
	in1 := make([]int, len(nums))
	in2 := make([]int, len(nums))
	copy(in1, nums)
	copy(in2, nums)

	start1 := time.Now()
	wiggled_nums_brute := wiggleSort_brute(nums)

	fmt.Printf("Initial array: %v, wiggle sorted, bruteforce: %v, execution time: %s\n", in1, wiggled_nums_brute, time.Since(start1))

	start2 := time.Now()
	wiggled_nums_nlogn := wiggleSort_nlogn(in1)
	fmt.Printf("Initial array: %v, wiggle sorted, nlogn: %v, execution time: %s\n\n", in2, wiggled_nums_nlogn, time.Since(start2))
}
