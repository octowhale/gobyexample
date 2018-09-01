// 选择 if else

package main

import (
	"fmt"
	"log"
)

func main() {

	// basic
	if 7%2 == 0 {
		log.Println("7 is even")
	} else {
		log.Println("7 is odd")
	}

	// if without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// multipule if-else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")

	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
