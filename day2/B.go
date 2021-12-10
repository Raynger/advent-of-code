package main

import "fmt"

func main() {
	var height, length, aim int
	for {
		var dir string
		var dist int
		if _, err := fmt.Scanf("%s %d\n", &dir, &dist); err != nil {
			// fmt.Println(err);
			break
		}

		if dir == "forward" {
			length += dist
			height += aim * dist
		} else if dir == "down" {
			aim += dist
		} else {
			aim -= dist
		}
		// sl = append(sl, d)
	}

	fmt.Println(height * length)
}
