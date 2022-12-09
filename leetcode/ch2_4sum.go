/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 2

	Q			: Given an array nums of n integers, return an array of all the unique quadruplets
				  [nums[a], nums[b], nums[c], nums[d]] such that:
     			  0 <= a, b, c, d < n
     			  a, b, c, and d are distinct.
     			  nums[a] + nums[b] + nums[c] + nums[d] == target

	Bonus		: Make this generalize for kSum

	Link		: https://leetcode.com/problems/4sum/
*/

package leetcode

import (
	"fmt"

	"github.com/pratikluitel/o_riley_algorithm_design_manual_excercises/utils"
)

// check if a particular combination is in a slice composed of combinations
func inCombinations(combinations [][]int, combination []int) bool {
	for _, comb := range combinations {
		matches := 0
		for _, term1 := range comb {
			for _, term2 := range combination {
				if term1 == term2 {
					matches += 1
					break
				}
			}
		}
		if matches == len(combination) {
			return true
		}
	}
	return false
}

// TODO: this is working for 240/290 cases. It is slow. Need to fix

// given a slice `nums` with `n` integers
// and a target integer `target`
//
// pick `k` unique elements from nums such that their sum is exactly `target`
//
// return all possible unique combinations of elements
func kSum(nums []int, target int, k int) [][]int {
	nums, _ = utils.SortList(nums)
	combinations := [][]int{}

	if k == 2 { // terminating case
		for jdx, val1 := range nums {
			for idx := jdx + 1; idx < len(nums); idx++ {
				val2 := nums[idx]
				in_combinations := inCombinations(combinations, []int{val1, val2})
				// fmt.Printf("Combinations: %v, %v in combinations? %t\n", combinations, []int{val1, val2}, in_combinations)
				if val1+val2 == target && !in_combinations {
					combinations = append(combinations, []int{val1, val2})
				}
			}
		}
		fmt.Printf("\nTarget: %d, Combinations: %v\n", target, combinations)
		return combinations
	}

	for idx, val := range nums {
		nums_temp := make([]int, len(nums))
		copy(nums_temp, nums)
		new_nums := append(nums_temp[:idx], nums_temp[idx+1:]...)
		fmt.Printf("\nNew nums: %v\n", new_nums)
		combs := kSum(new_nums, target-val, k-1)
		if len(combs) != 0 {
			for _, comb := range combs {
				combn := append(comb, val)
				in_combinations := inCombinations(combinations, combn)
				if !in_combinations {
					combinations = append(combinations, combn)
					fmt.Printf("\nTarget: %d, Combinations: %v\n", target, combinations)
				}
			}
		}
	}
	return combinations
}

func Run_6() {
	combs := kSum([]int{2, -4, -5, -2, -3, -5, 0, 4, -2}, -14, 4)
	fmt.Printf("%v\n", combs)
}
