// functions 函数

package main

import "log"

// 函数的位置可以任意
// 形参类型相同，可以简写
func minus(a, b int) int {
	return a - b
}

func main() {
	// plus
	log.Println("1 + 2 = ", plus(1, 2))

	// minus
	log.Println("1 - 2 = ", minus(1, 2))
}

// 函数的位置可以任意
func plus(a int, b int) int {
	return a + b
}
