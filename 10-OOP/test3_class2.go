package main

import "fmt"

type Human struct {
	name string
	age  int
}

type SuperMan struct {
	Human
	level int
}

func (this *Human) say() {
	fmt.Println("I'm", this.name, "and", this.age, "years old")
}

func (this *SuperMan) Fly() {
	fmt.Println("i can fly")
}

func main() {

	var superMan SuperMan
	var human Human
	human.name = "human"
	human.age = 18
	human.say()

	superMan.name = "superMan"
	superMan.age = 18
	superMan.level = 1

	superMan.say()
	superMan.Fly()
}
