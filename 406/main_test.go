package main

import "testing"

// var result [][]int

func Benchmark_mysort(b *testing.B) {
	// var r [][]int
	input := [][]int{[]int{6, 0}, []int{1, 4}, []int{2, 2}, []int{5, 0}, []int{4, 0}, []int{3, 2}}
	for i := 0; i < b.N; i++ {
		reconstructQueue(input)
		// r = reconstructQueue(input)
	}
	// result = r
}

func Benchmark_packagesort(b *testing.B) {
	// var r [][]int
	input := [][]int{[]int{6, 0}, []int{1, 4}, []int{2, 2}, []int{5, 0}, []int{4, 0}, []int{3, 2}}
	for i := 0; i < b.N; i++ {
		// r = reconstructQueue2(input)
		reconstructQueue2(input)
	}
	// result = r
}
