package util

import (
	"fmt"
	"strconv"
	"strings"
)

func Compare(lOperand interface{}, rOperand interface{}) (int, error) {
	if lOperand == nil && rOperand == nil {
		return 0, nil
	}
	if lOperand == nil || rOperand == nil {
		return 0, fmt.Errorf("One of parameters is nil!")
	}
	lt, err := GetOperandDataType(lOperand)
	if err != nil {
		return 0, err
	}
	rt, err := GetOperandDataType(rOperand)
	if err != nil {
		return 0, err
	}
	return branchCompare(lOperand, lt, rOperand, rt)
}

func branchCompare(lOperand interface{}, lt OperandDataType, rOperand interface{}, rt OperandDataType) (int, error) {
	if lt != rt {
		if lt == STR || rt == STR {
			if lt == STR && rt == BOOL {
				return strAndBoolCompare(lOperand, rOperand)
			} else if lt == STR && rt == INT {
				return strAndIntCompare(lOperand, rOperand)
			} else if lt == STR && rt == FLOAT {
				return strAndFloatCompare(lOperand, rOperand)
			} else if lt == BOOL && rt == STR {
				return boolAndStrCompare(lOperand, rOperand)
			} else if lt == INT && rt == STR {
				return intAndStrCompare(lOperand, rOperand)
			} else if lt == FLOAT && rt == STR {
				return floatAndStrCompare(lOperand, rOperand)
			}
		}
		if lt == INT && rt == FLOAT {
			return compareFloat64(lOperand.(float64), rOperand.(float64))
		} else if lt == FLOAT && rt == INT {
			return compareFloat64(lOperand.(float64), rOperand.(float64))
		}
	} else {
		if lt == BOOL {
			return compareBool(lOperand.(bool), rOperand.(bool))
		} else if lt == INT {
			li, _ := ToInt64(lOperand)
			ri, _ := ToInt64(rOperand)
			return compareInt64(li, ri)
		} else if lt == FLOAT {
			lf, _ := ToFloat64(lOperand)
			rf, _ := ToFloat64(rOperand)
			return compareFloat64(lf, rf)
		} else if lt == STR {
			return compareString(lOperand.(string), rOperand.(string))
		}
	}
	return -1, fmt.Errorf("Unsupported comparable data types!")
}

func strAndBoolCompare(lOperand interface{}, rOperand interface{}) (int, error) {
	ls := lOperand.(string)
	lb, err := strconv.ParseBool(ls)
	if err != nil {
		return -1, err
	}
	rb := rOperand.(bool)
	return compareBool(lb, rb)
}

func boolAndStrCompare(lOperand interface{}, rOperand interface{}) (int, error) {
	lb := lOperand.(bool)
	rs := rOperand.(string)
	rb, err := strconv.ParseBool(rs)
	if err != nil {
		return -1, err
	}
	return compareBool(lb, rb)
}

func strAndIntCompare(lOperand interface{}, rOperand interface{}) (int, error) {
	ls := lOperand.(string)
	i, err := strconv.Atoi(ls)
	li := int64(i)
	if err != nil {
		return -1, err
	}
	ri, err := ToInt64(rOperand)
	if err != nil {
		return -1, err
	}
	return compareInt64(li, ri)
}

func intAndStrCompare(lOperand interface{}, rOperand interface{}) (int, error) {
	li := lOperand.(int64)
	rs := rOperand.(string)
	ri, err := strconv.Atoi(rs)
	if err != nil {
		return -1, err
	}
	return compareInt64(li, int64(ri))
}

func strAndFloatCompare(lOperand interface{}, rOperand interface{}) (int, error) {
	ls := lOperand.(string)
	lf, err := strconv.ParseFloat(ls, 64)
	if err != nil {
		return -1, err
	}
	rf, err := ToFloat64(rOperand)
	if err != nil {
		return -1, err
	}
	return compareFloat64(lf, rf)
}

func floatAndStrCompare(lOperand interface{}, rOperand interface{}) (int, error) {
	lf := lOperand.(float64)
	rs := rOperand.(string)
	rf, err := strconv.ParseFloat(rs, 64)
	if err != nil {
		return -1, err
	}
	return compareFloat64(lf, rf)
}

func compareInt64(li int64, ri int64) (int, error) {
	if li == ri {
		return 0, nil
	} else if li > ri {
		return 1, nil
	}
	return -1, nil
}

func compareFloat64(lf float64, rf float64) (int, error) {
	if lf == rf {
		return 0, nil
	} else if lf > rf {
		return 1, nil
	}
	return -1, nil
}

func compareBool(lb bool, rb bool) (int, error) {
	if lb == rb {
		return 0, nil
	} else if lb == true {
		return 1, nil
	}
	return -1, nil
}

func compareString(ls string, rs string) (int, error) {
	return strings.Compare(ls, rs), nil
}
