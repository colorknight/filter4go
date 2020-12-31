package operand

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTransType(t *testing.T) {
	var a int64 = 123
	var b float64 = 23.2

	c := float64(a)
	fmt.Printf("%f\n", c)
	d := int64(b)
	fmt.Printf("%d\n", d)

	var a1 interface{} = a
	var b1 interface{} = 23.2
	var e int8 = 3
	c1, ok := a1.(float64)
	fmt.Printf("%f,%t\n", c1, ok)
	d1, ok := b1.(int64)
	fmt.Printf("%d,%t\n", d1, ok)

	c2, ok := a1.(int64)
	fmt.Printf("%d,%t\n", c2, ok)
	d2, ok := b1.(float64)
	fmt.Printf("%f,%t\n", d2, ok)
	ty := reflect.TypeOf(a1)
	fmt.Println("get Type is :", ty.Name())
	fmt.Println(a - int64(e))
	ty = reflect.TypeOf(23)
	fmt.Println("get Type is :", ty.Name())
	//var p int64 = 3
	//var q uint64 = 5
	//fmt.Println(q > p)

}

func TestTransInt(t *testing.T) {
	var a int = 32
	var b interface{} = a
	a1, ok := b.(int)
	a2 := int64(a1)
	fmt.Printf("%d,%t\n", a1, ok)
	fmt.Printf("%d\n", a2)
	ss := fmt.Sprintf("%s:%d", "a2", 1)
	fmt.Println(ss)
}