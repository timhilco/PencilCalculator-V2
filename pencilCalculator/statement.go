package pencilCalculator

import (
	"context"
	"fmt"
	"strings"
)

type Statement struct {
	StatementKey       string
	Expression         string
	SubscriberChannels map[string]chan string
	ComponentChannels  map[string]chan string
	StatementChannel   chan string
	Result             PencilResult
	ResultType         string
	Variables          []WorksheetVariable
	//Callers            []Statement
}

func (s Statement) String() string {
	var sb strings.Builder
	s1 := fmt.Sprintf("- Statement: %s --\n", s.StatementKey)
	sb.WriteString(s1)
	s1 = fmt.Sprintf("Expression: %s \n", s.Expression)
	sb.WriteString(s1)
	s1 = fmt.Sprintf("Statement(my) Channel: %T -> %p \n", s.StatementChannel, s.StatementChannel)
	sb.WriteString(s1)

	sb.WriteString("Component Channels ->\n")
	for key, entry := range s.ComponentChannels {
		s1 = fmt.Sprintf("%s -> %T:%p \n", key, entry, entry)
		sb.WriteString(s1)
	}
	sb.WriteString("Subscriber Channels ->\n")
	for key, c := range s.SubscriberChannels {
		s1 = fmt.Sprintf("%s -> %T:%p \n", key, c, c)
		sb.WriteString(s1)
	}
	sb.WriteString("Variables ->\n")
	for _, c := range s.Variables {
		s1 = fmt.Sprintf("%v\n", c)
		sb.WriteString(s1)
	}
	s1 = fmt.Sprintf("Result: = %v \n", s.Result)
	sb.WriteString(s1)

	return sb.String()
}
func (s Statement) Evaluate(ctx context.Context, memo *StatementsMemo) PencilResult {
	result := evaluateStatementAsElemental(ctx, memo, s)
	return result
}
func evaluateStatementAsElemental(ctx context.Context, memo *StatementsMemo, s Statement) PencilResult {
	expression := s.Expression
	//sf:= fmt.Sprintf("Evaluating expression: %s\n", expression)
	result := Evaluate(ctx, memo, expression)
	//sf:= fmt.Sprintf("%s=%v\n", expression, result)
	return result
}
