package main

import "fmt"

func main() {
	var s int
	fmt.Scanf("%d", &s)
	s3 := s % 2
	s /= 10
	s2 := s % 2
	s /= 10
	s1 := s % 2
	fmt.Println(s1 + s2 + s3)
}
