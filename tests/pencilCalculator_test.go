package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"go-PencilCalculator/api"
	"go-PencilCalculator/pencilCalculator"
	"go-PencilCalculator/util/logger"
	"math"
	"testing"

	"github.com/rs/zerolog"
)

func Test1(t *testing.T) {
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
	ctx := context.Background()
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)
	statement := "@Now()"

	result := pencilCalculator.Evaluate(ctx, nil, statement)
	want := 1.524156
	fmt.Printf("Result:%v -> Want:%f", result, want)
}
func TestRounding(t *testing.T) {
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
	ctx := context.Background()
	inputData := make(map[string]string)
	inputData["Employee"] = empJson
	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)

	f := 0.33 * 0.33
	g := (3333 * 3333) / math.Pow10(2*2)
	fmt.Printf("%.10f", f)
	//statement := "@RoundToDecimalPlaces(@RoundToDecimalPlaces(0.3333,2)*@RoundToDecimalPlaces(0.3333,2),2)"
	statement := "@NewFloatIntegerBased (3333,2)*@NewFloatIntegerBased (3333,2)"
	result := pencilCalculator.Evaluate(ctx, nil, statement)
	//rf := result.PrValue.(float64)
	fmt.Printf("Result:%v -> Want:%.7f", result, g)
}

