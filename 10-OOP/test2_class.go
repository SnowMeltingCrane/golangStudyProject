package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) GetName() string {
	return this.Name
}

func (this *Hero) GetAd() int {
	return this.Ad
}

func (this *Hero) GetLevel() int {
	return this.Level
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) SetAd(ad int) {
	this.Ad = ad
}

func (this *Hero) setLevel(level int) {
	this.Level = level
}

func main() {

	hero := Hero{Name: "hero", Ad: 100, Level: 1}

	fmt.Println(hero)

	hero.SetName("hero1")
	fmt.Println(hero)

}
