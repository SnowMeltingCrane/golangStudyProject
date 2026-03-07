package main

import "fmt"

type myint int

type Book struct {
	title string
	auth  string
}

func changeBook2(book *Book) {
	book.auth = "xue"
}

func main() {
	var a myint = 10

	fmt.Printf("type of a = %T\n", a)

	var book Book

	book.title = "Go语言"
	book.auth = "雪融鹤"
	fmt.Println("book = ", book)

	changeBook2(&book)
	fmt.Println("book = ", book)
}
