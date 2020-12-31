grammar oper;

options {
language = Go;
}
//
//tokens {
//LPAREN = '(';
//RPAREN : ')';
//PLUS = '+';
//MINUS = '-';
//ASTERRISK = '*';
//SOLIDUS = '/';
//PERCENT = '%';
//AMPERSAND = '&';
//VERTICAL = '|';
//CIRCUMFLEX = '^';
//}

@header {
import . "filter4go/operand"
}

oper returns[Operand operand]
	: '(' expList = expressionList {$operand, _ = CreateOperandsExpression("", $expList.operands)} ')'
	| exp = expression {$operand = $exp.operand}
	;

expressionList returns[[]Operand operands]
@init {
$operands = []Operand{}
}
    	: exp = expression {$operands = append($operands, $exp.operand)} (',' exp = expression {$operands = append($operands, $exp.operand)})*
    	;

parExpression returns[Operand operand]
    	:  '(' exp = expression ')' {$operand, _ = CreateParenExpression("(", []Operand{$exp.operand})}
    	;

expression returns[Operand operand]
    	: exp = exclusiveOrExpression {$operand = $exp.operand} ( t = '|' exp = exclusiveOrExpression {$operand, _= CreateBitwiseOrExpression($t.GetText(), []Operand{$operand, $exp.operand})})*
    	;

exclusiveOrExpression returns[Operand operand]
    	: exp = andExpression {$operand = $exp.operand} ( t = '^' exp = andExpression {
    	$operand, _ = CreateBitwiseXorExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	})*
    	;

andExpression returns[Operand operand]
    	: exp = shiftExpression {$operand = $exp.operand} ( t = '&'  exp = shiftExpression {
    	$operand, _ = CreateBitwiseAndExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	})*
    	;

shiftExpression returns[Operand operand]
    	: exp = additiveExpression {$operand = $exp.operand} ( t = ('<<'|'>>') exp = additiveExpression {
    	if $t.GetText() == "<<" {
    	    $operand, _ = CreateLeftShiftExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	} else {
            $operand, _ = CreateRightShiftExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	}
    	})*
    	;

additiveExpression returns[Operand operand]
    	: exp = multiplicativeExpression {$operand = $exp.operand} ( t = ('+'|'-')  exp = multiplicativeExpression {
    	if $t.GetText() == "+" {
    	    $operand, _ = CreateAddExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	} else {
    	    $operand, _ = CreateSubtractExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	}
    	})*
    	;

multiplicativeExpression returns[Operand operand]
    	: exp = notExpression {$operand = $exp.operand} ( t = ('*'|'/'|'%' ) exp = notExpression {
    	if $t.GetText() == "*" {
            $operand, _ = CreateMultiplyExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	} else if $t.GetText() == "/" {
            $operand, _ = CreateDivideExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	} else {
            $operand, _ = CreateModularExpression($t.GetText(), []Operand{$operand, $exp.operand})
    	}
    	})*
    	;

notExpression returns[Operand operand]
    	: t = '~'? exp = primary
    	{
    	    if $t != nil {
    	        $operand, _ = CreateBitwiseNotExpression($t.GetText(), []Operand{$exp.operand})
    	    } else {
    	        $operand = $exp.operand
    	    }
    	}
    	;

primary returns[Operand operand]
    	: exp = parExpression {$operand = $exp.operand}
    	| fexp = function   {$operand = $fexp.operand}
    	| mexp = member   {$operand = $mexp.operand}
    	| cexp = constant {$operand = $cexp.operand}

    	;

member returns[Operand operand]
	: exp = variable {$operand = $exp.operand}
	('[' index = expression? {$operand, _ = CreateArrayExpression("[", []Operand{$operand, $index.operand})}']')*
	('.' (vexp = variable {$operand, _ = CreateMemberVariableExpression("", []Operand{$operand, $vexp.operand})}
	| fexp = function {$operand, _ = CreateMemberMethodExpression("", []Operand{$operand, $fexp.operand})})
	('[' index = expression? {$operand, _ = CreateArrayExpression("[", []Operand{$operand, $index.operand})}']')*)*
	;

function returns[Operand operand]
@init {
    var err error
    var parameters []Operand
}
	: t = Identifier '(' (expList = expressionList {parameters = $expList.operands})? ')' {
	$operand, err = FuncFactory.CreateFunction($t.GetText(), parameters)
	if err != nil {
	    panic(err)
	}
	}
	;

variable returns[Operand operand]
	: t = Identifier {$operand = CreateVariable($t.GetText());}
	;

constant returns[Operand operand]
@init {}
	: t = IntegerLiteral {$operand = CreateLong($t.GetText())}
	| t = FloatingPointLiteral {$operand = CreateDouble($t.GetText())}
	| t = StringLiteral {$operand = CreateString($t.GetText())}
	| t = NULL {$operand = CreateNull($t.GetText())}
	| t = TRUE {$operand = CreateBoolean($t.GetText())}
	| t = FALSE {$operand = CreateBoolean($t.GetText());}
	;

//Lexer

NULL		:	N_ U_ L_ L_;
TRUE		:	T_ R_ U_ E_;
FALSE		:	F_ A_ L_ S_ E_;

IntegerLiteral
	: HexLiteral
	| DecimalLiteral
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

fragment
IntegerTypeSuffix : ('l'|'L') ;

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
