// maps 映射 字典

package main

import "log"

func main() {

	// 初始化一个 map
	// key: string type
	// value: int type
	m := make(map[string]int)

	// 赋值
	m["k1"] = 13
	m["k2"] = 7
	log.Println(m)

	v1 := m["k1"]
	log.Println("v1: ", v1)

	// len
	log.Println("len: ", len(m))

	// func delete(m map[Type]Type1, key Type)
	delete(m, "k2")
	log.Println("after del , m: ", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.

	// map 取值时， 有两个返回值。 第一个是 value。 第二个是可选参数，表示key是否存在。
	//   存在, 返回 value 和 true
	// 不存在, 返回 0 值 和 false

	value1, prs1 := m["k2"]
	log.Println("optional v: ", value1, prs1)
	value2, prs2 := m["k1"]
	log.Println("optional v: ", value2, prs2)

	// 初始化并赋值
	m2 := map[string]int{"foo": 1, "bar": 2}
	log.Println("m2: ", m2)
}
