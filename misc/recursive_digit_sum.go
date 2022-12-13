package misc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func super_digit(digit string, k int) int {
	if len(digit) <= 1 {
		dig_int, _ := strconv.Atoi(digit)
		return dig_int
	}

	digits := strings.Split(digit, "")
	sum := 0
	for _, dig := range digits {
		dig_int, _ := strconv.Atoi(dig)
		sum += dig_int
	}
	sum_str := fmt.Sprintf("%d", k*sum)
	return super_digit(sum_str, 1)
}

func superDigit(n string, k int32) int32 {
	converted_n := ""
	for _, digit := range n {
		if _, err := strconv.Atoi(string(digit)); err == nil { // the digit is an integer
			converted_n += string(digit)
		}
	}
	return int32(super_digit(converted_n, int(k)))
}

func Run() {
	digit := "593"
	k := 10
	start := time.Now()
	super := superDigit(digit, int32(k))
	fmt.Printf("\ndigit: %s, k: %d, super digit: %d, execution time: %s\n", digit, k, super, time.Since(start))
}
