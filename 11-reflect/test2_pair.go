package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type Book struct {
	title string
	auth  string
}

func (book *Book) Read() {
	fmt.Println("Reading...")
}

func (book *Book) Write() {
	fmt.Println("Writing...")
}

func main() {

	b := Book{"Go语言", "Go语言"}

	var r Reader
	r = &b
	r.Read()

	var w Writer
	w = r.(Writer)
	w.Write()
}
