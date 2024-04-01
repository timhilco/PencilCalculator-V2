package tests

import (
	"context"
	"fmt"
	"go-PencilCalculator/generated/a1234"
	"go-PencilCalculator/generated/codeGenerator"
	"go-PencilCalculator/generated/x9999"
	"go-PencilCalculator/pencilCalculator"
	"testing"

	"github.com/rs/zerolog"
)

func Test_A1234_Generate(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, pencilCalculator.LoggingLevelContextKey{}, zerolog.DebugLevel)
	program := "w1:a1 :=  100;"
	client := "a1234"
	status, _ := codeGenerator.GenerateProgram(ctx, client, program)
	fmt.Println(status)

}

func Test_A1234_Run(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, pencilCalculator.LoggingLevelContextKey{}, zerolog.DebugLevel)

	input := buildInputData()
	results, _ := a1234.RunProgram(ctx, input)
	printResults(results)

}
func Test_x9999(t *testing.T) {

	ctx := context.Background()
	input := buildInputData()
	results, _ := x9999.RunProgram(ctx, input)
	printResults(results)
}
func buildInputData() []byte {
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
	return []byte(empJson)
}
func printResults(output map[string]pencilCalculator.PencilResult) {
	for key, result := range output {
		fmt.Printf("%s = %v\n", key, result.PrValue)
	}
}
