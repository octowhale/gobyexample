// closure 闭包

package main

import "fmt"

// func() int 是一个整体，是一个匿名函数
// intSeq 函数返回的是一个匿名函数，
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println("nextInt = ", nextInt())
	fmt.Println("nextInt = ", nextInt())
	fmt.Println("nextInt = ", nextInt())

	newInts := intSeq()
	fmt.Println("newInts = ", newInts())

	fmt.Println("nextInt = ", nextInt())

	// nextInt 和 newInts 都指向了各自的函数内存地址。 因此，他们的 i 不共用。
}
