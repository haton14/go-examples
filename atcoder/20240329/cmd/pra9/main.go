package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	for i1 := 0; i1 <= n; i1++ {
		for i2 := 0; i2 <= n-i1; i2++ {
			i3 := n - i1 - i2
			if i1*10000+i2*5000+i3*1000 == y {
				fmt.Println(i1, i2, i3)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
