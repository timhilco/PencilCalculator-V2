package tests

import (
	"context"
	"fmt"
	"go-PencilCalculator/pencilCalculator"
	"os"
	"testing"
	"time"
)

func TestBuildAProgram(t *testing.T) {
	num := 10
	name := "program1"
	BuildProgram(num, name)
}
func TestProgram1(t *testing.T) {
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

	b, err := os.ReadFile("./data/data.txt")
	program := string(b)
	if err != nil {
		fmt.Println(err)
	}
	statements := pencilCalculator.EvaluateProgram(ctx, program)
	fmt.Println("---------------Program Results-------------------------------------")
	for _, statement := range statements {
		fmt.Printf("Statement: %s = %v\n", statement.StatementKey, statement.Result)
	}

}
func TestProgram2(t *testing.T) {
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

	b, _ := os.ReadFile("./data/data.txt")
	//b := buildProgram1()
	program := string(b)
	statements := pencilCalculator.EvaluateProgram(ctx, program)
	fmt.Println("---------------Program Results-------------------------------------")
	for _, statement := range statements {
		fmt.Printf("Statement: %s = %v\n", statement.StatementKey, statement.Result)
	}

}

func TestProgram3(t *testing.T) {
	//program := buildRandomProgram()
	//fmt.Println(program)
}

func TestChannel(t *testing.T) {
	cell1 := Cell{
		formula:            10,
		SubscriberChannels: make([]chan int, 0),
		//postingChannel:      make(chan chan int),
	}
	for i := 0; i < 200; i++ {
		cj := make(chan int)
		fmt.Printf("GoRoutine %d Posting\n", i)
		cell1.AddSubscriberChannel(cj)
		go func(j int) {
			fmt.Printf("GoRoutine %d Waiting\n", j)
			value := <-cj * j
			fmt.Printf("J: %d = %d\n", j, value)
		}(i)
	}
	cell1.Evaluate()
	time.Sleep(10 * time.Second)
	fmt.Println("Program Ending")

}

type Cell struct {
	formula            int
	SubscriberChannels []chan int
	//postingChannel      chan chan int
}

func (c *Cell) AddSubscriberChannel(ch chan int) {

	fmt.Printf("Cell adding subscriber\n")
	c.SubscriberChannels = append(c.SubscriberChannels, ch)
	fmt.Printf("Cell adding subscriber length: %d\n", len(c.SubscriberChannels))

}

func (c *Cell) Evaluate() {
	fmt.Printf("--> Evaluate \n")
	value := c.formula
	for _, channel := range c.SubscriberChannels {
		fmt.Printf("Evaluate posting to channel\n")
		//time.Sleep(3 * time.Second)
		channel <- value
	}
}
