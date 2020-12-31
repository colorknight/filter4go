// Code generated from ./filter/filter.g4 by ANTLR 4.7.1. DO NOT EDIT.

package filter // filter
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import . "filter4go/metadata"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 255,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 57,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 65, 10, 3, 3, 4, 5, 4,
	68, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 83, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 100, 10, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 108, 10, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 117, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 5, 10, 126, 10, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 5, 11, 134, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 141,
	10, 12, 12, 12, 14, 12, 144, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 155, 10, 14, 12, 14, 14, 14, 158, 11,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 7, 15, 165, 10, 15, 12, 15, 14,
	15, 168, 11, 15, 3, 16, 3, 16, 3, 16, 7, 16, 173, 10, 16, 12, 16, 14, 16,
	176, 11, 16, 3, 17, 3, 17, 3, 17, 7, 17, 181, 10, 17, 12, 17, 14, 17, 184,
	11, 17, 3, 18, 3, 18, 3, 18, 7, 18, 189, 10, 18, 12, 18, 14, 18, 192, 11,
	18, 3, 19, 3, 19, 3, 19, 7, 19, 197, 10, 19, 12, 19, 14, 19, 200, 11, 19,
	3, 20, 5, 20, 203, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5,
	21, 211, 10, 21, 3, 22, 3, 22, 3, 22, 5, 22, 216, 10, 22, 3, 22, 7, 22,
	219, 10, 22, 12, 22, 14, 22, 222, 11, 22, 3, 22, 3, 22, 3, 22, 5, 22, 227,
	10, 22, 3, 22, 3, 22, 5, 22, 231, 10, 22, 3, 22, 7, 22, 234, 10, 22, 12,
	22, 14, 22, 237, 11, 22, 7, 22, 239, 10, 22, 12, 22, 14, 22, 242, 11, 22,
	3, 23, 3, 23, 3, 23, 5, 23, 247, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 7, 3, 2, 5, 11, 3, 2,
	16, 17, 3, 2, 18, 19, 3, 2, 20, 22, 6, 2, 30, 30, 35, 35, 37, 37, 39, 41,
	2, 261, 2, 50, 3, 2, 2, 2, 4, 58, 3, 2, 2, 2, 6, 67, 3, 2, 2, 2, 8, 82,
	3, 2, 2, 2, 10, 99, 3, 2, 2, 2, 12, 101, 3, 2, 2, 2, 14, 105, 3, 2, 2,
	2, 16, 114, 3, 2, 2, 2, 18, 123, 3, 2, 2, 2, 20, 130, 3, 2, 2, 2, 22, 137,
	3, 2, 2, 2, 24, 147, 3, 2, 2, 2, 26, 151, 3, 2, 2, 2, 28, 161, 3, 2, 2,
	2, 30, 169, 3, 2, 2, 2, 32, 177, 3, 2, 2, 2, 34, 185, 3, 2, 2, 2, 36, 193,
	3, 2, 2, 2, 38, 202, 3, 2, 2, 2, 40, 210, 3, 2, 2, 2, 42, 212, 3, 2, 2,
	2, 44, 243, 3, 2, 2, 2, 46, 250, 3, 2, 2, 2, 48, 252, 3, 2, 2, 2, 50, 51,
	5, 4, 3, 2, 51, 56, 8, 2, 1, 2, 52, 53, 7, 36, 2, 2, 53, 54, 5, 2, 2, 2,
	54, 55, 8, 2, 1, 2, 55, 57, 3, 2, 2, 2, 56, 52, 3, 2, 2, 2, 56, 57, 3,
	2, 2, 2, 57, 3, 3, 2, 2, 2, 58, 59, 5, 6, 4, 2, 59, 64, 8, 3, 1, 2, 60,
	61, 7, 27, 2, 2, 61, 62, 5, 4, 3, 2, 62, 63, 8, 3, 1, 2, 63, 65, 3, 2,
	2, 2, 64, 60, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 5, 3, 2, 2, 2, 66, 68,
	7, 34, 2, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 70, 5, 8, 5, 2, 70, 7, 3, 2, 2, 2, 71, 72, 5, 10, 6, 2, 72, 73, 8,
	5, 1, 2, 73, 83, 3, 2, 2, 2, 74, 75, 7, 3, 2, 2, 75, 76, 5, 2, 2, 2, 76,
	77, 7, 4, 2, 2, 77, 78, 8, 5, 1, 2, 78, 83, 3, 2, 2, 2, 79, 80, 5, 26,
	14, 2, 80, 81, 8, 5, 1, 2, 81, 83, 3, 2, 2, 2, 82, 71, 3, 2, 2, 2, 82,
	74, 3, 2, 2, 2, 82, 79, 3, 2, 2, 2, 83, 9, 3, 2, 2, 2, 84, 85, 5, 12, 7,
	2, 85, 86, 8, 6, 1, 2, 86, 100, 3, 2, 2, 2, 87, 88, 5, 14, 8, 2, 88, 89,
	8, 6, 1, 2, 89, 100, 3, 2, 2, 2, 90, 91, 5, 16, 9, 2, 91, 92, 8, 6, 1,
	2, 92, 100, 3, 2, 2, 2, 93, 94, 5, 18, 10, 2, 94, 95, 8, 6, 1, 2, 95, 100,
	3, 2, 2, 2, 96, 97, 5, 20, 11, 2, 97, 98, 8, 6, 1, 2, 98, 100, 3, 2, 2,
	2, 99, 84, 3, 2, 2, 2, 99, 87, 3, 2, 2, 2, 99, 90, 3, 2, 2, 2, 99, 93,
	3, 2, 2, 2, 99, 96, 3, 2, 2, 2, 100, 11, 3, 2, 2, 2, 101, 102, 5, 26, 14,
	2, 102, 103, 9, 2, 2, 2, 103, 104, 5, 26, 14, 2, 104, 13, 3, 2, 2, 2, 105,
	107, 5, 26, 14, 2, 106, 108, 7, 34, 2, 2, 107, 106, 3, 2, 2, 2, 107, 108,
	3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 7, 28, 2, 2, 110, 111, 5, 26,
	14, 2, 111, 112, 7, 27, 2, 2, 112, 113, 5, 26, 14, 2, 113, 15, 3, 2, 2,
	2, 114, 116, 5, 26, 14, 2, 115, 117, 7, 34, 2, 2, 116, 115, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 7, 31, 2, 2, 119,
	120, 7, 3, 2, 2, 120, 121, 5, 22, 12, 2, 121, 122, 7, 4, 2, 2, 122, 17,
	3, 2, 2, 2, 123, 125, 5, 26, 14, 2, 124, 126, 7, 34, 2, 2, 125, 124, 3,
	2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 7, 33, 2,
	2, 128, 129, 5, 26, 14, 2, 129, 19, 3, 2, 2, 2, 130, 131, 5, 26, 14, 2,
	131, 133, 7, 32, 2, 2, 132, 134, 7, 34, 2, 2, 133, 132, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 7, 35, 2, 2, 136, 21,
	3, 2, 2, 2, 137, 142, 5, 26, 14, 2, 138, 139, 7, 12, 2, 2, 139, 141, 5,
	26, 14, 2, 140, 138, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2,
	2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2,
	145, 146, 8, 12, 1, 2, 146, 23, 3, 2, 2, 2, 147, 148, 7, 3, 2, 2, 148,
	149, 5, 26, 14, 2, 149, 150, 7, 4, 2, 2, 150, 25, 3, 2, 2, 2, 151, 156,
	5, 28, 15, 2, 152, 153, 7, 13, 2, 2, 153, 155, 5, 28, 15, 2, 154, 152,
	3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2,
	2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 8, 14, 1, 2,
	160, 27, 3, 2, 2, 2, 161, 166, 5, 30, 16, 2, 162, 163, 7, 14, 2, 2, 163,
	165, 5, 30, 16, 2, 164, 162, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164,
	3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 29, 3, 2, 2, 2, 168, 166, 3, 2,
	2, 2, 169, 174, 5, 32, 17, 2, 170, 171, 7, 15, 2, 2, 171, 173, 5, 32, 17,
	2, 172, 170, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174,
	175, 3, 2, 2, 2, 175, 31, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 182, 5,
	34, 18, 2, 178, 179, 9, 3, 2, 2, 179, 181, 5, 34, 18, 2, 180, 178, 3, 2,
	2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2,
	183, 33, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 190, 5, 36, 19, 2, 186,
	187, 9, 4, 2, 2, 187, 189, 5, 36, 19, 2, 188, 186, 3, 2, 2, 2, 189, 192,
	3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 35, 3, 2,
	2, 2, 192, 190, 3, 2, 2, 2, 193, 198, 5, 38, 20, 2, 194, 195, 9, 5, 2,
	2, 195, 197, 5, 38, 20, 2, 196, 194, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2,
	198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 37, 3, 2, 2, 2, 200, 198,
	3, 2, 2, 2, 201, 203, 7, 23, 2, 2, 202, 201, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 5, 40, 21, 2, 205, 39, 3, 2, 2, 2,
	206, 211, 5, 24, 13, 2, 207, 211, 5, 44, 23, 2, 208, 211, 5, 42, 22, 2,
	209, 211, 5, 48, 25, 2, 210, 206, 3, 2, 2, 2, 210, 207, 3, 2, 2, 2, 210,
	208, 3, 2, 2, 2, 210, 209, 3, 2, 2, 2, 211, 41, 3, 2, 2, 2, 212, 220, 5,
	46, 24, 2, 213, 215, 7, 24, 2, 2, 214, 216, 5, 26, 14, 2, 215, 214, 3,
	2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 219, 7, 25, 2,
	2, 218, 213, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220,
	221, 3, 2, 2, 2, 221, 240, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 226,
	7, 26, 2, 2, 224, 227, 5, 46, 24, 2, 225, 227, 5, 44, 23, 2, 226, 224,
	3, 2, 2, 2, 226, 225, 3, 2, 2, 2, 227, 235, 3, 2, 2, 2, 228, 230, 7, 24,
	2, 2, 229, 231, 5, 26, 14, 2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2,
	2, 231, 232, 3, 2, 2, 2, 232, 234, 7, 25, 2, 2, 233, 228, 3, 2, 2, 2, 234,
	237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 239,
	3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 223, 3, 2, 2, 2, 239, 242, 3, 2,
	2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 43, 3, 2, 2, 2,
	242, 240, 3, 2, 2, 2, 243, 244, 7, 42, 2, 2, 244, 246, 7, 3, 2, 2, 245,
	247, 5, 22, 12, 2, 246, 245, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248,
	3, 2, 2, 2, 248, 249, 7, 4, 2, 2, 249, 45, 3, 2, 2, 2, 250, 251, 7, 42,
	2, 2, 251, 47, 3, 2, 2, 2, 252, 253, 9, 6, 2, 2, 253, 49, 3, 2, 2, 2, 27,
	56, 64, 67, 82, 99, 107, 116, 125, 133, 142, 156, 166, 174, 182, 190, 198,
	202, 210, 215, 220, 226, 230, 235, 240, 246,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'='", "'<'", "'<='", "'>'", "'>='", "'<>'", "'!='",
	"','", "'|'", "'^'", "'&'", "'<<'", "'>>'", "'+'", "'-'", "'*'", "'/'",
	"'%'", "'~'", "'['", "']'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "AND", "BETWEEN", "EXISTS", "FALSE", "IN",
	"IS", "LIKE", "NOT", "NULL", "OR", "TRUE", "Percent", "IntegerLiteral",
	"FloatingPointLiteral", "StringLiteral", "Identifier", "WS",
}

