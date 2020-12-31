// Code generated from ./oper/oper.g4 by ANTLR 4.7.1. DO NOT EDIT.

package oper // oper
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import . "filter4go/operand"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 210,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 41, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 49, 10, 3, 12, 3, 14, 3, 52, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 65, 10, 5, 12, 5, 14, 5,
	68, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 76, 10, 6, 12, 6,
	14, 6, 79, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 87, 10, 7,
	12, 7, 14, 7, 90, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 98,
	10, 8, 12, 8, 14, 8, 101, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7,
	9, 109, 10, 9, 12, 9, 14, 9, 112, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 7, 10, 120, 10, 10, 12, 10, 14, 10, 123, 11, 10, 3, 11, 5, 11,
	126, 10, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 143, 10, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 149, 10, 13, 3, 13, 3, 13, 7, 13, 153, 10,
	13, 12, 13, 14, 13, 156, 11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 165, 10, 13, 3, 13, 3, 13, 5, 13, 169, 10, 13, 3, 13,
	3, 13, 7, 13, 173, 10, 13, 12, 13, 14, 13, 176, 11, 13, 7, 13, 178, 10,
	13, 12, 13, 14, 13, 181, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 188, 10, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 208, 10, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 2, 5, 3, 2, 9, 10, 3, 2, 11, 12, 3, 2, 13, 15, 2, 218,
	2, 40, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 53, 3, 2, 2, 2, 8, 58, 3, 2, 2,
	2, 10, 69, 3, 2, 2, 2, 12, 80, 3, 2, 2, 2, 14, 91, 3, 2, 2, 2, 16, 102,
	3, 2, 2, 2, 18, 113, 3, 2, 2, 2, 20, 125, 3, 2, 2, 2, 22, 142, 3, 2, 2,
	2, 24, 144, 3, 2, 2, 2, 26, 182, 3, 2, 2, 2, 28, 192, 3, 2, 2, 2, 30, 207,
	3, 2, 2, 2, 32, 33, 7, 3, 2, 2, 33, 34, 5, 4, 3, 2, 34, 35, 8, 2, 1, 2,
	35, 36, 7, 4, 2, 2, 36, 41, 3, 2, 2, 2, 37, 38, 5, 8, 5, 2, 38, 39, 8,
	2, 1, 2, 39, 41, 3, 2, 2, 2, 40, 32, 3, 2, 2, 2, 40, 37, 3, 2, 2, 2, 41,
	3, 3, 2, 2, 2, 42, 43, 5, 8, 5, 2, 43, 50, 8, 3, 1, 2, 44, 45, 7, 5, 2,
	2, 45, 46, 5, 8, 5, 2, 46, 47, 8, 3, 1, 2, 47, 49, 3, 2, 2, 2, 48, 44,
	3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2,
	51, 5, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 3, 2, 2, 54, 55, 5, 8,
	5, 2, 55, 56, 7, 4, 2, 2, 56, 57, 8, 4, 1, 2, 57, 7, 3, 2, 2, 2, 58, 59,
	5, 10, 6, 2, 59, 66, 8, 5, 1, 2, 60, 61, 7, 6, 2, 2, 61, 62, 5, 10, 6,
	2, 62, 63, 8, 5, 1, 2, 63, 65, 3, 2, 2, 2, 64, 60, 3, 2, 2, 2, 65, 68,
	3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 9, 3, 2, 2, 2,
	68, 66, 3, 2, 2, 2, 69, 70, 5, 12, 7, 2, 70, 77, 8, 6, 1, 2, 71, 72, 7,
	7, 2, 2, 72, 73, 5, 12, 7, 2, 73, 74, 8, 6, 1, 2, 74, 76, 3, 2, 2, 2, 75,
	71, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2,
	2, 78, 11, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 5, 14, 8, 2, 81, 88,
	8, 7, 1, 2, 82, 83, 7, 8, 2, 2, 83, 84, 5, 14, 8, 2, 84, 85, 8, 7, 1, 2,
	85, 87, 3, 2, 2, 2, 86, 82, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 13, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91,
	92, 5, 16, 9, 2, 92, 99, 8, 8, 1, 2, 93, 94, 9, 2, 2, 2, 94, 95, 5, 16,
	9, 2, 95, 96, 8, 8, 1, 2, 96, 98, 3, 2, 2, 2, 97, 93, 3, 2, 2, 2, 98, 101,
	3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 15, 3, 2, 2,
	2, 101, 99, 3, 2, 2, 2, 102, 103, 5, 18, 10, 2, 103, 110, 8, 9, 1, 2, 104,
	105, 9, 3, 2, 2, 105, 106, 5, 18, 10, 2, 106, 107, 8, 9, 1, 2, 107, 109,
	3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2,
	2, 2, 110, 111, 3, 2, 2, 2, 111, 17, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2,
	113, 114, 5, 20, 11, 2, 114, 121, 8, 10, 1, 2, 115, 116, 9, 4, 2, 2, 116,
	117, 5, 20, 11, 2, 117, 118, 8, 10, 1, 2, 118, 120, 3, 2, 2, 2, 119, 115,
	3, 2, 2, 2, 120, 123, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2,
	2, 2, 122, 19, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 124, 126, 7, 16, 2, 2,
	125, 124, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	128, 5, 22, 12, 2, 128, 129, 8, 11, 1, 2, 129, 21, 3, 2, 2, 2, 130, 131,
	5, 6, 4, 2, 131, 132, 8, 12, 1, 2, 132, 143, 3, 2, 2, 2, 133, 134, 5, 26,
	14, 2, 134, 135, 8, 12, 1, 2, 135, 143, 3, 2, 2, 2, 136, 137, 5, 24, 13,
	2, 137, 138, 8, 12, 1, 2, 138, 143, 3, 2, 2, 2, 139, 140, 5, 30, 16, 2,
	140, 141, 8, 12, 1, 2, 141, 143, 3, 2, 2, 2, 142, 130, 3, 2, 2, 2, 142,
	133, 3, 2, 2, 2, 142, 136, 3, 2, 2, 2, 142, 139, 3, 2, 2, 2, 143, 23, 3,
	2, 2, 2, 144, 145, 5, 28, 15, 2, 145, 154, 8, 13, 1, 2, 146, 148, 7, 17,
	2, 2, 147, 149, 5, 8, 5, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 151, 8, 13, 1, 2, 151, 153, 7, 18, 2, 2, 152,
	146, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 155,
	3, 2, 2, 2, 155, 179, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 157, 164, 7, 19,
	2, 2, 158, 159, 5, 28, 15, 2, 159, 160, 8, 13, 1, 2, 160, 165, 3, 2, 2,
	2, 161, 162, 5, 26, 14, 2, 162, 163, 8, 13, 1, 2, 163, 165, 3, 2, 2, 2,
	164, 158, 3, 2, 2, 2, 164, 161, 3, 2, 2, 2, 165, 174, 3, 2, 2, 2, 166,
	168, 7, 17, 2, 2, 167, 169, 5, 8, 5, 2, 168, 167, 3, 2, 2, 2, 168, 169,
	3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 8, 13, 1, 2, 171, 173, 7, 18,
	2, 2, 172, 166, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2,
	174, 175, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177,
	157, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180,
	3, 2, 2, 2, 180, 25, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 182, 183, 7, 26,
	2, 2, 183, 187, 7, 3, 2, 2, 184, 185, 5, 4, 3, 2, 185, 186, 8, 14, 1, 2,
	186, 188, 3, 2, 2, 2, 187, 184, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188,
	189, 3, 2, 2, 2, 189, 190, 7, 4, 2, 2, 190, 191, 8, 14, 1, 2, 191, 27,
	3, 2, 2, 2, 192, 193, 7, 26, 2, 2, 193, 194, 8, 15, 1, 2, 194, 29, 3, 2,
	2, 2, 195, 196, 7, 23, 2, 2, 196, 208, 8, 16, 1, 2, 197, 198, 7, 24, 2,
	2, 198, 208, 8, 16, 1, 2, 199, 200, 7, 25, 2, 2, 200, 208, 8, 16, 1, 2,
	201, 202, 7, 20, 2, 2, 202, 208, 8, 16, 1, 2, 203, 204, 7, 21, 2, 2, 204,
	208, 8, 16, 1, 2, 205, 206, 7, 22, 2, 2, 206, 208, 8, 16, 1, 2, 207, 195,
	3, 2, 2, 2, 207, 197, 3, 2, 2, 2, 207, 199, 3, 2, 2, 2, 207, 201, 3, 2,
	2, 2, 207, 203, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 31, 3, 2, 2, 2,
	20, 40, 50, 66, 77, 88, 99, 110, 121, 125, 142, 148, 154, 164, 168, 174,
	179, 187, 207,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','", "'|'", "'^'", "'&'", "'<<'", "'>>'", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'~'", "'['", "']'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"NULL", "TRUE", "FALSE", "IntegerLiteral", "FloatingPointLiteral", "StringLiteral",
	"Identifier", "WS",
}

