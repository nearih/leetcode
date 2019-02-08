package main

import "fmt"

func main() {
	// count all zero
	grid := [][]int{[]int{1, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 2, -1}}
	n := uniquePathsIII(grid)
	fmt.Println(n)
}

func uniquePathsIII(grid [][]int) int {
	zero := 0
	start := []int{}
	// get total number of zero
	for y, slice := range grid {
		for x, num := range slice {
			if num == 0 {
				zero++
			}
			if num == 1 {
				start = append(start, y)
				start = append(start, x)
			}
		}
	}
	n := recur(cparray(grid), start[0], start[1], start[0], start[1], 0, zero)
	pg(cparray(grid))

	return n
}

func cparray(grid [][]int) [][]int {

	var mat = make([][]int, len(grid))
	for i := range mat {
		mat[i] = make([]int, len(grid[0]))
	}

	for n := range grid {
		for m := range grid[0] {
			mat[n][m] = grid[n][m]
		}

	}
	return mat
}

func pg(grid [][]int) {
	for _, n := range grid {
		for _, j := range n {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkGrid(grid [][]int) int {
	for _, n := range grid {
		for _, j := range n {
			if j == 0 {
				return 0
			}
		}
	}
	return 1
}

func recur(grid [][]int, y, x, oy, ox, c, zero int) int {

	pg(grid)
	if grid[y][x] == 2 {
		return checkGrid(grid)

	}
	grid[y][x] = -1

	success := 0
	if (y < len(grid)-1) && (grid[y+1][x] != -1) {
		fmt.Println("case y+1")
		success += recur(cparray(grid), y+1, x, y, x, c, zero)

	}
	if (y > 0) && (grid[y-1][x] != -1) {
		fmt.Println("case y-1")
		success += recur(cparray(grid), y-1, x, y, x, c, zero)

	}
	if (x < len(grid[0])-1) && (grid[y][x+1] != -1) {
		fmt.Println("case x+1")
		success += recur(cparray(grid), y, x+1, y, x, c, zero)

	}
	if (x > 0) && (grid[y][x-1] != -1) {
		fmt.Println("case x-1")
		success += recur(cparray(grid), y, x-1, y, x, c, zero)

	}

	return success
}
