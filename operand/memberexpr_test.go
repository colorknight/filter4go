package operand

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectField(t *testing.T) {
	s := "this is a test"
	ty := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	fmt.Println("get Type is :", ty.Name())
	fmt.Println("get Value Type is :", v.Type().Name())
	fmt.Println("get Value is :", v.String())
}
