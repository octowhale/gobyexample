// 常量 constant
package main

import (
	"fmt"
	"log"
	"math"
)

// 常量可以在任意位置定义
const s string = "constant"

func main() {

	log.Println("https://gobyexample.com/constants")
	log.Println("http://www.runoob.com/go/go-constants.html")

	fmt.Println(s)

	const n = 500000
	const b = true
	fmt.Println(n)
	log.Println(b)

	// 常量表达式以任意精度执行算术。
	const d = 3e20 / n
	fmt.Println(d)

	// 数字常量在给定之前没有类型，例如通过显式强制转换
	fmt.Println(int64(n))
	fmt.Println(float64(n))

	// 常量可以作为传参
	fmt.Println(math.Sin(n))

	f_iota()
}

func f_iota() {

}
