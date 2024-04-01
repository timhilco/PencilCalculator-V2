package pencilCalculator

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

var functionMap = map[string]interface{}{

	"Long":                 Long,
	"Max":                  Max,
	"Now":                  Now,
	"RoundToDecimalPlaces": RoundToDecialPlaces_AF,
	"NewFloatIntegerBased": NewFloatIntegerBased,
	"Fibonacci":            Fibonacci,
}

func convertBinaryElementsToFloat(a PencilResult, b PencilResult) (float64, float64) {
	binaryType, _ := determineBinaryOperationType(a, b)
	var leftNumberFLOAT float64 = 0
	var rightNumberFLOAT float64 = 0
	switch binaryType {
	case LeftIntRightInt:
		leftNumberFLOAT = float64(a.PrValue.(int64))
		rightNumberFLOAT = float64(b.PrValue.(int64))
	case LeftIntRightFloat:
		leftNumberFLOAT = float64(a.PrValue.(int64))
		rightNumberFLOAT = b.PrValue.(float64)
	case LeftFloatRightInt:
		rightNumberFLOAT = float64(b.PrValue.(int64))
		leftNumberFLOAT = a.PrValue.(float64)
	case LeftFloatRightFloat:
		leftNumberFLOAT = a.PrValue.(float64)
		rightNumberFLOAT = b.PrValue.(float64)
	case LeftIntRightIntFloat:
		leftNumberFLOAT = float64(a.PrValue.(int64))
		fi := b.PrValue.(FloatIntegerNumber)
		rightInteger := float64(fi.IntegerValue)
		divisor := math.Pow10(int(fi.Precision))
		rightNumberFLOAT = rightInteger / divisor
	case InvalidTypes:
		return 0.0, 0.0
	default:

	}
	return leftNumberFLOAT, rightNumberFLOAT
}
func Max(a, b interface{}) (PencilResult, error) {
	left := a.(PencilResult)
	right := b.(PencilResult)
	leftNumberFLOAT, rightNumberFLOAT := convertBinaryElementsToFloat(left, right)
	var result interface{}
	result = leftNumberFLOAT
	if rightNumberFLOAT > leftNumberFLOAT {
		result = rightNumberFLOAT
	}
	switch result.(type) {
	case int64:
		result = result.(float64)
	default:
		// Leave as float64
	}
	return PencilResult{
		Type:    PencilTypeFloat,
		PrValue: result,
	}, nil
}
func Now() (PencilResult, error) {
	result := time.Now()
	return PencilResult{
		Type:    PencilTypeDateTime,
		PrValue: result,
	}, nil
}
func Fibonacci(a interface{}) (PencilResult, error) {
	n := a.(PencilResult).PrValue.(int64)
	result := fibonacciRecursion(n)
	return PencilResult{
		Type:    PencilTypeDateTime,
		PrValue: result,
	}, nil
}
func fibonacciRecursion(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}
func Long(a interface{}, b interface{}) (PencilResult, error) {
	r := a.(PencilResult).PrValue.(int64)
	i := a.(PencilResult).PrValue.(int64)
	rand.Seed(time.Now().UnixNano())
	ri := int(r)
	n := rand.Intn(ri) + int(i)
	d := time.Duration(n) * time.Second
	time.Sleep(d)
	return PencilResult{
		Type:    PencilTypeInteger,
		PrValue: int64(n),
	}, nil
}
func NewFloatIntegerBased(integerInt, precisionInt interface{}) (PencilResult, error) {
	integer := integerInt.(PencilResult).PrValue.(int64)
	precision := precisionInt.(PencilResult).PrValue.(int64)

	return PencilResult{
		Type: PencilTypeIntegerFloat,
		PrValue: FloatIntegerNumber{
			IntegerValue: int64(integer),
			Precision:    precision,
		},
	}, nil
}
func RoundToDecialPlaces_AF(inputFloat, precisionInt interface{}) (PencilResult, error) {
	var input float64
	inputType := inputFloat.(PencilResult).Type
	switch inputType {
	case PencilTypeFloat:
		input = inputFloat.(PencilResult).PrValue.(float64)
	case PencilTypeIntegerFloat:
		fin := inputFloat.(PencilResult).PrValue.(FloatIntegerNumber)
		input = fin.ConvertFloatIntToFloat64()

	}

	precision := precisionInt.(PencilResult).PrValue.(int64)
	result, _ := RoundToDecialPlaces(input, precision)
	return PencilResult{
		Type:    PencilTypeFloat,
		PrValue: result,
	}, nil
}
func RoundToDecialPlaces(input float64, places int64) (rounded float64, err error) {

	// If the float is not a number
	if math.IsNaN(input) {
		return math.NaN(), errors.New("not a number")
	}

	// Find out the actual sign and correct the input for later
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	// Use the places arg to get the amount of precision wanted
	precision := math.Pow(10, float64(places))

	// Find the decimal place we are looking to round
	digit := input * precision

	// Get the actual decimal number as a fraction to be compared
	_, decimal := math.Modf(digit)

	// If the decimal is less than .5 we round down otherwise up
	if decimal >= 0.5 {
		rounded = math.Ceil(digit)
	} else {
		rounded = math.Floor(digit)
	}

	// Finally we do the math to actually create a rounded number
	return rounded / precision * sign, nil
}
