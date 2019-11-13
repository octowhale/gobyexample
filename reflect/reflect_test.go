package reflectdemo

import (
	"fmt"
	"testing"
)

func Test_reflectdemo(t *testing.T) {
	fmt.Println(Add(1, 1))
	fmt.Println(Add("a", "b"))
	fmt.Println(Add(1, "b"))
}
