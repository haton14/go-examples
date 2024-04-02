package main

import "fmt"

func main() {
	var cnt500, cnt100, cnt50, x int
	fmt.Scan(&cnt500, &cnt100, &cnt50, &x)
	count := 0
	for i := cnt500; i >= 0; i-- {
		for j := cnt100; j >= 0; j-- {
			for k := cnt50; k >= 0; k-- {
				if i*500+j*100+k*50 == x {
					count++
					continue
				}
			}
		}
	}
	fmt.Println(count)
}
