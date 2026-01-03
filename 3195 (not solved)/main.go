package main

import "fmt"

/*
Input: grid = [[0,1,0],[1,0,1]]

Output: 6
*/

func minimumArea(grid [][]int) int {

	var height, width int

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				height = i
				width = j
			}
		}
	}
	return (height + 1) * (width + 1)

}

func main() {
	grid1 := [][]int{{0, 1, 0}, {1, 0, 1}}
	grid2 := [][]int{{1, 0}, {0, 0}}
	fmt.Println(minimumArea(grid1))
	fmt.Println(minimumArea(grid2))

}
