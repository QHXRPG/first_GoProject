package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")        //不加分号, Println自动换行
	time.Sleep(1 * time.Second) //睡眠1秒

	var a int
	fmt.Println("a=", a) //默认为0

	var b int = 100
	fmt.Println("b=", b)

	var c = 100
	fmt.Println("c=", c) //让编译器猜类型

	fmt.Printf("type of c = %T\n", c) //查看变量类型

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	e := 100 //省略var关键字, 声明全局变量不可用该方法
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := 3.14 //省略var关键字
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	// 多变量赋值
	var xx, yy int = 100, 200
	fmt.Println("xx=", xx, "yy=", yy)

	var kk, ll = 100, "qsc"
	fmt.Println("kk=", kk, "ll=", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv=", vv, "jj=", jj)
}
