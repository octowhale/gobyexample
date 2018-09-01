// 切片 slices

package main

import (
	"fmt"
	"log"
)

func main() {

	// 使用 make 定义一个切片
	s := make([]string, 3)
	log.Println("emp: ", s)

	// 切片赋值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	log.Println("s: ", s)
	log.Println("len(s): ", len(s))

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片copy
	c := make([]string, len(s))

	// func Copy(dst Writer, src Reader)
	copy(c, s)
	log.Println("cpy c: ", c)

	// 切片的切片操作
	l := s[2:5]
	log.Println("sl1: ", l)
	l = s[:5]
	log.Println("sl2: ", l)
	l = s[2:]
	log.Println("sl3: ", l)

	// 定义并初始化切片
	t := []string{"p", "q", "m"}
	log.Println("t: ", t)

	// 多维切片，
	// 与多维数组不同， 内部切片可以变长
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = j
		}
	}
	log.Println(twoD)
}
