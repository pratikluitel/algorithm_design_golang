/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Given an array of integers temperatures represents the daily temperatures,
				  return an array answer such that answer[i] is the number of days you have to wait
				  after the ith day to get a warmer temperature.
				  If there is no future day for which this is possible, keep answer[i] == 0 instead.

	Link		: https://leetcode.com/problems/daily-temperatures/
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

// custom integer stack implementation
type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// TODO see other approaches
// ref: https://leetcode.com/problems/daily-temperatures/solutions/1574808/c-python-3-simple-solutions-w-explanation-examples-images-2-monotonic-stack-approaches/
func dailyTemperatures_stack(temperatures []int) []int {
	indices_stack := stack{}
	wait_days := []int{}
	for day := 0; day < len(temperatures); day++ {
		wait_days = append(wait_days, 0) // the default case, where there is no future day for which this is possible
		for len(indices_stack) != 0 && temperatures[day] > temperatures[indices_stack[len(indices_stack)-1]] {
			diff := day - indices_stack[len(indices_stack)-1]
			wait_days[indices_stack[len(indices_stack)-1]] = diff
			indices_stack.Pop()
		}
		indices_stack.Push(day)
	}
	return wait_days
}

func Run_1() {
	start1 := time.Now()
	temperatures := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 50, 60, 70, 80, 90, 100} // the worst case scenario for a brute force solution
	wait_days := dailyTemperatures(temperatures)
	fmt.Printf("Brute Force approach -> wait days array = %v, execution time: %s\n\n", wait_days, time.Since(start1)) // this is approx O(n^2) time
	start2 := time.Now()
	wait_days2 := dailyTemperatures_stack(temperatures)
	fmt.Printf("Stack approach -> wait days array = %v, execution time: %s\n\n", wait_days2, time.Since(start2)) // this is approx O(n) time

}
