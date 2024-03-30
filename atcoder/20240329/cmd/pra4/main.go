package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	compressed := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
		compressed[s[i]] = s[i]
	}
	result := 0
GO:
	for {
		for k, v := range compressed {
			if v%2 != 0 {
				break GO
			}
			compressed[k] = v / 2
		}
		result++
	}
	fmt.Println(result)
}
