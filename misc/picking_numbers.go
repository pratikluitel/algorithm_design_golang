// Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to
// Link: https://www.hackerrank.com/challenges/picking-numbers/problem

package misc

func pickingNumbers(a []int32) int32 {
	counter_map := map[int32]int32{} //maps a number to its count in the a slice

	for _, val := range a {
		if _, ok := counter_map[val]; ok { // if the val exists
			counter_map[val]++
		} else {
			counter_map[val] = 1
		}
	}

	var max int32 = 0
	for _, value := range a {
		count := counter_map[value]

		count_low, ok1 := counter_map[value-1]
		count_high, ok2 := counter_map[value+1]
		if !ok1 {
			count_low = 0
		}
		if !ok2 {
			count_high = 0
		}

		var length int32
		if count+count_low > count+count_high {
			length = count + count_low
		} else {
			length = count + count_high
		}

		if length > max {
			max = length
		}
	}

	return max
}
