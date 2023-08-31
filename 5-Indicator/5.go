package main

import "fmt"

func fun1(p *int) {
	*p = 10 //解引用
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a int = 1
	fmt.Println("a = ", a)
	fun1(&a)
	fmt.Println("a = ", a)

	var m int = 10
	var n int = 20
	fmt.Println("m = ", m, "n = ", n)
	swap(&m, &n)
	fmt.Println("m = ", m, "n = ", n)
}
