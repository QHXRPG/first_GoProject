package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{Title: "A", Year: 2001, Price: 100, Actors: []string{"a", "b"}}
	fmt.Println(movie)
	//编码 结构体->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----------------")
	fmt.Printf("%s\n", jsonStr)
	fmt.Println(reflect.TypeOf(jsonStr))
	fmt.Println(jsonStr)

	//解码
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----------------")
	fmt.Println(myMovie)
}
