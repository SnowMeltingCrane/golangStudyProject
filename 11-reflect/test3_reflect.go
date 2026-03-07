package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type :", reflect.TypeOf(arg))
	fmt.Println("value :", reflect.ValueOf(arg))
}

func main() {
	var num = 1.123
	reflectNum(num)
}