var ruleNames = []string{
	"oper", "expressionList", "parExpression", "expression", "exclusiveOrExpression",
	"andExpression", "shiftExpression", "additiveExpression", "multiplicativeExpression",
	"notExpression", "primary", "member", "function", "variable", "constant",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type operParser struct {
	*antlr.BaseParser
}

func NewoperParser(input antlr.TokenStream) *operParser {
	this := new(operParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "oper.g4"

	return this
}

// operParser tokens.
const (
	operParserEOF                  = antlr.TokenEOF
	operParserT__0                 = 1
	operParserT__1                 = 2
	operParserT__2                 = 3
	operParserT__3                 = 4
	operParserT__4                 = 5
	operParserT__5                 = 6
	operParserT__6                 = 7
	operParserT__7                 = 8
	operParserT__8                 = 9
	operParserT__9                 = 10
	operParserT__10                = 11
	operParserT__11                = 12
	operParserT__12                = 13
	operParserT__13                = 14
	operParserT__14                = 15
	operParserT__15                = 16
	operParserT__16                = 17
	operParserNULL                 = 18
	operParserTRUE                 = 19
	operParserFALSE                = 20
	operParserIntegerLiteral       = 21
	operParserFloatingPointLiteral = 22
	operParserStringLiteral        = 23
	operParserIdentifier           = 24
	operParserWS                   = 25
)

// operParser rules.
const (
	operParserRULE_oper                     = 0
	operParserRULE_expressionList           = 1
	operParserRULE_parExpression            = 2
	operParserRULE_expression               = 3
	operParserRULE_exclusiveOrExpression    = 4
	operParserRULE_andExpression            = 5
	operParserRULE_shiftExpression          = 6
	operParserRULE_additiveExpression       = 7
	operParserRULE_multiplicativeExpression = 8
	operParserRULE_notExpression            = 9
	operParserRULE_primary                  = 10
	operParserRULE_member                   = 11
	operParserRULE_function                 = 12
	operParserRULE_variable                 = 13
	operParserRULE_constant                 = 14
)

// IOperContext is an interface to support dynamic dispatch.
type IOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpList returns the expList rule contexts.
	GetExpList() IExpressionListContext

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// SetExpList sets the expList rule contexts.
	SetExpList(IExpressionListContext)

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsOperContext differentiates from other interfaces.
	IsOperContext()
}

type OperContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	expList IExpressionListContext
	exp     IExpressionContext
}

