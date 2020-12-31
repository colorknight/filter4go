package operand

import (
	"filter4go/metadata"
	"filter4go/util"
	"strconv"
)

type ConstantType int

const (
	NULL ConstantType = iota
	LONG
	DOUBLE
	STRING
	BOOLEAN
)

type Constant struct {
	BaseOperand
	constantType ConstantType
	data         interface{}
}

func (op Constant) Operate(entityMap map[string]interface{}) (interface{}, error) {
	return op.data, nil
}

func (op Constant) GetValue() interface{} {
	return op.data
}

func (op Constant) Clear() {}

func (op *Constant) GetConstantType() ConstantType {
	return op.constantType
}

type createConstant func(value string) Operand

func CreateNull(value string) Operand {
	return Constant{
		BaseOperand:  BaseOperand{name: metadata.NULL, operandType: CONSTANT},
		constantType: NULL,
		data:         value,
	}
}

func CreateLong(value string) Operand {
	l, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return nil
	}
	return &Constant{
		BaseOperand:  BaseOperand{name: value, operandType: CONSTANT},
		constantType: LONG,
		data:         l,
	}
}

func CreateDouble(value string) Operand {
	d, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil
	}
	return &Constant{
		BaseOperand:  BaseOperand{name: value, operandType: CONSTANT},
		constantType: DOUBLE,
		data:         d,
	}
}

func CreateBoolean(value string) Operand {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	return &Constant{
		BaseOperand:  BaseOperand{name: value, operandType: CONSTANT},
		constantType: BOOLEAN,
		data:         b,
	}
}

func CreateString(value string) Operand {
	return &Constant{
		BaseOperand:  BaseOperand{name: value, operandType: CONSTANT},
		constantType: STRING,
		data:         util.UntranslateEscape(&value),
	}
}

var constFuncMap = map[ConstantType]createConstant{
	NULL:    CreateNull,
	LONG:    CreateLong,
	DOUBLE:  CreateDouble,
	STRING:  CreateString,
	BOOLEAN: CreateBoolean,
}

func CreateConstant(constantType ConstantType, value string) Operand {
	createConstant := constFuncMap[constantType]
	return createConstant(value)
}
