package operand

import (
	"filter4go/metadata"
	"fmt"
	"reflect"
	"strings"
)

type MemberVariableExpression struct {
	Expression
	target     Operand
	member     Variable
	targetType reflect.Type
	fieldIndex int
	fieldCache map[reflect.Type]int
}

func (e MemberVariableExpression) GetTarget() Operand {
	return e.target
}

func (e MemberVariableExpression) GetMember() Variable {
	return e.member
}

func (e *MemberVariableExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteString(e.target.ToString())
	bd.WriteByte(metadata.PERIOD)
	bd.WriteString(e.member.ToString())
	e.name = bd.String()
}

func (e MemberVariableExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	o, err := e.target.Operate(entityMap)
	if err != nil {
		return nil, err
	}
	m, ok := o.(map[string]interface{})
	if ok {
		return m[e.member.name], nil
	}
	return e.operate(o)
}

func (e *MemberVariableExpression) operate(o interface{}) (interface{}, error) {
	v := reflect.ValueOf(o)
	vType := v.Type()
	index := e.getFieldIndex(vType)
	if index == -1 {
		return nil, fmt.Errorf("Cann't find field named '%s'!", e.member.name)
	}
	return v.Field(index).Interface(), nil
}

func (e *MemberVariableExpression) getFieldIndex(vType reflect.Type) int {
	if e.targetType != nil && vType == e.targetType {
		return e.fieldIndex
	}
	index, ok := e.fieldCache[vType]
	if ok {
		return index
	}
	field, ok := vType.FieldByName(e.member.name)
	if ok {
		e.fieldCache[vType] = field.Index[0]
		e.targetType = vType
		e.fieldIndex = field.Index[0]
		return e.fieldIndex
	}
	return -1
}

type MemberMethodExpression struct {
	Expression
	target      Operand
	function    Function
	targetType  reflect.Type
	method      reflect.Method
	methodCache map[reflect.Type]reflect.Method
}

func (e MemberMethodExpression) GetTarget() Operand {
	return e.target
}

func (e MemberMethodExpression) GetFunction() Function {
	return e.function
}

func (e *MemberMethodExpression) buildNameString() {
	var bd strings.Builder
	bd.WriteString(e.target.ToString())
	bd.WriteByte(metadata.PERIOD)
	bd.WriteString(e.function.ToString())
	e.name = bd.String()
}

func (e MemberMethodExpression) Operate(entityMap map[string]interface{}) (interface{}, error) {
	o, err := e.target.Operate(entityMap)
	if err != nil {
		return nil, err
	}
	parameters := []reflect.Value{reflect.ValueOf(o)}
	parameterObjects := e.getParameterObjects(entityMap)
	parameters = append(parameters, parameterObjects...)
	return e.invoke(o, parameters)
}

func (e *MemberMethodExpression) getParameterObjects(entityMap map[string]interface{}) []reflect.Value {
	parameters := e.function.GetParameters()
	if parameters == nil {
		return make([]reflect.Value, 0)
	}
	parameterObjects := make([]reflect.Value, len(parameters))
	for i := 0; i < len(parameters); i++ {
		v, _ := parameters[i].Operate(entityMap)
		parameterObjects[i] = reflect.ValueOf(v)
	}
	return parameterObjects
}

func (e *MemberMethodExpression) invoke(target interface{}, parameters []reflect.Value) (interface{}, error) {
	v := reflect.ValueOf(target)
	vType := v.Type()
	method := e.getMethod(vType)
	if method == (reflect.Method{}) {
		return nil, fmt.Errorf("Get method '%s' from type '%s' failed!", e.function.GetName(), vType.Name())
	}
	ret := method.Func.Call(parameters)
	if len(ret) != 0 {
		return ret[0], nil
	}
	return nil, nil
}

func (e *MemberMethodExpression) getMethod(vType reflect.Type) reflect.Method {
	if e.targetType != nil && vType == e.targetType {
		return e.method
	}
	method, ok := e.methodCache[vType]
	if ok {
		return method
	}
	method, ok = vType.MethodByName(e.function.GetName())
	if ok {
		e.methodCache[vType] = method
		e.targetType = vType
		e.method = method
		return method
	}
	return reflect.Method{}
}

func CreateMemberVariableExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr MemberVariableExpression
	expr.operandType = EXPRESSION
	expr.expressionType = MEMBER
	expr.target = operands[0]
	expr.member = operands[1].(Variable)
	expr.buildNameString()
	expr.fieldCache = map[reflect.Type]int{}
	return expr, nil
}

func CreateMemberMethodExpression(operator string, operands []Operand) (Operand, error) {
	if operands == nil || len(operands) != 2 {
		return nil, fmt.Errorf("Invalid operands count!")
	}
	var expr MemberMethodExpression
	expr.operandType = EXPRESSION
	expr.expressionType = MEMBER
	expr.target = operands[0]
	expr.function = operands[1].(Function)
	expr.buildNameString()
	expr.methodCache = map[reflect.Type]reflect.Method{}
	return expr, nil
}

type MemberFunction struct {
	BaseFunction
}

func CreateMemberFunction(name string, parameters []Operand) (Function, error) {
	function := MemberFunction{}
	function.name = name
	function.parameterCount = -1
	function.BuildFunction(parameters)
	return function, nil
}
