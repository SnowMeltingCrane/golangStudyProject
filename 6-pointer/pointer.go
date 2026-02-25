package main

import "fmt"

func changeValue(p *int) {
	*p = 10
}

func swap(pa, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	a := 20
	b := 30
	var c int
	swap(&a, &b)

	changeValue(&c)

	fmt.Println(a, b)
	fmt.Println(c)
}
