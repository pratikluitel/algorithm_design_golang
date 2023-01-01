/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 2

	Q			: We define a magic square to be an nxn matrix of distinct positive integers from 1 to n^2
				  where the sum of any row, column, or diagonal of length n is always equal to
				  the same number: the magic constant.

				  You will be given a 3x3 matrix of s integers in the inclusive range [1,9].
				  We can convert any digit a to any other digit b in the range [1,9] at cost of |a-b|.
				  Given s, convert it into a magic square at minimal cost. Print this cost on a new line.

	Link		: https://www.hackerrank.com/challenges/magic-square-forming/problem
*/

package hackerrank

import "math"

func formingMagicSquare(s [][]int32) int32 {
	costs := make([]int32, 8)

	options := [8][3][3]int32{ // slice of all possible magic square of 3x3, there is no way except brute force to do this, sorry
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	}

	for i, row := range s {
		for j, value := range row {
			for k := range costs {
				costs[k] += int32(math.Abs(float64(options[k][i][j] - value)))
			}
		}
	}

	var min int32 = 0

	for i, cost := range costs {
		if cost < min || i == 0 {
			min = cost
		}
	}

	return min
}
