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

func Test_reflectCompare(t *testing.T) {
	compare(1)
	compare("a")
}

func Test_String(t *testing.T) {
	fmt.Println(string("abc"))
}
