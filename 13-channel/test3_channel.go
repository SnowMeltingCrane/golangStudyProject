package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//
	//for {
	//	if date, ok := <-c; ok {
	//		fmt.Println(date)
	//	} else {
	//		break
	//	}
	//}

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Finished...")
}