var ruleNames = []string{
	"searchCondition", "booleanTerm", "booleanFactor", "booleanPrimary", "predicate",
	"comparisonPredicate", "betweenPredicate", "inPredicate", "likePredicate",
	"nullPredicate", "expressionList", "parExpression", "expression", "exclusiveOrExpression",
	"andExpression", "shiftExpression", "additiveExpression", "multiplicativeExpression",
	"notExpression", "primary", "member", "function", "variable", "constant",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type filterParser struct {
	*antlr.BaseParser
}

func NewfilterParser(input antlr.TokenStream) *filterParser {
	this := new(filterParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "filter.g4"

	return this
}

// filterParser tokens.
const (
	filterParserEOF                  = antlr.TokenEOF
	filterParserT__0                 = 1
	filterParserT__1                 = 2
	filterParserT__2                 = 3
	filterParserT__3                 = 4
	filterParserT__4                 = 5
	filterParserT__5                 = 6
	filterParserT__6                 = 7
	filterParserT__7                 = 8
	filterParserT__8                 = 9
	filterParserT__9                 = 10
	filterParserT__10                = 11
	filterParserT__11                = 12
	filterParserT__12                = 13
	filterParserT__13                = 14
	filterParserT__14                = 15
	filterParserT__15                = 16
	filterParserT__16                = 17
	filterParserT__17                = 18
	filterParserT__18                = 19
	filterParserT__19                = 20
	filterParserT__20                = 21
	filterParserT__21                = 22
	filterParserT__22                = 23
	filterParserT__23                = 24
	filterParserAND                  = 25
	filterParserBETWEEN              = 26
	filterParserEXISTS               = 27
	filterParserFALSE                = 28
	filterParserIN                   = 29
	filterParserIS                   = 30
	filterParserLIKE                 = 31
	filterParserNOT                  = 32
	filterParserNULL                 = 33
	filterParserOR                   = 34
	filterParserTRUE                 = 35
	filterParserPercent              = 36
	filterParserIntegerLiteral       = 37
	filterParserFloatingPointLiteral = 38
	filterParserStringLiteral        = 39
	filterParserIdentifier           = 40
	filterParserWS                   = 41
)

// filterParser rules.
const (
	filterParserRULE_searchCondition          = 0
	filterParserRULE_booleanTerm              = 1
	filterParserRULE_booleanFactor            = 2
	filterParserRULE_booleanPrimary           = 3
	filterParserRULE_predicate                = 4
	filterParserRULE_comparisonPredicate      = 5
	filterParserRULE_betweenPredicate         = 6
	filterParserRULE_inPredicate              = 7
	filterParserRULE_likePredicate            = 8
	filterParserRULE_nullPredicate            = 9
	filterParserRULE_expressionList           = 10
	filterParserRULE_parExpression            = 11
	filterParserRULE_expression               = 12
	filterParserRULE_exclusiveOrExpression    = 13
	filterParserRULE_andExpression            = 14
	filterParserRULE_shiftExpression          = 15
	filterParserRULE_additiveExpression       = 16
	filterParserRULE_multiplicativeExpression = 17
	filterParserRULE_notExpression            = 18
	filterParserRULE_primary                  = 19
	filterParserRULE_member                   = 20
	filterParserRULE_function                 = 21
	filterParserRULE_variable                 = 22
	filterParserRULE_constant                 = 23
)

// ISearchConditionContext is an interface to support dynamic dispatch.
type ISearchConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpCtx returns the opCtx token.
	GetOpCtx() antlr.Token

	// SetOpCtx sets the opCtx token.
	SetOpCtx(antlr.Token)

	// GetLCtx returns the lCtx rule contexts.
	GetLCtx() IBooleanTermContext

	// GetRCtx returns the rCtx rule contexts.
	GetRCtx() ISearchConditionContext

	// SetLCtx sets the lCtx rule contexts.
	SetLCtx(IBooleanTermContext)

	// SetRCtx sets the rCtx rule contexts.
	SetRCtx(ISearchConditionContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsSearchConditionContext differentiates from other interfaces.
	IsSearchConditionContext()
}

type SearchConditionContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lCtx              IBooleanTermContext
	opCtx             antlr.Token
	rCtx              ISearchConditionContext
}

func NewEmptySearchConditionContext() *SearchConditionContext {
	var p = new(SearchConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_searchCondition
	return p
}

func (*SearchConditionContext) IsSearchConditionContext() {}

func NewSearchConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchConditionContext {
	var p = new(SearchConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_searchCondition

	return p
}

func (s *SearchConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchConditionContext) GetOpCtx() antlr.Token { return s.opCtx }

func (s *SearchConditionContext) SetOpCtx(v antlr.Token) { s.opCtx = v }

func (s *SearchConditionContext) GetLCtx() IBooleanTermContext { return s.lCtx }

func (s *SearchConditionContext) GetRCtx() ISearchConditionContext { return s.rCtx }

func (s *SearchConditionContext) SetLCtx(v IBooleanTermContext) { s.lCtx = v }

func (s *SearchConditionContext) SetRCtx(v ISearchConditionContext) { s.rCtx = v }

func (s *SearchConditionContext) GetOperationMetadata() OperatableMetadata {
	return s.operationMetadata
}

func (s *SearchConditionContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *SearchConditionContext) BooleanTerm() IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *SearchConditionContext) OR() antlr.TerminalNode {
	return s.GetToken(filterParserOR, 0)
}

func (s *SearchConditionContext) SearchCondition() ISearchConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearchConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearchConditionContext)
}

