package main

import "fmt"

func printArray(myArray []int) {
	for index, value := range myArray {
		fmt.Println("index=", index, "value=", value)
	}
	myArray[0] = 1000 //传参不再是值拷贝
}

func main() {
	myArray1 := [10]int{1, 2, 3}
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := []int{4, 4, 4, 4} //用切片表示动态数组，传参不再是值拷贝
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index=", index, "value=", value)
	}
	printArray(myArray3) //传参不再是值拷贝
	fmt.Println("------------------")
	for index, value := range myArray3 {
		fmt.Println("index=", index, "value=", value)
	}

	myArray3 = make([]int, 10) //make开辟空间

	var myArray4 []int = make([]int, 4)
	printArray(myArray4)

	fmt.Println(myArray2[0:2])
}
