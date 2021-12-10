package main

import "fmt"

func main() {
	var s []string
	for {
		var in string
		if _, err := fmt.Scanf("%s\n", &in); err != nil {
			// fmt.Println(err);
			break
		}

		s = append(s, in)
	}

	var most, least int
	for i := range s[0] {
		cnt := 0
		for j := range s {
			if string(s[j][i]) == "1" {
				cnt++
			} else {
				cnt--
			}
		}
		most <<= 1
		least <<= 1
		if cnt >= 0 {
			most |= 1
		} else {
			least |= 1
		}
	}

	fmt.Println(most * least)
}
