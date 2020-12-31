grammar filter;

options {
language = Go;
}

@header {
import . "filter4go/metadata"
}

//tokens
//{
//}

searchCondition returns[OperatableMetadata operationMetadata]
	: lCtx = booleanTerm {$operationMetadata = $lCtx.operationMetadata} (opCtx = OR rCtx = searchCondition {
	   $operationMetadata = OperationMetadata{BINARY, $opCtx.GetText(), $operationMetadata, $rCtx.operationMetadata}
	})?
	;

booleanTerm returns[OperatableMetadata operationMetadata]
@after {
if $operationMetadata == nil {
	$operationMetadata = $lCtx.operationMetadata
}
}
	: lCtx = booleanFactor {$operationMetadata = $lCtx.operationMetadata} (opCtx = AND rCtx = booleanTerm {
	    $operationMetadata = OperationMetadata{BINARY, $opCtx.GetText(), $operationMetadata, $rCtx.operationMetadata}
	})?
	;

booleanFactor returns[OperatableMetadata operationMetadata]
@after {
if $notCtx != nil {
	$operationMetadata = OperationMetadata{UNARY, $notCtx.GetText(), nil, $rCtx.operationMetadata}
} else {
	$operationMetadata = $rCtx.operationMetadata
}
}
	: notCtx = NOT? rCtx = booleanPrimary
	;

booleanPrimary returns[OperatableMetadata operationMetadata]
	: predicateCtx = predicate {$operationMetadata = $predicateCtx.operationMetadata;}
	| '(' conditionCtx = searchCondition ')' {$operationMetadata = OperationMetadata{UNARY, PAREN, nil, $conditionCtx.operationMetadata}}
	| exprCtx = expression {$operationMetadata = RelationOperationMetadata{UNARY, EXPR, "", $exprCtx.expressionText}}
	;

predicate returns[OperatableMetadata operationMetadata]
	: comparisonMetadata = comparisonPredicate {$operationMetadata = $comparisonMetadata.operationMetadata}
	| betweenMetadata = betweenPredicate {$operationMetadata = $betweenMetadata.operationMetadata}
	| inMetadata = inPredicate {$operationMetadata = $inMetadata.operationMetadata}
	| likeMetadata = likePredicate {$operationMetadata = $likeMetadata.operationMetadata}
	| nullMetadata = nullPredicate {$operationMetadata = $nullMetadata.operationMetadata}
	;

comparisonPredicate returns[OperatableMetadata operationMetadata]
@after {
$operationMetadata = RelationOperationMetadata{BINARY, $opCtx.GetText(), $lExprCtx.expressionText, $rExprCtx.expressionText};
}
	: lExprCtx = expression opCtx = ('=' | '<' | '<=' | '>' | '>=' | '<>' | '!=') rExprCtx = expression
	;

betweenPredicate returns[OperatableMetadata operationMetadata]
@after {
rExpr := fmt.Sprintf("(%s,%s)", $rExpr1Ctx.expressionText, $rExpr2Ctx.expressionText);
$operationMetadata = RelationOperationMetadata{BINARY, BETWEEN, $lExprCtx.expressionText, rExpr};
if $notCtx != nil {
	$operationMetadata = OperationMetadata{UNARY, $notCtx.GetText(), nil, $operationMetadata};
}
}
	: lExprCtx = expression notCtx = NOT? BETWEEN rExpr1Ctx = expression AND rExpr2Ctx = expression
	;

inPredicate returns[OperatableMetadata operationMetadata]
@after {
inExpressionText := fmt.Sprintf("%s%s%s", $lrCtx.GetText(), $rExprCtx.expressionText, $rrCtx.GetText())
$operationMetadata = RelationOperationMetadata{BINARY, IN, $lExprCtx.expressionText, inExpressionText}
if $notCtx != nil {
	$operationMetadata = OperationMetadata{UNARY, $notCtx.GetText(), nil, $operationMetadata}
}
}
	: lExprCtx = expression notCtx = NOT? IN lrCtx = '(' rExprCtx = expressionList rrCtx = ')'
	;

likePredicate returns[OperatableMetadata operationMetadata]
@after {
$operationMetadata = RelationOperationMetadata{BINARY, LIKE, $lExprCtx.expressionText, $rExprCtx.expressionText};
if $notCtx != nil {
	$operationMetadata = OperationMetadata{UNARY, $notCtx.GetText(), nil, $operationMetadata}
}
}
	: lExprCtx = expression notCtx=NOT? LIKE rExprCtx = expression
	;

nullPredicate returns[OperatableMetadata operationMetadata]
@after {
$operationMetadata = RelationOperationMetadata{BINARY, IS, $lExprCtx.expressionText, $rExprCtx.GetText()}
if $notCtx != nil {
	$operationMetadata = OperationMetadata{UNARY, $notCtx.GetText(), nil, $operationMetadata}
}
}
	: lExprCtx = expression IS notCtx = NOT? rExprCtx = NULL
	;

expressionList returns[string expressionText]
    	: expression (',' expression)* {$expressionText = $text}
    	;

parExpression
    	:  '('expression')'
    	;

expression returns[string expressionText]
    	: exclusiveOrExpression ( '|' exclusiveOrExpression)* {$expressionText = $text}
    	;

