package operand

import (
	"errors"
	"filter4go/metadata"
	"fmt"
	"strings"
)

type OperandType int

const (
	UNKNOWN OperandType = iota
	CONSTANT
	VARIABLE
	FUNCTION
	EXPRESSION
)

type Operand interface {
	GetName() string

	Operate(entityMap map[string]interface{}) (interface{}, error)

	BooleanOperate(entityMap map[string]interface{}) (bool, error)

	GetOperandType() OperandType

	ToString() string
}

type BaseOperand struct {
	name        string
	operandType OperandType
}

func isTrue(object interface{}) bool {
	if object == nil {
		return false
	}
	t, b := object.(bool)
	if b {
		return t
	}
	return true
}

func (op BaseOperand) GetName() string {
	return op.name
}

func (op BaseOperand) GetOperandType() OperandType {
	return op.operandType
}

func (op BaseOperand) Operate(entityMap map[string]interface{}) (interface{}, error) {
	return nil, errors.New("Unsupported operation!")
}

func (op BaseOperand) BooleanOperate(entityMap map[string]interface{}) (bool, error) {
	obj, err := op.Operate(entityMap)
	if err != nil {
		return false, err
	}
	return isTrue(obj), nil
}

func (op BaseOperand) ToString() string {
	return op.name
}

type Variable struct {
	BaseOperand
}

func (v Variable) Operate(entityMap map[string]interface{}) (interface{}, error) {
	return entityMap[v.name], nil
}

func CreateVariable(name string) Operand {
	var v Variable
	v.name = name
	v.operandType = VARIABLE
	return v
}

type Function interface {
	GetName() string

	Operate(entityMap map[string]interface{}) (interface{}, error)

	BooleanOperate(entityMap map[string]interface{}) (bool, error)

	GetOperandType() OperandType

	ToString() string

	GetParameterCount() int

	GetParameters() []Operand
}

type BaseFunction struct {
	BaseOperand
	parameterCount int
	parameters     []Operand
	functionString string
}

func (f *BaseFunction) BuildFunction(parameters []Operand) error {
	if f.parameterCount >= 0 {
		if parameters == nil {
			if f.parameterCount > 0 {
				return fmt.Errorf("Function '%s' need %d parameters!", f.GetName(), f.parameterCount)
			}
		} else if len(parameters) != f.parameterCount {
			return fmt.Errorf("Function '%s' need %d parameters!", f.GetName(), f.parameterCount)
		}
	}
	f.operandType = FUNCTION
	f.parameters = parameters
	f.buildFunctionString()
	return nil
}

func (f *BaseFunction) buildFunctionString() {
	var bd strings.Builder
	bd.WriteString(f.name)
	bd.WriteByte(metadata.LPAREN)
	if f.parameters != nil {
		for i := 0; i < len(f.parameters); i++ {
			if i != 0 {
				bd.WriteByte(metadata.COMMA)
			}
			bd.WriteString(f.parameters[i].ToString())
		}
	}
	bd.WriteByte(metadata.RPAREN)
	f.functionString = bd.String()
}

func (f *BaseFunction) operateParameters(entityMap map[string]interface{}) ([]interface{}, error) {
	pvs := []interface{}{}
	if f.parameters != nil && len(f.parameters) != 0 {
		pvs = make([]interface{}, len(f.parameters))
	}
	var err error
	for i, parameter := range f.parameters {
		pvs[i], err = parameter.Operate(entityMap)
		if err != nil {
			return nil, err
		}
	}
	return pvs, nil
}

func (f BaseFunction) GetParameterCount() int {
	return f.parameterCount
}

func (f BaseFunction) GetParameters() []Operand {
	return f.parameters
}

func (f BaseFunction) ToString() string {
	return f.functionString
}
