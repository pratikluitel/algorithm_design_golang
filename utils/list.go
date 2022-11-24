package utils

// TODO dtype independent sorting
// sorts list, also returns an array with the indices from the old list
func SortList(nums []int) ([]int, []int) {
	sorted_list := []int{}
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
