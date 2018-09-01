// 分支 switch

package main

import (
	"fmt"
	"time"
)

func main() {

	// 基本用法
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 使用逗号分割多个条件,
	// 使用默认退出值

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	// switch 可以不设置条件
	// 实现 if/else 逻辑
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.
	// 类型switch比较类型而不是值。您可以使用它来发现接口值的类型。
	// 在此示例中，变量t将具有与其子句对应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
