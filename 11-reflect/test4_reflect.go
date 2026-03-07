package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user call")
	fmt.Println("Name:", this.Name, "Age:", this.Age)
}

func main() {
	user := User{"张三", 20}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType:", inputType)
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue:", inputValue)

	// 获取input的成员变量
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取input的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}

}
