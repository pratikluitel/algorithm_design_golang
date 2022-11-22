/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: You are choreographing a circus show with various animals.
				  For one act, you are given two kangaroos on a number line ready to
				  jump in the positive direction (i.e, toward positive infinity).
    			  The first kangaroo starts at location x1 and moves at a rate of
				  v1 meters per jump. The second kangaroo starts at location x2 and moves at a rate of
				  v2 meters per jump.
				  You have to figure out a way to get both kangaroos at the same location at the same time
				  as part of the show. If it is possible, return YES, otherwise return NO.

	Link		: https://www.hackerrank.com/challenges/kangaroo/problem
*/

package hackerrank

import (
	"fmt"
	"math"
	"time"
)

func kangaroo(x1 int, v1 int, x2 int, v2 int) string {
	prev_diff := math.Inf(1)
	for pos1, pos2 := x1, x2; ; pos1, pos2 = pos1+v1, pos2+v2 {
		diff := math.Abs(float64(pos1 - pos2))
		if diff == 0.0 {
			return "YES"
		} else if diff >= prev_diff { // if the difference starts increasing or keeps staying the same, the second kangaroo is already past the point of catching up
			return "NO"
		}
		prev_diff = diff
	}
}

func Run_2() {
	start := time.Now()
	can := kangaroo(0, 3, 4, 2)
	fmt.Printf("Can the two kangaroos meet? ->%s, execution time: %s\n\n", can, time.Since(start))
}
