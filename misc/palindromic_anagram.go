// Given a string, determine if it can be rearranged into a palindrome. Return the string YES or NO.
package misc

func palindromic_anagram(s string) bool {
	letter_map := map[rune]int{}
	len_s := len(s)
	for _, let := range s {
		_, exists := letter_map[let]
		if !exists {
			letter_map[let] = 1
		} else {
			letter_map[let] += 1
		}
	}
	if len_s%2 == 0 { // if the string has even number of letters, for it to be a palindrome, it must have even number of all letters
		for _, val := range letter_map {
			if val%2 != 0 {
				return false
			}
		}
	} else { // else, all but one even
		odd := 0
		for _, val := range letter_map {
			if val%2 != 0 && odd > 1 {
				return false
			} else if val%2 != 0 {
				odd += 1
			}
		}
	}
	return true
}
