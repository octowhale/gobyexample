// variadic functions 不定参数函数
// 函数参数个数不定

package main

import (
	"fmt"
)

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	// 以 函数名... 格式传入 切片
	sum(nums...)
}

func sum(nums ...int) {
	// 以 ...type 表示不定长度参数
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	// return total
	fmt.Println(total)
}
