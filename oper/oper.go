package oper

import (
	. "filter4go/operand"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseOperand(operand string) Operand {
	is := antlr.NewInputStream(operand)
	// Create the Lexer
	lexer := NewoperLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the filter Parser
	parser := NewoperParser(tokens)
	return parser.Oper().GetOperand()
}
