package utils

import "fmt"

// TODO dtype independent sorting, combine list and string to one package

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

// TODO implement properly
func QuickSort(str string) string {
	_, _, sorted_str := partition(str, "") // no pivots initially
	return sorted_str
}

func SortString(str string) string {
	sorted_string := ""
	for _, ch := range str {
		sorted_string += string(ch)
		for idx := len(sorted_string) - 1; idx > 0; idx-- {
			if sorted_string[idx] < sorted_string[idx-1] {
				temp_num := sorted_string[idx]
				sorted_string = sorted_string[:idx] + string(sorted_string[idx-1]) + sorted_string[idx+1:]
				sorted_string = sorted_string[:idx-1] + string(temp_num) + sorted_string[idx:]
			}
		}
	}
	return sorted_string
}
