package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//用go创建一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		//匿名函数
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() //退出当前goroutine
			fmt.Println("B")
		}() //在{}后面接一个()表示定义并调用该函数，无形参

		fmt.Println("A")
	}() //在{}后面接一个()表示定义并调用该函数，无形参

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
