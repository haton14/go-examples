package main

import (
	"fmt"
	"iter"
)

// $ go run .
// unko
// unko unko
// true
func main() {
	for range f {
		fmt.Println("unko unko")
	}

	for i := range getList([]int{4, 2, 32, 44, 5}) {
		fmt.Println(i)
		if i == 44 {
			break
		}
	}
}

func f(yield func() bool) {
	fmt.Println("unko")
	// yiled()が呼ばれるとloopに戻る
	// loopが一回終わるとyiled()の戻り値boolにloopを続行するか否かの結果を返した上で、fに戻る
	fmt.Println(yield())
	// yiledが呼ばれずに関数fが終了すると、loopも終了する
}

func getList(list []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range list {
			if !yield(v) {
				return
			}
		}
	}
}