func (s *SearchConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) SearchCondition() (localctx ISearchConditionContext) {
	localctx = NewSearchConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, filterParserRULE_searchCondition)
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
		p.SetState(48)

		var _x = p.BooleanTerm()

		localctx.(*SearchConditionContext).lCtx = _x
	}
	localctx.(*SearchConditionContext).operationMetadata = localctx.(*SearchConditionContext).GetLCtx().GetOperationMetadata()
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserOR {
		{
			p.SetState(50)

			var _m = p.Match(filterParserOR)

			localctx.(*SearchConditionContext).opCtx = _m
		}
		{
			p.SetState(51)

			var _x = p.SearchCondition()

			localctx.(*SearchConditionContext).rCtx = _x
		}

		localctx.(*SearchConditionContext).operationMetadata = OperationMetadata{BINARY, localctx.(*SearchConditionContext).GetOpCtx().GetText(), localctx.(*SearchConditionContext).operationMetadata, localctx.(*SearchConditionContext).GetRCtx().GetOperationMetadata()}

	}

	return localctx
}

// IBooleanTermContext is an interface to support dynamic dispatch.
type IBooleanTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpCtx returns the opCtx token.
	GetOpCtx() antlr.Token

	// SetOpCtx sets the opCtx token.
	SetOpCtx(antlr.Token)

	// GetLCtx returns the lCtx rule contexts.
	GetLCtx() IBooleanFactorContext

	// GetRCtx returns the rCtx rule contexts.
	GetRCtx() IBooleanTermContext

	// SetLCtx sets the lCtx rule contexts.
	SetLCtx(IBooleanFactorContext)

	// SetRCtx sets the rCtx rule contexts.
	SetRCtx(IBooleanTermContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsBooleanTermContext differentiates from other interfaces.
	IsBooleanTermContext()
}

type BooleanTermContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lCtx              IBooleanFactorContext
	opCtx             antlr.Token
	rCtx              IBooleanTermContext
}

func NewEmptyBooleanTermContext() *BooleanTermContext {
	var p = new(BooleanTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_booleanTerm
	return p
}

func (*BooleanTermContext) IsBooleanTermContext() {}

func NewBooleanTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTermContext {
	var p = new(BooleanTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_booleanTerm

	return p
}

func (s *BooleanTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanTermContext) GetOpCtx() antlr.Token { return s.opCtx }

func (s *BooleanTermContext) SetOpCtx(v antlr.Token) { s.opCtx = v }

func (s *BooleanTermContext) GetLCtx() IBooleanFactorContext { return s.lCtx }

func (s *BooleanTermContext) GetRCtx() IBooleanTermContext { return s.rCtx }

func (s *BooleanTermContext) SetLCtx(v IBooleanFactorContext) { s.lCtx = v }

func (s *BooleanTermContext) SetRCtx(v IBooleanTermContext) { s.rCtx = v }

func (s *BooleanTermContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *BooleanTermContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *BooleanTermContext) BooleanFactor() IBooleanFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanFactorContext)
}

func (s *BooleanTermContext) AND() antlr.TerminalNode {
	return s.GetToken(filterParserAND, 0)
}

func (s *BooleanTermContext) BooleanTerm() IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) BooleanTerm() (localctx IBooleanTermContext) {
	localctx = NewBooleanTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, filterParserRULE_booleanTerm)
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

		var _x = p.BooleanFactor()

		localctx.(*BooleanTermContext).lCtx = _x
	}
	localctx.(*BooleanTermContext).operationMetadata = localctx.(*BooleanTermContext).GetLCtx().GetOperationMetadata()
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserAND {
		{
			p.SetState(58)

			var _m = p.Match(filterParserAND)

			localctx.(*BooleanTermContext).opCtx = _m
		}
		{
			p.SetState(59)

			var _x = p.BooleanTerm()

			localctx.(*BooleanTermContext).rCtx = _x
		}

		localctx.(*BooleanTermContext).operationMetadata = OperationMetadata{BINARY, localctx.(*BooleanTermContext).GetOpCtx().GetText(), localctx.(*BooleanTermContext).operationMetadata, localctx.(*BooleanTermContext).GetRCtx().GetOperationMetadata()}

	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	if localctx.(*BooleanTermContext).operationMetadata == nil {
		localctx.(*BooleanTermContext).operationMetadata = localctx.(*BooleanTermContext).GetLCtx().GetOperationMetadata()
	}

	return localctx
}

// IBooleanFactorContext is an interface to support dynamic dispatch.
type IBooleanFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNotCtx returns the notCtx token.
	GetNotCtx() antlr.Token

	// SetNotCtx sets the notCtx token.
	SetNotCtx(antlr.Token)

	// GetRCtx returns the rCtx rule contexts.
	GetRCtx() IBooleanPrimaryContext

	// SetRCtx sets the rCtx rule contexts.
	SetRCtx(IBooleanPrimaryContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsBooleanFactorContext differentiates from other interfaces.
	IsBooleanFactorContext()
}

type BooleanFactorContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	notCtx            antlr.Token
	rCtx              IBooleanPrimaryContext
}

