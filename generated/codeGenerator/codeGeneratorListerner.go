package codeGenerator

import (
	"context"
	"fmt"
	"go-PencilCalculator/parser"
	"go-PencilCalculator/pencilCalculator"
	"go-PencilCalculator/util/logger"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rs/zerolog"
)

type CodeGeneratorListener struct {
	*parser.BaseHilcoPencilGrammarParserListener
	logger            zerolog.Logger
	lexer             *parser.HilcoPencilGrammarLexer
	processingContext context.Context
	statementMap      map[string]CodeGenStatement
	worksheetVariable []WorksheetVariable
	stack             []StatementElement
}
type WorksheetVariable struct {
	Name string
}
type ElementType int

const (
	ElementTypeOperand ElementType = iota
	ElementTypeOperatorBinary
	ElementTypeOperatorRelational
	ElementTypeOperatorLogical
	ElementTypeOperatorSeparator
	ElementTypeControl
)

type StatementElement struct {
	Type    ElementType
	SeValue interface{}
}
type CodeGenStatement struct {
	StatementKey     string
	Expression       string
	StatementChannel chan string
	ResultType       string
	Variables        []WorksheetVariable
	VariableKey      string
	FunctionName     string
	//Callers            []Statement
}
type Function struct {
	FunctionName   string
	ReturnType     string
	StatementKey   string
	FunctionLogic  []string
	ReturnVariable string
}

func (l *CodeGeneratorListener) push(i StatementElement) {
	l.stack = append(l.stack, i)
}

/*
func (p *HilcoPencilGrammarParserListener) getTop() PencilResult {
	return p.stack[(len(p.stack) - 1)]
}
*/

func (l *CodeGeneratorListener) pop() StatementElement {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}
func (p *CodeGeneratorListener) logStack() {
	/*
		if len(p.stack) == 0 {
			p.logger.Debug().Msg("Stack is empty")
			return
		}
		for i, pr := range p.stack {
				text := fmt.Sprintf("%d:%s", i, pr.String())
			p.logger.Debug().Msg(text)
		}
	*/
}
func (p *CodeGeneratorListener) GetStatementMap() map[string]CodeGenStatement {
	return p.statementMap
}

func (p *CodeGeneratorListener) SetLexer(lexer *parser.HilcoPencilGrammarLexer) {
	p.lexer = lexer
}
func (p *CodeGeneratorListener) SetContext(ctx context.Context) {
	p.processingContext = ctx

}

func (p *CodeGeneratorListener) EnterProgram(ctx *parser.ProgramContext) {

	loggingLevel := p.processingContext.Value(pencilCalculator.LoggingLevelContextKey{}).(int)
	//p.logger = logger.NewMultiWithFile(loggingLevel, false, "../logs/codeGeneratorListener.txt", true)
	p.logger = logger.NewMultiWithFile(loggingLevel, false, "", false)
	p.logger.Debug().Msg("EnterProgram")
	p.statementMap = make(map[string]CodeGenStatement, 0)
}
func (p *CodeGeneratorListener) ExitProgram(ctx *parser.ProgramContext) {

	p.logger.Debug().Msg("------------------ Exit Program Listener -------------------------")

}

// EnterStatement is called when production statement is entered.
func (s *CodeGeneratorListener) EnterStatement(ctx *parser.StatementContext) {
	s.logger.Debug().Msg("<<Enter EnterStatement: ")
	s.worksheetVariable = make([]WorksheetVariable, 0)

}

// ExitStatement is called when production statement is exited.
func (s *CodeGeneratorListener) ExitStatement(ctx *parser.StatementContext) {
	s.logger.Debug().Msg(">>Enter ExitStatement: ")

	//s.statementMap[key] = statement
	//s.logger.Debug().Msg(">>Exit ExitStatement: " + statement.String())

}
func (s CodeGenStatement) String() string {
	var sb strings.Builder
	s1 := fmt.Sprintf("- Statement: %s --\n", s.StatementKey)
	sb.WriteString(s1)
	s1 = fmt.Sprintf("Expression: %s \n", s.Expression)
	sb.WriteString(s1)
	s1 = fmt.Sprintf("Statement(my) Channel: %T -> %p \n", s.StatementChannel, s.StatementChannel)
	sb.WriteString(s1)

	sb.WriteString("Component Channels ->\n")

	sb.WriteString("Variables ->\n")
	for _, c := range s.Variables {
		s1 = fmt.Sprintf("%v\n", c)
		sb.WriteString(s1)
	}

	return sb.String()
}
func (s *CodeGeneratorListener) EnterWorksheetVariable(ctx *parser.WorksheetVariableContext) {
	s.logger.Debug().Msg("Enter EnterWorksheetVariable: ")
}

// ExitWorksheetVariable is called when production worksheetVariable is exited.
func (s *CodeGeneratorListener) ExitWorksheetVariable(ctx *parser.WorksheetVariableContext) {
	s.logger.Debug().Msg("Exit  EnterWorksheetVariable: ")
	t := ctx.GetText()
	variable := WorksheetVariable{
		Name: t,
	}
	s.worksheetVariable = append(s.worksheetVariable, variable)

}
func (p *CodeGeneratorListener) VisitTerminal(node antlr.TerminalNode) {
	p.logger.Debug().Msg("Enter VisitTerminal: " + node.GetText())
	token := node.GetSymbol()
	symbol := token.GetTokenType()
	name := p.lexer.SymbolicNames[token.GetTokenType()]
	value := node.GetText()
	text := fmt.Sprintf("%s(%d) -> %s", name, symbol, value)
	switch name {

	case "AND":
		p.logger.Debug().Msg("VisitTerminal <> Boolean Logical Operator: " + text)
		se := StatementElement{
			Type:    ElementTypeOperatorLogical,
			SeValue: "&&",
		}
		p.push(se)

	case "OR":
		p.logger.Debug().Msg("VisitTerminal <> Boolean Logical Operator: " + text)
		se := StatementElement{
			Type:    ElementTypeOperatorLogical,
			SeValue: "||",
		}
		p.push(se)

	case "ADD",
		"MINUS",
		"MULTIPLY",
		"DIVIDE":
		p.logger.Debug().Msg("VisitTerminal <> Binary Operator: " + text)
		se := StatementElement{
			Type:    ElementTypeOperatorBinary,
			SeValue: value,
		}
		p.push(se)
	case
		"LESS_THAN",
		"GREATER_THAN",
		"EQUAL",
		"NOT_EQUAL":
		p.logger.Debug().Msg("VisitTerminal <> Relational Operator: " + text)
		se := StatementElement{
			Type:    ElementTypeOperatorRelational,
			SeValue: value,
		}
		p.push(se)

	case "COMMA",
		"LPAREN",
		"RPAREN":
		p.logger.Debug().Msg("VisitTerminal <> Separator: " + text)
		se := StatementElement{
			Type:    ElementTypeOperatorSeparator,
			SeValue: value,
		}
		p.push(se)
	case "INT",
		"FLOAT",
		"STRING":
		p.logger.Debug().Msg("VisitTerminal <> Operand: " + text)
		se := StatementElement{
			Type:    ElementTypeOperand,
			SeValue: value,
		}
		p.push(se)

	default:
		p.logger.Debug().Msg("VisitTerminal <> Ignoring Terminal:  " + name)
	}

	p.logger.Debug().Msg("Exit VisitTerminal: " + text)

}
