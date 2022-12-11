/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 2
	Q			: A pangram is a string that contains every letter of the alphabet.
				  Given a sentence determine whether it is a pangram in the English alphabet.
				  Ignore case. Return either pangram or not pangram as appropriate.

	Link		: https://www.hackerrank.com/challenges/pangrams/problem
*/

package hackerrank

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func inSlice(slice []rune, char rune) bool {
	for _, ch := range slice {
		if ch == char {
			return true
		}
	}
	return false
}

// returns 'pangram' if the input s is a pangram
// else returns 'not pangram'
func pangrams(s string) string {
	s_upper := strings.ToUpper(s)
	cont_chars := []rune{}
	if len(s) < 26 {
		return "not pangram"
	}
	for idx, char := range s_upper {
		if unicode.IsLetter(char) && !inSlice(cont_chars, char) {
			cont_chars = append(cont_chars, char)
		}
		if len(cont_chars)+len(s[idx:]) < 26 { // the rest of the string cannot have all remaining letters
			return "not pangram"
		}
	}
	if len(cont_chars) == 26 {
		return "pangram"
	} else {
		return "not pangram"
	}
}

func Run_4() {
	str := "The quick brown fox jumps over the lay dog"
	start := time.Now()
	pan := pangrams(str)
	fmt.Printf("\npangram? - %s, execution time: %s\n", pan, time.Since(start))
}
