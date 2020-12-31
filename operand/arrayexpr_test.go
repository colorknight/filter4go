package operand

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAryTransType(t *testing.T) {
	ary := []string{"a", "b"}
	var a interface{} = ary
	typ := reflect.TypeOf(a)
	var b interface{} = []interface{}{}
	if typ.ConvertibleTo(reflect.TypeOf(b)) {
		fmt.Println("It's ok!")
	} else {
		fmt.Println("It's not ok!")
	}
}

func TestAryAccess(t *testing.T) {
	ary := []string{"a", "b"}
	var a interface{} = ary
	v := reflect.ValueOf(a)
	for i := 0; i < v.Len(); i++ {
		fmt.Println(v.Index(i))
	}
}

func TestMapAccess(t *testing.T) {
	ary := map[string]string{"n1": "a1", "n2": "a2"}
	var a interface{} = ary
	v := reflect.ValueOf(a)
	keys := v.MapKeys()
	for _, key := range keys {
		fmt.Println(v.MapIndex(key))
	}
}
