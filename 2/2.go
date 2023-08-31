package main

import "fmt"

// const 配合iota 定义枚举类型
const (
	Beijing  = 10 * iota // 0
	Shanghai             // 10
	Shenzhen             // 20
)

const (
	a, b = iota + 1, iota + 2 // iota = 0
	c, d                      // iota = 1
	e, f = iota * 3, iota * 2 // iota = 2
	g, h                      // iota = 3
)

func main() {
	const lenght int = 10 //常量， 只读
	fmt.Println("lenght=", lenght)

	fmt.Println("Beijing = ", Beijing)
	fmt.Println("Shanghai = ", Shanghai)
	fmt.Println("Shenzhen", Shenzhen)
}