func NewEmptyBooleanFactorContext() *BooleanFactorContext {
	var p = new(BooleanFactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_booleanFactor
	return p
}

func (*BooleanFactorContext) IsBooleanFactorContext() {}

func NewBooleanFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFactorContext {
	var p = new(BooleanFactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_booleanFactor

	return p
}

func (s *BooleanFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanFactorContext) GetNotCtx() antlr.Token { return s.notCtx }

func (s *BooleanFactorContext) SetNotCtx(v antlr.Token) { s.notCtx = v }

func (s *BooleanFactorContext) GetRCtx() IBooleanPrimaryContext { return s.rCtx }

func (s *BooleanFactorContext) SetRCtx(v IBooleanPrimaryContext) { s.rCtx = v }

func (s *BooleanFactorContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *BooleanFactorContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *BooleanFactorContext) BooleanPrimary() IBooleanPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanPrimaryContext)
}

func (s *BooleanFactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(filterParserNOT, 0)
}

func (s *BooleanFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) BooleanFactor() (localctx IBooleanFactorContext) {
	localctx = NewBooleanFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, filterParserRULE_booleanFactor)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserNOT {
		{
			p.SetState(64)

			var _m = p.Match(filterParserNOT)

			localctx.(*BooleanFactorContext).notCtx = _m
		}

	}
	{
		p.SetState(67)

		var _x = p.BooleanPrimary()

		localctx.(*BooleanFactorContext).rCtx = _x
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	if localctx.(*BooleanFactorContext).GetNotCtx() != nil {
		localctx.(*BooleanFactorContext).operationMetadata = OperationMetadata{UNARY, localctx.(*BooleanFactorContext).GetNotCtx().GetText(), nil, localctx.(*BooleanFactorContext).GetRCtx().GetOperationMetadata()}
	} else {
		localctx.(*BooleanFactorContext).operationMetadata = localctx.(*BooleanFactorContext).GetRCtx().GetOperationMetadata()
	}

	return localctx
}

// IBooleanPrimaryContext is an interface to support dynamic dispatch.
type IBooleanPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPredicateCtx returns the predicateCtx rule contexts.
	GetPredicateCtx() IPredicateContext

	// GetConditionCtx returns the conditionCtx rule contexts.
	GetConditionCtx() ISearchConditionContext

	// GetExprCtx returns the exprCtx rule contexts.
	GetExprCtx() IExpressionContext

	// SetPredicateCtx sets the predicateCtx rule contexts.
	SetPredicateCtx(IPredicateContext)

	// SetConditionCtx sets the conditionCtx rule contexts.
	SetConditionCtx(ISearchConditionContext)

	// SetExprCtx sets the exprCtx rule contexts.
	SetExprCtx(IExpressionContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsBooleanPrimaryContext differentiates from other interfaces.
	IsBooleanPrimaryContext()
}

type BooleanPrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	predicateCtx      IPredicateContext
	conditionCtx      ISearchConditionContext
	exprCtx           IExpressionContext
}

func NewEmptyBooleanPrimaryContext() *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_booleanPrimary
	return p
}

func (*BooleanPrimaryContext) IsBooleanPrimaryContext() {}

func NewBooleanPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_booleanPrimary

	return p
}

func (s *BooleanPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanPrimaryContext) GetPredicateCtx() IPredicateContext { return s.predicateCtx }

func (s *BooleanPrimaryContext) GetConditionCtx() ISearchConditionContext { return s.conditionCtx }

func (s *BooleanPrimaryContext) GetExprCtx() IExpressionContext { return s.exprCtx }

func (s *BooleanPrimaryContext) SetPredicateCtx(v IPredicateContext) { s.predicateCtx = v }

func (s *BooleanPrimaryContext) SetConditionCtx(v ISearchConditionContext) { s.conditionCtx = v }

func (s *BooleanPrimaryContext) SetExprCtx(v IExpressionContext) { s.exprCtx = v }

func (s *BooleanPrimaryContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *BooleanPrimaryContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *BooleanPrimaryContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *BooleanPrimaryContext) SearchCondition() ISearchConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearchConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearchConditionContext)
}

func (s *BooleanPrimaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BooleanPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) BooleanPrimary() (localctx IBooleanPrimaryContext) {
	localctx = NewBooleanPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, filterParserRULE_booleanPrimary)

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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)

			var _x = p.Predicate()

			localctx.(*BooleanPrimaryContext).predicateCtx = _x
		}
		localctx.(*BooleanPrimaryContext).SetOperationMetadata(localctx.(*BooleanPrimaryContext).GetPredicateCtx().GetOperationMetadata())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(filterParserT__0)
		}
		{
			p.SetState(73)

			var _x = p.SearchCondition()

			localctx.(*BooleanPrimaryContext).conditionCtx = _x
		}
		{
			p.SetState(74)
			p.Match(filterParserT__1)
		}
		localctx.(*BooleanPrimaryContext).operationMetadata = OperationMetadata{UNARY, PAREN, nil, localctx.(*BooleanPrimaryContext).GetConditionCtx().GetOperationMetadata()}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)

			var _x = p.Expression()

			localctx.(*BooleanPrimaryContext).exprCtx = _x
		}
		localctx.(*BooleanPrimaryContext).operationMetadata = RelationOperationMetadata{UNARY, EXPR, "", localctx.(*BooleanPrimaryContext).GetExprCtx().GetExpressionText()}

	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComparisonMetadata returns the comparisonMetadata rule contexts.
	GetComparisonMetadata() IComparisonPredicateContext

	// GetBetweenMetadata returns the betweenMetadata rule contexts.
	GetBetweenMetadata() IBetweenPredicateContext

	// GetInMetadata returns the inMetadata rule contexts.
	GetInMetadata() IInPredicateContext

	// GetLikeMetadata returns the likeMetadata rule contexts.
	GetLikeMetadata() ILikePredicateContext

	// GetNullMetadata returns the nullMetadata rule contexts.
	GetNullMetadata() INullPredicateContext

	// SetComparisonMetadata sets the comparisonMetadata rule contexts.
	SetComparisonMetadata(IComparisonPredicateContext)

	// SetBetweenMetadata sets the betweenMetadata rule contexts.
	SetBetweenMetadata(IBetweenPredicateContext)

	// SetInMetadata sets the inMetadata rule contexts.
	SetInMetadata(IInPredicateContext)

	// SetLikeMetadata sets the likeMetadata rule contexts.
	SetLikeMetadata(ILikePredicateContext)

	// SetNullMetadata sets the nullMetadata rule contexts.
	SetNullMetadata(INullPredicateContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	operationMetadata  OperatableMetadata
	comparisonMetadata IComparisonPredicateContext
	betweenMetadata    IBetweenPredicateContext
	inMetadata         IInPredicateContext
	likeMetadata       ILikePredicateContext
	nullMetadata       INullPredicateContext
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) GetComparisonMetadata() IComparisonPredicateContext {
	return s.comparisonMetadata
}

func (s *PredicateContext) GetBetweenMetadata() IBetweenPredicateContext { return s.betweenMetadata }

func (s *PredicateContext) GetInMetadata() IInPredicateContext { return s.inMetadata }

func (s *PredicateContext) GetLikeMetadata() ILikePredicateContext { return s.likeMetadata }

func (s *PredicateContext) GetNullMetadata() INullPredicateContext { return s.nullMetadata }

func (s *PredicateContext) SetComparisonMetadata(v IComparisonPredicateContext) {
	s.comparisonMetadata = v
}

func (s *PredicateContext) SetBetweenMetadata(v IBetweenPredicateContext) { s.betweenMetadata = v }

func (s *PredicateContext) SetInMetadata(v IInPredicateContext) { s.inMetadata = v }

func (s *PredicateContext) SetLikeMetadata(v ILikePredicateContext) { s.likeMetadata = v }

func (s *PredicateContext) SetNullMetadata(v INullPredicateContext) { s.nullMetadata = v }

