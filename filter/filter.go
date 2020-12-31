package filter

import (
	"filter4go/metadata"
	"filter4go/oper"
	. "filter4go/operand"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strings"
)

func ParseFilterMetadata(filter string) metadata.OperatableMetadata {
	is := antlr.NewInputStream(filter)
	// Create the Lexer
	lexer := NewfilterLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the filter Parser
	parser := NewfilterParser(tokens)
	return parser.SearchCondition().GetOperationMetadata()
}

func ParseFilter(filter string) (Operand, error) {
	metadata := ParseFilterMetadata(filter)
	return buildOperand(metadata)
}

func BuildFilter(metadata metadata.OperatableMetadata) (Operand, error) {
	return buildOperand(metadata)
}

func buildOperand(meta metadata.OperatableMetadata) (Operand, error) {
	typ := reflect.TypeOf(meta)
	if typ.Name() == "RelationOperationMetadata" {
		rom := meta.(metadata.RelationOperationMetadata)
		return buildByRelationOperationMetadata(rom)
	} else {
		om := meta.(metadata.OperationMetadata)
		return buildByOperationMetadata(om)
	}
}

func buildByOperationMetadata(meta metadata.OperationMetadata) (Operand, error) {
	operands := []Operand{}
	index := 0
	if meta.OperatorType == metadata.BINARY {
		operand, err := buildOperand(meta.LOperand)
		if err != nil {
			return nil, err
		}
		operands = append(operands, operand)
		index++
	}
	operand, err := buildOperand(meta.ROperand)
	if err != nil {
		return nil, err
	}
	operands = append(operands, operand)
	return createOperand(meta.Operator, operands)
}

func buildByRelationOperationMetadata(meta metadata.RelationOperationMetadata) (Operand, error) {
	operands := []Operand{nil, nil}
	index := 0
	if meta.OperatorType == metadata.BINARY {
		operands[index] = oper.ParseOperand(meta.LOperand)
		index++
	}
	operands[index] = oper.ParseOperand(meta.ROperand)
	return createOperand(meta.Operator, operands)
}

var exprFuncMap = map[string]CreateExpression{
	metadata.AND:     CreateAndExpression,
	metadata.OR:      CreateOrExpression,
	metadata.NOT:     CreateNotExpression,
	metadata.EQ:      CreateEqualExpression,
	metadata.GE:      CreateGreatAndEqualExpression,
	metadata.GT:      CreateGreatThanExpression,
	metadata.LE:      CreateLittleAndEqualExpression,
	metadata.LT:      CreateLittleAndEqualExpression,
	metadata.NE:      CreateNotEqualExpression,
	metadata.BETWEEN: CreateBetweenExpression,
	metadata.IN:      CreateInExpression,
	metadata.LIKE:    CreateLikeExpression,
	metadata.IS:      CreateIsExpression,
	metadata.EXPR:    CreateExprExpression,
	metadata.PAREN:   CreateParenExpression,
}

func createOperand(operator string, operands []Operand) (Operand, error) {
	if operator != "" {
		op := strings.ToLower(operator)
		createFunc := exprFuncMap[op]
		if createFunc != nil {
			return createFunc(operator, operands)
		}
	}
	return nil, fmt.Errorf("Unsupported operator '%s'", operator)
}

func CreateExprExpression(operator string, operands []Operand) (Operand, error) {
	return operands[0], nil
}
