package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	switch a * b % 2 {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	}
}
