package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	ans := 0
	for {
		pass := sc.Scan()
		if !pass {
			break
		}
		s := strings.Split(strings.Split(sc.Text(), "|")[1], " ")

		for _, v := range s {
			n := len(v)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
