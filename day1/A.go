package main

import "fmt"
import "io"

func main() {
	var cur, next, ans int
	fmt.Scan(&cur);
	for {
		_, err := fmt.Scan(&next);
		if err == io.EOF {
			break;
		}
		if next > cur {
			ans++;
		}
		cur = next;
	} 
	fmt.Println(ans);
}