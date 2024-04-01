// Code generated from HilcoPencilGrammarParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // HilcoPencilGrammarParser

import "github.com/antlr4-go/antlr/v4"

// HilcoPencilGrammarParserListener is a complete listener for a parse tree produced by HilcoPencilGrammarParser.
type HilcoPencilGrammarParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterNameCalculator is called when entering the NameCalculator production.
	EnterNameCalculator(c *NameCalculatorContext)

	// EnterUnaryNotCalculator is called when entering the UnaryNotCalculator production.
	EnterUnaryNotCalculator(c *UnaryNotCalculatorContext)

	// EnterPercent is called when entering the Percent production.
	EnterPercent(c *PercentContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterBinaryRelationalCalculator is called when entering the BinaryRelationalCalculator production.
	EnterBinaryRelationalCalculator(c *BinaryRelationalCalculatorContext)

	// EnterWorkSheetVariableCalculator is called when entering the WorkSheetVariableCalculator production.
	EnterWorkSheetVariableCalculator(c *WorkSheetVariableCalculatorContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterExpressionAtFunction is called when entering the ExpressionAtFunction production.
	EnterExpressionAtFunction(c *ExpressionAtFunctionContext)

	// EnterDate is called when entering the Date production.
	EnterDate(c *DateContext)

	// EnterCase is called when entering the Case production.
	EnterCase(c *CaseContext)

	// EnterInteger is called when entering the Integer production.
	EnterInteger(c *IntegerContext)

	// EnterBinaryArithmeticCalculator is called when entering the BinaryArithmeticCalculator production.
	EnterBinaryArithmeticCalculator(c *BinaryArithmeticCalculatorContext)

	// EnterBinaryLogicalCalculator is called when entering the BinaryLogicalCalculator production.
	EnterBinaryLogicalCalculator(c *BinaryLogicalCalculatorContext)

	// EnterExpressionKeyword is called when entering the ExpressionKeyword production.
	EnterExpressionKeyword(c *ExpressionKeywordContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterUnaryMinusCalculator is called when entering the UnaryMinusCalculator production.
	EnterUnaryMinusCalculator(c *UnaryMinusCalculatorContext)

	// EnterBinaryExponentialCalculator is called when entering the BinaryExponentialCalculator production.
	EnterBinaryExponentialCalculator(c *BinaryExponentialCalculatorContext)

	// EnterExpressionDataAccess is called when entering the ExpressionDataAccess production.
	EnterExpressionDataAccess(c *ExpressionDataAccessContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterCaseList is called when entering the caseList production.
	EnterCaseList(c *CaseListContext)

	// EnterCaseItem is called when entering the caseItem production.
	EnterCaseItem(c *CaseItemContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterWorksheetVariable is called when entering the worksheetVariable production.
	EnterWorksheetVariable(c *WorksheetVariableContext)

	// EnterAtFunction is called when entering the atFunction production.
	EnterAtFunction(c *AtFunctionContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterDataAccessor is called when entering the dataAccessor production.
	EnterDataAccessor(c *DataAccessorContext)

	// EnterAccessorList is called when entering the accessorList production.
	EnterAccessorList(c *AccessorListContext)

	// EnterAccessorObjectOrArray is called when entering the accessorObjectOrArray production.
	EnterAccessorObjectOrArray(c *AccessorObjectOrArrayContext)

	// EnterAccessorObject is called when entering the accessorObject production.
	EnterAccessorObject(c *AccessorObjectContext)

	// EnterAccessorArray is called when entering the accessorArray production.
	EnterAccessorArray(c *AccessorArrayContext)

	// EnterSpecialKeyword is called when entering the specialKeyword production.
	EnterSpecialKeyword(c *SpecialKeywordContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitNameCalculator is called when exiting the NameCalculator production.
	ExitNameCalculator(c *NameCalculatorContext)

	// ExitUnaryNotCalculator is called when exiting the UnaryNotCalculator production.
	ExitUnaryNotCalculator(c *UnaryNotCalculatorContext)

	// ExitPercent is called when exiting the Percent production.
	ExitPercent(c *PercentContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitBinaryRelationalCalculator is called when exiting the BinaryRelationalCalculator production.
	ExitBinaryRelationalCalculator(c *BinaryRelationalCalculatorContext)

	// ExitWorkSheetVariableCalculator is called when exiting the WorkSheetVariableCalculator production.
	ExitWorkSheetVariableCalculator(c *WorkSheetVariableCalculatorContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitExpressionAtFunction is called when exiting the ExpressionAtFunction production.
	ExitExpressionAtFunction(c *ExpressionAtFunctionContext)

	// ExitDate is called when exiting the Date production.
	ExitDate(c *DateContext)

	// ExitCase is called when exiting the Case production.
	ExitCase(c *CaseContext)

	// ExitInteger is called when exiting the Integer production.
	ExitInteger(c *IntegerContext)

	// ExitBinaryArithmeticCalculator is called when exiting the BinaryArithmeticCalculator production.
	ExitBinaryArithmeticCalculator(c *BinaryArithmeticCalculatorContext)

	// ExitBinaryLogicalCalculator is called when exiting the BinaryLogicalCalculator production.
	ExitBinaryLogicalCalculator(c *BinaryLogicalCalculatorContext)

	// ExitExpressionKeyword is called when exiting the ExpressionKeyword production.
	ExitExpressionKeyword(c *ExpressionKeywordContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitUnaryMinusCalculator is called when exiting the UnaryMinusCalculator production.
	ExitUnaryMinusCalculator(c *UnaryMinusCalculatorContext)

	// ExitBinaryExponentialCalculator is called when exiting the BinaryExponentialCalculator production.
	ExitBinaryExponentialCalculator(c *BinaryExponentialCalculatorContext)

	// ExitExpressionDataAccess is called when exiting the ExpressionDataAccess production.
	ExitExpressionDataAccess(c *ExpressionDataAccessContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitCaseList is called when exiting the caseList production.
	ExitCaseList(c *CaseListContext)

	// ExitCaseItem is called when exiting the caseItem production.
	ExitCaseItem(c *CaseItemContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitWorksheetVariable is called when exiting the worksheetVariable production.
	ExitWorksheetVariable(c *WorksheetVariableContext)

	// ExitAtFunction is called when exiting the atFunction production.
	ExitAtFunction(c *AtFunctionContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitDataAccessor is called when exiting the dataAccessor production.
	ExitDataAccessor(c *DataAccessorContext)

	// ExitAccessorList is called when exiting the accessorList production.
	ExitAccessorList(c *AccessorListContext)

	// ExitAccessorObjectOrArray is called when exiting the accessorObjectOrArray production.
	ExitAccessorObjectOrArray(c *AccessorObjectOrArrayContext)

	// ExitAccessorObject is called when exiting the accessorObject production.
	ExitAccessorObject(c *AccessorObjectContext)

	// ExitAccessorArray is called when exiting the accessorArray production.
	ExitAccessorArray(c *AccessorArrayContext)

	// ExitSpecialKeyword is called when exiting the specialKeyword production.
	ExitSpecialKeyword(c *SpecialKeywordContext)
}
