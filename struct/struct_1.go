package main

import "fmt"

type named_person struct {
	Name string
	Age  int
}

func main() {
	tang := named_person{"tangxin", 20}
	fmt.Println(tang)

	tang2 := named_person{Name: "tangxin", Age: 30}
	fmt.Println(tang2)
}
