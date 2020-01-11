package main

import "math"

func main() {
	titleToNumber("ZY")
}

func titleToNumber(s string) int {
	a := []int{}
	for _, r := range s {
		a = append(a, int(r)-64)
	}

	res := 0
	for i := len(a) - 1; i > -1; i-- {
		res += a[i] * int(math.Pow(26, float64(len(a)-i-1)))
	}
	return 0
}
