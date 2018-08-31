// 变量 variables
package main

import "fmt"

func main() {

	//
	var a = "initial"
	fmt.Println(a)

	// declare multiple
	var b, c int = 1, 2
	fmt.Printf("b = %v, c=%v", b, c)

	// infer value
	var d = true
	fmt.Println(d)

	// zero value
	var e int
	fmt.Println(e)

	// :=
	f := "short"
	fmt.Println(f)
}