func (s *PredicateContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *PredicateContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *PredicateContext) ComparisonPredicate() IComparisonPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonPredicateContext)
}

func (s *PredicateContext) BetweenPredicate() IBetweenPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBetweenPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBetweenPredicateContext)
}

func (s *PredicateContext) InPredicate() IInPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInPredicateContext)
}

func (s *PredicateContext) LikePredicate() ILikePredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikePredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikePredicateContext)
}

func (s *PredicateContext) NullPredicate() INullPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullPredicateContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, filterParserRULE_predicate)

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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)

			var _x = p.ComparisonPredicate()

			localctx.(*PredicateContext).comparisonMetadata = _x
		}
		localctx.(*PredicateContext).operationMetadata = localctx.(*PredicateContext).GetComparisonMetadata().GetOperationMetadata()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)

			var _x = p.BetweenPredicate()

			localctx.(*PredicateContext).betweenMetadata = _x
		}
		localctx.(*PredicateContext).operationMetadata = localctx.(*PredicateContext).GetBetweenMetadata().GetOperationMetadata()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)

			var _x = p.InPredicate()

			localctx.(*PredicateContext).inMetadata = _x
		}
		localctx.(*PredicateContext).operationMetadata = localctx.(*PredicateContext).GetInMetadata().GetOperationMetadata()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)

			var _x = p.LikePredicate()

			localctx.(*PredicateContext).likeMetadata = _x
		}
		localctx.(*PredicateContext).operationMetadata = localctx.(*PredicateContext).GetLikeMetadata().GetOperationMetadata()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)

			var _x = p.NullPredicate()

			localctx.(*PredicateContext).nullMetadata = _x
		}
		localctx.(*PredicateContext).operationMetadata = localctx.(*PredicateContext).GetNullMetadata().GetOperationMetadata()

	}

	return localctx
}

// IComparisonPredicateContext is an interface to support dynamic dispatch.
type IComparisonPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpCtx returns the opCtx token.
	GetOpCtx() antlr.Token

	// SetOpCtx sets the opCtx token.
	SetOpCtx(antlr.Token)

	// GetLExprCtx returns the lExprCtx rule contexts.
	GetLExprCtx() IExpressionContext

	// GetRExprCtx returns the rExprCtx rule contexts.
	GetRExprCtx() IExpressionContext

	// SetLExprCtx sets the lExprCtx rule contexts.
	SetLExprCtx(IExpressionContext)

	// SetRExprCtx sets the rExprCtx rule contexts.
	SetRExprCtx(IExpressionContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsComparisonPredicateContext differentiates from other interfaces.
	IsComparisonPredicateContext()
}

type ComparisonPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lExprCtx          IExpressionContext
	opCtx             antlr.Token
	rExprCtx          IExpressionContext
}

func NewEmptyComparisonPredicateContext() *ComparisonPredicateContext {
	var p = new(ComparisonPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_comparisonPredicate
	return p
}

func (*ComparisonPredicateContext) IsComparisonPredicateContext() {}

func NewComparisonPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonPredicateContext {
	var p = new(ComparisonPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_comparisonPredicate

	return p
}

func (s *ComparisonPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonPredicateContext) GetOpCtx() antlr.Token { return s.opCtx }

func (s *ComparisonPredicateContext) SetOpCtx(v antlr.Token) { s.opCtx = v }

func (s *ComparisonPredicateContext) GetLExprCtx() IExpressionContext { return s.lExprCtx }

func (s *ComparisonPredicateContext) GetRExprCtx() IExpressionContext { return s.rExprCtx }

func (s *ComparisonPredicateContext) SetLExprCtx(v IExpressionContext) { s.lExprCtx = v }

func (s *ComparisonPredicateContext) SetRExprCtx(v IExpressionContext) { s.rExprCtx = v }

func (s *ComparisonPredicateContext) GetOperationMetadata() OperatableMetadata {
	return s.operationMetadata
}

func (s *ComparisonPredicateContext) SetOperationMetadata(v OperatableMetadata) {
	s.operationMetadata = v
}

func (s *ComparisonPredicateContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ComparisonPredicateContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) ComparisonPredicate() (localctx IComparisonPredicateContext) {
	localctx = NewComparisonPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, filterParserRULE_comparisonPredicate)
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
		p.SetState(99)

		var _x = p.Expression()

		localctx.(*ComparisonPredicateContext).lExprCtx = _x
	}
	{
		p.SetState(100)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ComparisonPredicateContext).opCtx = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserT__2)|(1<<filterParserT__3)|(1<<filterParserT__4)|(1<<filterParserT__5)|(1<<filterParserT__6)|(1<<filterParserT__7)|(1<<filterParserT__8))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ComparisonPredicateContext).opCtx = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(101)

		var _x = p.Expression()

		localctx.(*ComparisonPredicateContext).rExprCtx = _x
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	localctx.(*ComparisonPredicateContext).SetOperationMetadata(RelationOperationMetadata{BINARY, localctx.(*ComparisonPredicateContext).GetOpCtx().GetText(), localctx.(*ComparisonPredicateContext).GetLExprCtx().GetExpressionText(), localctx.(*ComparisonPredicateContext).GetRExprCtx().GetExpressionText()})

	return localctx
}

// IBetweenPredicateContext is an interface to support dynamic dispatch.
type IBetweenPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNotCtx returns the notCtx token.
	GetNotCtx() antlr.Token

	// SetNotCtx sets the notCtx token.
	SetNotCtx(antlr.Token)

	// GetLExprCtx returns the lExprCtx rule contexts.
	GetLExprCtx() IExpressionContext

	// GetRExpr1Ctx returns the rExpr1Ctx rule contexts.
	GetRExpr1Ctx() IExpressionContext

	// GetRExpr2Ctx returns the rExpr2Ctx rule contexts.
	GetRExpr2Ctx() IExpressionContext

	// SetLExprCtx sets the lExprCtx rule contexts.
	SetLExprCtx(IExpressionContext)

	// SetRExpr1Ctx sets the rExpr1Ctx rule contexts.
	SetRExpr1Ctx(IExpressionContext)

	// SetRExpr2Ctx sets the rExpr2Ctx rule contexts.
	SetRExpr2Ctx(IExpressionContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsBetweenPredicateContext differentiates from other interfaces.
	IsBetweenPredicateContext()
}

type BetweenPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lExprCtx          IExpressionContext
	notCtx            antlr.Token
	rExpr1Ctx         IExpressionContext
	rExpr2Ctx         IExpressionContext
}

func NewEmptyBetweenPredicateContext() *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_betweenPredicate
	return p
}

func (*BetweenPredicateContext) IsBetweenPredicateContext() {}

func NewBetweenPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_betweenPredicate

	return p
}

func (s *BetweenPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BetweenPredicateContext) GetNotCtx() antlr.Token { return s.notCtx }

func (s *BetweenPredicateContext) SetNotCtx(v antlr.Token) { s.notCtx = v }

func (s *BetweenPredicateContext) GetLExprCtx() IExpressionContext { return s.lExprCtx }

func (s *BetweenPredicateContext) GetRExpr1Ctx() IExpressionContext { return s.rExpr1Ctx }

func (s *BetweenPredicateContext) GetRExpr2Ctx() IExpressionContext { return s.rExpr2Ctx }

func (s *BetweenPredicateContext) SetLExprCtx(v IExpressionContext) { s.lExprCtx = v }

func (s *BetweenPredicateContext) SetRExpr1Ctx(v IExpressionContext) { s.rExpr1Ctx = v }

