// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilGrammarParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link HilcoPencilGrammarParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface HilcoPencilGrammarParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#program}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProgram(HilcoPencilGrammarParser.ProgramContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatement(HilcoPencilGrammarParser.StatementContext ctx);
	/**
	 * Visit a parse tree produced by the {@code NameCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNameCalculator(HilcoPencilGrammarParser.NameCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code UnaryNotCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnaryNotCalculator(HilcoPencilGrammarParser.UnaryNotCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Percent}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPercent(HilcoPencilGrammarParser.PercentContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParens(HilcoPencilGrammarParser.ParensContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BinaryRelationalCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryRelationalCalculator(HilcoPencilGrammarParser.BinaryRelationalCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code WorkSheetVariableCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWorkSheetVariableCalculator(HilcoPencilGrammarParser.WorkSheetVariableCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code String}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitString(HilcoPencilGrammarParser.StringContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ExpressionAtFunction}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionAtFunction(HilcoPencilGrammarParser.ExpressionAtFunctionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Date}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDate(HilcoPencilGrammarParser.DateContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Case}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCase(HilcoPencilGrammarParser.CaseContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInteger(HilcoPencilGrammarParser.IntegerContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BinaryArithmeticCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryArithmeticCalculator(HilcoPencilGrammarParser.BinaryArithmeticCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BinaryLogicalCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryLogicalCalculator(HilcoPencilGrammarParser.BinaryLogicalCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ExpressionKeyword}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionKeyword(HilcoPencilGrammarParser.ExpressionKeywordContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Float}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFloat(HilcoPencilGrammarParser.FloatContext ctx);
	/**
	 * Visit a parse tree produced by the {@code UnaryMinusCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnaryMinusCalculator(HilcoPencilGrammarParser.UnaryMinusCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BinaryExponentialCalculator}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryExponentialCalculator(HilcoPencilGrammarParser.BinaryExponentialCalculatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ExpressionDataAccess}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionDataAccess(HilcoPencilGrammarParser.ExpressionDataAccessContext ctx);
	/**
	 * Visit a parse tree produced by the {@code If}
	 * labeled alternative in {@link HilcoPencilGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIf(HilcoPencilGrammarParser.IfContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#caseStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseStatement(HilcoPencilGrammarParser.CaseStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#caseList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseList(HilcoPencilGrammarParser.CaseListContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#caseItem}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseItem(HilcoPencilGrammarParser.CaseItemContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIfStatement(HilcoPencilGrammarParser.IfStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitName(HilcoPencilGrammarParser.NameContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#worksheetVariable}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWorksheetVariable(HilcoPencilGrammarParser.WorksheetVariableContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#atFunction}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAtFunction(HilcoPencilGrammarParser.AtFunctionContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#argList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgList(HilcoPencilGrammarParser.ArgListContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#dataAccessor}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDataAccessor(HilcoPencilGrammarParser.DataAccessorContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#accessorList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAccessorList(HilcoPencilGrammarParser.AccessorListContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#accessorObjectOrArray}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAccessorObjectOrArray(HilcoPencilGrammarParser.AccessorObjectOrArrayContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#accessorObject}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAccessorObject(HilcoPencilGrammarParser.AccessorObjectContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#accessorArray}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAccessorArray(HilcoPencilGrammarParser.AccessorArrayContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilGrammarParser#specialKeyword}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSpecialKeyword(HilcoPencilGrammarParser.SpecialKeywordContext ctx);
}