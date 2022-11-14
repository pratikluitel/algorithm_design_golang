/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Given an array of integers temperatures represents the daily temperatures,
				  return an array answer such that answer[i] is the number of days you have to wait
				  after the ith day to get a warmer temperature.
				  If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/

package leetcode

import (
	"fmt"
	"time"
)

func dailyTemperatures(temperatures []int) []int {

	wait_days := []int{}

	for day := 0; day < len(temperatures); day++ {
		wait_days = append(wait_days, 0) // the default case, where there is no future day for which this is possible
		if day == len(temperatures)-1 {  // since you can't wait for the last day
			break
		}
		for idx := day + 1; idx < len(temperatures); idx++ { // day+1 won't produce an error because the last index won't reach this block
			if temperatures[idx] > temperatures[day] {
				wait_days[day] = idx - day
				break
			}
		}
	}

	return wait_days
}

// ref: https://leetcode.com/problems/daily-temperatures/solutions/1574808/c-python-3-simple-solutions-w-explanation-examples-images-2-monotonic-stack-approaches/
func dailyTemperatures_stack(temperatures []int) []int {
	indices_stack := []int{}
	wait_days := []int{}
	stack_start := 0
	for day := 1; day < len(temperatures); day++ {
		wait_days = append(wait_days, 0) // the default case, where there is no future day for which this is possible
		if temperatures[day-1] > temperatures[day] {
			indices_stack = append(indices_stack, day-1)
		} else {
			for idx := stack_start; idx < len(indices_stack); idx++ {
				diff := day - 1 - indices_stack[idx]
				wait_days = append(wait_days, diff)
			}
			stack_start = len(indices_stack) - 1
		}
	}
	return wait_days
}

func Run() {
	start1 := time.Now()
	temperatures := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160} // the worst case scenario for a brute force solution
	wait_days := dailyTemperatures(temperatures)
	fmt.Printf("Brute Force approach -> wait days array = %v, execution time: %s\n\n", wait_days, time.Since(start1)) // this is approx n^2 time
	start2 := time.Now()
	wait_days2 := dailyTemperatures(temperatures)
	fmt.Printf("Stack approach -> wait days array = %v, execution time: %s\n\n", wait_days2, time.Since(start2)) // this is approx n^2 time

}