func NewEmptyOperContext() *OperContext {
	var p = new(OperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_oper
	return p
}

func (*OperContext) IsOperContext() {}

func NewOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperContext {
	var p = new(OperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_oper

	return p
}

func (s *OperContext) GetParser() antlr.Parser { return s.parser }

func (s *OperContext) GetExpList() IExpressionListContext { return s.expList }

func (s *OperContext) GetExp() IExpressionContext { return s.exp }

func (s *OperContext) SetExpList(v IExpressionListContext) { s.expList = v }

func (s *OperContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *OperContext) GetOperand() Operand { return s.operand }

func (s *OperContext) SetOperand(v Operand) { s.operand = v }

func (s *OperContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *OperContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Oper() (localctx IOperContext) {
	localctx = NewOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, operParserRULE_oper)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(operParserT__0)
		}
		{
			p.SetState(31)

			var _x = p.ExpressionList()

			localctx.(*OperContext).expList = _x
		}
		localctx.(*OperContext).operand, _ = CreateOperandsExpression("", localctx.(*OperContext).GetExpList().GetOperands())
		{
			p.SetState(33)
			p.Match(operParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)

			var _x = p.Expression()

			localctx.(*OperContext).exp = _x
		}
		localctx.(*OperContext).operand = localctx.(*OperContext).GetExp().GetOperand()

	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// GetOperands returns the operands attribute.
	GetOperands() []Operand

	// SetOperands sets the operands attribute.
	SetOperands([]Operand)

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operands []Operand
	exp      IExpressionContext
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) GetExp() IExpressionContext { return s.exp }

func (s *ExpressionListContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *ExpressionListContext) GetOperands() []Operand { return s.operands }

func (s *ExpressionListContext) SetOperands(v []Operand) { s.operands = v }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, operParserRULE_expressionList)

	localctx.(*ExpressionListContext).operands = []Operand{}

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)

		var _x = p.Expression()

		localctx.(*ExpressionListContext).exp = _x
	}
	localctx.(*ExpressionListContext).operands = append(localctx.(*ExpressionListContext).operands, localctx.(*ExpressionListContext).GetExp().GetOperand())
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__2 {
		{
			p.SetState(42)
			p.Match(operParserT__2)
		}
		{
			p.SetState(43)

			var _x = p.Expression()

			localctx.(*ExpressionListContext).exp = _x
		}
		localctx.(*ExpressionListContext).operands = append(localctx.(*ExpressionListContext).operands, localctx.(*ExpressionListContext).GetExp().GetOperand())

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IExpressionContext
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *ParExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *ParExpressionContext) GetOperand() Operand { return s.operand }

func (s *ParExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *ParExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, operParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(operParserT__0)
	}
	{
		p.SetState(52)

		var _x = p.Expression()

		localctx.(*ParExpressionContext).exp = _x
	}
	{
		p.SetState(53)
		p.Match(operParserT__1)
	}
	localctx.(*ParExpressionContext).operand, _ = CreateParenExpression("(", []Operand{localctx.(*ParExpressionContext).GetExp().GetOperand()})

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IExclusiveOrExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IExclusiveOrExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IExclusiveOrExpressionContext
	t       antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetT() antlr.Token { return s.t }

func (s *ExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *ExpressionContext) GetExp() IExclusiveOrExpressionContext { return s.exp }

func (s *ExpressionContext) SetExp(v IExclusiveOrExpressionContext) { s.exp = v }

func (s *ExpressionContext) GetOperand() Operand { return s.operand }

