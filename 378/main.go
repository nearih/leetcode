package main

import "fmt"

func main() {
	n := kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8)
	fmt.Println(n)
}

func kthSmallest(matrix [][]int, k int) int {
	return recur(matrix, k, 0, 0)
}

func recur(matrix [][]int, k int, i, j int) int {
	// fmt.Println(i, j, k, i*len(matrix)+j)
	if i*len(matrix)+j > k-1 {
		return matrix[i][j]
	}

	if (j == len(matrix[0])-1) && (i == len(matrix)-1) {

	}

	if (j == len(matrix[0])-1) && (i < len(matrix)-1) {
		return recur(matrix, k, i+1, 0)
	}

	if (i == len(matrix)-1) && (j < len(matrix[0])-1) {
		return recur(matrix, k, i, j+1)
	}

	if matrix[i][j+1] > matrix[i+1][j] && (i < len(matrix)-1) {
		return recur(matrix, k, i+1, j)
	}
	if matrix[i][j+1] > matrix[i+1][0] && (i < len(matrix)-1) {
		return recur(matrix, k, i+1, 0)
	}

	return recur(matrix, k, i, j+1)
}
