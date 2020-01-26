package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{1, -3, 5, -2, 4, -6}

	c1 := make(chan int)
	c2 := make(chan int)
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)
	x, y := <-c1, <-c2 // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
