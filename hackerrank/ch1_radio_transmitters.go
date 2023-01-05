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

import (
	"fmt"

	"github.com/pratikluitel/o_riley_algorithm_design_manual_excercises/utils"
)

// TODO: doesn't work for large inputs, need to review
func hackerlandRadioTransmitters(x []int, k int) int {
	if x == nil {
		return 0
	}
	x, _ = utils.Sort(x)
	fmt.Printf("Sorted array: %v\n", x)
	nodes := []int{}

	last_node_location := x[0]
	for idx := 0; idx < len(x); idx++ {
		distance := x[idx] - last_node_location
		if len(nodes) == 0 && distance > k {
			nodes = append(nodes, x[idx-1])
			last_node_location = x[idx]
		} else if len(nodes) != 0 && distance > k {
			nodes = append(nodes, x[idx])
			last_node_location = x[idx]
		}
		fmt.Printf("At node: %d, last node: %d, Transmitters placed: %v\n", x[idx], last_node_location, nodes)
	}

	fmt.Printf("Transmitters at: %v\n", nodes)
	return len(nodes)
}

func Run_3() {
	x := []int{7, 2, 4, 6, 5, 9, 11, 12}
	k := 2
	fmt.Printf("Original array: %v, range: %d\n", x, k)
	pos := hackerlandRadioTransmitters(x, k)
	fmt.Printf("Min number of transmitters needed: %v\n\n", pos)
}
