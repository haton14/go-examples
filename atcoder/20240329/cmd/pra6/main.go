package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var result int
	for ; n >= a; n-- {
		if valid(n, a, b) {
			result += n
		}
	}
	fmt.Println(result)
}

func valid(target, a, b int) bool {
	v := 0
	for target > 9 {
		v += target % 10
		target = target / 10
	}
	v += target
	return a <= v && v <= b
}
