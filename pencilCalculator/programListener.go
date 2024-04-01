package pencilCalculator

import (
	"context"
	"strings"

	"go-PencilCalculator/parser"
	"go-PencilCalculator/util/logger"

	"github.com/rs/zerolog"
)

type HilcoPencilStatementListener struct {
	*parser.BaseHilcoPencilGrammarParserListener
	logger            zerolog.Logger
	lexer             *parser.HilcoPencilGrammarLexer
	processingContext context.Context
	statementMap      map[string]*Statement
	worksheetVariable []WorksheetVariable
}
type WorksheetVariable struct {
	Name string
}

func (p *HilcoPencilStatementListener) GetStatementMap() map[string]*Statement {
	return p.statementMap
}

func (p *HilcoPencilStatementListener) SetLexer(lexer *parser.HilcoPencilGrammarLexer) {
	p.lexer = lexer
}
func (p *HilcoPencilStatementListener) SetContext(ctx context.Context) {
	p.processingContext = ctx

}
func (p *HilcoPencilStatementListener) Result() PencilResult {
	return PencilResult{}
}
func (p *HilcoPencilStatementListener) EnterProgram(ctx *parser.ProgramContext) {
	loggingLevel := p.processingContext.Value(LoggingLevelContextKey{})
	loggingLevel2, _ := loggingLevel.(int)
	p.logger = logger.NewMultiWithFile(loggingLevel2, false, "./logs/hilcoPencilStatementListener.txt", true)
	p.logger.Debug().Msg("EnterProgram")
	p.statementMap = make(map[string]*Statement, 0)
}
func (p *HilcoPencilStatementListener) ExitProgram(ctx *parser.ProgramContext) {

	p.logger.Debug().Msg("------------------ Exit Program Listener -------------------------")

}

// EnterStatement is called when production statement is entered.
func (s *HilcoPencilStatementListener) EnterStatement(ctx *parser.StatementContext) {
	s.logger.Debug().Msg("<<Enter EnterStatement: ")
	s.worksheetVariable = make([]WorksheetVariable, 0)

}

// ExitStatement is called when production statement is exited.
func (s *HilcoPencilStatementListener) ExitStatement(ctx *parser.StatementContext) {
	s.logger.Debug().Msg(">>Enter ExitStatement: ")
	s0 := ctx.GetText()
	value := strings.Split(s0, ":=")
	s1 := value[1]
	v := strings.Trim(s1, ";")
	children := ctx.GetChildren()
	var variables []WorksheetVariable
	key := children[0].(*parser.WorkSheetVariableCalculatorContext).GetText()
	cCH := make(map[string]chan string)
	for i := 1; i < len(s.worksheetVariable); i++ {
		variables = append(variables, s.worksheetVariable[i])
	}
	r := PencilResult{
		Type:    PencilTypeNil,
		PrValue: nil,
	}

	subCH := make(map[string]chan string)
	sCH := make(chan string)

	statement := Statement{
		StatementKey:       key,
		Expression:         v,
		Variables:          variables,
		StatementChannel:   sCH,
		SubscriberChannels: subCH,
		ComponentChannels:  cCH,
		Result:             r,
		ResultType:         "PencilTypeInteger",
	}
	s.statementMap[key] = &statement
	//s.logger.Debug().Msg(">>Exit ExitStatement: " + statement.String())

}

// ExitWorksheetVariable is called when production worksheetVariable is exited.
func (s *HilcoPencilStatementListener) ExitWorksheetVariable(ctx *parser.WorksheetVariableContext) {
	s.logger.Debug().Msg("Exit  EnterWorksheetVariable: ")
	t := ctx.GetText()
	variable := WorksheetVariable{
		Name: t,
	}
	s.worksheetVariable = append(s.worksheetVariable, variable)

}

func (s *HilcoPencilStatementListener) EnterWorksheetVariable(ctx *parser.WorksheetVariableContext) {
	s.logger.Debug().Msg("Enter EnterWorksheetVariable: ")
}
