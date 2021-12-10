package main

import (
	"fmt"
	"strconv"
)

func filter(m map[string]bool, size int, count_most bool) int64 {
	for i := 0; i < size; i++ {
		cnt := 0
		for key, _ := range m {
			if string(key[i]) == "1" {
				cnt++
			} else {
				cnt--
			}
		}

		var keep_ones = false
		if count_most {
			if cnt >= 0 {
				keep_ones = true
			}
		} else {
			if cnt < 0 {
				keep_ones = true
			}
		}

		for key, _ := range m {
			if string(key[i]) == "1" && !keep_ones {
				delete(m, key)
			} else if string(key[i]) == "0" && keep_ones {
				delete(m, key)
			}
		}
		if len(m) == 1 {
			break
		}
	}
	for key := range m {
		if i, err := strconv.ParseInt(key, 2, 32); err == nil {
			return i
		}
	}
	return -1
}

func main() {
	m := make(map[string]bool)
	l := make(map[string]bool)
	var size int
	for {
		var in string
		if _, err := fmt.Scanf("%s\n", &in); err != nil {
			// fmt.Println(err);
			break
		}

		m[in] = true
		l[in] = true
		size = len(in)
	}

	// var most, least int
	fmt.Println(filter(m, size, true) * filter(l, size, false))
}