func (s *ExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *ExpressionContext) AllExclusiveOrExpression() []IExclusiveOrExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExclusiveOrExpressionContext)(nil)).Elem())
	var tst = make([]IExclusiveOrExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExclusiveOrExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) ExclusiveOrExpression(i int) IExclusiveOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusiveOrExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExclusiveOrExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, operParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)

		var _x = p.ExclusiveOrExpression()

		localctx.(*ExpressionContext).exp = _x
	}
	localctx.(*ExpressionContext).operand = localctx.(*ExpressionContext).GetExp().GetOperand()
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__3 {
		{
			p.SetState(58)

			var _m = p.Match(operParserT__3)

			localctx.(*ExpressionContext).t = _m
		}
		{
			p.SetState(59)

			var _x = p.ExclusiveOrExpression()

			localctx.(*ExpressionContext).exp = _x
		}
		localctx.(*ExpressionContext).operand, _ = CreateBitwiseOrExpression(localctx.(*ExpressionContext).GetT().GetText(), []Operand{localctx.(*ExpressionContext).operand, localctx.(*ExpressionContext).GetExp().GetOperand()})

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IExclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IAndExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IAndExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsExclusiveOrExpressionContext differentiates from other interfaces.
	IsExclusiveOrExpressionContext()
}

type ExclusiveOrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IAndExpressionContext
	t       antlr.Token
}

func NewEmptyExclusiveOrExpressionContext() *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_exclusiveOrExpression
	return p
}

func (*ExclusiveOrExpressionContext) IsExclusiveOrExpressionContext() {}

func NewExclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_exclusiveOrExpression

	return p
}

func (s *ExclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusiveOrExpressionContext) GetT() antlr.Token { return s.t }

func (s *ExclusiveOrExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *ExclusiveOrExpressionContext) GetExp() IAndExpressionContext { return s.exp }

func (s *ExclusiveOrExpressionContext) SetExp(v IAndExpressionContext) { s.exp = v }

func (s *ExclusiveOrExpressionContext) GetOperand() Operand { return s.operand }

func (s *ExclusiveOrExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *ExclusiveOrExpressionContext) AllAndExpression() []IAndExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem())
	var tst = make([]IAndExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAndExpressionContext)
		}
	}

	return tst
}

func (s *ExclusiveOrExpressionContext) AndExpression(i int) IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *ExclusiveOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusiveOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) ExclusiveOrExpression() (localctx IExclusiveOrExpressionContext) {
	localctx = NewExclusiveOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, operParserRULE_exclusiveOrExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)

		var _x = p.AndExpression()

		localctx.(*ExclusiveOrExpressionContext).exp = _x
	}
	localctx.(*ExclusiveOrExpressionContext).operand = localctx.(*ExclusiveOrExpressionContext).GetExp().GetOperand()
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__4 {
		{
			p.SetState(69)

			var _m = p.Match(operParserT__4)

			localctx.(*ExclusiveOrExpressionContext).t = _m
		}
		{
			p.SetState(70)

			var _x = p.AndExpression()

			localctx.(*ExclusiveOrExpressionContext).exp = _x
		}

		localctx.(*ExclusiveOrExpressionContext).operand, _ = CreateBitwiseXorExpression(localctx.(*ExclusiveOrExpressionContext).GetT().GetText(), []Operand{localctx.(*ExclusiveOrExpressionContext).operand, localctx.(*ExclusiveOrExpressionContext).GetExp().GetOperand()})

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IShiftExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IShiftExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IShiftExpressionContext
	t       antlr.Token
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) GetT() antlr.Token { return s.t }

func (s *AndExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *AndExpressionContext) GetExp() IShiftExpressionContext { return s.exp }

func (s *AndExpressionContext) SetExp(v IShiftExpressionContext) { s.exp = v }

func (s *AndExpressionContext) GetOperand() Operand { return s.operand }

func (s *AndExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *AndExpressionContext) AllShiftExpression() []IShiftExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShiftExpressionContext)(nil)).Elem())
	var tst = make([]IShiftExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShiftExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) ShiftExpression(i int) IShiftExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionContext)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, operParserRULE_andExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)

		var _x = p.ShiftExpression()

		localctx.(*AndExpressionContext).exp = _x
	}
	localctx.(*AndExpressionContext).operand = localctx.(*AndExpressionContext).GetExp().GetOperand()
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__5 {
		{
			p.SetState(80)

			var _m = p.Match(operParserT__5)

			localctx.(*AndExpressionContext).t = _m
		}
		{
			p.SetState(81)

			var _x = p.ShiftExpression()

			localctx.(*AndExpressionContext).exp = _x
		}

		localctx.(*AndExpressionContext).operand, _ = CreateBitwiseAndExpression(localctx.(*AndExpressionContext).GetT().GetText(), []Operand{localctx.(*AndExpressionContext).operand, localctx.(*AndExpressionContext).GetExp().GetOperand()})

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IShiftExpressionContext is an interface to support dynamic dispatch.
type IShiftExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IAdditiveExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IAdditiveExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsShiftExpressionContext differentiates from other interfaces.
	IsShiftExpressionContext()
}

type ShiftExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IAdditiveExpressionContext
	t       antlr.Token
}

