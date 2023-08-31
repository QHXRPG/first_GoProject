package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func deferfunc() int {
	fmt.Println("defer....")
	return 0
}

func returnfunc() int {
	fmt.Println("return ....")
	return 0
}

func returnanddefer() int {
	defer deferfunc()
	return returnfunc()
}

func main() {
	defer func1()
	defer func2()
	defer func3()

	//return 还是defer先执行
	returnanddefer()
}
