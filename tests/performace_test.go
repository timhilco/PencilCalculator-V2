package tests

import (
	"context"
	"fmt"
	"go-PencilCalculator/pencilCalculator"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func getTwoRandomIntegers(min int, max int) (int64, int64) {
	var i1, i2 int64
	rand.Seed(time.Now().UnixNano())

	i1 = int64(rand.Intn(max-min+1) + min)
	i2 = int64(rand.Intn(max-min+1) + min)

	return i1, i2
}
func TestFloatRounding(t *testing.T) {
	f, _ := getTwoRandomFloats(10000000.0, 20000000.0, false)
	f2 := math.RoundToEven(f * 1000000)
	fmt.Printf("%.10f\n", f2)
	f3 := f2 / 1000000
	fmt.Printf("%.10f\n", f)
	fmt.Printf("%.10f\n", f3)
}

func getTwoRandomFloats(min float64, max float64, isLessThanOne bool) (float64, float64) {

	rand.Seed(time.Now().UnixNano())

	f1 := (rand.Float64() * (max - min)) + min
	var f2 float64
	if isLessThanOne {
		f2 = rand.Float64()
	} else {
		f2 = (rand.Float64() * (max - min)) + min
	}
	return f1, f2
}

/*
type tests struct {
	expression string
	want       interface{}
}
*/

func TestFloatPerformance(t *testing.T) {
	ctx := context.Background()
	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000
		}, {
			"effectiveDate": "01/01/2022",
			"amount": 50000
		}]
	
	}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)
	ctx = context.WithValue(ctx, pencilCalculator.LoggingLevelContextKey{}, zerolog.Disabled)
	var expression string
	var want float64
	//var err error
	for i := 1; i < 25; i++ {
		i1, i2 := getTwoRandomFloats(0.0, 10000.0, true)
		i1s := fmt.Sprintf("%.10f", i1)
		i2s := fmt.Sprintf("%.10f", i2)

		for j := 1; j < 5; j++ {
			switch j {
			case 1:
				expression = fmt.Sprintf("%s+%s", i1s, i2s)
				want = i1 + i2
			case 2:
				expression = fmt.Sprintf("%.6f-%.6f", i1, i2)
				want = i1 - i2
			case 3:
				expression = fmt.Sprintf("%.6f*%.6f", i1, i2)
				want = i1 * i2
			case 4:
				expression = fmt.Sprintf("%.6f/%.6f", i1, i2)
				want = i1 / i2
			}

			testname := fmt.Sprintf("Expression: %s", expression)

			t.Run(testname, func(t *testing.T) {
				ans := pencilCalculator.Evaluate(ctx, nil, expression)
				isPass, sResult, sWant, message := DoesTestPass(ans, want)
				if !isPass {
					t.Errorf("got %s, want %s  --> %s", sResult, sWant, message)
				} else if message != "N/A" {
					t.Logf("%s -> %s", expression, message)
				}

			})

		}
	}

}
func TestIntegerPerformance(t *testing.T) {
	ctx := context.Background()
	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000
		}, {
			"effectiveDate": "01/01/2022",
			"amount": 50000
		}]
	
	}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)
	ctx = context.WithValue(ctx, pencilCalculator.LoggingLevelContextKey{}, zerolog.Disabled)
	var expression string
	var want int64
	for i := 1; i < 100; i++ {
		i1, i2 := getTwoRandomIntegers(1, 100000)
		for j := 1; j < 5; j++ {
			switch j {
			case 1:
				expression = fmt.Sprintf("%d+%d", i1, i2)
				want = i1 + i2
			case 2:
				expression = fmt.Sprintf("%d-%d", i1, i2)
				want = i1 - i2
			case 3:
				expression = fmt.Sprintf("%d*%d", i1, i2)
				want = i1 * i2
			case 4:
				expression = fmt.Sprintf("%d/%d", i1, i2)
				want = i1 / i2
			}
			testname := fmt.Sprintf("Expression: %s", expression)
			t.Run(testname, func(t *testing.T) {
				ans := pencilCalculator.Evaluate(ctx, nil, expression)
				v := ans.PrValue
				if v != want {
					//t.Errorf("got %d, want %d", ans, want)
					t.Fatalf("got %d, want %d", ans, want)
				}
			})
		}
	}

}
