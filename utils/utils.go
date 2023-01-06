package utils

import (
	"fmt"
)

type Sortable interface {
	int | int32 | int64 | float32 | float64
}

// sorts list, also returns an array with the indices from the old list.
//
// if you pass a string to this, make sure that it is a slice of runes, and not of the string type.
func Sort[S Sortable](nums []S) ([]S, []int) {
	sorted_list := []S{}
	sorted_indices := []int{}
	for jdx, num := range nums {
		sorted_list = append(sorted_list, num)
		sorted_indices = append(sorted_indices, jdx)
		for idx := len(sorted_list) - 1; idx > 0; idx-- {
			if sorted_list[idx] < sorted_list[idx-1] {
				temp_num := sorted_list[idx]
				temp_idx := sorted_indices[idx]
				sorted_list[idx] = sorted_list[idx-1]
				sorted_indices[idx] = sorted_indices[idx-1]
				sorted_list[idx-1] = temp_num
				sorted_indices[idx-1] = temp_idx
			}
		}
	}
	return sorted_list, sorted_indices
}

func partition(str string, pivots string) (string, string, string) {
	low_partition := ""
	high_partition := ""
	if len(str) == 1 {
		// for a string of length 1, the only character is the pivot, for empty string, answer is the empty string
		return low_partition, high_partition, pivots + str
	}
	if len(str) == 0 {
		return low_partition, high_partition, pivots
	}
	pivot := string(str[len(str)-1])
	for idx, val := range str {
		if string(val) <= string(pivot) && idx != len(str)-1 {
			low_partition += string(val)
		} else if string(val) > string(pivot) && idx != len(str)-1 {
			high_partition += string(val)
		}
	}

	fmt.Printf("(%s, (%s, %s))\n", pivot, low_partition, high_partition)
	if len(high_partition) == 1 {
		pivots = pivot + high_partition
	} else if len(low_partition) == 1 {
		pivots = low_partition + pivot
	}

	// pivots += pivot
	_, _, pivots_low := partition(low_partition, pivot)
	fmt.Printf("\t")
	_, _, pivots_high := partition(high_partition, pivot)
	fmt.Printf("\n")
	fmt.Printf("%s %s\n", pivots_low, pivots_high)

	return low_partition, high_partition, pivots_low + pivots_high // the final sorted array is low_pivots + high_pivots

}

// TODO implement properly and make type agnostic
func QuickSort(str string) string {
	_, _, sorted_str := partition(str, "") // no pivots initially
	return sorted_str
}

// returns true if an element is in a slice.
// also returns the index the element is in. -1 if not present
func IsIn[S Sortable](slice []S, to_find S) (bool, int) {
	for idx, elem := range slice {
		if elem == to_find {
			return true, idx
		}
	}
	return false, -1
}
