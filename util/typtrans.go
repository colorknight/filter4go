package util

import (
	"fmt"
	"reflect"
)

type OperandDataType int

const (
	UNKNOWN OperandDataType = iota
	BOOL
	INT
	FLOAT
	STR
)

var typeDataTypeMap = map[reflect.Kind]OperandDataType{
	reflect.Bool:    BOOL,
	reflect.Int:     INT,
	reflect.Int8:    INT,
	reflect.Uint8:   INT,
	reflect.Int16:   INT,
	reflect.Uint16:  INT,
	reflect.Int32:   INT,
	reflect.Uint32:  INT,
	reflect.Int64:   INT,
	reflect.Uint64:  INT,
	reflect.Float32: FLOAT,
	reflect.Float64: FLOAT,
	reflect.String:  STR,
}

func GetOperandDataType(operand interface{}) (OperandDataType, error) {
	tp := reflect.TypeOf(operand)
	k := tp.Kind()
	dataType, ok := typeDataTypeMap[k]
	if !ok {
		return UNKNOWN, fmt.Errorf("Unsupported data type %s", dataType)
	}
	return dataType, nil
}

func ToInt64(operand interface{}) (int64, error) {
	tp := reflect.TypeOf(operand)
	k := tp.Kind()
	dataType, ok := typeDataTypeMap[k]
	if dataType != INT || !ok {
		return 0, fmt.Errorf("Unsupported data type %s", dataType)
	}
	switch k {
	case reflect.Int:
		i := operand.(int)
		return int64(i), nil
	case reflect.Int8:
		i := operand.(int8)
		return int64(i), nil
	case reflect.Int16:
		i := operand.(int16)
		return int64(i), nil
	case reflect.Int32:
		i := operand.(int32)
		return int64(i), nil
	case reflect.Uint:
		i := operand.(uint)
		return int64(i), nil
	case reflect.Uint8:
		i := operand.(uint8)
		return int64(i), nil
	case reflect.Uint16:
		i := operand.(uint16)
		return int64(i), nil
	case reflect.Uint32:
		i := operand.(uint32)
		return int64(i), nil
	case reflect.Uint64:
		i := operand.(uint64)
		return int64(i), nil
	default:
		return operand.(int64), nil
	}
}

func ToFloat64(operand interface{}) (float64, error) {
	tp := reflect.TypeOf(operand)
	k := tp.Kind()
	dataType, ok := typeDataTypeMap[k]
	if dataType != FLOAT || !ok {
		return 0, fmt.Errorf("Unsupported data type %s", dataType)
	}
	switch k {
	case reflect.Float32:
		f := operand.(float32)
		return float64(f), nil
	default:
		return operand.(float64), nil
	}
}
