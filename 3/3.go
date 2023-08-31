package main

import "fmt"

// 匿名返回值
func foo1(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 99
	return c, c*2 + 1
}

// 有名返回值
func foo2(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 100
	r2 = 200
	return
}

func main() {
	c1, c2 := foo1("abc", 4)
	fmt.Println("c1 = ", c1, " c2 = ", c2)

	c1, c2 = foo2("abc", 4)
	fmt.Println("c1 = ", c1, " c2 = ", c2)
}
