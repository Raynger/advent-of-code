package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var grid [][]int
var dx [8]int = [8]int{-1, 0, 1, 0, -1, 1, 1, -1}
var dy [8]int = [8]int{0, 1, 0, -1, -1, 1, -1, 1}

func dfs(i int, j int, vis [][]bool) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	grid[i][j]++

	if grid[i][j] <= 9 || vis[i][j] {
		return 0
	}

	vis[i][j] = true

	ans := 1
	for k := 0; k < 8; k++ {
		nx := i + dx[k]
		ny := j + dy[k]
		ans += dfs(nx, ny, vis)
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		pass := sc.Scan()
		if !pass {
			break
		}
		var row []int

		for _, num := range sc.Text() {
			n, _ := strconv.Atoi(string(num))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	ans := 0
	for k := 0; k < 100; k++ {
		var vis [][]bool
		for i := 0; i < 10; i++ {
			vis = append(vis, make([]bool, 10))
		}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				ans += dfs(i, j, vis)
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] > 9 {
					grid[i][j] = 0
				}
			}
		}
	}
	fmt.Println(ans)
}
