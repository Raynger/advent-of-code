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
		value["]"] = 57
		value["}"] = 1197
		value[">"] = 25137
		value[")"] = 3

		var stack []string

		for _, ch := range A {
			if ch == "[" || ch == "{" || ch == "<" || ch == "(" {
				stack = append(stack, ch)
			} else {
				if len(stack) == 0 || m[ch] != stack[len(stack)-1] {
					ans += value[ch]
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	fmt.Println(ans)
}
