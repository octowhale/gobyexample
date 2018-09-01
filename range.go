// range 遍历

package main

import "log"

func main() {
	nums := []int{1, 3, 4, 5}

	sum := 0

	for _, num := range nums {
		sum += num
	}
	log.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			log.Println("index :", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		log.Printf("k: %v , v: %v", k, v)
	}

	for k := range kvs {
		log.Println("keys: ", k)
	}

	for _, v := range kvs {
		log.Println("values: ", v)
	}

	for i, c := range "go" {
		log.Printf("i: %v -> c: %v", i, c)
	}

}
