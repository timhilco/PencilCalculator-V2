package main

import (
	"context"
	"fmt"
	"go-PencilCalculator/pencilCalculator"
	"go-PencilCalculator/tests"
	"os"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
	var programName string = "program1"
	num := 1000
	tests.BuildProgram(num, programName)
	if len(os.Args) > 1 {
		programName = os.Args[1]
	}
	ctx := context.Background()
	empJson := `{
		"id": 11,
		"name": { "first": "John", "last" : "Sample"},
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000,
			"units" : 100
			}, {
				"effectiveDate": "01/01/2022",
				"amount": 50000,
				"units" : 110
				}]
				
				}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)
	var program string
	b, err := os.ReadFile("./programs/program1.pp")
	switch programName {
	case "program1":
		program = string(b)
		if err != nil {
			fmt.Println(err)
		}
	}
	statements := pencilCalculator.EvaluateProgram(ctx, program)
	fmt.Println("---------------Program Results-------------------------------------")
	for _, statement := range statements {
		fmt.Printf("Statement: %s = %v\n", statement.StatementKey, statement.Result)
	}

}
