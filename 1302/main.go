package main

import "sync"

var m map[int][]int
var max int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	deepestLeavesSum(nil)
}

func deepestLeavesSum(root *TreeNode) int {
	m = make(map[int][]int)
	c := 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		walker(root, c, &wg)
		wg.Done()
	}()

	wg.Wait()

	res := 0
	for _, v := range m[max] {
		res += v
	}

	return res
}

func walker(t *TreeNode, c int, wg *sync.WaitGroup) {
	if t == nil {
		return
	}

	m[c] = append(m[c], t.Val)

	if max < c {
		max = c
	}

	c++
	wg.Add(2)
	go func() {
		walker(t.Left, c, wg)
		wg.Done()
	}()
	go func() {
		walker(t.Right, c, wg)
		wg.Done()
	}()
}
