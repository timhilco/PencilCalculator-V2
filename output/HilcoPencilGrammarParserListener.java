// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilGrammarParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link HilcoPencilGrammarParser}.
 */
public interface HilcoPencilGrammarParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(HilcoPencilGrammarParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(HilcoPencilGrammarParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(HilcoPencilGrammarParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(HilcoPencilGrammarParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NameCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterNameCalculator(HilcoPencilGrammarParser.NameCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NameCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitNameCalculator(HilcoPencilGrammarParser.NameCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code UnaryNotCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterUnaryNotCalculator(HilcoPencilGrammarParser.UnaryNotCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnaryNotCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitUnaryNotCalculator(HilcoPencilGrammarParser.UnaryNotCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Percent}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterPercent(HilcoPencilGrammarParser.PercentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Percent}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitPercent(HilcoPencilGrammarParser.PercentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterParens(HilcoPencilGrammarParser.ParensContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitParens(HilcoPencilGrammarParser.ParensContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinaryRelationalCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterBinaryRelationalCalculator(HilcoPencilGrammarParser.BinaryRelationalCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinaryRelationalCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitBinaryRelationalCalculator(HilcoPencilGrammarParser.BinaryRelationalCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code WorkSheetVariableCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterWorkSheetVariableCalculator(HilcoPencilGrammarParser.WorkSheetVariableCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code WorkSheetVariableCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitWorkSheetVariableCalculator(HilcoPencilGrammarParser.WorkSheetVariableCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code String}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterString(HilcoPencilGrammarParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code String}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitString(HilcoPencilGrammarParser.StringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExpressionAtFunction}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionAtFunction(HilcoPencilGrammarParser.ExpressionAtFunctionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExpressionAtFunction}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionAtFunction(HilcoPencilGrammarParser.ExpressionAtFunctionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Date}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterDate(HilcoPencilGrammarParser.DateContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Date}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitDate(HilcoPencilGrammarParser.DateContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Case}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterCase(HilcoPencilGrammarParser.CaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Case}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitCase(HilcoPencilGrammarParser.CaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterInteger(HilcoPencilGrammarParser.IntegerContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitInteger(HilcoPencilGrammarParser.IntegerContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinaryArithmeticCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterBinaryArithmeticCalculator(HilcoPencilGrammarParser.BinaryArithmeticCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinaryArithmeticCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitBinaryArithmeticCalculator(HilcoPencilGrammarParser.BinaryArithmeticCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinaryLogicalCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterBinaryLogicalCalculator(HilcoPencilGrammarParser.BinaryLogicalCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinaryLogicalCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitBinaryLogicalCalculator(HilcoPencilGrammarParser.BinaryLogicalCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExpressionKeyword}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionKeyword(HilcoPencilGrammarParser.ExpressionKeywordContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExpressionKeyword}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionKeyword(HilcoPencilGrammarParser.ExpressionKeywordContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Float}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterFloat(HilcoPencilGrammarParser.FloatContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Float}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitFloat(HilcoPencilGrammarParser.FloatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code UnaryMinusCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterUnaryMinusCalculator(HilcoPencilGrammarParser.UnaryMinusCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnaryMinusCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitUnaryMinusCalculator(HilcoPencilGrammarParser.UnaryMinusCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinaryExponentialCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterBinaryExponentialCalculator(HilcoPencilGrammarParser.BinaryExponentialCalculatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinaryExponentialCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitBinaryExponentialCalculator(HilcoPencilGrammarParser.BinaryExponentialCalculatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExpressionDataAccess}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionDataAccess(HilcoPencilGrammarParser.ExpressionDataAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExpressionDataAccess}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionDataAccess(HilcoPencilGrammarParser.ExpressionDataAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code If}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterIf(HilcoPencilGrammarParser.IfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code If}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitIf(HilcoPencilGrammarParser.IfContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#caseStatement}.
	 * @param ctx the parse tree
	 */
	void enterCaseStatement(HilcoPencilGrammarParser.CaseStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#caseStatement}.
	 * @param ctx the parse tree
	 */
	void exitCaseStatement(HilcoPencilGrammarParser.CaseStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#caseList}.
	 * @param ctx the parse tree
	 */
	void enterCaseList(HilcoPencilGrammarParser.CaseListContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#caseList}.
	 * @param ctx the parse tree
	 */
	void exitCaseList(HilcoPencilGrammarParser.CaseListContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#caseItem}.
	 * @param ctx the parse tree
	 */
	void enterCaseItem(HilcoPencilGrammarParser.CaseItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#caseItem}.
	 * @param ctx the parse tree
	 */
	void exitCaseItem(HilcoPencilGrammarParser.CaseItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(HilcoPencilGrammarParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(HilcoPencilGrammarParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#name}.
	 * @param ctx the parse tree
	 */
	void enterName(HilcoPencilGrammarParser.NameContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#name}.
	 * @param ctx the parse tree
	 */
	void exitName(HilcoPencilGrammarParser.NameContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#worksheetVariable}.
	 * @param ctx the parse tree
	 */
	void enterWorksheetVariable(HilcoPencilGrammarParser.WorksheetVariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#worksheetVariable}.
	 * @param ctx the parse tree
	 */
	void exitWorksheetVariable(HilcoPencilGrammarParser.WorksheetVariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#atFunction}.
	 * @param ctx the parse tree
	 */
	void enterAtFunction(HilcoPencilGrammarParser.AtFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#atFunction}.
	 * @param ctx the parse tree
	 */
	void exitAtFunction(HilcoPencilGrammarParser.AtFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#argList}.
	 * @param ctx the parse tree
	 */
	void enterArgList(HilcoPencilGrammarParser.ArgListContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#argList}.
	 * @param ctx the parse tree
	 */
	void exitArgList(HilcoPencilGrammarParser.ArgListContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#dataAccessor}.
	 * @param ctx the parse tree
	 */
	void enterDataAccessor(HilcoPencilGrammarParser.DataAccessorContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#dataAccessor}.
	 * @param ctx the parse tree
	 */
	void exitDataAccessor(HilcoPencilGrammarParser.DataAccessorContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#accessorList}.
	 * @param ctx the parse tree
	 */
	void enterAccessorList(HilcoPencilGrammarParser.AccessorListContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#accessorList}.
	 * @param ctx the parse tree
	 */
	void exitAccessorList(HilcoPencilGrammarParser.AccessorListContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#accessorObjectOrArray}.
	 * @param ctx the parse tree
	 */
	void enterAccessorObjectOrArray(HilcoPencilGrammarParser.AccessorObjectOrArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#accessorObjectOrArray}.
	 * @param ctx the parse tree
	 */
	void exitAccessorObjectOrArray(HilcoPencilGrammarParser.AccessorObjectOrArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#accessorObject}.
	 * @param ctx the parse tree
	 */
	void enterAccessorObject(HilcoPencilGrammarParser.AccessorObjectContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#accessorObject}.
	 * @param ctx the parse tree
	 */
	void exitAccessorObject(HilcoPencilGrammarParser.AccessorObjectContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#accessorArray}.
	 * @param ctx the parse tree
	 */
	void enterAccessorArray(HilcoPencilGrammarParser.AccessorArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#accessorArray}.
	 * @param ctx the parse tree
	 */
	void exitAccessorArray(HilcoPencilGrammarParser.AccessorArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilGrammarParser#specialKeyword}.
	 * @param ctx the parse tree
	 */
	void enterSpecialKeyword(HilcoPencilGrammarParser.SpecialKeywordContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilGrammarParser#specialKeyword}.
	 * @param ctx the parse tree
	 */
	void exitSpecialKeyword(HilcoPencilGrammarParser.SpecialKeywordContext ctx);
}