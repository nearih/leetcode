package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("main")
	permute([]int{1, 2, 3, 4, 5})
}

func permute(nums []int) [][]int {
	sort.Ints(nums)
	result := perm(nums)
	fmt.Println(result)
	fmt.Println(len(result))
	return result
}

func perm(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for i, num := range nums {
		temp := []int{}
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)
		resrec := perm(temp)
		for _, i := range resrec {
			ans := []int{}
			ans = append(ans, num)
			ans = append(ans, i...)
			res = append(res, ans)
		}
	}
	return res
}
