package main

import (
	"fmt"
	"time"
)

// 从goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newTask Goroutine", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	//创建一个Go程 去执行newTask()
	go newTask()

	i := 0
	fmt.Println("---------")
	for {
		i++
		fmt.Println("main Goroutine", i)
		time.Sleep(1 * time.Second)
	}
}
