package main

import (
	"fmt"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	const N = 10
	var grid [N][N]int

	for {
		var a, b, c, d int
		_, err := fmt.Scanf("%d,%d -> %d,%d\n", &a, &b, &c, &d)
		if err != nil {
			break
		}
		if a != c && b != d {
			continue
		}

		if c < a {
			a, c = c, a
		}

		if d < b {
			b, d = d, b
		}

		for i := a; i <= c; i++ {
			for j := b; j <= d; j++ {
				grid[i][j]++
			}
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] > 1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
	// sc.Scan()
	// sc.Text()
}
