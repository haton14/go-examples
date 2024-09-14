package main

import "fmt"

// $ go run .
// unko
// unko unko
// true
func main() {
	for range f {
		fmt.Println("unko unko")
	}
}

func f(yield func() bool) {
	fmt.Println("unko")
	// yiled()が呼ばれるとloopに戻る
	// loopが一回終わるとyiled()の戻り値boolにloopを続行するか否かの結果を返した上で、fに戻る
	fmt.Println(yield())
	// yiledが呼ばれずに関数fが終了すると、loopも終了する
}
