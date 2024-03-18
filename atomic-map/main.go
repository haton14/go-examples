package main

import (
	"fmt"
	"maps"
	"sync/atomic"
	"time"
)

type cache struct {
	Value string
}

func main() {
	config := atomic.Pointer[map[string]cache]{}
	config.Store(&map[string]cache{"key": {Value: "value"}})
	go func() {
		for range 10 {
			bb(&config)
		}
	}()
	go func() {
		for range 10 {
			cc(&config)
		}
	}()
	time.Sleep(2 * time.Second)

}

func bb(config *atomic.Pointer[map[string]cache]) {
	for {
		b := config.Load()
		newb := maps.Clone(*b)
		newb["key"] = cache{Value: "value2"}
		if config.CompareAndSwap(b, &newb) {
			fmt.Println("b: ", b)
			break
		}
		fmt.Println("他で更新していたので、b retry")
	}
}

func cc(config *atomic.Pointer[map[string]cache]) {
	for {
		c := config.Load()
		newc := maps.Clone(*c)
		newc["key"] = cache{Value: "value1"}
		if config.CompareAndSwap(c, &newc) {
			fmt.Println("c: ", c)
			break
		}
		fmt.Println("他で更新していたので、c retry")
	}
}
