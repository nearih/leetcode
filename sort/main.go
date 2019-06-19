package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{[]int{6, 0}, []int{1, 4}, []int{2, 2}, []int{5, 0}, []int{4, 0}, []int{3, 2}}
	res := sortlist(input)
	fmt.Println(res)
}

func sortlist(people [][]int) [][]int {
	for {
		count := 0
		for i := 0; i < len(people)-1; i++ {
			if (people[i][1]) > (people[i+1][1]) {
				tmp := people[i]
				count++
				people[i] = people[i+1]
				people[i+1] = tmp
			}
			// if (people[i][1]) == (people[i+1][1]) && (people[i][0]) > (people[i+1][0]) {
			// 	tmp := people[i]
			// 	count++
			// 	people[i] = people[i+1]
			// 	people[i+1] = tmp
			// }
		}
		if count == 0 {
			break
		}
	}
	sort.SliceStable(people, func(i, j int) bool { return people[i][1] < people[j][1] })
	return people
}
