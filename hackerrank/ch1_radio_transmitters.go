/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Hackerland is a one-dimensional city with houses aligned at integral locations along a road.
				  The Mayor wants to install radio transmitters on the roofs of the city's houses.
				  Each transmitter has a fixed range meaning it can transmit a signal to all houses within
				  that number of units distance away.
				  Given a map of Hackerland and the transmission range,
				  determine the minimum number of transmitters so that every house is
				  within range of at least one transmitter.
				  Each transmitter must be installed on top of an existing house.

	Link		: https://www.hackerrank.com/challenges/hackerland-radio-transmitters/
*/

package hackerrank

import "fmt"

func sortList(nums []int32) []int32 {
	sorted_list := []int32{}
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

// TODO this is incorrect heuristic for large lists, need to optimize
func hackerlandRadioTransmitters(x []int32, k int32) int32 {
	x = sortList(x)
	min_idx := 0
	nodes := []int32{}
	for idx, val := range x {
		max_val := val + k
		for jdx := min_idx; jdx < len(x) && x[jdx] <= max_val; jdx++ {
			min_idx = jdx
		}

		if idx == 0 || nodes[len(nodes)-1]+k < x[min_idx] {
			nodes = append(nodes, x[min_idx])
		}
	}
	fmt.Printf("Transmitters at: %v\n", nodes)
	return int32(len(nodes))
}

func Run_3() {
	x := []int32{7, 2, 4, 6, 5, 9, 12, 11}
	var k int32 = 2
	fmt.Printf("Original array: %v, range: %d\n", x, k)
	pos := hackerlandRadioTransmitters(x, k)
	fmt.Printf("Min number of transmitters needed: %v\n\n", pos)
}
