package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ans []int64
	for {
		ok := sc.Scan()
		if !ok {
			break
		}

		A := strings.Split(sc.Text(), "")

		m := make(map[string]string)
		m["]"] = "["
		m["}"] = "{"
		m[">"] = "<"
		m[")"] = "("

		value := make(map[string]int)
		value["("] = 1
		value["["] = 2
		value["{"] = 3
		value["<"] = 4

		var stack []string
		invalid := false
		for _, ch := range A {
			if ch == "[" || ch == "{" || ch == "<" || ch == "(" {
				stack = append(stack, ch)
			} else {
				if len(stack) == 0 || m[ch] != stack[len(stack)-1] {
					invalid = true
					break
				}
				stack = stack[:len(stack)-1]
			}
		}

		if !invalid && len(stack) > 0 {
			var n int64
			for k := len(stack) - 1; k >= 0; k-- {
				n *= 5
				n += int64(value[stack[k]])
			}
			ans = append(ans, n)
		}
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })

	fmt.Println(ans[len(ans)/2])
}
