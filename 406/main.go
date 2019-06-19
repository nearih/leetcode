package main

import (
	"fmt"
	"sort"
)

func main() {
	// input := [][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}}
	input := [][]int{[]int{8, 2}, []int{4, 2}, []int{4, 5}, []int{2, 0}, []int{1, 4}, []int{7, 2}, []int{9, 1}, []int{3, 1}, []int{9, 0}, []int{1, 0}}
	// input := [][]int{[]int{6, 0}, []int{1, 4}, []int{2, 2}, []int{5, 0}, []int{4, 0}, []int{3, 2}}
	// input := [][]int{[]int{2, 4}, []int{3, 4}, []int{9, 0}, []int{0, 6}, []int{6, 0}, []int{7, 1}, []int{7, 3}, []int{2, 5}, []int{1, 1}, []int{8, 0}}
	res := reconstructQueue(input)
	fmt.Println(res)
	fmt.Println()
}

func reconstructQueue(people [][]int) [][]int {
	withk := [][]int{}
	withh := [][]int{}
	for _, n := range people {
		if n[1] != 0 {
			withk = append(withk, n)
		}
		if n[1] == 0 {
			withh = append(withh, n)
		}
	}
	sort.SliceStable(withh, func(i, j int) bool { return withh[i][0] < withh[j][0] })
	sort.SliceStable(withk, func(i, j int) bool {
		if withk[i][1] < withk[j][1] {
			return true
		}
		return (withk[i][0]) > (withk[j][0]) && (withk[i][1]) == (withk[j][1])
	})
	for _, k := range withk {
		tmp := 0
		c := 0
		for i := 0; i < len(withh); i++ {
			if k[0] <= withh[i][0] {
				tmp++
			}
			c++
			if tmp == k[1] {
				if c == len(withh) {
					withh = append(withh, k)
					c = 0
					continue
				}
				withh = append(withh, nil)
				copy(withh[i+1:], withh[i:len(withh)-1])
				withh[i+1] = k
				c = 0
				continue

			}
		}
	}
	return withh
}

func reconstructQueue3(people [][]int) [][]int {
	withk := [][]int{}
	withh := [][]int{}
	for i := range people {
		if people[i][1] == 0 {
			withh = append(withh, people[i])
			continue
		}
		withk = append(withk, people[i])
	}
	sort.SliceStable(withh, func(i, j int) bool { return withh[i][0] < withh[j][0] })
	sort.SliceStable(withk, func(i, j int) bool {
		if withk[i][1] < withk[j][1] {
			return true
		}
		return (withk[i][0]) > (withk[j][0]) && (withk[i][1]) == (withk[j][1])
	})
	for k := range withk {
		tmp := 0
		for i := 0; i < len(withh); i++ {
			if withk[k][0] <= withh[i][0] {
				tmp++
			}
			if tmp == withk[k][1] {
				tmpls2 := [][]int{}
				tmpls2 = append(tmpls2, withh[i+1:]...)
				withh = append(withh[:i+1], withk[k])
				withh = append(withh, tmpls2...)
				fmt.Println(withh)
				continue
			}
		}
	}
	return withh
}

func reconstructQueueOld(people [][]int) [][]int {
	withk := [][]int{}
	withh := [][]int{}
	for {
		count := 0
		for i := 0; i < len(people)-1; i++ {
			if (people[i][1]) > (people[i+1][1]) {
				tmp := people[i]
				count++
				people[i] = people[i+1]
				people[i+1] = tmp
			}
			if (people[i][1]) == (people[i+1][1]) && (people[i][0]) > (people[i+1][0]) {
				tmp := people[i]
				count++
				people[i] = people[i+1]
				people[i+1] = tmp
			}
		}
		if count == 0 {
			break
		}
	}
	// fmt.Println("11", people)
	for _, n := range people {
		if n[1] != 0 {
			withk = append(withk, n)
		}
		if n[1] == 0 {
			withh = append(withh, n)
		}
	}
	for {
		count := 0
		for i := 0; i < len(withk)-1; i++ {
			if withk[i][0] < withk[i+1][0] {
				count++
				tmp := withk[i]
				withk[i] = withk[i+1]
				withk[i+1] = tmp
			}
		}
		if count == 0 {
			break
		}
	}
	for _, k := range withk {
		tmp := 0
		for i := 0; i < len(withh); i++ {
			if k[0] <= withh[i][0] {
				tmp++
			}
			if tmp == k[1] {
				tmpls2 := [][]int{}
				tmpls2 = append(tmpls2, withh[i+1:]...)
				withh = append(withh[:i+1], k)
				withh = append(withh, tmpls2...)
				continue
			}
		}
	}
	return withh
}

func reconstructQueue2(people [][]int) [][]int {
	withk := [][]int{}
	withh := [][]int{}
	sort.SliceStable(people, func(i, j int) bool { return people[i][1] < people[j][1] })
	sort.SliceStable(people, func(i, j int) bool { return (people[i][1]) == (people[j][1]) && (people[i][0]) < (people[j][0]) })
	// fmt.Println("22", people)
	for _, n := range people {
		if n[1] != 0 {
			withk = append(withk, n)
		}
		if n[1] == 0 {
			withh = append(withh, n)
		}
	}

	sort.SliceStable(withk, func(i, j int) bool { return (withk[i][0] > withk[j][0]) })
	for _, k := range withk {
		tmp := 0
		for i := 0; i < len(withh); i++ {
			if k[0] <= withh[i][0] {
				tmp++
			}
			if tmp == k[1] {
				tmpls2 := [][]int{}
				tmpls2 = append(tmpls2, withh[i+1:]...)
				withh = append(withh[:i+1], k)
				withh = append(withh, tmpls2...)
				continue
			}
		}
	}
	return withh
}
