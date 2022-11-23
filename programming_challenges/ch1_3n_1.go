/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1
	Title		: The 3n+1 problem
	Link		: https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=29&page=show_problem&problem=36
*/

package programming_challenges

import (
	"fmt"
	"math"
	"time"
)

func algorithm(n int) int { // returns the cycle length of
	cycle := 1 // even for n=1, the cycle length is 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		cycle++
	}
	return cycle
}

func max_cycle(i int, j int) {
	start := time.Now()
	max := math.Inf(-1)

	for idx := i; idx <= j; idx++ {
		cycle := algorithm(idx)
		if float64(cycle) > max {
			max = float64(cycle)
		}
	}
	fmt.Printf("%d %d %d, execution time: %s\n", i, j, int(max), time.Since(start))
}

func Run_1() {
	max_cycle(1, 1000)
}
