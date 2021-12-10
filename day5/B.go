package main

import (
	"fmt"
	"math"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	const N = 1000
	var grid [N][N]int

	for {
		var a, b, c, d int
		_, err := fmt.Scanf("%d,%d -> %d,%d\n", &a, &b, &c, &d)
		if err != nil {
			break
		}

		var hD, vD int
		if c != a {
			hD = (c - a) / int(math.Abs(float64(c-a)))
		}
		if d != b {
			vD = (d - b) / int(math.Abs(float64(d-b)))
		}

		for {
			grid[a][b]++

			if a == c && b == d {
				break
			}

			a += hD
			b += vD
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
