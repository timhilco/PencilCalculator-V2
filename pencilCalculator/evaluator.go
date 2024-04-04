package pencilCalculator

import (
	"context"
	"fmt"
	"go-PencilCalculator/parser"
	"sort"
	"sync"
	"time"

	"go-PencilCalculator/util/logger"

	"github.com/antlr4-go/antlr/v4"
)

func Evaluate(ctx context.Context, memo *StatementsMemo, input string) PencilResult {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewHilcoPencilGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewHilcoPencilGrammarParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener HilcoPencilGrammarParserListener

	listener.SetLexer(lexer)
	//listener.SetInputData(jsonObject)
	listener.SetContext(ctx)
	listener.SetStatementsMemo(memo)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	value := listener.Result()
	if value.Type == PencilTypeIntegerFloat {
		v := value.PrValue.(FloatIntegerNumber)
		value = PencilResult{
			Type:    PencilTypeFloat,
			PrValue: v.ConvertFloatIntToFloatInputPlaces(v.Precision),
		}
	}
	return value

}
func EvaluateProgram(ctx context.Context, program string) map[string]*Statement {
	// Setup the input
	is := antlr.NewInputStream(program) //
	loggingLevel := ctx.Value(LoggingLevelContextKey{})
	loggingLevel2, _ := loggingLevel.(int)

	logger := logger.NewMultiWithFile(loggingLevel2, false, "../logs/evaluateProgram.txt", true)

	// Create the Lexer
	lexer := parser.NewHilcoPencilGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewHilcoPencilGrammarParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener HilcoPencilStatementListener

	listener.SetLexer(lexer)
	//listener.SetInputData(jsonObject)
	listener.SetContext(ctx)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	statements := listener.GetStatementMap()
	for sKey := range statements {
		statement := statements[sKey]
		variables := statement.Variables
		for _, variable := range variables {
			key := variable.Name
			componentStatement := statements[key]
			// Make Statement Changes - Components
			components := statement.ComponentChannels
			components[key] = componentStatement.StatementChannel
			statement.ComponentChannels = components
			statements[sKey] = statement
			// Subscriber Channels
			subscribers := componentStatement.SubscriberChannels
			subscribers[sKey] = statement.StatementChannel
			componentStatement.SubscriberChannels = subscribers
			statements[key] = componentStatement
		}

	}
	memo := NewStatementsMemo(statements)
	ctx = context.WithValue(ctx, StatementMapContextKey{}, statements)
	ctx = context.WithValue(ctx, StatementsMemoContextKey{}, &memo)
	ctx = context.WithValue(ctx, ProcessingPassContextKey{}, "Elemental")
	elementalStatements := make([]*Statement, 0)
	dependentStatements := make([]*Statement, 0)
	statementsCountMap := make(map[int][]*Statement)
	for key, statement := range statements {
		//logger.Info().Msg(statement.String())
		var msg string
		if len(statement.ComponentChannels) != 0 {
			msg = key + " is a dependent statement"
			dependentStatements = append(dependentStatements, statement)
		} else {
			msg = key + " is an elemental statement"
			elementalStatements = append(elementalStatements, statement)
		}
		logger.Debug().Msg(msg)
		count := len(statement.ComponentChannels)
		sa := statementsCountMap[count]
		if sa == nil {
			a := make([]*Statement, 0)
			a = append(a, statement)
			statementsCountMap[count] = a
		} else {
			sa = append(sa, statement)
			statementsCountMap[count] = sa
		}
	}

	var wg sync.WaitGroup
	keys := make([]int, 0, len(statementsCountMap))
	for k := range statementsCountMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := range keys {
		statements := statementsCountMap[keys[i]]
		for _, statement := range statements {
			wg.Add(1)
			go func(s *Statement) {
				defer wg.Done()
				startTime := time.Now()
				skey := s.StatementKey
				sf := fmt.Sprintf("** Evaluating statement: %s", skey)
				logger.Info().Msg(sf)
				result, _ := memo.Evaluate(skey, ctx)
				sf = fmt.Sprintf("-- Evaluated statement: %v = %v", skey, result)
				logger.Debug().Msg(sf)

				s.Result = result
				endTime := time.Now()
				d := endTime.Sub(startTime)
				sf = fmt.Sprintf("** Statement processing complete: %s Time: %v", skey, d)
				logger.Info().Msg(sf)
			}(statement)
		}
	}
	wg.Wait()
	/*
		for _, statement := range elementalStatements {
			wg.Add(1)
			go func(s *Statement) {
				defer wg.Done()
				startTime := time.Now()
				key := s.StatementKey
				sf := fmt.Sprintf("** Evaluating elemental statement: %s", key)
				logger.Debug().Msg(sf)
				result, _ := memo.Evaluate(key, ctx)
				sf = fmt.Sprintf("-- Evaluated elemental statement: %v = %v", key, result)
				logger.Debug().Msg(sf)
				/* Update statements from memo cache - Hopefully dependents are the same
				s.Result = result
				statements[key] = s

					for key, subscriber := range s.SubscriberChannels {
						sf = fmt.Sprintf(" Calling Elemental Subscribers: %s->%s:%v - Result:%v\n", s.StatementKey, key, subscriber, result)
						logger.Debug().Msg(sf)
						subscriber <- s.StatementKey
						sf = fmt.Sprintf(" Called  Elemental Subscribers: %s->%s:%v \n", s.StatementKey, key, subscriber)
						logger.Debug().Msg(sf)

				endTime := time.Now()
				d := endTime.Sub(startTime)
				sf = fmt.Sprintf("** Elemental statement processing complete: %s Time: %v", key, d)
				logger.Info().Msg(sf)
			}(statement)

		}
		wg.Wait()
		for _, statement := range dependentStatements {
			wg.Add(1)
			go func(s *Statement) {
				defer wg.Done()
				startTime := time.Now()
				key := s.StatementKey
				sf := fmt.Sprintf("** Evaluating dependent statement: %s", key)
				logger.Debug().Msg(sf)
				result, _ := memo.Evaluate(key, ctx)
				sf = fmt.Sprintf("-- Evaluated dependent statement: %v = %v", key, result)
				logger.Debug().Msg(sf)
				/* Update statements from memo cache - Hopefully dependents are the same
				s.Result = result
				statements[key] = s

					for key, subscriber := range s.SubscriberChannels {
						sf = fmt.Sprintf(" Calling Elemental Subscribers: %s->%s:%v - Result:%v\n", s.StatementKey, key, subscriber, result)
						logger.Debug().Msg(sf)
						subscriber <- s.StatementKey
						sf = fmt.Sprintf(" Called  Elemental Subscribers: %s->%s:%v \n", s.StatementKey, key, subscriber)
						logger.Debug().Msg(sf)

				endTime := time.Now()
				d := endTime.Sub(startTime)
				sf = fmt.Sprintf("** Dependent statement processing complete: %s Time: %v", s.StatementKey, d)
				logger.Info().Msg(sf)
			}(statement)

		}
		wg.Wait()
	*/
	/*
		ctx2 := context.Background()
		ctx2 = context.WithValue(ctx2, ProcessingPassContextKey{}, "Dependent")
		ctx2 = context.WithValue(ctx2, StatementMapContextKey{}, statements)
		inputData := ctx.Value(InputDataContextKey{}).(map[string]string)
		ctx2 = context.WithValue(ctx2, InputDataContextKey{}, inputData)

		for _, statement := range dependentStatements {
			wg.Add(1)
			go func(s Statement) {
				defer wg.Done()
				startTime := time.Now()
				count := len(s.ComponentChannels)
				a := float64(rand.Intn(10)) + 5.00
				n := math.Abs(a)
				if n == 0.0 {
					n = 1.0
				}
				d := time.Duration(n) * time.Second
				ticker := time.NewTicker(d)
				run := true
				i := 0
				c := s.StatementChannel
				sf := fmt.Sprintf("-> Going to wait for: %s:%v\n", s.StatementKey, c)
				logger.Debug().Msg(sf)
				for run {
					select {
					case <-ticker.C:
						t := time.Now()
						fTime := t.Format(time.RFC3339)
						sf := fmt.Sprintf("-> Hit ticker %s for %s\n", fTime, s.StatementKey)
						logger.Debug().Msg(sf)
						//run = false
					case <-c:
						sf := fmt.Sprintf("-> Subscriber published for statement: %s", s.StatementKey)
						logger.Debug().Msg(sf)
						i++
						if i == count {
							run = false
						}
					default:
						//sf := fmt.Sprintf("-> Hit default: %s", s.StatementKey)
						//logger.Debug().Msg(sf)
						//d := time.Duration(1) * time.Second
						//time.Sleep(d)

					}
				}

				sf = fmt.Sprintf("** Evaluating dependent statement: %s\n", s.StatementKey)
				logger.Debug().Msg(sf)
				result := evaluateStatement(ctx2, s)
				s.Result = result
				key := s.StatementKey
				statements[key] = s
				sf = fmt.Sprintf("-- Evaluated dependent statement: %s = %v\n", s.StatementKey, result)
				logger.Debug().Msg(sf)
				for key, subscriber := range s.SubscriberChannels {
					sf := fmt.Sprintf("Calling Dependent Subscriber: %s->%s:%v - Result:%v\n", s.StatementKey, key, subscriber, result)
					logger.Debug().Msg(sf)
					subscriber <- s.StatementKey
					sf = fmt.Sprintf("Called   Dependent Subscriber: %s->%s%v\n", s.StatementKey, key, subscriber)
					logger.Debug().Msg(sf)
				}
				endTime := time.Now()
				dt := endTime.Sub(startTime)
				sf = fmt.Sprintf("** Dependent statement processing complete: %s Time: %v\n", s.StatementKey, dt)
				logger.Info().Msg(sf)
			}(statement)
			//sf:= fmt.Sprintf("Ignoring statement: %v\n", statement.VariableName)

		}
	*/

	return statements
}
func evaluateStatement(ctx context.Context, s Statement, memo *StatementsMemo) PencilResult {
	expression := s.Expression
	//sf:= fmt.Sprintf("Evaluating expression: %s\n", expression)
	result := Evaluate(ctx, memo, expression)
	//sf:= fmt.Sprintf("%s=%v\n", expression, result)
	return result
}
