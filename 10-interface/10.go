package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string //猫的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 万能类型
func myfunc(arg interface{}) {
	fmt.Println("myfunc ....")
	fmt.Println("arg = ", arg)
}

func main() {
	var animal AnimalIF
	animal = &Cat{"green"}
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())
	fmt.Println("-------------")
	animal = &Dog{"red"}
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())

	fmt.Println("-------------")
	myfunc("asd")
	myfunc(12)
}
