// pointer 指针

package main

import "fmt"

func main() {

	fmt.Println("the-way-to-go: \n\t https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.9.md")
	var i1 = 5

	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

	// 这个地址可以存储在一个叫做指针的特殊数据类型中，在本例中这是一个指向 int 的指针，即 i1：此处使用 *int 表示。
	// 如果我们想调用指针 intP，我们可以这样声明它：
	var intP *int

	// 使用 intP = &i1 是合法的，此时 intP 指向 i1
	// intP 存储了 i1 的内存地址；它指向了 i1 的位置，它引用了变量 i1。
	intP = &i1

	// 指针的格式化标识符为 %p
	// 你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。
	// 使用一个指针引用一个值被称为间接引用。
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
