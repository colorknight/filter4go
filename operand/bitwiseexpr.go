package operand

import (
	"filter4go/util"
	"fmt"
	"reflect"
)

type calc func(lOperand int64, rOperand int64) int64

type BitwiseExpression struct {
	OperationExpression
	calc calc
}

func (e BitwiseExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	lRet, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return nil, err
	}
	rRet, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return nil, err
	}
	il, ilerr := util.ToInt64(lRet)
	ir, irerr := util.ToInt64(rRet)
	if ilerr != nil || irerr != nil {
		return nil, fmt.Errorf("Invalid operand type '%s' and '%s'", reflect.TypeOf(lRet).String(), reflect.TypeOf(rRet).String())
	}
	return e.calc(il, ir), nil
}

type BitwiseAndExpression struct {
	BitwiseExpression
}

func (e *BitwiseAndExpression) calcFunc(lOperand int64, rOperand int64) int64 {
	return lOperand & rOperand
}

func CreateBitwiseAndExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr BitwiseAndExpression
	expr.operandType = EXPRESSION
	expr.expressionType = BITWISE
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calc = expr.calcFunc
	return expr, nil
}

type BitwiseOrExpression struct {
	BitwiseExpression
}

func (e *BitwiseOrExpression) calcFunc(lOperand int64, rOperand int64) int64 {
	return lOperand | rOperand
}

func CreateBitwiseOrExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr BitwiseOrExpression
	expr.operandType = EXPRESSION
	expr.expressionType = BITWISE
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calc = expr.calcFunc
	return expr, nil
}

type BitwiseXorExpression struct {
	BitwiseExpression
}

func (e *BitwiseXorExpression) calcFunc(lOperand int64, rOperand int64) int64 {
	return lOperand ^ rOperand
}

func CreateBitwiseXorExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr BitwiseXorExpression
	expr.operandType = EXPRESSION
	expr.expressionType = BITWISE
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calc = expr.calcFunc
	return expr, nil
}

type BitwiseNotExpression struct {
	BitwiseExpression
}

func (e *BitwiseNotExpression) calcFunc(lOperand int64, rOperand int64) int64 {
	return ^rOperand
}

func CreateBitwiseNotExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 1 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr BitwiseNotExpression
	expr.operandType = EXPRESSION
	expr.expressionType = BITWISE
	expr.operator = operator
	expr.rOperand = operands[0]
	expr.buildNameString()
	expr.calc = expr.calcFunc
	return expr, nil
}

type LeftShiftExpression struct {
	BitwiseExpression
}

func (e *LeftShiftExpression) calcFunc(lOperand int64, rOperand int64) int64 {
	return lOperand << rOperand
}

func CreateLeftShiftExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr LeftShiftExpression
	expr.operandType = EXPRESSION
	expr.expressionType = BITWISE
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calc = expr.calcFunc
	return expr, nil
}

type RightShiftExpression struct {
	BitwiseExpression
}

func (e *RightShiftExpression) calcFunc(lOperand int64, rOperand int64) int64 {
	return lOperand >> rOperand
}

func CreateRightShiftExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr RightShiftExpression
	expr.operandType = EXPRESSION
	expr.expressionType = BITWISE
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	expr.calc = expr.calcFunc
	return expr, nil
}
