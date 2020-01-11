package main

func main() {

}

func search(grid [][]int, i, j int) int {
	sum := 0
	if i+j == len(grid)+len(grid[0])-2 {
		return sum + grid[i][j]
	} else if i == len(grid)-1 {
		return sum + grid[i][j] + search(grid, i, j+1)
	} else if j == len(grid[0])-1 {
		return sum + grid[i][j] + search(grid, i+1, j)
	}

	sumA := grid[i][j] + search(grid, i+1, j)
	sumB := grid[i][j] + search(grid, i, j+1)

	if sumA > sumB {
		return sum + sumB
	}
	return sum + sumA
}

func minPathSum(grid [][]int) int {
	return search(grid, len(grid)-1, len(grid[0])-1)
}

func searchMinus(grid [][]int, i, j int) int {
	sum := 0
	if i == 0 && j == 0 {
		return sum + grid[i][j]
	}
	if i == 0 && j > 0 {
		return sum + grid[i][j] + search(grid, i, j-1)
	} else if j == 0 && i > 0 {
		return sum + grid[i][j] + search(grid, i-1, j)
	}

	if i > 0 && j > 0 {
		sumA := grid[i][j] + search(grid, i-1, j)
		sumB := grid[i][j] + search(grid, i, j-1)

		if sumA > sumB {
			return sum + sumB
		} else {
			return sum + sumA
		}
	}
	return sum
}

func searchLopp(grid [][]int) int {
	ver := len(grid)
	hor := len(grid[0])
	// vertical
	for i := 1; i < ver; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// horizontal
	for j := 1; j < hor; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < ver; i++ {
		for j := 1; j < hor; j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
		}
	}

	return grid[ver-1][hor-1]
}
