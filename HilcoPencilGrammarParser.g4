parser grammar HilcoPencilGrammarParser;
options {tokenVocab= HilcoPencilGrammarLexer;} 

/*    ----------------------------------------------------------------fxParse.grm  ---------------------------------------------
Expr   : CaseStatement
            | MatchedExpr
            | UnmatchedExpr;
CaseStatement : 'case' Expr 'is' CaseList 'endcase' {caseExpr:list:};
CaseList : CaseItem CaseList {liftRightChild}
            | "epsilon" {OrderedChildren};
CaseItem : Base '::' Expr ';' {caseItem:expr:};
MatchedExpr : 'if'  Logical  'then'  MatchedExpr  'else'  MatchedExpr {if:then:else:}
            |  Logical  ;
UnmatchedExpr : 'if' Logical 'then'  Expr {if:then:}
            | 'if' Logical 'then'  MatchedExpr 'else' UnmatchedExpr {if:then:else:};
Logical :     Logical 'AND' Relative {logicalAND:and:}
            | Logical 'OR' Relative {logicalOR:and:}
            | 'NOT' Relative {logicalNOT:}
            | Relative;
Relative : Relative '>' Arith {greaterThan:and:}
            | Relative '>=' Arith {greaterThanEqual:and:}
            | Relative '<' Arith {lessThan:and:}
            | Relative '<=' Arith {lessThanEqual:and:}
            | Relative '=' Arith {equalTo:and:}
            | Relative '~=' Arith {notEqualTo:and:}
            | Arith;
Arith   :     Arith '+' Term {add:and:}
            |  Arith '-' Term {subtract:and:}
            |  Term;
Term  :    Term '*' Factor {multiply:and:}
            | Term '/' Factor {divide:and:}
            | Factor;
Factor :  '-' Factor {negate:}
           | Base;
Base : '(' Expr ')'
            | Name
            | <percent> {percent:}
            | <number> {number:}
            | <float> {number:}
            | <string>   {string:}
            | <date> {date:}
            | AtFunction
            | ContextAccessor
            | SpecialKeyword;
Name  : Calc
            | Sheet ':' Calc {sheet:calc:};
Sheet : <uppercasename> {sheetNamed:};
Calc : <lowercasename> {named:};
AtFunction : AtFunctionName  '(' ArgList ')'{createatFunctionCalculator:withParms:};
ArgList : Logical Args  {liftRightChild}
            |   NullArg Args {liftRightChild};
Args    : ',' ArgList
            | "epsilon" {OrderedChildren};
NullArg : "epsilon" {nil};
AtFunctionName : <atfunction> {atFunctionNamed:};
ContextAccessor : ObjectType AccessList {createContextAccessorCalculator:withFields:};
AccessList : '.' AccessorMessageOrObjectType AccessList {liftRightChild}
            | "epsilon"	 {OrderedChildren};
AccessorMessageOrObjectType : AccessorMessage | ObjectType {createAccessorMessageOrObjectType:} ;
AccessorMessage : AccessName {createAccessorMessage:}
            | AccessName '(' ArgList ')' {createAccessorMessage:withArguments:};
ObjectType : <uppercasename> {objectNamed:};
AccessName : <lowercasename> {accessNamed:};
SpecialKeyword : 'true' {booleanTrueConstant}
            | 'false' {booleanFalseConstant}
            | 'nil' {nilConstant}  ; 
-----------------------------------------------------------------------------------------------------------------------------------
*/ 

program : expression 
        | statement* |;

statement: expression ASSIGNMENT   expression SEMI_COLON;      

expression   :  caseStatement       	     			#Case
         //   | switchStatement      	 					#ExpressionSwitch
            | ifStatement           					#If
            | LPAREN expression RPAREN          		#Parens
            | MINUS expression              			#UnaryMinusCalculator
            | NOT expression             				#UnaryNotCalculator
            | <assoc=right> expression EXPONENTIAL  expression #BinaryExponentialCalculator
            | expression MULTIPLY expression         	#BinaryArithmeticCalculator
            | expression DIVIDE expression         		#BinaryArithmeticCalculator
            | expression ADD expression         	  	#BinaryArithmeticCalculator
            | expression MINUS expression         		#BinaryArithmeticCalculator
            | expression GREATER_THAN expression        #BinaryRelationalCalculator
            | expression GREATER_THAN_EQUAL expression 	#BinaryRelationalCalculator
            | expression LESS_THAN expression         	#BinaryRelationalCalculator
            | expression LESS_THAN_EQUAL expression     #BinaryRelationalCalculator
            | expression EQUAL expression         		#BinaryRelationalCalculator
            | expression NOT_EQUAL expression        	#BinaryRelationalCalculator
            | expression AND expression       			#BinaryLogicalCalculator
            | expression OR expression        			#BinaryLogicalCalculator
            | name                  					#NameCalculator
            | worksheetVariable               #WorkSheetVariableCalculator
            | atFunction           				 		#ExpressionAtFunction
            | dataAccessor         				 		#ExpressionDataAccess
            | specialKeyword   				     		#ExpressionKeyword
            | PERCENT             				  		#Percent
            | INT                   					#Integer
            | FLOAT                 					#Float
            | STRING                					#String
            | DATE                				  		#Date   
            ;

caseStatement : CASE expression IS caseList END_CASE ;
//switchStatement : SWITCH  caseList END_SWITCH ;  
caseList : caseItem caseList 
           | ;
caseItem : expression DOUBLE_COLON expression SEMI_COLON ;
ifStatement : IF expression THEN expression (ELSE expression)? 
 ;

name  :    ID;
worksheetVariable : CLASSNAME COLON ID ;

atFunction : ATFUNCTION  LPAREN argList? RPAREN;
argList : expression (COMMA expression? )*;

dataAccessor : CLASSNAME accessorList+ ;
  accessorList : DOT (accessorObjectOrArray) ; 
            
//accessList : DOT accessorMessageOrObjectType accessList 
// ;
//accessorMessageOrObjectType : accessorMessage | CLASSNAME ;
accessorObjectOrArray : accessorObject
           | accessorArray;

accessorObject: ID;
accessorArray: ID LPAREN argList RPAREN;
specialKeyword : KEYWORD_TRUE
              | KEYWORD_FALSE 
              | KEYWORD_NIL 
               ; 

