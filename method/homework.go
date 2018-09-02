// 根据为结构增加方法的知识， 尝试声明一个底层类型为 int 的类型，并实现调用该方法就自增 100
// 例如， a.Increase() 之后，  a 从 0 变成 100
package main

import "fmt"

type TZ int

func main() {

	var tz TZ
	// tz.Increase(100)
	// tz.Increase(100)
	// tz.Increase(100)

	fmt.Println(tz.Increase(100))
	fmt.Println(tz.Increase(200))
	fmt.Println(tz.Increase(300))

	fmt.Println("func increase: ", increase(100))
	fmt.Println("func increase: ", increase(100))
	fmt.Println("func increase: ", increase(100))
}

// method 方法
// func (接收者名称 接收者类型) 方法名(方法参数 方法类型, 方法参数2 方法类型) (返回参数 返回类型, 返回参数2 返回类型)
func (tz *TZ) Increase(num int) (val TZ) {
	*tz += TZ(num)
	// fmt.Println(*tz)
	return *tz
}

// function 函数
// func 函数名(函数参数 函数类型, 函数参数2 函数类型) (返回参数 返回类型, 返回参数2 返回类型)
func increase(num int) (val int) {
	val += num
	return val
}
