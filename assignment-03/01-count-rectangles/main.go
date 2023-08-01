package main

import "fmt"

func countRectangles(matrix [][]int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	count := 0

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || matrix[row][col] == 0 || visited[row][col] {
			return
		}

		visited[row][col] = true

		dfs(row, col+1)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row-1, col)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}

func main() {
	// Test input
	matrix := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	result := countRectangles(matrix)
	fmt.Printf("%v\n", result)
}
