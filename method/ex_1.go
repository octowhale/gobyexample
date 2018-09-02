package main

import "fmt"

type Human struct {
	name string
	city string
}

func main() {
	zhangfei := Human{"张翼德", "燕人"}
	zhangfei.Shout("俺乃" + zhangfei.city + zhangfei.name)
	zhangfei.Kill("吕布")
}

// method: shout
func (h *Human) Shout(words string) {
	fmt.Println(words)
}

// method: kill
func (h *Human) Kill(enemyName string) {
	fmt.Println(h.name, "杀死了敌将", enemyName, ", 功绩 +1")
}
