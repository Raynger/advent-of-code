package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(st string) string {
	word := strings.Split(st, "")
	sort.Strings(word)
	return strings.Join(word, "")
}

func common(a, b string) int {
	ans := 0
	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				ans++
				break
			}
		}
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	ans := 0
	for {
		pass := sc.Scan()
		if !pass {
			break
		}
		A := strings.Split(sc.Text(), " | ")

		s := strings.Split(A[0], " ")
		test := strings.Split(A[1], " ")

		sort.Slice(s, func(i, j int) bool {
			return len(s[i]) < len(s[j])
		})

		m := make(map[string]int)

		m[sortString(s[0])] = 1
		m[sortString(s[1])] = 7
		m[sortString(s[2])] = 4
		m[sortString(s[9])] = 8

		for i := 6; i < 9; i++ {
			if common(s[2], s[i]) == 4 {
				m[sortString(s[i])] = 9
			} else if common(s[0], s[i]) == 2 {
				m[sortString(s[i])] = 0
			} else {
				m[sortString(s[i])] = 6
			}
		}
		for i := 3; i < 6; i++ {
			if common(s[0], s[i]) == 2 {
				m[sortString(s[i])] = 3
			} else if common(s[2], s[i]) == 3 {
				m[sortString(s[i])] = 5
			} else {
				m[sortString(s[i])] = 2
			}
		}

		n := 0
		for _, v := range test {
			n *= 10
			n += m[sortString(v)]
		}
		ans += n
	}
	fmt.Println(ans)
}
