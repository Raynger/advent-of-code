package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	points := make(map[Point]bool)
	for {
		sc.Scan()

		s := sc.Text()
		if len(s) == 0 {
			break
		}
		p := strings.Split(s, ",")

		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])

		points[Point{x, y}] = true
	}
	for {
		good := sc.Scan()
		if !good {
			break
		}
		s := strings.Split(sc.Text(), "=")

		isX := s[0][len(s[0])-1] == 'x'
		num, _ := strconv.Atoi(s[1])

		for point := range points {
			if isX && point.x < num {
				continue
			}

			if !isX && point.y < num {
				continue
			}

			delete(points, point)

			if isX {
				point.x = num - (point.x - num)
			}

			if !isX {
				point.y = num - (point.y - num)
			}

			points[point] = true
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 100; j++ {
			if points[Point{j, i}] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
