// keep going https://www.runoob.com/go/go-constants.html
package main

import "fmt"

func main() {
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			fmt.Printf("%1d x %1d = %2d | ", x, y, x*y)
		}
		fmt.Printf("\n")
	}
}
