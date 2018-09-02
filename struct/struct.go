// struct 结构

package main

import "fmt"

type person struct {
	// 命名字段
	Name string
	Age  int
}

type person1 struct {
	// 命名字段
	Name string
	Age  int
}

type person_2 struct {
	// 匿名字段
	string
	int
}

func main() {

	// 命名字段
	fmt.Println(person{Name: "zhangfei", Age: 10})
	fmt.Println(person{Age: 20, Name: "guanyu"})

	// 不指定字段名称时，按值类型顺序
	fmt.Println(person{"liubei", 30})
	// fmt.Println(person{30, "wangwu"})

	// 0 值
	// 不给字段赋值的时候，默认使用 0 值
	zhaoyun := person{Name: "zhaoyun"}
	fmt.Println(zhaoyun)
	zhaoyun.Age = 30
	fmt.Println("after birthday: zhaoyun is ", zhaoyun)

	// 取内存地址
	fmt.Println(&person{"huangzong", 80})

	// 结构是值传递
	// 如果要改变值的话，可以通过指针进行操作
	machao := &person{Name: "machao", Age: 31}
	fmt.Println(machao)
	machao.Age = 32
	fmt.Println("after birthday: zhaoyun is ", machao)

	// 类型的比较
	caocao := person{Name: "caocao", Age: 35}

	// 一千个人眼中有一千哈姆雷特
	// 不同结构类型之间没有可比性，会报错。 即使看起来都差不多，甚至一样。
	// caocao1 := person1{Name: "caocao", Age: 35}
	// fmt.Println(caocao == caocao1)
	// invalid operation: caocao == caocao1 (mismatched types person and person1)

	// 年年岁岁花相似，岁岁年年人不同
	// 同一个结构类型定义的两个类型可以比较。
	caocao35 := person{Name: "caocao", Age: 35}
	fmt.Println(caocao == caocao35) // true
	caocao36 := person{Name: "caocao", Age: 36}
	fmt.Println(caocao35 == caocao36) // false
}
