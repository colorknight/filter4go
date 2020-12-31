package operand

import (
	"filter4go/metadata"
	"fmt"
	"reflect"
	"strings"
)

type collectionDataAccessor func(collection interface{}, index interface{}) (interface{}, error)

func mapAccessor(collection interface{}, index interface{}) (interface{}, error) {
	v := reflect.ValueOf(collection)
	return v.MapIndex(reflect.ValueOf(index)).Interface(), nil
}

func arrayAccessor(collection interface{}, index interface{}) (interface{}, error) {
	v := reflect.ValueOf(collection)
	inx64, ok := index.(int64)
	inx := int(inx64)
	if ok && inx > -1 && inx < v.Len() {
		return v.Index(inx).Interface(), nil
	}
	return nil, fmt.Errorf("Invalid array index!")
}

type ArrayExpression struct {
	Expression
	array Operand
	index Operand
}

func (e *ArrayExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteString(e.array.ToString())
	bd.WriteByte(metadata.LBRACKET)
	if e.index != nil {
		bd.WriteString(e.index.ToString())
	}
	bd.WriteByte(metadata.RBRACKET)
	e.name = bd.String()
}

func (e ArrayExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	o, err := e.array.Operate(entityMap)
	if o == nil || err != nil {
		return nil, err
	}
	inx, err := e.index.Operate(entityMap)
	tp := reflect.TypeOf(o)
	arrayAccessor := e.getAccessor(tp.String())
	if arrayAccessor != nil {
		return arrayAccessor(o, inx)
	}
	return nil, fmt.Errorf("Unsupported array type!")
}

func (e *ArrayExpression) getAccessor(name string) collectionDataAccessor {
	if strings.HasPrefix(name, "[]") {
		return arrayAccessor
	} else if strings.HasPrefix(name, "map") {
		return mapAccessor
	}
	return nil
}

func CreateArrayExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr ArrayExpression
	expr.operandType = EXPRESSION
	expr.expressionType = ARRAY
	expr.array = operands[0]
	expr.index = operands[1]
	expr.buildNameString()
	return expr, nil
}
