package main

import "fmt"

func main() {
	// 第一种声明方式
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用前需要进行内存空间分配
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "golang"

	fmt.Println(myMap1)

	// 第二种声明方式
	myMap2 := make(map[int]string)

	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "golang"
	fmt.Println(myMap2)

	// 第三种声明方式
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "python",
		"three": "golang",
	}
	fmt.Println(myMap3)
}
