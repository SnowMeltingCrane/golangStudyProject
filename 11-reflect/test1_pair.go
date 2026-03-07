package main

import "fmt"

func main() {

	var a string = "abcld"

	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
