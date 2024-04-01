package codeGenerator

import (
	"context"
	"fmt"
	"go-PencilCalculator/parser"
	"os"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
)

type programModel struct {
	PackageName string
	Statements  map[string]CodeGenStatement
	Functions   []Function
}

func buildVariableNameFromWorksheetName(key string) string {
	s := "prValue_" + key
	v := strings.ReplaceAll(s, ":", "_")
	return v
}
func GenerateProgram(ctx context.Context, client string, program string) (string, error) {
	fileName := "../generated/" + client + "/" + client + ".go"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	statementMap := parseProgram(ctx, program)
	functions := buildFunctions(statementMap)

	defer f.Close()
	pm := programModel{
		PackageName: "a1234",
		Statements:  statementMap,
		Functions:   functions,
	}
	dir := "../generated/templates/*.tmpl"
	t := parseTemplates(dir)
	err = t.ExecuteTemplate(f, "goFile", pm)
	if err != nil {
		fmt.Println(err)
	}
	return "OK", nil
}
func parseTemplates(directory string) *template.Template {
	t, err := template.ParseGlob(directory)
	if err != nil {
		fmt.Println(err)
	}
	return t

}
func parseProgram(ctx context.Context, program string) map[string]CodeGenStatement {

	is := antlr.NewInputStream(program) //
	//logger := logger.NewMultiWithFile(false, false, "./logs/evaluateProgram.txt")

	// Create the Lexer
	lexer := parser.NewHilcoPencilGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewHilcoPencilGrammarParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener CodeGeneratorListener

	listener.SetLexer(lexer)
	//listener.SetInputData(jsonObject)
	listener.SetContext(ctx)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	//listenerstatements := listener.GetStatementMap()
	sm := make(map[string]CodeGenStatement)
	c1 := CodeGenStatement{
		StatementKey:     "W1:a1",
		Expression:       "",
		StatementChannel: nil,
		ResultType:       "PencilTypeInteger",
		Variables:        nil,
		VariableKey:      "W1a1",
		FunctionName:     "W1_a1",
		//Callers            []Statement
	}
	sm["W1:a1"] = c1
	c2 := CodeGenStatement{
		StatementKey:     "W1:a2",
		Expression:       "",
		StatementChannel: nil,
		ResultType:       "PencilTypeInteger",
		Variables:        nil,
		VariableKey:      "W1a2",
		FunctionName:     "W1_a2",
		//Callers            []Statement
	}
	sm["W1:a2"] = c2
	c3 := CodeGenStatement{
		StatementKey:     "W1:a3",
		Expression:       "",
		StatementChannel: nil,
		ResultType:       "PencilTypeInteger",
		Variables:        nil,
		VariableKey:      "W1a3",
		FunctionName:     "W1_a3",
		//Callers            []Statement
	}
	sm["W1:a3"] = c3
	return sm
}
func buildFunctions(statements map[string]CodeGenStatement) []Function {
	f := make([]Function, 0)
	fl := make([]string, 1)
	fl = append(fl, "x = 100")
	f1 := Function{
		FunctionName:   "W1_a1",
		ReturnType:     "int64",
		StatementKey:   "W1:a1",
		FunctionLogic:  fl,
		ReturnVariable: "x",
	}
	f = append(f, f1)
	fl = make([]string, 1)
	fl = append(fl, "x = 200")
	f2 := Function{
		FunctionName:   "W1_a2",
		ReturnType:     "int64",
		StatementKey:   "W1:a2",
		FunctionLogic:  fl,
		ReturnVariable: "x",
	}
	f = append(f, f2)

	fl = make([]string, 1)
	fl = append(fl, `y := results["W1:a2"].PrValue.(int64) `)
	fl = append(fl, `x := results["W1:a1"].PrValue.(int64) `)
	fl = append(fl, "z = x + y ")

	f3 := Function{
		FunctionName:   "W1_a3",
		ReturnType:     "int64",
		StatementKey:   "W1:a3",
		ReturnVariable: "z",
		FunctionLogic:  fl,
	}
	f = append(f, f3)
	return f
}
