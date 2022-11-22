/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Given an integer d, rotate an array arr of size n that many steps left and return the result.

	Link		: https://www.hackerrank.com/challenges/array-left-rotation/problem
	Bonus		: https://leetcode.com/problems/rotate-array/
*/

package hackerrank

import (
	"fmt"
	"time"
)

func rotateLeft(d int, arr []int) []int {
	effective_rotations := d % len(arr) // since rotating len(arr) times leads to the same arr
	if effective_rotations == 0 {
		return arr
	}

	arr = append(arr[effective_rotations:], arr[:effective_rotations]...) // inplace for O(1) space
	return arr
}

func rotateRight(d int, arr []int) []int {
	effective_rotations := d % len(arr) // since rotating len(arr) times leads to the same arr
	if effective_rotations == 0 {
		return arr
	}
	new_first_idx := len(arr) - effective_rotations

	arr = append(arr[new_first_idx:], arr[:new_first_idx]...) // inplace for O(1) space
	return arr
}

// an alternate way by consecutive flipping, in practice very slightly slower than the above method
func rotate_alternate(d int, arr []int, direction string) []int { //O(n) time and O(1) space
	effective_rotations := d % len(arr) // since rotating len(arr) times leads to the same arr
	if effective_rotations == 0 {
		return arr
	}
	if direction == "right" {
		effective_rotations = len(arr) - effective_rotations
	}
	for idx, jdx := 0, effective_rotations-1; idx < jdx; idx, jdx = idx+1, jdx-1 { // flipping the first effective_rotations elements of the slice about the mid point
		temp := arr[idx]
		arr[idx] = arr[jdx]
		arr[jdx] = temp
	}
	for idx, jdx := effective_rotations, len(arr)-1; idx < jdx; idx, jdx = idx+1, jdx-1 { // flipping the last effective_rotations elements of the slice about the mid point
		temp := arr[idx]
		arr[idx] = arr[jdx]
		arr[jdx] = temp
	}
	for idx, jdx := 0, len(arr)-1; idx < jdx; idx, jdx = idx+1, jdx-1 { // flipping the new array
		temp := arr[idx]
		arr[idx] = arr[jdx]
		arr[jdx] = temp
	}
	return arr
}

func Run_1() {
	arr := []int{1, 2, 3, 4, 5, 17, 3498, 239, 32324, 324, 32, 32, 31, 32132, 31, 0, 2, 123, 2}
	arr_c := make([]int, len(arr))
	copy(arr_c, arr)
	d := 10
	start1 := time.Now()
	new_arr := rotateRight(d, arr)
	fmt.Printf("initial: %v, rotated %v, execution time: %s\n", arr, new_arr, time.Since(start1))
	start2 := time.Now()
	new_arr_alt := rotate_alternate(d, arr_c, "right")
	fmt.Printf("initial: %v, rotated %v, execution time: %s\n", arr, new_arr_alt, time.Since(start2))
}
