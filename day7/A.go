package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func f(A []int, n int) int {
	total := 0
	for _, v := range A {
		total += int(math.Abs(float64(v - n)))
	}
	return total
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var A []int

	sc.Scan()
	s := strings.Split(sc.Text(), ",")

	min := 0x3f3f3f3f
	max := -1
	for i := range s {
		n, _ := strconv.Atoi(s[i])
		A = append(A, n)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	for min != max {
		mid := (min + max) / 2
		if f(A, mid) < f(A, mid+1) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	fmt.Println(f(A, min))
}
