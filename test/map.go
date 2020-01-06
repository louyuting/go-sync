package main

import (
	"fmt"
	"sync"
)

func main() {

	m := sync.Map{}

	m.Store("a", "av")
	m.Store("b", "bv")
	m.Store("b", "bbv")
	m.Store("b", "bbbbv")

	for i := 0; i < 100; i++ {
		m.Load("a")
	}
	m.Store("c", "cv")

	fmt.Println(m.Load("a"))
	m.Delete("a")

	m.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, ",value:", value, "\n")
		return true
	})
}
