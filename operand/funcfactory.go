package operand

import (
	"fmt"
	"reflect"
	"strings"
)

type CreateFunction func(parameters []Operand) (Operand, error)

type FunctionFactory struct {
	funcMap map[string]CreateFunction
}

var FuncFactory FunctionFactory = FunctionFactory{
	funcMap: map[string]CreateFunction{
		"sprintf": CreateSprintfFunction,
	},
}

func (ff *FunctionFactory) registCreateFunction(name string, function CreateFunction) {
	ff.funcMap[name] = function
}

func (ff *FunctionFactory) CreateFunction(function string, parameters []Operand) (Operand, error) {
	function = strings.ToLower(function)
	createFunc := ff.funcMap[function]
	if createFunc == nil {
		return CreateMemberFunction(function, parameters)
	}
	return createFunc(parameters)
}

type Sprintf struct {
	BaseFunction
	funcPtr reflect.Value
}

func (f Sprintf) Operate(entityMap map[string]interface{}) (interface{}, error) {
	pvs, err := f.operateParameters(entityMap)
	if err != nil {
		return nil, err
	}
	pvv := make([]reflect.Value, len(pvs))
	for i := 0; i < len(pvs); i++ {
		pvv[i] = reflect.ValueOf(pvs[i])
	}
	ret := f.funcPtr.Call(pvv)
	return ret[0].Interface(), nil
}

func CreateSprintfFunction(parameters []Operand) (Operand, error) {
	sprintf := Sprintf{}
	sprintf.name = "sprintf"
	sprintf.parameterCount = -1
	err := sprintf.BuildFunction(parameters)
	if err != nil {
		return nil, err
	}
	sprintf.funcPtr = reflect.ValueOf(fmt.Sprintf)
	return sprintf, nil
}