exclusiveOrExpression
    	: andExpression ('^' andExpression)*
    	;

andExpression
    	: shiftExpression ('&' shiftExpression)*
    	;
shiftExpression
    	: additiveExpression (('<<'|'>>') additiveExpression)*
    	;
additiveExpression
    	: multiplicativeExpression (('+' | '-')  multiplicativeExpression)*
    	;

multiplicativeExpression
    	: notExpression (( '*' | '/' | '%' ) notExpression)*
    	;

notExpression
    	: '~'? exp = primary;

primary
    	: parExpression
    	| function
    	| member
    	| constant
    	;

member
	: variable('['expression?']')* (('.' (variable| function) ('['expression?']')*)*)
	;

function
	: Identifier '('expressionList?')'
	;

variable
	: Identifier
	;

constant
	: IntegerLiteral
	| FloatingPointLiteral
	| StringLiteral
	| NULL
	| TRUE
	| FALSE
	;

//Lexer

AND 		:	A_ N_ D_;
BETWEEN		:	B_ E_ T_ W_ E_ E_ N_;
EXISTS		:	E_ X_ I_ S_ T_ S_;
FALSE		:	F_ A_ L_ S_ E_;
IN			:	I_ N_;
IS			:	I_ S_;
LIKE		:	L_ I_ K_ E_;
NOT			:	N_ O_ T_;
NULL		:	N_ U_ L_ L_;
OR			:	O_ R_;
TRUE		:	T_ R_ U_ E_;

Percent
	: ('100' | '1'..'9' '0'..'9'?) '%'
	;

IntegerLiteral
	: DecimalLiteral
	| HexLiteral
	| OctalLiteral;

fragment
HexLiteral : '0' ('x'|'X') HexDigit+ ;

fragment
DecimalLiteral : ('+'|'-')? ('0' | '1'..'9' '0'..'9'*) ;
//DecimalLiteral : ('0' | '1'..'9' '0'..'9'*) ;

fragment
OctalLiteral : '0' ('0'..'7')+ ;

fragment
HexDigit : ('0'..'9'|'a'..'f'|'A'..'F') ;


FloatingPointLiteral
    	:   ('+'|'-')? ('0'..'9')+ '.' ('0'..'9')* Exponent?
    	|   '.' ('0'..'9')+ Exponent?
    	|   ('+'|'-')? ('0'..'9')+ Exponent
    	;

fragment
Exponent : ('e'|'E') ('+'|'-')? ('0'..'9')+ ;

StringLiteral
    :  '\'' ( ~'\'' | Escape)*  '\''
    |  '"' (~'"' | '\\"')* '"'
    ;
fragment
Escape	: '\'' '\'';

Identifier
	: Letter (Letter | Digital)*;

fragment
Letter	:
	'\u0024' |
       	'\u0041'..'\u005a' |
       	'\u005f' |
       	'\u0061'..'\u007a' |
       	'\u00c0'..'\u00d6' |
       	'\u00d8'..'\u00f6' |
       	'\u00f8'..'\u00ff' |
       	'\u0100'..'\u1fff' |
       	'\u3040'..'\u318f' |
       	'\u3300'..'\u337f' |
       	'\u3400'..'\u3d2d' |
       	'\u4e00'..'\u9fff' |
       	'\uf900'..'\ufaff';
fragment
Digital	:
	'\u0030'..'\u0039' |
       	'\u0660'..'\u0669' |
       	'\u06f0'..'\u06f9' |
       	'\u0966'..'\u096f' |
       	'\u09e6'..'\u09ef' |
       	'\u0a66'..'\u0a6f' |
       	'\u0ae6'..'\u0aef' |
       	'\u0b66'..'\u0b6f' |
       	'\u0be7'..'\u0bef' |
       	'\u0c66'..'\u0c6f' |
       	'\u0ce6'..'\u0cef' |
       	'\u0d66'..'\u0d6f' |
       	'\u0e50'..'\u0e59' |
       	'\u0ed0'..'\u0ed9' |
       	'\u1040'..'\u1049';

fragment A_ : 'a' | 'A';
fragment B_ : 'b' | 'B';
fragment C_ : 'c' | 'C';
fragment D_ : 'd' | 'D';
fragment E_ : 'e' | 'E';
fragment F_ : 'f' | 'F';
fragment G_ : 'g' | 'G';
fragment H_ : 'h' | 'H';
fragment I_ : 'i' | 'I';
fragment J_ : 'j' | 'J';
fragment K_ : 'k' | 'K';
fragment L_ : 'l' | 'L';
fragment M_ : 'm' | 'M';
fragment N_ : 'n' | 'N';
fragment O_ : 'o' | 'O';
fragment P_ : 'p' | 'P';
fragment Q_ : 'q' | 'Q';
fragment R_ : 'r' | 'R';
fragment S_ : 's' | 'S';
fragment T_ : 't' | 'T';
fragment U_ : 'u' | 'U';
fragment V_ : 'v' | 'V';
fragment W_ : 'w' | 'W';
fragment X_ : 'x' | 'X';
fragment Y_ : 'y' | 'Y';
fragment Z_ : 'z' | 'Z';

WS  :  [ \r\t\u000C\n]+ -> channel(HIDDEN);

