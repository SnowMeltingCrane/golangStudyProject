package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程写入数据：", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("主go程读取数据：", num, "len(c) = ", len(c), "cap(c) = ", cap(c))
	}

	// 防止子go程还没结束 主go程就结束
	time.Sleep(1 * time.Second)

	fmt.Println("main goroutine 结束")
}
