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
	"time"

	"github.com/pratikluitel/o_riley_algorithm_design_manual_excercises/utils"
)

// check if a particular combination is in a slice composed of combinations
func inCombinations(combinations [][]int, combination []int) bool {
	var matched bool
	for _, comb := range combinations {
		// to check equality of 2 slices, sort them and compare element by element
		matched = true
		for idx, val := range comb {
			if val != combination[idx] {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}
	return false
}

func twoSum(nums []int, target int) [][]int {
	combinations := [][]int{}
	lo, hi := 0, len(nums)-1

	for lo < hi {
		if nums[lo]+nums[hi] == target {
			combn, _ := utils.SortList([]int{nums[lo], nums[hi]})
			combinations = append(combinations, combn)
			lo++
			hi--
		} else if nums[lo]+nums[hi] > target || (hi != len(nums)-1 && nums[hi] == nums[hi+1]) {
			hi--
		} else if nums[lo]+nums[hi] < target || (lo != 0 && nums[lo] == nums[lo-1]) {
			lo++
		}
	}
	return combinations
}

// TODO: working, but slow. Need to optimize

// given a slice `nums` with `n` integers
// and a target integer `target`
//
// pick `k` unique elements from nums such that their sum is exactly `target`
//
// return all possible unique combinations of elements
func kSum(nums []int, target int, k int) [][]int {
	// check for certainty that the number cannot reach/ will overshoot the target
	if len(nums) == 0 || nums[0] > target/k || nums[len(nums)-1] < target/k {
		return [][]int{}
	}

	if k == 2 { // terminating case
		return twoSum(nums, target)
	}

	combinations := [][]int{}

	for idx, val := range nums {
		new_nums := nums[idx+1:]
		combs := kSum(new_nums, target-val, k-1)
		if len(combs) != 0 {
			for _, comb := range combs {
				combn, _ := utils.SortList(append(comb, val))
				in_combinations := inCombinations(combinations, combn)
				if !in_combinations {
					combinations = append(combinations, combn)
				}
			}
		}
	}
	return combinations
}

func Run_6() {
	ip := []int{0, 1, 3, -5, 3, 0}
	start := time.Now()
	input, _ := utils.SortList(ip)
	t := 1
	combs := kSum(input, t, 4)
	fmt.Printf("\ninput: %v, target: %d, combinations: %v, execution time: %s\n", input, t, combs, time.Since(start))
}