func NewEmptyShiftExpressionContext() *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_shiftExpression
	return p
}

func (*ShiftExpressionContext) IsShiftExpressionContext() {}

func NewShiftExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_shiftExpression

	return p
}

func (s *ShiftExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExpressionContext) GetT() antlr.Token { return s.t }

func (s *ShiftExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *ShiftExpressionContext) GetExp() IAdditiveExpressionContext { return s.exp }

func (s *ShiftExpressionContext) SetExp(v IAdditiveExpressionContext) { s.exp = v }

func (s *ShiftExpressionContext) GetOperand() Operand { return s.operand }

func (s *ShiftExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *ShiftExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem())
	var tst = make([]IAdditiveExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditiveExpressionContext)
		}
	}

	return tst
}

func (s *ShiftExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) ShiftExpression() (localctx IShiftExpressionContext) {
	localctx = NewShiftExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, operParserRULE_shiftExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)

		var _x = p.AdditiveExpression()

		localctx.(*ShiftExpressionContext).exp = _x
	}
	localctx.(*ShiftExpressionContext).operand = localctx.(*ShiftExpressionContext).GetExp().GetOperand()
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__6 || _la == operParserT__7 {
		{
			p.SetState(91)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ShiftExpressionContext).t = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == operParserT__6 || _la == operParserT__7) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ShiftExpressionContext).t = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(92)

			var _x = p.AdditiveExpression()

			localctx.(*ShiftExpressionContext).exp = _x
		}

		if localctx.(*ShiftExpressionContext).GetT().GetText() == "<<" {
			localctx.(*ShiftExpressionContext).operand, _ = CreateLeftShiftExpression(localctx.(*ShiftExpressionContext).GetT().GetText(), []Operand{localctx.(*ShiftExpressionContext).operand, localctx.(*ShiftExpressionContext).GetExp().GetOperand()})
		} else {
			localctx.(*ShiftExpressionContext).operand, _ = CreateRightShiftExpression(localctx.(*ShiftExpressionContext).GetT().GetText(), []Operand{localctx.(*ShiftExpressionContext).operand, localctx.(*ShiftExpressionContext).GetExp().GetOperand()})
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IMultiplicativeExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IMultiplicativeExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IMultiplicativeExpressionContext
	t       antlr.Token
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) GetT() antlr.Token { return s.t }

func (s *AdditiveExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *AdditiveExpressionContext) GetExp() IMultiplicativeExpressionContext { return s.exp }

func (s *AdditiveExpressionContext) SetExp(v IMultiplicativeExpressionContext) { s.exp = v }

func (s *AdditiveExpressionContext) GetOperand() Operand { return s.operand }

func (s *AdditiveExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplicativeExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicativeExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, operParserRULE_additiveExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)

		var _x = p.MultiplicativeExpression()

		localctx.(*AdditiveExpressionContext).exp = _x
	}
	localctx.(*AdditiveExpressionContext).operand = localctx.(*AdditiveExpressionContext).GetExp().GetOperand()
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__8 || _la == operParserT__9 {
		{
			p.SetState(102)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AdditiveExpressionContext).t = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == operParserT__8 || _la == operParserT__9) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AdditiveExpressionContext).t = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(103)

			var _x = p.MultiplicativeExpression()

			localctx.(*AdditiveExpressionContext).exp = _x
		}

		if localctx.(*AdditiveExpressionContext).GetT().GetText() == "+" {
			localctx.(*AdditiveExpressionContext).operand, _ = CreateAddExpression(localctx.(*AdditiveExpressionContext).GetT().GetText(), []Operand{localctx.(*AdditiveExpressionContext).operand, localctx.(*AdditiveExpressionContext).GetExp().GetOperand()})
		} else {
			localctx.(*AdditiveExpressionContext).operand, _ = CreateSubtractExpression(localctx.(*AdditiveExpressionContext).GetT().GetText(), []Operand{localctx.(*AdditiveExpressionContext).operand, localctx.(*AdditiveExpressionContext).GetExp().GetOperand()})
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() INotExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(INotExpressionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     INotExpressionContext
	t       antlr.Token
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) GetT() antlr.Token { return s.t }

func (s *MultiplicativeExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *MultiplicativeExpressionContext) GetExp() INotExpressionContext { return s.exp }

func (s *MultiplicativeExpressionContext) SetExp(v INotExpressionContext) { s.exp = v }

func (s *MultiplicativeExpressionContext) GetOperand() Operand { return s.operand }

func (s *MultiplicativeExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *MultiplicativeExpressionContext) AllNotExpression() []INotExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INotExpressionContext)(nil)).Elem())
	var tst = make([]INotExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INotExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) NotExpression(i int) INotExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INotExpressionContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, operParserRULE_multiplicativeExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)

		var _x = p.NotExpression()

		localctx.(*MultiplicativeExpressionContext).exp = _x
	}
	localctx.(*MultiplicativeExpressionContext).operand = localctx.(*MultiplicativeExpressionContext).GetExp().GetOperand()
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<operParserT__10)|(1<<operParserT__11)|(1<<operParserT__12))) != 0 {
		{
			p.SetState(113)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MultiplicativeExpressionContext).t = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<operParserT__10)|(1<<operParserT__11)|(1<<operParserT__12))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MultiplicativeExpressionContext).t = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(114)

			var _x = p.NotExpression()

			localctx.(*MultiplicativeExpressionContext).exp = _x
		}

		if localctx.(*MultiplicativeExpressionContext).GetT().GetText() == "*" {
			localctx.(*MultiplicativeExpressionContext).operand, _ = CreateMultiplyExpression(localctx.(*MultiplicativeExpressionContext).GetT().GetText(), []Operand{localctx.(*MultiplicativeExpressionContext).operand, localctx.(*MultiplicativeExpressionContext).GetExp().GetOperand()})
		} else if localctx.(*MultiplicativeExpressionContext).GetT().GetText() == "/" {
			localctx.(*MultiplicativeExpressionContext).operand, _ = CreateDivideExpression(localctx.(*MultiplicativeExpressionContext).GetT().GetText(), []Operand{localctx.(*MultiplicativeExpressionContext).operand, localctx.(*MultiplicativeExpressionContext).GetExp().GetOperand()})
		} else {
			localctx.(*MultiplicativeExpressionContext).operand, _ = CreateModularExpression(localctx.(*MultiplicativeExpressionContext).GetT().GetText(), []Operand{localctx.(*MultiplicativeExpressionContext).operand, localctx.(*MultiplicativeExpressionContext).GetExp().GetOperand()})
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INotExpressionContext is an interface to support dynamic dispatch.
type INotExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IPrimaryContext

	// SetExp sets the exp rule contexts.
	SetExp(IPrimaryContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsNotExpressionContext differentiates from other interfaces.
	IsNotExpressionContext()
}

type NotExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	t       antlr.Token
	exp     IPrimaryContext
}

