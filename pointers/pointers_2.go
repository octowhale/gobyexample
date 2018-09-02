package main

import (
	"fmt"
)

func zeroval(ival int) {
	// 值拷贝传递
	fmt.Println("in zeroval, ival's mem-addr = ", &ival)
	ival = 0
}

func zeroptr(iptr *int) {
	// 内存地址传递
	fmt.Println("in zeroptr, iptr's mem-addr = ", iptr)
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial,  i = ", i)
	fmt.Println("initial,  i's mem-addr = ", &i)

	zeroval(i)
	fmt.Println("after zeroval, i = ", i)
	fmt.Println("after zeroval,  i's mem-addr = ", &i)

	zeroptr(&i)
	fmt.Println("after zeroptr, i = ", i)
	fmt.Println("after zeroptr,  i's mem-addr = ", &i)

}