func (s *BetweenPredicateContext) SetRExpr2Ctx(v IExpressionContext) { s.rExpr2Ctx = v }

func (s *BetweenPredicateContext) GetOperationMetadata() OperatableMetadata {
	return s.operationMetadata
}

func (s *BetweenPredicateContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *BetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(filterParserBETWEEN, 0)
}

func (s *BetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(filterParserAND, 0)
}

func (s *BetweenPredicateContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BetweenPredicateContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(filterParserNOT, 0)
}

func (s *BetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) BetweenPredicate() (localctx IBetweenPredicateContext) {
	localctx = NewBetweenPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, filterParserRULE_betweenPredicate)
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
		p.SetState(103)

		var _x = p.Expression()

		localctx.(*BetweenPredicateContext).lExprCtx = _x
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserNOT {
		{
			p.SetState(104)

			var _m = p.Match(filterParserNOT)

			localctx.(*BetweenPredicateContext).notCtx = _m
		}

	}
	{
		p.SetState(107)
		p.Match(filterParserBETWEEN)
	}
	{
		p.SetState(108)

		var _x = p.Expression()

		localctx.(*BetweenPredicateContext).rExpr1Ctx = _x
	}
	{
		p.SetState(109)
		p.Match(filterParserAND)
	}
	{
		p.SetState(110)

		var _x = p.Expression()

		localctx.(*BetweenPredicateContext).rExpr2Ctx = _x
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	rExpr := fmt.Sprintf("(%s,%s)", localctx.(*BetweenPredicateContext).GetRExpr1Ctx().GetExpressionText(), localctx.(*BetweenPredicateContext).GetRExpr2Ctx().GetExpressionText())
	localctx.(*BetweenPredicateContext).SetOperationMetadata(RelationOperationMetadata{BINARY, BETWEEN, localctx.(*BetweenPredicateContext).GetLExprCtx().GetExpressionText(), rExpr})
	if localctx.(*BetweenPredicateContext).GetNotCtx() != nil {
		localctx.(*BetweenPredicateContext).SetOperationMetadata(OperationMetadata{UNARY, localctx.(*BetweenPredicateContext).GetNotCtx().GetText(), nil, localctx.(*BetweenPredicateContext).operationMetadata})
	}

	return localctx
}

// IInPredicateContext is an interface to support dynamic dispatch.
type IInPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNotCtx returns the notCtx token.
	GetNotCtx() antlr.Token

	// GetLrCtx returns the lrCtx token.
	GetLrCtx() antlr.Token

	// GetRrCtx returns the rrCtx token.
	GetRrCtx() antlr.Token

	// SetNotCtx sets the notCtx token.
	SetNotCtx(antlr.Token)

	// SetLrCtx sets the lrCtx token.
	SetLrCtx(antlr.Token)

	// SetRrCtx sets the rrCtx token.
	SetRrCtx(antlr.Token)

	// GetLExprCtx returns the lExprCtx rule contexts.
	GetLExprCtx() IExpressionContext

	// GetRExprCtx returns the rExprCtx rule contexts.
	GetRExprCtx() IExpressionListContext

	// SetLExprCtx sets the lExprCtx rule contexts.
	SetLExprCtx(IExpressionContext)

	// SetRExprCtx sets the rExprCtx rule contexts.
	SetRExprCtx(IExpressionListContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsInPredicateContext differentiates from other interfaces.
	IsInPredicateContext()
}

type InPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lExprCtx          IExpressionContext
	notCtx            antlr.Token
	lrCtx             antlr.Token
	rExprCtx          IExpressionListContext
	rrCtx             antlr.Token
}

func NewEmptyInPredicateContext() *InPredicateContext {
	var p = new(InPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_inPredicate
	return p
}

func (*InPredicateContext) IsInPredicateContext() {}

func NewInPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InPredicateContext {
	var p = new(InPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_inPredicate

	return p
}

func (s *InPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *InPredicateContext) GetNotCtx() antlr.Token { return s.notCtx }

func (s *InPredicateContext) GetLrCtx() antlr.Token { return s.lrCtx }

func (s *InPredicateContext) GetRrCtx() antlr.Token { return s.rrCtx }

func (s *InPredicateContext) SetNotCtx(v antlr.Token) { s.notCtx = v }

func (s *InPredicateContext) SetLrCtx(v antlr.Token) { s.lrCtx = v }

func (s *InPredicateContext) SetRrCtx(v antlr.Token) { s.rrCtx = v }

func (s *InPredicateContext) GetLExprCtx() IExpressionContext { return s.lExprCtx }

func (s *InPredicateContext) GetRExprCtx() IExpressionListContext { return s.rExprCtx }

func (s *InPredicateContext) SetLExprCtx(v IExpressionContext) { s.lExprCtx = v }

func (s *InPredicateContext) SetRExprCtx(v IExpressionListContext) { s.rExprCtx = v }

func (s *InPredicateContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *InPredicateContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *InPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(filterParserIN, 0)
}

func (s *InPredicateContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InPredicateContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *InPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(filterParserNOT, 0)
}

func (s *InPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) InPredicate() (localctx IInPredicateContext) {
	localctx = NewInPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, filterParserRULE_inPredicate)
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
		p.SetState(112)

		var _x = p.Expression()

		localctx.(*InPredicateContext).lExprCtx = _x
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserNOT {
		{
			p.SetState(113)

			var _m = p.Match(filterParserNOT)

			localctx.(*InPredicateContext).notCtx = _m
		}

	}
	{
		p.SetState(116)
		p.Match(filterParserIN)
	}
	{
		p.SetState(117)

		var _m = p.Match(filterParserT__0)

		localctx.(*InPredicateContext).lrCtx = _m
	}
	{
		p.SetState(118)

		var _x = p.ExpressionList()

		localctx.(*InPredicateContext).rExprCtx = _x
	}
	{
		p.SetState(119)

		var _m = p.Match(filterParserT__1)

		localctx.(*InPredicateContext).rrCtx = _m
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	inExpressionText := fmt.Sprintf("%s%s%s", localctx.(*InPredicateContext).GetLrCtx().GetText(), localctx.(*InPredicateContext).GetRExprCtx().GetExpressionText(), localctx.(*InPredicateContext).GetRrCtx().GetText())
	localctx.(*InPredicateContext).operationMetadata = RelationOperationMetadata{BINARY, IN, localctx.(*InPredicateContext).GetLExprCtx().GetExpressionText(), inExpressionText}
	if localctx.(*InPredicateContext).GetNotCtx() != nil {
		localctx.(*InPredicateContext).operationMetadata = OperationMetadata{UNARY, localctx.(*InPredicateContext).GetNotCtx().GetText(), nil, localctx.(*InPredicateContext).operationMetadata}
	}

	return localctx
}

// ILikePredicateContext is an interface to support dynamic dispatch.
type ILikePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNotCtx returns the notCtx token.
	GetNotCtx() antlr.Token

	// SetNotCtx sets the notCtx token.
	SetNotCtx(antlr.Token)

	// GetLExprCtx returns the lExprCtx rule contexts.
	GetLExprCtx() IExpressionContext

	// GetRExprCtx returns the rExprCtx rule contexts.
	GetRExprCtx() IExpressionContext

	// SetLExprCtx sets the lExprCtx rule contexts.
	SetLExprCtx(IExpressionContext)

	// SetRExprCtx sets the rExprCtx rule contexts.
	SetRExprCtx(IExpressionContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsLikePredicateContext differentiates from other interfaces.
	IsLikePredicateContext()
}

type LikePredicateContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lExprCtx          IExpressionContext
	notCtx            antlr.Token
	rExprCtx          IExpressionContext
}

func NewEmptyLikePredicateContext() *LikePredicateContext {
	var p = new(LikePredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_likePredicate
	return p
}

func (*LikePredicateContext) IsLikePredicateContext() {}

func NewLikePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_likePredicate

	return p
}

func (s *LikePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *LikePredicateContext) GetNotCtx() antlr.Token { return s.notCtx }

func (s *LikePredicateContext) SetNotCtx(v antlr.Token) { s.notCtx = v }

func (s *LikePredicateContext) GetLExprCtx() IExpressionContext { return s.lExprCtx }

func (s *LikePredicateContext) GetRExprCtx() IExpressionContext { return s.rExprCtx }

func (s *LikePredicateContext) SetLExprCtx(v IExpressionContext) { s.lExprCtx = v }

func (s *LikePredicateContext) SetRExprCtx(v IExpressionContext) { s.rExprCtx = v }

func (s *LikePredicateContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *LikePredicateContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(filterParserLIKE, 0)
}

func (s *LikePredicateContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LikePredicateContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(filterParserNOT, 0)
}

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) LikePredicate() (localctx ILikePredicateContext) {
	localctx = NewLikePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, filterParserRULE_likePredicate)
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
		p.SetState(121)

		var _x = p.Expression()

		localctx.(*LikePredicateContext).lExprCtx = _x
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserNOT {
		{
			p.SetState(122)

			var _m = p.Match(filterParserNOT)

			localctx.(*LikePredicateContext).notCtx = _m
		}

	}
	{
		p.SetState(125)
		p.Match(filterParserLIKE)
	}
	{
		p.SetState(126)

		var _x = p.Expression()

		localctx.(*LikePredicateContext).rExprCtx = _x
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	localctx.(*LikePredicateContext).SetOperationMetadata(RelationOperationMetadata{BINARY, LIKE, localctx.(*LikePredicateContext).GetLExprCtx().GetExpressionText(), localctx.(*LikePredicateContext).GetRExprCtx().GetExpressionText()})
	if localctx.(*LikePredicateContext).GetNotCtx() != nil {
		localctx.(*LikePredicateContext).operationMetadata = OperationMetadata{UNARY, localctx.(*LikePredicateContext).GetNotCtx().GetText(), nil, localctx.(*LikePredicateContext).operationMetadata}
	}

	return localctx
}

