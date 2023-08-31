package lib1

import "fmt"

// Lib1F 大写表示外部调用
func Lib1F() {
	fmt.Println("Lib1F ... ")
}

func init() {
	fmt.Println("lib1 init() ...")
}
