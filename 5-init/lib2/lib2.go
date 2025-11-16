package lib2

import "fmt"

// Lib1Test 当前lib1包提供的API
func Lib2Test() {
	fmt.Println("lib2.Lib1Test()...")
}

func init() {
	fmt.Println("lib2 init()...")
}
