package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var fish [9]int

	sc.Scan()
	s := strings.Split(sc.Text(), ",")

	for i := range s {
		n, _ := strconv.Atoi(s[i])
		fish[n]++
	}

	for i := 0; i < 80; i++ {
		newFish := fish[0]
		for j := 1; j < 9; j++ {
			fish[j-1] = fish[j]
		}
		fish[8] = newFish
		fish[6] += newFish
	}
	total := 0
	for i := 0; i < 9; i++ {
		total += fish[i]
	}
	fmt.Println(total)
}