// INullPredicateContext is an interface to support dynamic dispatch.
type INullPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNotCtx returns the notCtx token.
	GetNotCtx() antlr.Token

	// GetRExprCtx returns the rExprCtx token.
	GetRExprCtx() antlr.Token

	// SetNotCtx sets the notCtx token.
	SetNotCtx(antlr.Token)

	// SetRExprCtx sets the rExprCtx token.
	SetRExprCtx(antlr.Token)

	// GetLExprCtx returns the lExprCtx rule contexts.
	GetLExprCtx() IExpressionContext

	// SetLExprCtx sets the lExprCtx rule contexts.
	SetLExprCtx(IExpressionContext)

	// GetOperationMetadata returns the operationMetadata attribute.
	GetOperationMetadata() OperatableMetadata

	// SetOperationMetadata sets the operationMetadata attribute.
	SetOperationMetadata(OperatableMetadata)

	// IsNullPredicateContext differentiates from other interfaces.
	IsNullPredicateContext()
}

type NullPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	operationMetadata OperatableMetadata
	lExprCtx          IExpressionContext
	notCtx            antlr.Token
	rExprCtx          antlr.Token
}

func NewEmptyNullPredicateContext() *NullPredicateContext {
	var p = new(NullPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_nullPredicate
	return p
}

func (*NullPredicateContext) IsNullPredicateContext() {}

func NewNullPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullPredicateContext {
	var p = new(NullPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_nullPredicate

	return p
}

func (s *NullPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *NullPredicateContext) GetNotCtx() antlr.Token { return s.notCtx }

func (s *NullPredicateContext) GetRExprCtx() antlr.Token { return s.rExprCtx }

func (s *NullPredicateContext) SetNotCtx(v antlr.Token) { s.notCtx = v }

func (s *NullPredicateContext) SetRExprCtx(v antlr.Token) { s.rExprCtx = v }

func (s *NullPredicateContext) GetLExprCtx() IExpressionContext { return s.lExprCtx }

func (s *NullPredicateContext) SetLExprCtx(v IExpressionContext) { s.lExprCtx = v }

func (s *NullPredicateContext) GetOperationMetadata() OperatableMetadata { return s.operationMetadata }

func (s *NullPredicateContext) SetOperationMetadata(v OperatableMetadata) { s.operationMetadata = v }

func (s *NullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(filterParserIS, 0)
}

func (s *NullPredicateContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NullPredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(filterParserNULL, 0)
}

func (s *NullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(filterParserNOT, 0)
}

func (s *NullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) NullPredicate() (localctx INullPredicateContext) {
	localctx = NewNullPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, filterParserRULE_nullPredicate)
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
		p.SetState(128)

		var _x = p.Expression()

		localctx.(*NullPredicateContext).lExprCtx = _x
	}
	{
		p.SetState(129)
		p.Match(filterParserIS)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserNOT {
		{
			p.SetState(130)

			var _m = p.Match(filterParserNOT)

			localctx.(*NullPredicateContext).notCtx = _m
		}

	}
	{
		p.SetState(133)

		var _m = p.Match(filterParserNULL)

		localctx.(*NullPredicateContext).rExprCtx = _m
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	localctx.(*NullPredicateContext).operationMetadata = RelationOperationMetadata{BINARY, IS, localctx.(*NullPredicateContext).GetLExprCtx().GetExpressionText(), localctx.(*NullPredicateContext).GetRExprCtx().GetText()}
	if localctx.(*NullPredicateContext).GetNotCtx() != nil {
		localctx.(*NullPredicateContext).operationMetadata = OperationMetadata{UNARY, localctx.(*NullPredicateContext).GetNotCtx().GetText(), nil, localctx.(*NullPredicateContext).operationMetadata}
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpressionText returns the expressionText attribute.
	GetExpressionText() string

	// SetExpressionText sets the expressionText attribute.
	SetExpressionText(string)

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	expressionText string
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) GetExpressionText() string { return s.expressionText }

func (s *ExpressionListContext) SetExpressionText(v string) { s.expressionText = v }

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

func (p *filterParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, filterParserRULE_expressionList)
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
		p.SetState(135)
		p.Expression()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__9 {
		{
			p.SetState(136)
			p.Match(filterParserT__9)
		}
		{
			p.SetState(137)
			p.Expression()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	localctx.(*ExpressionListContext).expressionText = p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1))

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, filterParserRULE_parExpression)

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
		p.SetState(145)
		p.Match(filterParserT__0)
	}
	{
		p.SetState(146)
		p.Expression()
	}
	{
		p.SetState(147)
		p.Match(filterParserT__1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpressionText returns the expressionText attribute.
	GetExpressionText() string

	// SetExpressionText sets the expressionText attribute.
	SetExpressionText(string)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	expressionText string
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetExpressionText() string { return s.expressionText }

func (s *ExpressionContext) SetExpressionText(v string) { s.expressionText = v }

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

func (p *filterParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, filterParserRULE_expression)
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
		p.SetState(149)
		p.ExclusiveOrExpression()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__10 {
		{
			p.SetState(150)
			p.Match(filterParserT__10)
		}
		{
			p.SetState(151)
			p.ExclusiveOrExpression()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	localctx.(*ExpressionContext).expressionText = p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1))

	return localctx
}

// IExclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IExclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExclusiveOrExpressionContext differentiates from other interfaces.
	IsExclusiveOrExpressionContext()
}

type ExclusiveOrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusiveOrExpressionContext() *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_exclusiveOrExpression
	return p
}

