package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func not_done(nums [][]string, marks [][]bool) int {
	total := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !marks[i][j] {
				n, err := strconv.Atoi(nums[i][j])
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				total += n
			}
		}
	}
	return total
}

func check(marks [][]bool) bool {
	for i := 0; i < 5; i++ {
		good := true
		for j := 0; j < 5; j++ {
			if marks[i][j] == false {
				good = false
				break
			}
		}
		if good {
			return true
		}
	}

	for j := 0; j < 5; j++ {
		good := true
		for i := 0; i < 5; i++ {
			if marks[i][j] == false {
				good = false
				break
			}
		}
		if good {
			return true
		}
	}
	return false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var values []int

	sc.Scan()
	s := strings.Split(sc.Text(), ",")
	for i := range s {
		n, _ := strconv.Atoi(s[i])
		values = append(values, n)
	}
	done := false

	best_index := 0x3f3f3f3f
	ans := 0

	for !done {
		var nums [][]string
		var marks [][]bool

		sc.Scan()
		for i := 0; i < 5; i++ {
			sc.Scan()
			line := sc.Text()

			if line == "" {
				done = true
				break
			}
			nums = append(nums, strings.Fields(sc.Text()))
			marks = append(marks, make([]bool, 5))
		}
		if done {
			break
		}

		for index, v := range s {
			if index >= best_index {
				break
			}
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if nums[i][j] == v {
						marks[i][j] = true
					}
				}
			}
			if check(marks) {
				best_index = index
				ans = values[index] * not_done(nums, marks)
			}
		}
	}
	fmt.Println(ans)
}
