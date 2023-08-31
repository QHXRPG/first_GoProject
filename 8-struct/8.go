package main

import "fmt"

type myint int

type Book struct {
	title string
	auth  string
}

func func1(book *Book) {
	//指针传递
	book.title = "mmm"
}

func main() {
	var book1 Book
	book1.title = "qqq"
	book1.auth = "zzz"
	fmt.Println(book1)
	func1(&book1) //传地址
	fmt.Println(book1)
}
