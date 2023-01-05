/*
Code Author	: github.com/pratikluitel
Book		: The Algorithm Design Manual, written by Steven S Skiena
Chapter 	: 1
Title		: The Trip
Problem		: Find the minimum amount of total exchange that has to occur between N people with N different expenses
			  such that their expenses become equal.
Link		: https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=29&page=show_problem&problem=1078
*/

// TODO
package programming_challenges

import (
	"fmt"
	"math"
	"time"
)

// inputs:
//   - a slice where each element is the cost for 1 person
//
// returns:
//   - the minimum amount that needs to be transferred to maintain a balance
func min_exchange(costs []float64) float64 {
	to_get := 0.0 // the average amount each person has to get
	for _, val := range costs {
		to_get += val
	}
	to_get = math.Floor((to_get / float64(len(costs)) * 100)) / 100 // multiplying, flooring and dividing by 100 because amt<1/100 can't be paid

	// model this problem as: pooling money from the higher than average, and distributing what is pooled to the lower than avg - communism
	taken := 0.0
	given := 0.0
	for i := 0; i < len(costs); i++ {
		if costs[i] > to_get {
			taken += math.Floor((costs[i]-to_get)*100) / 100
		} else {
			given += math.Floor((to_get-costs[i])*100) / 100
		}
	}
	if taken > given {
		return taken
	} else {
		return given
	}
}

func Run_2() {
	amounts := []float64{1.01, 0.99, 0.99}
	min := min_exchange(amounts)
	start := time.Now()
	fmt.Printf("Min exchange for the amounts: %v, execution time: %s\n", min, time.Since(start))
}