/*
func TestTypeConversations(t *testing.T) {

	f64, _ := pencilCalculator.ConvertFloatStringToFloatIntegerNumber("123.45")
	fi1 := pencilCalculator.PencilResult{
		Type:    pencilCalculator.PencilTypeIntegerFloat,
		PrValue: f64,
	}
	f64, _ = pencilCalculator.ConvertFloatStringToFloatIntegerNumber("123.456")
	fi2 := pencilCalculator.PencilResult{
		Type:    pencilCalculator.PencilTypeIntegerFloat,
		PrValue: f64,
	}
	var tests = []struct {
		left      pencilCalculator.PencilResult
		right     pencilCalculator.PencilResult
		operation string
		want      interface{}
	}{
		{fi1, fi2, "ADD", float64(246.906)},
		{pencilCalculator.PencilResult{pencilCalculator.PencilTypeInteger, int64(10)}, fi1, "ADD", float64(133.45)},
		{pencilCalculator.PencilResult{pencilCalculator.PencilTypeFloat, 10.0}, pencilCalculator.PencilResult{pencilCalculator.PencilTypeFloat, 20.0}, "ADD", float64(30.0)},
		{pencilCalculator.PencilResult{pencilCalculator.PencilTypeInteger, int64(10)}, pencilCalculator.PencilResult{pencilCalculator.PencilTypeFloat, 20.0}, "ADD", float64(30.0)},
		{pencilCalculator.PencilResult{pencilCalculator.PencilTypeFloat, 10.0}, pencilCalculator.PencilResult{pencilCalculator.PencilTypeInteger, int64(20)}, "ADD", float64(30.0)},
		{pencilCalculator.PencilResult{pencilCalculator.PencilTypeInteger, int64(10)}, pencilCalculator.PencilResult{pencilCalculator.PencilTypeInteger, int64(20)}, "ADD", int64(30)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.left, tt.right)
		t.Run(testname, func(t *testing.T) {
			ans, _ := pencilCalculator.DoBinaryArithmetic(tt.left, tt.right, tt.operation)
			if ans.PrValue != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
*/
func TestExpressions(t *testing.T) {
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
	var tests = []struct {
		expression string
		want       interface{}
	}{
		// Passed test below
		{"@Max (100,@Max (200, 400))", float64(400.0)},
		{"@Max (500,@Max (200, 400))", float64(500.0)},
		{"10000 * (9 + 10000)", int64(100090000)},
		{"10000 * 9 + 10000", int64(100000)},
		{"(9 + 10000)", int64(10009)},
		{"@Max (500, 10000) * (9 + 10000)", float64(100090000.0)},
		{"@Max (500,(5000*2)) * (9 + 10000)", float64(100090000)},
		{"Employee.salary * @Max(1,2)", float64(100000.0)},
		{"Employee.salary + @Max(1,2)", float64(50002.0)},
		{"Employee.salary - @Max(1,2)", float64(49998.0)},
		{"Employee.salary / @Max(1,2)", float64(25000.0)},
		{"Employee.salary * 2", float64(100000.0)},
		{"Employee.salary + 2", float64(50002.0)},
		{"Employee.salary - 2", float64(49998.0)},
		{"Employee.salary / 2", float64(25000.0)},
		{"Employee.salary * 1", float64(50000.0)},
		{"Employee.salary + 1", float64(50001.0)},
		{"Employee.salary - 1", float64(49999.0)},
		{"Employee.salary / 1", float64(50000.0)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,1)/@NewFloatIntegerBased (3333,2),2)", float64(10.0)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,1)+@NewFloatIntegerBased (3333,2),2)", float64(366.63)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,1)-@NewFloatIntegerBased (3333,2),2)", float64(299.97)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,1)*@NewFloatIntegerBased (3333,2),2)", float64(11108.89)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)*@NewFloatIntegerBased (3333,1),2)", float64(11108.89)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)/@NewFloatIntegerBased (3333,1),2)", float64(0.1)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)+@NewFloatIntegerBased (3333,1),2)", float64(366.63)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)-@NewFloatIntegerBased (3333,1),2)", float64(-299.97)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)/@NewFloatIntegerBased (3333,2),2)", float64(1.0)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)+@NewFloatIntegerBased (3333,2),2)", float64(66.66)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)-@NewFloatIntegerBased (3333,2),2)", float64(0.0)},
		{"@RoundToDecimalPlaces(@NewFloatIntegerBased (3333,2)*@NewFloatIntegerBased (3333,2),2)", float64(1110.89)},
		{"@NewFloatIntegerBased (3333,2)*@NewFloatIntegerBased (3333,2)", float64(1110.8889)},
		{"Employee.payHistory(effectiveDate={01-01-2022},amount=50000).units", float64(110)},
		{"100+200+300", int64(600)},
		{"@Max(1,2)", float64(2.0)},
		{"@Now()", "Time"},
		{"Employee.salary", float64(50000.0)},
		{"Employee.salary > 10", true},
		{"Employee.salary > 10 AND 10 > 1", true},
		{"case Employee.department is 'IT' :: 1; 'default' :: 2 ; endcase", int64(1)},
		{"case Employee.department is 'ITS' :: 1; 'default' :: 2 ; endcase", int64(2)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Expression: %s", tt.expression)
		t.Run(testname, func(t *testing.T) {
			ans := pencilCalculator.Evaluate(ctx, nil, tt.expression)
			isPass, sResult, sWant, message := DoesTestPass(ans, tt.want)
			if !isPass {
				t.Errorf("got %s, want %s --> %s", sResult, sWant, message)
			}
		})
	}

}
func DoesTestPass(result pencilCalculator.PencilResult, want interface{}) (bool, string, string, string) {
	var pass bool = false
	var sResult string
	var sWant string
	var message string = "N/A"
	switch result.Type {
	case pencilCalculator.PencilTypeInteger:
		sResult = fmt.Sprintf("%d", result.PrValue)
		sWant = fmt.Sprintf("%d", want)
		if sResult == sWant {
			pass = true
		}
	case pencilCalculator.PencilTypeFloat:
		sResult = fmt.Sprintf("%.7f", result.PrValue)
		sWant = fmt.Sprintf("%.7f", want)
		if sResult == sWant {
			pass = true
		} else {
			sResult = fmt.Sprintf("%.10f", result.PrValue)
			sWant = fmt.Sprintf("%.10f", want)
			f1 := result.PrValue.(float64)
			f2 := want.(float64)
			diff := f1 - f2
			message = fmt.Sprintf("ERROR: Difference: %.10f over 0.00009", diff)
			if math.Abs(diff) < 0.0009 {
				message = fmt.Sprintf("Difference: %.10f", diff)
				pass = true
			}

		}
	case pencilCalculator.PencilTypeString:
		sResult = fmt.Sprintf("%s", result.PrValue)
		sWant = fmt.Sprintf("%s", want)
		if sResult == sWant {
			pass = true
		}
	case pencilCalculator.PencilTypeBoolean:
		sResult = fmt.Sprintf("%t", result.PrValue)
		sWant = fmt.Sprintf("%t", want)
		if sResult == sWant {
			pass = true
		}
	case pencilCalculator.PencilTypeDateTime:
		if want == "Time" {
			pass = true
		}
		sResult = fmt.Sprintf("%v", result.PrValue)
		sWant = fmt.Sprintf("%v", want)
		if sResult == sWant {
			pass = true
		}

	default:
		pass = false
		sResult = " UnKnown Result"
		sWant = fmt.Sprintf("%v", want)
	}
	return pass, sResult, sWant, message
}
func TestFloat(t *testing.T) {

	var tests = []struct {
		expression string
		want       float64
	}{

		{"123.45", float64(123.45)},
		{"123.1234567", float64(123.123457)},
		{"123.1234565", float64(123.123456)},
		{"123.1234564", float64(123.123456)},
		{"123.123456", float64(123.123456)},
		{"123.12345", float64(123.12345)},
		{"123.1234", float64(123.1234)},
		{"123.123", float64(123.123)},
		{"123.12", float64(123.12)},
		{"123.1", float64(123.1)},
		{"123", float64(123)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Expression: %s", tt.expression)
		t.Run(testname, func(t *testing.T) {
			fi, _ := pencilCalculator.ConvertFloatStringToFloatIntegerNumber(tt.expression)
			ans := fi.ConvertFloatIntToFloat6Decimal()
			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}

}

/*
	tests := [][2]float64{
		{-0.49999999999999994, negZero}, // -0.5+epsilon
		{-0.5, -1},
		{-0.5000000000000001, -1}, // -0.5-epsilon
		{0, 0},
		{0.49999999999999994, 0}, // 0.5-epsilon
		{0.5, 1},
		{0.5000000000000001, 1},                         // 0.5+epsilon
		{1.390671161567e-309, 0},                        //
		{2.2517998136852485e+15, 2.251799813685249e+15}, // 1 bit fraction
		{4.503599627370497e+15, 4.503599627370497e+15},  // large integer
		{math.Inf(-1), math.Inf(-1)},
		{math.Inf(1), math.Inf(1)},
		{math.NaN(), math.NaN()},
		{negZero, negZero},
	}
*/
func TestJSON(t *testing.T) {
	//Simple Employee JSON object which we will parse
	empJson := `{
        "id" : 11,
        "name" : "Irshad",
        "department" : "IT",
        "designation" : "Product Manager",
		"salary": 50000
	}`

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nSalary :", result["salary"])
}
func TestServerStartup(t *testing.T) {
	/*
		log := createLogger()
		processingContext := processor.ProcessingContext{
			Logger: log,
		}
	*/
	ctx := context.Background()
	logger := logger.NewMultiWithFile(logger.LOG_LEVEL_INFO, false, "../logs/serverLog.txt", true)
	options := api.ServerOptions{
		Logger:            logger,
		ServerHostName:    "localhost",
		ServerHostPort:    8080,
		ApiServerHostName: "http://localhost:8080/",
		LocalEvaluate:     true,
		LogGlobalLevel:    zerolog.Disabled,
	}
	api.Start(ctx, options)
}

/*
type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

func TestCall(t *testing.T) {
	StubStorage = map[string]interface{}{
		"funcA": funcA,
		"funcB": funcB,
		"Max":   Max,
	}

		resA, _ := Call("funcA", "value", "keyword1")
		var  A string
		prntA = resA.(string)
		fmt.Println(prntA)

		resB, _ := Call("funcB", 10)
		var prntB int
		prntB = resB.(int)
		fmt.Println(prntB)

	f := reflect.ValueOf(Max)
	resB, _ := Call(f, 10, 20)

	fmt.Println(resB)
}
func Call(f reflect.Value, params ...interface{}) (result interface{}, err error) {
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params are out of index")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}
func funcA(arg0 string, arg1 string) string {
	return fmt.Sprintf("Function one ! %v %v", arg0, arg1)
}
func funcB(args0 int) int {
	return args0
}
*/
