package operand

import (
	"filter4go/metadata"
	"fmt"
	"strings"
)

type AndExpression struct {
	OperationExpression
}

func (e AndExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.BooleanOperate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.BooleanOperate(entityMap)
	if err != nil {
		return false, err
	}
	return lValue && rValue, nil
}

func CreateAndExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr AndExpression
	expr.operandType = EXPRESSION
	expr.expressionType = LOGIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type OrExpression struct {
	OperationExpression
}

func (e OrExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.BooleanOperate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.BooleanOperate(entityMap)
	if err != nil {
		return false, err
	}
	return lValue || rValue, nil
}

func CreateOrExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr OrExpression
	expr.operandType = EXPRESSION
	expr.expressionType = LOGIC
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type NotExpression struct {
	OperationExpression
}

func (e NotExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	rValue, err := e.rOperand.BooleanOperate(entityMap)
	if err != nil {
		return false, err
	}
	return !rValue, nil
}

func CreateNotExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 1 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr NotExpression
	expr.operandType = EXPRESSION
	expr.expressionType = LOGIC
	expr.operator = operator
	expr.rOperand = operands[0]
	expr.buildNameString()
	return expr, nil
}

var logicFuncMap = map[string]CreateExpression{
	metadata.AND: CreateAndExpression,
	metadata.OR:  CreateOrExpression,
	metadata.NOT: CreateNotExpression,
}

func CreateLogic(operator string, operands []Operand) (Operand, error) {
	createExpression := logicFuncMap[strings.ToLower(operator)]
	if createExpression == nil {
		return nil, fmt.Errorf("Unsupported logic operator '%s'!", operator)
	}
	return createExpression(operator, operands)
}

func IsLogicOperator(operator string) bool {
	operator = strings.ToLower(operator)
	switch operator {
	case metadata.AND:
	case metadata.OR:
	case metadata.NOT:
		return true
	}
	return false
}
