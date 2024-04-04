package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	type empty struct{}
	result := make(map[int]empty)
	for _, v := range d {
		result[v] = empty{}
	}
	fmt.Println(len(result))
}
