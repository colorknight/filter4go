package operand

import (
	"filter4go/metadata"
	"filter4go/util"
	"fmt"
	"regexp"
	"strings"
)

type BetweenExpression struct {
	OperationExpression
	rOperands []Operand
}

func (e BetweenExpression) GetOperands() []Operand {
	return e.rOperands
}

func (e *BetweenExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteString(e.lOperand.ToString())
	bd.WriteByte(metadata.BLANKSPACE)
	bd.WriteString(e.operator)
	bd.WriteByte(metadata.BLANKSPACE)
	bd.WriteString(e.rOperands[0].ToString())
	bd.WriteByte(metadata.BLANKSPACE)
	bd.WriteString(metadata.AND)
	bd.WriteByte(metadata.BLANKSPACE)
	bd.WriteString(e.rOperands[1].ToString())
	e.name = bd.String()
}

func (e BetweenExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rLValue, err := e.rOperands[0].Operate(entityMap)
	if err != nil {
		return false, err
	}
	rGValue, err := e.rOperands[1].Operate(entityMap)
	if err != nil {
		return false, err
	}
	if lValue == rLValue || lValue == rGValue {
		return true, nil
	}
	c, err := util.Compare(rLValue, lValue)
	if err != nil {
		return false, err
	}
	if c > 0 {
		return false, nil
	}
	c, err = util.Compare(lValue, rGValue)
	if err != nil {
		return false, err
	}
	return c <= 0, nil
}

func CreateBetweenExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr BetweenExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	operandsExpr, ok := operands[1].(OperandsExpression)
	if !ok {
		return nil, fmt.Errorf("Invalid operands parameter!")
	}
	expr.rOperands = operandsExpr.operands
	expr.buildNameString()
	return expr, nil
}

type EqualExpression struct {
	OperationExpression
}

func (e EqualExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	c, err := util.Compare(lValue, rValue)
	if err != nil {
		return false, err
	}
	return c == 0, nil
}

func CreateEqualExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr EqualExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	//e := interface{}(expr).(Operand)
	return expr, nil
}

type GreatAndEqualExpression struct {
	OperationExpression
}

func (e GreatAndEqualExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	c, err := util.Compare(lValue, rValue)
	if err != nil {
		return false, err
	}
	return c >= 0, nil
}

func CreateGreatAndEqualExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr GreatAndEqualExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type GreatThanExpression struct {
	OperationExpression
}

func (e GreatThanExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	c, err := util.Compare(lValue, rValue)
	if err != nil {
		return false, err
	}
	return c > 0, nil
}

func CreateGreatThanExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr GreatThanExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type LittleAndEqualExpression struct {
	OperationExpression
}

func (e LittleAndEqualExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	c, err := util.Compare(lValue, rValue)
	if err != nil {
		return false, err
	}
	return c <= 0, nil
}

func CreateLittleAndEqualExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr LittleAndEqualExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type LittleThanExpression struct {
	OperationExpression
}

func (e LittleThanExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	c, err := util.Compare(lValue, rValue)
	if err != nil {
		return false, err
	}
	return c < 0, nil
}

func CreateLittleThanExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr LittleThanExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type NotEqualExpression struct {
	OperationExpression
}

func (e NotEqualExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	rValue, err := e.rOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	c, err := util.Compare(lValue, rValue)
	if err != nil {
		return false, err
	}
	return c != 0, nil
}

func CreateNotEqualExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr NotEqualExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type InExpression struct {
	OperationExpression
	rOperands []Operand
}

func (e InExpression) GetOperands() []Operand {
	return e.rOperands
}

func (e *InExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteString(e.lOperand.ToString())
	bd.WriteByte(metadata.BLANKSPACE)
	bd.WriteString(e.operator)
	bd.WriteByte(metadata.BLANKSPACE)
	bd.WriteByte(metadata.LPAREN)
	for i, r := range e.rOperands {
		if i != 0 {
			bd.WriteByte(metadata.COMMA)
		}
		bd.WriteString(r.ToString())
	}
	bd.WriteByte(metadata.RPAREN)
	e.name = bd.String()
}

func (e InExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	for _, r := range e.rOperands {
		rValue, err := r.Operate(entityMap)
		if err != nil {
			return false, err
		}
		c, err := util.Compare(lValue, rValue)
		if err == nil && c == 0 {
			return true, nil
		}
	}
	return false, nil
}

func CreateInExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr InExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	operandsExpr, ok := operands[1].(OperandsExpression)
	if !ok {
		return nil, fmt.Errorf("Invalid operands parameter!")
	}
	expr.rOperands = operandsExpr.operands
	expr.buildNameString()
	return expr, nil
}

type IsExpression struct {
	OperationExpression
}

func (e IsExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	return lValue == nil, nil
}

func CreateIsExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr LikeExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildNameString()
	return expr, nil
}

type LikeExpression struct {
	OperationExpression
	reg *regexp.Regexp
}

func (e *LikeExpression) buildRegExp() error {
	rValue, err := e.rOperand.Operate(nil)
	if err != nil {
		return err
	}
	rs, ok := rValue.(string)
	if !ok {
		return fmt.Errorf("It's not a string!")
	}
	rs, err = e.translatePattern2Regex(rs)
	if err != nil {
		return err
	}
	e.reg, err = regexp.Compile(rs)
	if err != nil {
		return err
	}
	return nil
}

func (e *LikeExpression) translatePattern2Regex(pattern string) (string, error) {
	var bd strings.Builder
	var bs bool
	for _, c := range pattern {
		if !bs {
			if c == '%' {
				bd.WriteString("[^\f]*")
				continue
			} else if c == metadata.UNDERSCORE {
				bd.WriteByte(metadata.PERIOD)
				continue
			} else if c == metadata.BACKSLASH {
				bs = true
				continue
			}
		} else {
			if c == '%' {
				bd.WriteByte('%')
				continue
			} else if c == metadata.UNDERSCORE {
				bd.WriteByte(metadata.UNDERSCORE)
				continue
			} else if c == metadata.BACKSLASH {
				bd.WriteString("\\\\")
				continue
			} else {
				return "", fmt.Errorf("Invalid escape char of search pattern '%s'!", pattern)
			}
			bs = false
		}
		bd.WriteRune(c)
	}
	return bd.String(), nil
}

func (e LikeExpression) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	lValue, err := e.lOperand.Operate(entityMap)
	if err != nil {
		return false, err
	}
	s, ok := lValue.(string)
	if !ok {
		return false, nil
	}
	return e.reg.MatchString(s), nil
}

func CreateLikeExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr LikeExpression
	expr.operandType = EXPRESSION
	expr.expressionType = RELATION
	expr.operator = operator
	expr.lOperand = operands[0]
	expr.rOperand = operands[1]
	expr.buildRegExp()
	expr.buildNameString()
	return expr, nil
}
