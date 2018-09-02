/*
练习 11.1 simple_interface.go：
定义一个接口 Simpler，它有一个 Get() 方法和一个 Set()，Get()返回一个整型值，Set() 有一个整型参数。创建一个结构体类型 Simple 实现这个接口。
*/

package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	num int
}

// method
func (sip *Simple) Get() int {
	return sip.num
}

func (sip *Simple) Set(v int) {
	sip.num = v
}

func main() {
	sip := Simple{5}
	fmt.Println(sip.Get())
	sip.Set(10)
	fmt.Println(sip.num)
	fmt.Println(sip.Get())
}
