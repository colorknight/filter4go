package oper

import (
	"fmt"
	"testing"
)

type st struct {
	Name string
}

func (p st) ShowName() {
	fmt.Printf("name is %s\n", p.Name)
}

func buildData(id int) map[string]interface{} {
	mp := map[string]interface{}{
		"num":  id,
		"name": "test",
		"ary":  []string{"a1", "a2"},
		"bean": st{"bean1"},
		"map": map[string]int{
			"id": id,
		},
	}
	return mp
}

func TestMemberFieldOperand(t *testing.T) {
	s1 := "bean.Name"
	operand := ParseOperand(s1)
	mp := buildData(1)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", result)
	mp = buildData(2)
	result, err = operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", result)
}

func TestMemberMethodOperand(t *testing.T) {
	s1 := "bean.ShowName()"
	operand := ParseOperand(s1)
	mp := buildData(1)
	_, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestAddIntOperand(t *testing.T) {
	s1 := "num + 5"
	operand := ParseOperand(s1)
	mp := buildData(1)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := result.(int64)
	if n != 6 {
		t.Error("invalid result")
	}
}

func TestAddFloatOperand(t *testing.T) {
	s1 := "5.1 - num"
	operand := ParseOperand(s1)
	mp := buildData(1)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := result.(float64)
	if n != 4.1 {
		t.Error("invalid result")
	}
}

func TestMultiplicativeOperand(t *testing.T) {
	s1 := "num * 10 / 2 % 3"
	operand := ParseOperand(s1)
	mp := buildData(1)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := result.(int64)
	if n != 2 {
		t.Error("invalid result")
	}
}

func TestLeftShiftOperand(t *testing.T) {
	s1 := "num << 2"
	operand := ParseOperand(s1)
	mp := buildData(1)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := result.(int64)
	if n != 4 {
		t.Error("invalid result")
	}
}

func TestRightShiftOperand(t *testing.T) {
	s1 := "num >> 1"
	operand := ParseOperand(s1)
	mp := buildData(4)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := result.(int64)
	if n != 2 {
		t.Error("invalid result")
	}
}

func TestWiseOrOperand(t *testing.T) {
	s1 := "num | 2"
	operand := ParseOperand(s1)
	mp := buildData(4)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := result.(int64)
	if n != 6 {
		t.Error("invalid result")
	}
}

func TestArrayOperand(t *testing.T) {
	s1 := "ary[1]"
	operand := ParseOperand(s1)
	mp := buildData(4)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := result.(string)
	if s != "a2" {
		t.Error("invalid result")
	}
}

func TestMapOperand(t *testing.T) {
	s1 := "map['id']"
	operand := ParseOperand(s1)
	mp := buildData(4)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := result.(int)
	if s != 4 {
		t.Error("invalid result")
	}
}

func TestSprintfOperand(t *testing.T) {
	s1 := "sprintf('%s:%d', ary[1] ,map['id'])"
	operand := ParseOperand(s1)
	mp := buildData(4)
	result, err := operand.Operate(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := result.(string)
	if s != "a2:4" {
		t.Error("invalid result")
	}
}

func TestOperandsOperand(t *testing.T) {
	s1 := "(50,80)"
	operand := ParseOperand(s1)
	if operand == nil {
		t.Error("invalid result")
	}
}

