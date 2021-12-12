package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfs(node string, adjList map[string][]string, vis map[string]bool) int {
	if node == "end" {
		return 1
	}

	ans := 0
	for _, next := range adjList[node] {
		if next == strings.ToLower(next) && vis[next] {
			continue
		}
		vis[next] = true
		ans += dfs(next, adjList, vis)
		vis[next] = false
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	adjList := make(map[string][]string)
	vis := make(map[string]bool)

	for {
		ok := sc.Scan()
		if !ok {
			break
		}

		A := strings.Split(sc.Text(), "-")

		adjList[A[0]] = append(adjList[A[0]], A[1])
		adjList[A[1]] = append(adjList[A[1]], A[0])

		vis[A[0]] = false
		vis[A[1]] = false
	}
	vis["start"] = true
	fmt.Println(dfs("start", adjList, vis))
}
