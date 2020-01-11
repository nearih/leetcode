package main

import "sort"

import "fmt"

func main() {
	x := isAnagram("qwyert", "yteqwr")
	fmt.Println(x)
	// fmt.Println(string(rune('A')))
}

func isAnagram(s string, t string) bool {
	sl := []int{}
	tl := []int{}

	for _, n := range s {
		sl = append(sl, int(n))
	}

	for _, n := range t {
		tl = append(tl, int(n))
	}

	sort.Ints(sl)
	sort.Ints(tl)
	for i, n := range sl {
		if n != tl[i] {
			return false
		}
	}
	return true
}
