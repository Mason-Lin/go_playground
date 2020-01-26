package main

import "fmt"

func main() {
	AddFunc := add(1, 2)
	fmt.Println(AddFunc(1, 1))
	fmt.Println(AddFunc(0, 0))
	fmt.Println(AddFunc(2, 2))
}

// 闭包使用方法
func add(x1, x2 int) func(x3 int, x4 int) (int, int, int) {
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}
