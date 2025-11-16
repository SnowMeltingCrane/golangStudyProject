package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 10
	return c
}

// 返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 10
	d := 20
	return c, d
}

// 返回多个返回值，命名
func foo3(a string, b int) (ret1 int, ret2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值赋值
	ret1 = 1000
	ret2 = 2000
	return
}

// 返回多个返回值，命名且类型相同
func foo4(a string, b int) (ret1, ret2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	ret1 = 1000
	ret2 = 2000
	return
}

func main() {

	c := foo1("hello", 10)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("hello", 10)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("hello", 10)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("hello", 10)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
