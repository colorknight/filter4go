package operand

import (
	"filter4go/metadata"
	"fmt"
	"strings"
)

type ExpressionType int

const (
	PAREN ExpressionType = iota
	OPERANDS
	LOGIC
	RELATION
	ARITHMETIC
	BITWISE
	ARRAY
	MEMBER
)

type Expression struct {
	BaseOperand
	expressionType ExpressionType
}

func (e Expression) GetExpressionType() ExpressionType {
	return e.expressionType
}

func (e *Expression) buildNameString() string {
	return e.name
}

type ParenExpression struct {
	Expression
	operand Operand
}

func (e *ParenExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteByte(metadata.LPAREN)
	bd.WriteString(e.operand.ToString())
	bd.WriteByte(metadata.RPAREN)
	e.name = bd.String()
}

func (e ParenExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	return e.operand.Operate(entityMap)
}

func (e ParenExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	return e.operand.BooleanOperate(entityMap)
}

type OperandsExpression struct {
	Expression
	operands []Operand
}

func (e *OperandsExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteByte(metadata.LPAREN)
	for i, o := range e.operands {
		if i != 0 {
			bd.WriteByte(',')
		}
		bd.WriteString(o.ToString())
	}
	bd.WriteByte(metadata.RPAREN)
	e.name = bd.String()
}

func (e OperandsExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	return e.operands, nil
}

type OperatorType int

const (
	UNARY OperatorType = iota
	BINARY
)

type OperationExpression struct {
	Expression
	operatorType OperatorType
	operator     string
	lOperand     Operand
	rOperand     Operand
}

func (e OperationExpression) GetOperatorType() OperatorType {
	return e.operatorType
}

func (e OperationExpression) GetOperator() string {
	return e.operator
}

func (e OperationExpression) GetLeftOperand() Operand {
	return e.lOperand
}

func (e OperationExpression) GetRightOperand() Operand {
	return e.rOperand
}

func (e *OperationExpression) buildNameString() {
	var bd strings.Builder
	if e.lOperand != nil {
		bd.WriteString(e.lOperand.ToString())
		bd.WriteByte(metadata.BLANKSPACE)
		bd.WriteString(e.operator)
		bd.WriteByte(metadata.BLANKSPACE)
		bd.WriteString(e.rOperand.ToString())
	} else {
		bd.WriteString(e.operator)
		bd.WriteByte(metadata.LPAREN)
		bd.WriteString(e.rOperand.ToString())
		bd.WriteByte(metadata.RPAREN)
	}
	e.name = bd.String()
}

type CreateExpression func(operator string, operands []Operand) (Operand, error)

func CreateParenExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 1 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr ParenExpression
	expr.operandType = EXPRESSION
	expr.expressionType = PAREN
	expr.operand = operands[0]
	expr.buildNameString()
	return expr, nil
}

func CreateOperandsExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr OperandsExpression
	expr.operandType = EXPRESSION
	expr.expressionType = OPERANDS
	expr.operands = operands
	expr.buildNameString()
	return expr, nil
}
