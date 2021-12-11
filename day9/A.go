package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	ans := 0
	var grid [][]int
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
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			shorter := true
			for k := 0; k < 4; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
					continue
				}
				if grid[i][j] >= grid[nx][ny] {
					shorter = false
				}
			}
			if shorter {
				ans += grid[i][j] + 1
			}
		}
	}
	fmt.Println(ans)
}
