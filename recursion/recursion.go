// recursion 递归

package main

import "log"

func main() {

	log.Println(fact(7))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
