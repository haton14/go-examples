package main

import "log"

func main() {
	pool, err := initDB()
	if err != nil {
		log.Fatalf("initDB(): err %s", err)
	}
	Select(pool)
}
