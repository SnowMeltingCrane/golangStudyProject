package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetName() string
}

type Cat struct {
	name  string
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("I'm cat, I'm sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetName() string {
	return this.name
}

func show(animal AnimalIF) {
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetName())
}

type Dog struct {
	name  string
	color string
}

func (d Dog) Sleep() {
	fmt.Println("I'm dog, I'm sleep")
}

func (d Dog) GetColor() string {
	return d.color
}

func (d Dog) GetName() string {
	return d.name
}

func main() {
	//var animal AnimalIF
	//animal = &Cat{"kitty", "white"}
	//animal.Sleep()
	//show(animal)
	//animal = &Dog{"doggy", "black"}
	//animal.Sleep()
	//show(animal)

	cat := Cat{"kitty", "white"}
	cat.Sleep()
	dog := Dog{"doggy", "black"}
	dog.Sleep()
}
