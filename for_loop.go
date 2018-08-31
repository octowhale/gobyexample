// 循环 for

/*
	for 是 golang
*/
package main

import "log"

func main() {
	i := 1

	// while
	for i <= 3 {
		log.Println(i)
		i++
	}

	// 三段式
	for j := 7; 7 <= 10; i++ {
		log.Println(7)
	}

	// 死循环
	/*
		for {
			log.Println("loop")
		}
	*/

	// break

	for {
		log.Println("loop")
		break
	}

	// continue
	for n := 1; n <= 5; n++ {
		if n == 3 {
			log.Println("loop")
			continue
		}
		log.Println(n)
	}
}
