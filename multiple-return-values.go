// function 多返回值

package main

import (
	"fmt"
	"log"
)

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	m, n := name_return_1()
	log.Printf("m = %v -- n = %v", m, n)

	m, n = name_return_2()
	log.Printf("m = %v -- n = %v", m, n)
}

func vals() (int, int) {
	// 无顺序返回
	return 3, 7
}
func name_return_1() (x int, y int) {
	// 有顺序返回

	// x, y 已经在定义函数时指定了
	// no new variables on left side of :=
	// x := 1

	x = 1
	y = 2

	// 函数返回的时候，已经指定了顺序，因此直接用 return 将按照定义的变量名顺序
	return
}

func name_return_2() (x int, y int) {
	// 有顺序返回
	x = 1
	y = 2

	// 函数返回的时候, 也可以打乱顺序
	return y, x
}
