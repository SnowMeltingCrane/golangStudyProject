package main

import "fmt"

func main() {

	// 变量make
	var slice []int
	slice = make([]int, 3)

	// 直接make
	slice1 := make([]int, 3)

	// 通过变量make
	var slice2 []int = make([]int, 3)

	// 直接定义
	slice3 := []int{1, 2, 3}

	fmt.Println("slice =", slice)
	fmt.Println("slice =", slice1)
	fmt.Println("slice =", slice2)
	fmt.Println("slice =", slice3)
}
