package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	tmp := strings.Split(s, "eraser")
	s = strings.Join(tmp, "")
	tmp = strings.Split(s, "erase")
	s = strings.Join(tmp, "")
	tmp = strings.Split(s, "dreamer")
	s = strings.Join(tmp, "")
	tmp = strings.Split(s, "dream")
	s = strings.Join(tmp, "")
	if s == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
