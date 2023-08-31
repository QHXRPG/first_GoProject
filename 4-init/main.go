package main

import (
	mylib1 "first_GoProject/4-init/lib1" //取名字 mylib1
	_ "first_GoProject/4-init/lib2"      //导入不使用，只调用它的init函数
)

func main() {
	mylib1.Lib1F()
	//lib2.Lib2F()
}
