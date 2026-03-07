package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	// interface{] 区分类型

	//给 interface{} 提供了类型断言机制
	//switch arg.(type) {
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//case Book:
	//	fmt.Println("Book")
	//default:
	//	fmt.Println("unknow")
	//}

	value, ok := arg.(string)
	if ok {
		fmt.Println("string:", value)
	} else {
		fmt.Println("arg is not string")
	}
}

func main() {
	book := Book{auth: "xue", title: "Go语言"}

	myFunc(book)
}
