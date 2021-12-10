package main


import "fmt"

func main() {
	var ans int
	var sl []int
	
	for {
		d := 0
		if _, err := fmt.Scanf("%d\n", &d); err != nil {
			// fmt.Println(err);
			break;
		}
		sl = append(sl, d);
	}
	
	cur := sl[0] + sl[1] + sl[2]
	for i := 3; i < len(sl); i++ {
		next := cur + sl[i] - sl[i-3];
		if cur < next {
			ans++;
		}
		cur = next;
		// fmt.Println(cur);
		// cur = next;
	} 
	fmt.Println(ans);
}