func NewEmptyNotExpressionContext() *NotExpressionContext {
	var p = new(NotExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_notExpression
	return p
}

func (*NotExpressionContext) IsNotExpressionContext() {}

func NewNotExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_notExpression

	return p
}

func (s *NotExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NotExpressionContext) GetT() antlr.Token { return s.t }

func (s *NotExpressionContext) SetT(v antlr.Token) { s.t = v }

func (s *NotExpressionContext) GetExp() IPrimaryContext { return s.exp }

func (s *NotExpressionContext) SetExp(v IPrimaryContext) { s.exp = v }

func (s *NotExpressionContext) GetOperand() Operand { return s.operand }

func (s *NotExpressionContext) SetOperand(v Operand) { s.operand = v }

func (s *NotExpressionContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) NotExpression() (localctx INotExpressionContext) {
	localctx = NewNotExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, operParserRULE_notExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == operParserT__13 {
		{
			p.SetState(122)

			var _m = p.Match(operParserT__13)

			localctx.(*NotExpressionContext).t = _m
		}

	}
	{
		p.SetState(125)

		var _x = p.Primary()

		localctx.(*NotExpressionContext).exp = _x
	}

	if localctx.(*NotExpressionContext).GetT() != nil {
		localctx.(*NotExpressionContext).operand, _ = CreateBitwiseNotExpression(localctx.(*NotExpressionContext).GetT().GetText(), []Operand{localctx.(*NotExpressionContext).GetExp().GetOperand()})
	} else {
		localctx.(*NotExpressionContext).operand = localctx.(*NotExpressionContext).GetExp().GetOperand()
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IParExpressionContext

	// GetFexp returns the fexp rule contexts.
	GetFexp() IFunctionContext

	// GetMexp returns the mexp rule contexts.
	GetMexp() IMemberContext

	// GetCexp returns the cexp rule contexts.
	GetCexp() IConstantContext

	// SetExp sets the exp rule contexts.
	SetExp(IParExpressionContext)

	// SetFexp sets the fexp rule contexts.
	SetFexp(IFunctionContext)

	// SetMexp sets the mexp rule contexts.
	SetMexp(IMemberContext)

	// SetCexp sets the cexp rule contexts.
	SetCexp(IConstantContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IParExpressionContext
	fexp    IFunctionContext
	mexp    IMemberContext
	cexp    IConstantContext
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) GetExp() IParExpressionContext { return s.exp }

func (s *PrimaryContext) GetFexp() IFunctionContext { return s.fexp }

func (s *PrimaryContext) GetMexp() IMemberContext { return s.mexp }

func (s *PrimaryContext) GetCexp() IConstantContext { return s.cexp }

func (s *PrimaryContext) SetExp(v IParExpressionContext) { s.exp = v }

func (s *PrimaryContext) SetFexp(v IFunctionContext) { s.fexp = v }

func (s *PrimaryContext) SetMexp(v IMemberContext) { s.mexp = v }

func (s *PrimaryContext) SetCexp(v IConstantContext) { s.cexp = v }

func (s *PrimaryContext) GetOperand() Operand { return s.operand }

func (s *PrimaryContext) SetOperand(v Operand) { s.operand = v }

func (s *PrimaryContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *PrimaryContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *PrimaryContext) Member() IMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *PrimaryContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, operParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)

			var _x = p.ParExpression()

			localctx.(*PrimaryContext).exp = _x
		}
		localctx.(*PrimaryContext).operand = localctx.(*PrimaryContext).GetExp().GetOperand()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)

			var _x = p.Function()

			localctx.(*PrimaryContext).fexp = _x
		}
		localctx.(*PrimaryContext).operand = localctx.(*PrimaryContext).GetFexp().GetOperand()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)

			var _x = p.Member()

			localctx.(*PrimaryContext).mexp = _x
		}
		localctx.(*PrimaryContext).operand = localctx.(*PrimaryContext).GetMexp().GetOperand()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)

			var _x = p.Constant()

			localctx.(*PrimaryContext).cexp = _x
		}
		localctx.(*PrimaryContext).operand = localctx.(*PrimaryContext).GetCexp().GetOperand()

	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IVariableContext

	// GetIndex returns the index rule contexts.
	GetIndex() IExpressionContext

	// GetVexp returns the vexp rule contexts.
	GetVexp() IVariableContext

	// GetFexp returns the fexp rule contexts.
	GetFexp() IFunctionContext

	// SetExp sets the exp rule contexts.
	SetExp(IVariableContext)

	// SetIndex sets the index rule contexts.
	SetIndex(IExpressionContext)

	// SetVexp sets the vexp rule contexts.
	SetVexp(IVariableContext)

	// SetFexp sets the fexp rule contexts.
	SetFexp(IFunctionContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	exp     IVariableContext
	index   IExpressionContext
	vexp    IVariableContext
	fexp    IFunctionContext
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) GetExp() IVariableContext { return s.exp }

