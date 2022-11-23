/*
The power of an integer x is defined as the number of steps needed to transform x into 1 using the following steps:

    if x is even then x = x / 2
    if x is odd then x = 3 * x + 1

Given three integers lo, hi and k. The task is to sort all integers in the interval [lo, hi] by the power value in ascending order,
if two or more integers have the same power value sort them by ascending order.

Return the kth integer in the range [lo, hi] sorted by the power value
*/

package misc

import "github.com/pratikluitel/o_riley_algorithm_design_manual_excercises/utils"

func power(n int) int { // returns the cycle length of
	power := 1 // even for n=1, the cycle length is 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		power++
	}
	return power
}

func getKth(lo int, hi int, k int) int {
	power_arr := []int{}
	nums := []int{}
	for num := lo; num <= hi; num++ {
		power_arr = append(power_arr, power(num))
		nums = append(nums, num)
	}
	_, indices_sorted := utils.SortList(power_arr)
	nums_cpy := make([]int, len(nums))
	copy(nums_cpy, nums)
	for idx, val := range indices_sorted {
		nums[idx] = nums_cpy[val]
	}
	return nums[k-1]
}
