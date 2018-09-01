// 数组 arrays

package main

import (
	"fmt"
	"log"
)

func main() {

	// 定义一个 5 个元素的 数字数组
	// 不进行初始化， 默认 0 值
	var a [5]int

	log.Println("emp: ", a)

	a[4] = 100
	log.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	// 初始化一个数组，并赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// 定义一个数组，并赋值
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	log.Println("2d: ", twoD)
}