func (*ExclusiveOrExpressionContext) IsExclusiveOrExpressionContext() {}

func NewExclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_exclusiveOrExpression

	return p
}

func (s *ExclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) ExclusiveOrExpression() (localctx IExclusiveOrExpressionContext) {
	localctx = NewExclusiveOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, filterParserRULE_exclusiveOrExpression)
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
		p.SetState(159)
		p.AndExpression()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__11 {
		{
			p.SetState(160)
			p.Match(filterParserT__11)
		}
		{
			p.SetState(161)
			p.AndExpression()
		}

		p.SetState(166)
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

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, filterParserRULE_andExpression)
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
		p.SetState(167)
		p.ShiftExpression()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__12 {
		{
			p.SetState(168)
			p.Match(filterParserT__12)
		}
		{
			p.SetState(169)
			p.ShiftExpression()
		}

		p.SetState(174)
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

	// IsShiftExpressionContext differentiates from other interfaces.
	IsShiftExpressionContext()
}

type ShiftExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftExpressionContext() *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_shiftExpression
	return p
}

func (*ShiftExpressionContext) IsShiftExpressionContext() {}

func NewShiftExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_shiftExpression

	return p
}

func (s *ShiftExpressionContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) ShiftExpression() (localctx IShiftExpressionContext) {
	localctx = NewShiftExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, filterParserRULE_shiftExpression)
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
		p.SetState(175)
		p.AdditiveExpression()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__13 || _la == filterParserT__14 {
		{
			p.SetState(176)
			_la = p.GetTokenStream().LA(1)

			if !(_la == filterParserT__13 || _la == filterParserT__14) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(177)
			p.AdditiveExpression()
		}

		p.SetState(182)
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

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, filterParserRULE_additiveExpression)
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
		p.SetState(183)
		p.MultiplicativeExpression()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__15 || _la == filterParserT__16 {
		{
			p.SetState(184)
			_la = p.GetTokenStream().LA(1)

			if !(_la == filterParserT__15 || _la == filterParserT__16) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(185)
			p.MultiplicativeExpression()
		}

		p.SetState(190)
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

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, filterParserRULE_multiplicativeExpression)
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
		p.SetState(191)
		p.NotExpression()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserT__17)|(1<<filterParserT__18)|(1<<filterParserT__19))) != 0 {
		{
			p.SetState(192)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserT__17)|(1<<filterParserT__18)|(1<<filterParserT__19))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(193)
			p.NotExpression()
		}

		p.SetState(198)
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

	// GetExp returns the exp rule contexts.
	GetExp() IPrimaryContext

	// SetExp sets the exp rule contexts.
	SetExp(IPrimaryContext)

	// IsNotExpressionContext differentiates from other interfaces.
	IsNotExpressionContext()
}

type NotExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	exp    IPrimaryContext
}

func NewEmptyNotExpressionContext() *NotExpressionContext {
	var p = new(NotExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_notExpression
	return p
}

func (*NotExpressionContext) IsNotExpressionContext() {}

func NewNotExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_notExpression

	return p
}

func (s *NotExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NotExpressionContext) GetExp() IPrimaryContext { return s.exp }

func (s *NotExpressionContext) SetExp(v IPrimaryContext) { s.exp = v }

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

func (p *filterParser) NotExpression() (localctx INotExpressionContext) {
	localctx = NewNotExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, filterParserRULE_notExpression)
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
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserT__20 {
		{
			p.SetState(199)
			p.Match(filterParserT__20)
		}

	}
	{
		p.SetState(202)

		var _x = p.Primary()

		localctx.(*NotExpressionContext).exp = _x
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

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

func (p *filterParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, filterParserRULE_primary)

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

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.ParExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Function()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
			p.Member()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(207)
			p.Constant()
		}

	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

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

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, filterParserRULE_member)
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
		p.SetState(210)
		p.Variable()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__21 {
		{
			p.SetState(211)
			p.Match(filterParserT__21)
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserT__0)|(1<<filterParserT__20)|(1<<filterParserFALSE))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(filterParserNULL-33))|(1<<(filterParserTRUE-33))|(1<<(filterParserIntegerLiteral-33))|(1<<(filterParserFloatingPointLiteral-33))|(1<<(filterParserStringLiteral-33))|(1<<(filterParserIdentifier-33)))) != 0) {
			{
				p.SetState(212)
				p.Expression()
			}

		}
		{
			p.SetState(215)
			p.Match(filterParserT__22)
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == filterParserT__23 {
		{
			p.SetState(221)
			p.Match(filterParserT__23)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(222)
				p.Variable()
			}

		case 2:
			{
				p.SetState(223)
				p.Function()
			}

		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == filterParserT__21 {
			{
				p.SetState(226)
				p.Match(filterParserT__21)
			}
			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserT__0)|(1<<filterParserT__20)|(1<<filterParserFALSE))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(filterParserNULL-33))|(1<<(filterParserTRUE-33))|(1<<(filterParserIntegerLiteral-33))|(1<<(filterParserFloatingPointLiteral-33))|(1<<(filterParserStringLiteral-33))|(1<<(filterParserIdentifier-33)))) != 0) {
				{
					p.SetState(227)
					p.Expression()
				}

			}
			{
				p.SetState(230)
				p.Match(filterParserT__22)
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(240)
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

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(filterParserIdentifier, 0)
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

func (p *filterParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, filterParserRULE_function)
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
		p.SetState(241)
		p.Match(filterParserIdentifier)
	}
	{
		p.SetState(242)
		p.Match(filterParserT__0)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserT__0)|(1<<filterParserT__20)|(1<<filterParserFALSE))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(filterParserNULL-33))|(1<<(filterParserTRUE-33))|(1<<(filterParserIntegerLiteral-33))|(1<<(filterParserFloatingPointLiteral-33))|(1<<(filterParserStringLiteral-33))|(1<<(filterParserIdentifier-33)))) != 0) {
		{
			p.SetState(243)
			p.ExpressionList()
		}

	}
	{
		p.SetState(246)
		p.Match(filterParserT__1)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(filterParserIdentifier, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, filterParserRULE_variable)

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
		p.SetState(248)
		p.Match(filterParserIdentifier)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(filterParserIntegerLiteral, 0)
}

func (s *ConstantContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(filterParserFloatingPointLiteral, 0)
}

func (s *ConstantContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(filterParserStringLiteral, 0)
}

func (s *ConstantContext) NULL() antlr.TerminalNode {
	return s.GetToken(filterParserNULL, 0)
}

func (s *ConstantContext) TRUE() antlr.TerminalNode {
	return s.GetToken(filterParserTRUE, 0)
}

func (s *ConstantContext) FALSE() antlr.TerminalNode {
	return s.GetToken(filterParserFALSE, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *filterParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, filterParserRULE_constant)
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
		p.SetState(250)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(filterParserFALSE-28))|(1<<(filterParserNULL-28))|(1<<(filterParserTRUE-28))|(1<<(filterParserIntegerLiteral-28))|(1<<(filterParserFloatingPointLiteral-28))|(1<<(filterParserStringLiteral-28)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
