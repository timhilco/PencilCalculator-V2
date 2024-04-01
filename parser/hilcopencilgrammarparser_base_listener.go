// Code generated from HilcoPencilGrammarParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // HilcoPencilGrammarParser

import "github.com/antlr4-go/antlr/v4"

// BaseHilcoPencilGrammarParserListener is a complete listener for a parse tree produced by HilcoPencilGrammarParser.
type BaseHilcoPencilGrammarParserListener struct{}

var _ HilcoPencilGrammarParserListener = &BaseHilcoPencilGrammarParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHilcoPencilGrammarParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHilcoPencilGrammarParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitStatement(ctx *StatementContext) {}

// EnterNameCalculator is called when production NameCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterNameCalculator(ctx *NameCalculatorContext) {}

// ExitNameCalculator is called when production NameCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitNameCalculator(ctx *NameCalculatorContext) {}

// EnterUnaryNotCalculator is called when production UnaryNotCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterUnaryNotCalculator(ctx *UnaryNotCalculatorContext) {
}

// ExitUnaryNotCalculator is called when production UnaryNotCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitUnaryNotCalculator(ctx *UnaryNotCalculatorContext) {
}

// EnterPercent is called when production Percent is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterPercent(ctx *PercentContext) {}

// ExitPercent is called when production Percent is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitPercent(ctx *PercentContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitParens(ctx *ParensContext) {}

// EnterBinaryRelationalCalculator is called when production BinaryRelationalCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterBinaryRelationalCalculator(ctx *BinaryRelationalCalculatorContext) {
}

// ExitBinaryRelationalCalculator is called when production BinaryRelationalCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitBinaryRelationalCalculator(ctx *BinaryRelationalCalculatorContext) {
}

// EnterWorkSheetVariableCalculator is called when production WorkSheetVariableCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterWorkSheetVariableCalculator(ctx *WorkSheetVariableCalculatorContext) {
}

// ExitWorkSheetVariableCalculator is called when production WorkSheetVariableCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitWorkSheetVariableCalculator(ctx *WorkSheetVariableCalculatorContext) {
}

// EnterString is called when production String is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitString(ctx *StringContext) {}

// EnterExpressionAtFunction is called when production ExpressionAtFunction is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterExpressionAtFunction(ctx *ExpressionAtFunctionContext) {
}

// ExitExpressionAtFunction is called when production ExpressionAtFunction is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitExpressionAtFunction(ctx *ExpressionAtFunctionContext) {
}

// EnterDate is called when production Date is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production Date is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitDate(ctx *DateContext) {}

// EnterCase is called when production Case is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production Case is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitCase(ctx *CaseContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterBinaryArithmeticCalculator is called when production BinaryArithmeticCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterBinaryArithmeticCalculator(ctx *BinaryArithmeticCalculatorContext) {
}

// ExitBinaryArithmeticCalculator is called when production BinaryArithmeticCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitBinaryArithmeticCalculator(ctx *BinaryArithmeticCalculatorContext) {
}

// EnterBinaryLogicalCalculator is called when production BinaryLogicalCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterBinaryLogicalCalculator(ctx *BinaryLogicalCalculatorContext) {
}

// ExitBinaryLogicalCalculator is called when production BinaryLogicalCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitBinaryLogicalCalculator(ctx *BinaryLogicalCalculatorContext) {
}

// EnterExpressionKeyword is called when production ExpressionKeyword is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterExpressionKeyword(ctx *ExpressionKeywordContext) {
}

// ExitExpressionKeyword is called when production ExpressionKeyword is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitExpressionKeyword(ctx *ExpressionKeywordContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitFloat(ctx *FloatContext) {}

// EnterUnaryMinusCalculator is called when production UnaryMinusCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterUnaryMinusCalculator(ctx *UnaryMinusCalculatorContext) {
}

// ExitUnaryMinusCalculator is called when production UnaryMinusCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitUnaryMinusCalculator(ctx *UnaryMinusCalculatorContext) {
}

// EnterBinaryExponentialCalculator is called when production BinaryExponentialCalculator is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterBinaryExponentialCalculator(ctx *BinaryExponentialCalculatorContext) {
}

// ExitBinaryExponentialCalculator is called when production BinaryExponentialCalculator is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitBinaryExponentialCalculator(ctx *BinaryExponentialCalculatorContext) {
}

// EnterExpressionDataAccess is called when production ExpressionDataAccess is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterExpressionDataAccess(ctx *ExpressionDataAccessContext) {
}

// ExitExpressionDataAccess is called when production ExpressionDataAccess is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitExpressionDataAccess(ctx *ExpressionDataAccessContext) {
}

// EnterIf is called when production If is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitIf(ctx *IfContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterCaseList is called when production caseList is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterCaseList(ctx *CaseListContext) {}

// ExitCaseList is called when production caseList is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitCaseList(ctx *CaseListContext) {}

// EnterCaseItem is called when production caseItem is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterCaseItem(ctx *CaseItemContext) {}

// ExitCaseItem is called when production caseItem is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitCaseItem(ctx *CaseItemContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterName is called when production name is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitName(ctx *NameContext) {}

// EnterWorksheetVariable is called when production worksheetVariable is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterWorksheetVariable(ctx *WorksheetVariableContext) {
}

// ExitWorksheetVariable is called when production worksheetVariable is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitWorksheetVariable(ctx *WorksheetVariableContext) {}

// EnterAtFunction is called when production atFunction is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterAtFunction(ctx *AtFunctionContext) {}

// ExitAtFunction is called when production atFunction is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitAtFunction(ctx *AtFunctionContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitArgList(ctx *ArgListContext) {}

// EnterDataAccessor is called when production dataAccessor is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterDataAccessor(ctx *DataAccessorContext) {}

// ExitDataAccessor is called when production dataAccessor is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitDataAccessor(ctx *DataAccessorContext) {}

// EnterAccessorList is called when production accessorList is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterAccessorList(ctx *AccessorListContext) {}

// ExitAccessorList is called when production accessorList is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitAccessorList(ctx *AccessorListContext) {}

// EnterAccessorObjectOrArray is called when production accessorObjectOrArray is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterAccessorObjectOrArray(ctx *AccessorObjectOrArrayContext) {
}

// ExitAccessorObjectOrArray is called when production accessorObjectOrArray is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitAccessorObjectOrArray(ctx *AccessorObjectOrArrayContext) {
}

// EnterAccessorObject is called when production accessorObject is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterAccessorObject(ctx *AccessorObjectContext) {}

// ExitAccessorObject is called when production accessorObject is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitAccessorObject(ctx *AccessorObjectContext) {}

// EnterAccessorArray is called when production accessorArray is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterAccessorArray(ctx *AccessorArrayContext) {}

// ExitAccessorArray is called when production accessorArray is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitAccessorArray(ctx *AccessorArrayContext) {}

// EnterSpecialKeyword is called when production specialKeyword is entered.
func (s *BaseHilcoPencilGrammarParserListener) EnterSpecialKeyword(ctx *SpecialKeywordContext) {}

// ExitSpecialKeyword is called when production specialKeyword is exited.
func (s *BaseHilcoPencilGrammarParserListener) ExitSpecialKeyword(ctx *SpecialKeywordContext) {}
