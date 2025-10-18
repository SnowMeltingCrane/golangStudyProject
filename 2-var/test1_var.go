package main

import (
	"fmt"
)

// 声明全局变量，方法1，2，3可以
var a int
var b string
var c float64

// 方法4 只能在函数体内声明

func main() {
	//方式一，声明一个变量，默认值为0
	var a int
	fmt.Println("a = ", a)
	//方式二，声明一个变量，并赋初值
	var b int = 10
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Println("bb = ", bb)
	fmt.Printf("type of bb = %T\n", bb)

	//方式三，声明变量省去数据类型的声明，默认为int，通过赋值语句自动推导数据类型
	var c = 20
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Println("cc = ", cc)
	fmt.Printf("type of cc = %T\n", cc)

	//方式四，(常用)省去var关键字，直接声明变量并赋值
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	ee := "abcd"
	fmt.Println("ee = ", ee)
	fmt.Printf("type of ee = %T\n", ee)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	//声明多个变量
	var xx, yy int
	fmt.Println("xx = ", xx)
	fmt.Println("yy = ", yy)
	var kk, ll = 100, "abcd"
	fmt.Println("kk = ", kk)
	fmt.Println("ll = ", ll)

	var (
		vv = 100
		jj = "abcd"
	)
	fmt.Println("vv = ", vv)
	fmt.Println("jj = ", jj)

}
