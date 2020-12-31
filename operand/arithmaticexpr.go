package operand

import (
	"filter4go/util"
	"fmt"
	"reflect"
)

type calcIi func(lOperand int64, rOperand int64) int64

type calcIf func(lOperand int64, rOperand float64) float64

type calcFi func(lOperand float64, rOperand int64) float64

type calcFf func(lOperand float64, rOperand float64) float64

type ArithmeticExpression struct {
	OperationExpression
	calcIi calcIi
	calcIf calcIf
	calcFi calcFi
	calcFf calcFf
}

func (e ArithmeticExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	lRet, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return nil, err
	}
	rRet, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return nil, err
	}
	il, ilerr := util.ToInt64(lRet)
	fl, flerr := util.ToFloat64(lRet)
	ir, irerr := util.ToInt64(rRet)
	fr, frerr := util.ToFloat64(rRet)
	if flerr != nil && frerr != nil {
		return e.calcIi(il, ir), nil
	} else if flerr != nil && irerr != nil {
		return e.calcIf(il, fr), nil
	} else if ilerr != nil && frerr != nil {
		return e.calcFi(fl, ir), nil
	} else if ilerr != nil && irerr != nil {
		return e.calcFf(fl, fr), nil
	}
	return nil, fmt.Errorf("Invalid operand type '%s' and '%s'", reflect.TypeOf(lRet).String(), reflect.TypeOf(rRet).String())
}

type AddExpression struct {
	ArithmeticExpression
}

func (e *AddExpression) calcIiFunc(lOperand int64, rOperand int64) int64 {
	return lOperand + rOperand
}

func (e *AddExpression) calcIfFunc(lOperand int64, rOperand float64) float64 {
	return float64(lOperand) + rOperand
}

func (e *AddExpression) calcFiFunc(lOperand float64, rOperand int64) float64 {
	return lOperand + float64(rOperand)
}

func (e *AddExpression) calcFfFunc(lOperand float64, rOperand float64) float64 {
	return lOperand + rOperand
}

func CreateAddExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr AddExpression
	expr.operandType = EXPRESSION
	expr.expressionType = ARITHMETIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calcIi = expr.calcIiFunc
	expr.calcIf = expr.calcIfFunc
	expr.calcFi = expr.calcFiFunc
	expr.calcFf = expr.calcFfFunc
	return expr, nil
}

type SubtractExpression struct {
	ArithmeticExpression
}

func (e *SubtractExpression) calcIiFunc(lOperand int64, rOperand int64) int64 {
	return lOperand - rOperand
}

func (e *SubtractExpression) calcIfFunc(lOperand int64, rOperand float64) float64 {
	return float64(lOperand) - rOperand
}

func (e *SubtractExpression) calcFiFunc(lOperand float64, rOperand int64) float64 {
	return lOperand - float64(rOperand)
}

func (e *SubtractExpression) calcFfFunc(lOperand float64, rOperand float64) float64 {
	return lOperand - rOperand
}

func CreateSubtractExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr SubtractExpression
	expr.operandType = EXPRESSION
	expr.expressionType = ARITHMETIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calcIi = expr.calcIiFunc
	expr.calcIf = expr.calcIfFunc
	expr.calcFi = expr.calcFiFunc
	expr.calcFf = expr.calcFfFunc
	return expr, nil
}

type MultiplyExpression struct {
	ArithmeticExpression
}

func (e *MultiplyExpression) calcIiFunc(lOperand int64, rOperand int64) int64 {
	return lOperand * rOperand
}

func (e *MultiplyExpression) calcIfFunc(lOperand int64, rOperand float64) float64 {
	return float64(lOperand) * rOperand
}

func (e *MultiplyExpression) calcFiFunc(lOperand float64, rOperand int64) float64 {
	return lOperand * float64(rOperand)
}

func (e *MultiplyExpression) calcFfFunc(lOperand float64, rOperand float64) float64 {
	return lOperand * rOperand
}

func CreateMultiplyExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr MultiplyExpression
	expr.operandType = EXPRESSION
	expr.expressionType = ARITHMETIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calcIi = expr.calcIiFunc
	expr.calcIf = expr.calcIfFunc
	expr.calcFi = expr.calcFiFunc
	expr.calcFf = expr.calcFfFunc
	return expr, nil
}

type DivideExpression struct {
	ArithmeticExpression
}

func (e *DivideExpression) calcIiFunc(lOperand int64, rOperand int64) int64 {
	if rOperand == 0 {
		return -1
	}
	return lOperand / rOperand
}

func (e *DivideExpression) calcIfFunc(lOperand int64, rOperand float64) float64 {
	if rOperand == 0.0 {
		return -1
	}
	return float64(lOperand) / rOperand
}

func (e *DivideExpression) calcFiFunc(lOperand float64, rOperand int64) float64 {
	if rOperand == 0.0 {
		return -1
	}
	return lOperand / float64(rOperand)
}

func (e *DivideExpression) calcFfFunc(lOperand float64, rOperand float64) float64 {
	if rOperand == 0.0 {
		return -1
	}
	return lOperand / rOperand
}

func CreateDivideExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr DivideExpression
	expr.operandType = EXPRESSION
	expr.expressionType = ARITHMETIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calcIi = expr.calcIiFunc
	expr.calcIf = expr.calcIfFunc
	expr.calcFi = expr.calcFiFunc
	expr.calcFf = expr.calcFfFunc
	return expr, nil
}

type ModularExpression struct {
	ArithmeticExpression
}

func (e *ModularExpression) calcIiFunc(lOperand int64, rOperand int64) int64 {
	return lOperand % rOperand
}

func (e *ModularExpression) calcIfFunc(lOperand int64, rOperand float64) float64 {
	return -1
}

func (e *ModularExpression) calcFiFunc(lOperand float64, rOperand int64) float64 {
	return -1
}

func (e *ModularExpression) calcFfFunc(lOperand float64, rOperand float64) float64 {
	return -1
}

func CreateModularExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr ModularExpression
	expr.operandType = EXPRESSION
	expr.expressionType = ARITHMETIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calcIi = expr.calcIiFunc
	expr.calcIf = expr.calcIfFunc
	expr.calcFi = expr.calcFiFunc
	expr.calcFf = expr.calcFfFunc
	return expr, nil
}
