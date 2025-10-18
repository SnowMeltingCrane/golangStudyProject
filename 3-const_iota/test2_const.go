package main

import "fmt"

// iota 只能出现在const 声明中
const (
	BEIJING = iota
	SHANGHAI
	FUJIAN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	// 常量（只读）
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("FUJIAN = ", FUJIAN)

	fmt.Println("a = ", a, " b = ", b)
	fmt.Println("c = ", c, " d = ", d)
	fmt.Println("e = ", e, " f = ", f)
	fmt.Println("g = ", g, " h = ", h)
	fmt.Println("i = ", i, " k = ", k)
}
