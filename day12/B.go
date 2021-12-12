package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfs(node string, adjList map[string][]string, vis map[string]int, used bool) int {
	if node == "end" {
		return 1
	}

	ans := 0
	for _, next := range adjList[node] {
		if next == strings.ToLower(next) && vis[next] >= 1 {
			if next == "start" || used {
				continue
			}
		}
		vis[next]++
		if next == strings.ToLower(next) && vis[next] >= 2 {
			used = true
		}
		ans += dfs(next, adjList, vis, used)
		if next == strings.ToLower(next) && vis[next] >= 2 {
			used = false
		}
		vis[next]--
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	adjList := make(map[string][]string)
	vis := make(map[string]int)

	for {
		ok := sc.Scan()
		if !ok {
			break
		}

		A := strings.Split(sc.Text(), "-")

		adjList[A[0]] = append(adjList[A[0]], A[1])
		adjList[A[1]] = append(adjList[A[1]], A[0])

		vis[A[0]] = 0
		vis[A[1]] = 0
	}
	vis["start"] = 2
	fmt.Println(dfs("start", adjList, vis, false))
}
