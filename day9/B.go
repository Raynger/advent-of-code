package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var grid [][]int
var vis [][]bool
var dx [4]int = [4]int{-1, 0, 1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}

func dfs(i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	if grid[i][j] == 9 || vis[i][j] {
		return 0
	}

	vis[i][j] = true

	ans := 1
	for k := 0; k < 4; k++ {
		nx := i + dx[k]
		ny := j + dy[k]
		ans += dfs(nx, ny)
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
		var visRow []bool

		for _, num := range sc.Text() {
			n, _ := strconv.Atoi(string(num))
			row = append(row, n)
			visRow = append(visRow, false)
		}
		grid = append(grid, row)
		vis = append(vis, visRow)
	}

	var ans []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !vis[i][j] {
				ans = append(ans, dfs(i, j))
			}
		}
	}
	sort.Ints(ans)
	fmt.Println(ans[len(ans)-1] * ans[len(ans)-2] * ans[len(ans)-3])
}
