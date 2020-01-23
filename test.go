// keep going https://www.runoob.com/go/go-constants.html
package main

import "fmt"

var (
	glo_a int
	glo_b bool
)

func main() {
	var loc_a int = 1
	co_a, co_b := 2, 3
	fmt.Println("Hello", glo_a, glo_b, loc_a, co_a, co_b)
}
