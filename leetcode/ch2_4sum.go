/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 2

	Q			: Given an array nums of n integers, return an array of all the unique quadruplets
				  [nums[a], nums[b], nums[c], nums[d]] such that:
     			  0 <= a, b, c, d < n
     			  a, b, c, and d are distinct.
     			  nums[a] + nums[b] + nums[c] + nums[d] == target

	Link		: https://leetcode.com/problems/4sum/
*/

package leetcode

// func kSum(nums []int, target int, k int) ([][]int, bool) {
// 	ksums := [][]int{}
// 	for idx, val := range nums{

// 		ksum_prev, success := kSum(nums, target-val, k-1)
// 		if success {
// 			ksum_prev := append(ksum_prev, val)
// 			ksums := append(ksums,ksum_prev)
// 		}
// 	}
// }

// func fourSum(nums []int, target int) [][]int {
// 	ksums, _ := kSum(nums, target, 4)
// 	return ksums
// }

// func Run_6() {

// }
