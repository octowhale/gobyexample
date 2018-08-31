// 值类型 values
package main

import "fmt"

func main() {

	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	// int and float
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)

	// bool
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