func (s *MemberContext) GetIndex() IExpressionContext { return s.index }

func (s *MemberContext) GetVexp() IVariableContext { return s.vexp }

func (s *MemberContext) GetFexp() IFunctionContext { return s.fexp }

func (s *MemberContext) SetExp(v IVariableContext) { s.exp = v }

func (s *MemberContext) SetIndex(v IExpressionContext) { s.index = v }

func (s *MemberContext) SetVexp(v IVariableContext) { s.vexp = v }

func (s *MemberContext) SetFexp(v IFunctionContext) { s.fexp = v }

func (s *MemberContext) GetOperand() Operand { return s.operand }

func (s *MemberContext) SetOperand(v Operand) { s.operand = v }

func (s *MemberContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *MemberContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *MemberContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *MemberContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *MemberContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MemberContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, operParserRULE_member)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)

		var _x = p.Variable()

		localctx.(*MemberContext).exp = _x
	}
	localctx.(*MemberContext).operand = localctx.(*MemberContext).GetExp().GetOperand()
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__14 {
		{
			p.SetState(144)
			p.Match(operParserT__14)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<operParserT__0)|(1<<operParserT__13)|(1<<operParserNULL)|(1<<operParserTRUE)|(1<<operParserFALSE)|(1<<operParserIntegerLiteral)|(1<<operParserFloatingPointLiteral)|(1<<operParserStringLiteral)|(1<<operParserIdentifier))) != 0 {
			{
				p.SetState(145)

				var _x = p.Expression()

				localctx.(*MemberContext).index = _x
			}

		}
		localctx.(*MemberContext).operand, _ = CreateArrayExpression("[", []Operand{localctx.(*MemberContext).operand, localctx.(*MemberContext).GetIndex().GetOperand()})
		{
			p.SetState(149)
			p.Match(operParserT__15)
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == operParserT__16 {
		{
			p.SetState(155)
			p.Match(operParserT__16)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(156)

				var _x = p.Variable()

				localctx.(*MemberContext).vexp = _x
			}
			localctx.(*MemberContext).operand, _ = CreateMemberVariableExpression("", []Operand{localctx.(*MemberContext).operand, localctx.(*MemberContext).GetVexp().GetOperand()})

		case 2:
			{
				p.SetState(159)

				var _x = p.Function()

				localctx.(*MemberContext).fexp = _x
			}
			localctx.(*MemberContext).operand, _ = CreateMemberMethodExpression("", []Operand{localctx.(*MemberContext).operand, localctx.(*MemberContext).GetFexp().GetOperand()})

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == operParserT__14 {
			{
				p.SetState(164)
				p.Match(operParserT__14)
			}
			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<operParserT__0)|(1<<operParserT__13)|(1<<operParserNULL)|(1<<operParserTRUE)|(1<<operParserFALSE)|(1<<operParserIntegerLiteral)|(1<<operParserFloatingPointLiteral)|(1<<operParserStringLiteral)|(1<<operParserIdentifier))) != 0 {
				{
					p.SetState(165)

					var _x = p.Expression()

					localctx.(*MemberContext).index = _x
				}

			}
			localctx.(*MemberContext).operand, _ = CreateArrayExpression("[", []Operand{localctx.(*MemberContext).operand, localctx.(*MemberContext).GetIndex().GetOperand()})
			{
				p.SetState(169)
				p.Match(operParserT__15)
			}

			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetExpList returns the expList rule contexts.
	GetExpList() IExpressionListContext

	// SetExpList sets the expList rule contexts.
	SetExpList(IExpressionListContext)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	t       antlr.Token
	expList IExpressionListContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetT() antlr.Token { return s.t }

