package main

import "fmt"

// Hero 用struct定义一个类
type Hero struct {
	Name  string
	Age   int
	Level int
}

func (this *Hero) show() {
	fmt.Println("name=", this.Name)
	fmt.Println("age=", this.Age)
	fmt.Println("level=", this.Level)
}

// 继承
type Superman struct {
	Hero        //继承Hero类
	is_superman bool
}

func (this Superman) show() {
	fmt.Println("name=", this.Name)
	fmt.Println("age=", this.Age)
	fmt.Println("level=", this.Level)
	fmt.Println("is_superman=", this.is_superman)
	fmt.Println("\n")
}

func main() {
	hero := Hero{Age: 20, Name: "qhx", Level: 100}
	hero.show()

	superman := Superman{Hero{Name: "qhx", Age: 123, Level: 100}, true}
	superman.show()

	superman.Name = "Qhx"
	superman.show()
}
