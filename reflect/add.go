package reflectdemo

import (
	"fmt"
	"reflect"
)

func Add(t1 interface{}, t2 interface{}) (ret interface{}) {
	t1Type := reflect.TypeOf(t1)
	t2Type := reflect.TypeOf(t2)

	if t1Type.Name() == t2Type.Name() {
		if t1Type.Name() == reflect.Int.String() {
			return reflect.ValueOf(t1).Interface().(int) + reflect.ValueOf(t2).Interface().(int)
		} else if t1Type.Name() == reflect.String.String() {
			return fmt.Sprintf("%s%s", reflect.ValueOf(t1).Interface().(string), reflect.ValueOf(t2).Interface().(string))
		}
	} else {
		return fmt.Errorf("类型不匹配")
	}
	return
}

func compare(t1 interface{}) {
	t1type := reflect.ValueOf(t1).Type()
	t2type := reflect.ValueOf(2).Type()
	if t1type == t2type {
		fmt.Printf("%v 是 int 类型\n", reflect.TypeOf(t1).Name())
	} else {
		fmt.Printf("%v 不是 int 类型\n", reflect.TypeOf(t1).Name())
	}

}