func (s *FunctionContext) SetT(v antlr.Token) { s.t = v }

func (s *FunctionContext) GetExpList() IExpressionListContext { return s.expList }

func (s *FunctionContext) SetExpList(v IExpressionListContext) { s.expList = v }

func (s *FunctionContext) GetOperand() Operand { return s.operand }

func (s *FunctionContext) SetOperand(v Operand) { s.operand = v }

func (s *FunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(operParserIdentifier, 0)
}

func (s *FunctionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, operParserRULE_function)

	var err error
	var parameters []Operand

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)

		var _m = p.Match(operParserIdentifier)

		localctx.(*FunctionContext).t = _m
	}
	{
		p.SetState(181)
		p.Match(operParserT__0)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<operParserT__0)|(1<<operParserT__13)|(1<<operParserNULL)|(1<<operParserTRUE)|(1<<operParserFALSE)|(1<<operParserIntegerLiteral)|(1<<operParserFloatingPointLiteral)|(1<<operParserStringLiteral)|(1<<operParserIdentifier))) != 0 {
		{
			p.SetState(182)

			var _x = p.ExpressionList()

			localctx.(*FunctionContext).expList = _x
		}
		parameters = localctx.(*FunctionContext).GetExpList().GetOperands()

	}
	{
		p.SetState(187)
		p.Match(operParserT__1)
	}

	localctx.(*FunctionContext).operand, err = FuncFactory.CreateFunction(localctx.(*FunctionContext).GetT().GetText(), parameters)
	if err != nil {
		panic(err)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	t       antlr.Token
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) GetT() antlr.Token { return s.t }

func (s *VariableContext) SetT(v antlr.Token) { s.t = v }

func (s *VariableContext) GetOperand() Operand { return s.operand }

func (s *VariableContext) SetOperand(v Operand) { s.operand = v }

func (s *VariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(operParserIdentifier, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, operParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)

		var _m = p.Match(operParserIdentifier)

		localctx.(*VariableContext).t = _m
	}
	localctx.(*VariableContext).SetOperand(CreateVariable(localctx.(*VariableContext).GetT().GetText()))

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetOperand returns the operand attribute.
	GetOperand() Operand

	// SetOperand sets the operand attribute.
	SetOperand(Operand)

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand Operand
	t       antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = operParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = operParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetT() antlr.Token { return s.t }

func (s *ConstantContext) SetT(v antlr.Token) { s.t = v }

func (s *ConstantContext) GetOperand() Operand { return s.operand }

func (s *ConstantContext) SetOperand(v Operand) { s.operand = v }

func (s *ConstantContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(operParserIntegerLiteral, 0)
}

func (s *ConstantContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(operParserFloatingPointLiteral, 0)
}

func (s *ConstantContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(operParserStringLiteral, 0)
}

func (s *ConstantContext) NULL() antlr.TerminalNode {
	return s.GetToken(operParserNULL, 0)
}

func (s *ConstantContext) TRUE() antlr.TerminalNode {
	return s.GetToken(operParserTRUE, 0)
}

func (s *ConstantContext) FALSE() antlr.TerminalNode {
	return s.GetToken(operParserFALSE, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *operParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, operParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(205)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case operParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)

			var _m = p.Match(operParserIntegerLiteral)

			localctx.(*ConstantContext).t = _m
		}
		localctx.(*ConstantContext).operand = CreateLong(localctx.(*ConstantContext).GetT().GetText())

	case operParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)

			var _m = p.Match(operParserFloatingPointLiteral)

			localctx.(*ConstantContext).t = _m
		}
		localctx.(*ConstantContext).operand = CreateDouble(localctx.(*ConstantContext).GetT().GetText())

	case operParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)

			var _m = p.Match(operParserStringLiteral)

			localctx.(*ConstantContext).t = _m
		}
		localctx.(*ConstantContext).operand = CreateString(localctx.(*ConstantContext).GetT().GetText())

	case operParserNULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)

			var _m = p.Match(operParserNULL)

			localctx.(*ConstantContext).t = _m
		}
		localctx.(*ConstantContext).operand = CreateNull(localctx.(*ConstantContext).GetT().GetText())

	case operParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)

			var _m = p.Match(operParserTRUE)

			localctx.(*ConstantContext).t = _m
		}
		localctx.(*ConstantContext).operand = CreateBoolean(localctx.(*ConstantContext).GetT().GetText())

	case operParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(203)

			var _m = p.Match(operParserFALSE)

			localctx.(*ConstantContext).t = _m
		}
		localctx.(*ConstantContext).SetOperand(CreateBoolean(localctx.(*ConstantContext).GetT().GetText()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
