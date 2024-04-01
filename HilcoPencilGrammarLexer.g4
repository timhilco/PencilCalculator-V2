lexer grammar HilcoPencilGrammarLexer;

 
/* -----------------------------------fxparse.tok ----------------------------------------------------------------------------------

<lowercasename> : [a-z] [a-zA-Z0-9\:\\\_]* ;
<number> : [0-9]+; 
<string> : '((~['] | '' )*)' {compactDoubleApostrophes}; 
<float> : [0-9]+\.[0-9]*;
<date> : \{[0-9][0-9] (\-|\/) ([0-9][0-9] | [a-zA-Z][a-zA-Z][a-zA-Z]) (\-|\/) [0-9][0-9] ([0-9][0-9])?\};
<atfunction> : \@ [a-zA-Z] [a-zA-Z0-9]*;
<uppercasename> : [A-Z] [a-zA-Z0-9\_]*; 
<percent> : ([0-9]+|[0-9]+\.[0-9]* ) \%; 
<comment> : \" ~[\"]* \" {ignoreComment};
<space> : [\s\t\r\n]+ {ignoreDelimiter};
---------------------------------------------------------------------------------------------------------------------------------------
*/
 ASSIGNMENT : ':=';
 CASE : 'case';
 END_CASE: 'endcase'; 
 IS: 'is';
 SWITCH : 'switch';
 END_SWITCH: 'endswitch'; 
 IF: 'if';
 THEN: 'then';
 ELSE: 'else';
 DOUBLE_COLON: '::';
 SEMI_COLON: ';';
 NOT : 'NOT'; 
 EXPONENTIAL: '^';
 MULTIPLY :'*' ;
 DIVIDE:  '/' ;
 ADD:  '+' ;
 MINUS :  '-';
 GREATER_THAN:  '>' ;
 GREATER_THAN_EQUAL: '>=';
 LESS_THAN :  '<';
 LESS_THAN_EQUAL: '<=';
 EQUAL :  '=';
 NOT_EQUAL:  '~=';
 AND : 'AND' ;
 OR:  'OR';
 COLON      : ':'   ;
 LBRACKET   : '['   ;
 RBRACKET   : ']'   ;
 CURLYLBRACKET   : '{'   ;
 CURLYRBRACKET   : '}'   ;
 LPAREN     : '('   ;
 RPAREN     : ')'   ;
 UNDERSCORE  : '_';
 PERCENT_SIGN : '%';
 AT_SIGN : '@';
 COMMA: ',';
 DOT: '.';
 KEYWORD_TRUE : 'true'; 
 KEYWORD_FALSE: 'false'; 
 KEYWORD_NIL: 'nil';  
 fragment DIGIT: '0'..'9';
 fragment LOWERCASELETTERS
 	:	 'a'..'z' ;
 fragment UPPERCASELETTERS 
 	:	'A'..'Z' ;
CLASSNAME : UPPERCASELETTERS (LOWERCASELETTERS | UPPERCASELETTERS | DIGIT | UNDERSCORE)*; 
ID : LOWERCASELETTERS (	  LOWERCASELETTERS 
			| UPPERCASELETTERS 
			|  DIGIT
			| UNDERSCORE
			| COLON
			| '\\'
			)* ; 
			 
ATFUNCTION : AT_SIGN (LOWERCASELETTERS  
		 | UPPERCASELETTERS
		 ) 
		 (LOWERCASELETTERS 
		 | UPPERCASELETTERS 
		 | DIGIT
		 )*;

INT: (DIGIT)+;			
STRING
    :  '\'' ( ESC_SEQ | ~('\\'|'\'') )* '\''
    ; 
FLOAT : (DIGIT)+ '.' (DIGIT)*;
PERCENT : (INT|FLOAT) PERCENT_SIGN ; 
DATE :  CURLYLBRACKET
         DIGIT DIGIT   
	('-' |'/') 
	(
		(DIGIT DIGIT)  
	| 
	   ( (LOWERCASELETTERS   | UPPERCASELETTERS  ) 
	     (LOWERCASELETTERS   | UPPERCASELETTERS  )
	     (LOWERCASELETTERS   | UPPERCASELETTERS )
	   )
	 )    
	 ('-' | '/' ) 
	 DIGIT DIGIT (DIGIT DIGIT)?
	 CURLYRBRACKET
	 ;
fragment
EXPONENT : ('e'|'E') ('+'|'-')? ('0'..'9')+ ;

fragment
HEX_DIGIT : ('0'..'9'|'a'..'f'|'A'..'F') ;

fragment
ESC_SEQ
    :   '\\' ('b'|'t'|'n'|'f'|'r'|'"'|'\''|'\\')
    |   UNICODE_ESC
    |   OCTAL_ESC
    ;

fragment
OCTAL_ESC
    :   '\\' ('0'..'3') ('0'..'7') ('0'..'7')
    |   '\\' ('0'..'7') ('0'..'7')
    |   '\\' ('0'..'7')
    ;

fragment
UNICODE_ESC
    :   '\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
    ;

NEWLINE :'\r'? '\n' -> channel(HIDDEN);
 
WS
  : ( ' '
    | '\t'
    | '\f'
    // handle newlines
    | ( '\r\n'  // DOS/Windows
      )
    | '\u0003'
    )
    -> channel(HIDDEN)
  ;
