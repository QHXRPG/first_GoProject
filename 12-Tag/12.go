package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name"`
	Sex  string `info:"sex"`
	Age  int    `info:"age"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() //找到当前结构体全部的元素
	fmt.Println(t.NumField())
	fmt.Println(t.Field(0), t.Field(1), t.Field(2))
	fmt.Println(t.Field(0).Tag, t.Field(1).Tag, t.Field(2).Tag)
	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println("info: ", tagString)
	}

}

func main() {
	var re resume
	findTag(&re)
}
