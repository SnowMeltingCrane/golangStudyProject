package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 运行中")

		c <- 666
	}()

	num := <-c

	fmt.Println("num=", num)

	fmt.Println("main 结束")
}
