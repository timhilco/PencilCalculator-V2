package pencilCalculator

import (
	"fmt"
	"math"
)

type PencilType int
type InputDataContextKey struct{}
type LoggingLevelContextKey struct{}
type StatementMapContextKey struct{}
type StatementsMemoContextKey struct{}
type WaitGroupContextKey struct{}
type ProcessingPassContextKey struct{}

const (
	PencilTypeString PencilType = iota
	PencilTypeInteger
	PencilTypeFloat
	PencilTypeIntegerFloat
	PencilTypeBoolean
	PencilTypeNil
	PencilTypeDateTime
	PencilTypeNameCalculator
	//PencilTypeClassName
	PencilTypeLeftParen
	PencilTypeComma
	PencilTypeSlice
	PencilTypeArray
	PencilTypeMap
	PencilTypeStruct
	PencilTypeOperation
	PencilTypeCaseStatement
	PencilTypeCaseItem
	PencilTypeDataAccessor
	PencilTypeDataAccessorElement
	PencilTypeError
)

func convertPencilTypeToString(r PencilResult) string {
	var s string
	switch r.Type {
	case PencilTypeInteger:
		s = "PencilTypeInteger"
	default:
		s = "PencilTypeString"
	}
	return s
}

type BinaryElementType int

const (
	LeftIntRightInt BinaryElementType = iota
	LeftIntRightFloat
	LeftIntRightIntFloat
	LeftFloatRightFloat
	LeftFloatRightInt
	LeftFloatRightIntFloat
	LeftFloatIntRightFloat
	LeftFloatIntRightInt
	LeftFloatIntRightIntFloat
	LeftBooleanRightBoolean
	LeftStringRightString
	InvalidTypes
)

type PencilResult struct {
	Type    PencilType //
	PrValue interface{}
}
type CaseItem struct {
	matchValue  PencilResult
	resultValue PencilResult
}
type FloatIntegerNumber struct {
	IntegerValue int64
	Precision    int64
}

func convertTwoFloatIntNumberToCommonPrecision(a FloatIntegerNumber, b FloatIntegerNumber) (FloatIntegerNumber, FloatIntegerNumber) {
	xPrecision := a.Precision
	yPrecision := b.Precision
	xInteger := a.IntegerValue
	yInteger := b.IntegerValue
	var difference int64 = 0
	var changeX bool = false
	if xPrecision > yPrecision {
		difference = xPrecision - yPrecision
		yPrecision = xPrecision
	} else {
		if xPrecision < yPrecision {
			difference = yPrecision - xPrecision
			xPrecision = yPrecision
			changeX = true
		}
	}
	power := (int64)(math.Pow10(int(difference)))
	if changeX {
		xInteger = xInteger * power
	} else {
		yInteger = yInteger * power

	}
	x := FloatIntegerNumber{
		IntegerValue: xInteger,
		Precision:    xPrecision,
	}
	y := FloatIntegerNumber{
		IntegerValue: yInteger,
		Precision:    yPrecision,
	}
	return x, y
}

func (fin FloatIntegerNumber) Add(input FloatIntegerNumber) FloatIntegerNumber {
	left, right := convertTwoFloatIntNumberToCommonPrecision(fin, input)

	newInteger := (left.IntegerValue + right.IntegerValue)
	return FloatIntegerNumber{
		IntegerValue: newInteger,
		Precision:    left.Precision,
	}
}
func (fin FloatIntegerNumber) Subtract(input FloatIntegerNumber) FloatIntegerNumber {
	left, right := convertTwoFloatIntNumberToCommonPrecision(fin, input)

	newInteger := (left.IntegerValue - right.IntegerValue)
	return FloatIntegerNumber{
		IntegerValue: newInteger,
		Precision:    left.Precision,
	}
}

func (fin FloatIntegerNumber) Multiply(input FloatIntegerNumber) FloatIntegerNumber {
	i := fin.IntegerValue * input.IntegerValue
	p := fin.Precision + input.Precision
	return FloatIntegerNumber{
		IntegerValue: i,
		Precision:    p,
	}
}
func (fin FloatIntegerNumber) Divide(input FloatIntegerNumber) FloatIntegerNumber {
	i := fin.IntegerValue / input.IntegerValue
	p := fin.Precision - input.Precision
	if p < 0 {
		p = p * -1
		f := math.Pow10(int(p))
		i = int64(f)
		p = 0
	}
	return FloatIntegerNumber{
		IntegerValue: i,
		Precision:    p,
	}

}
func (fin FloatIntegerNumber) Equal(i FloatIntegerNumber) bool {
	b := ((fin.IntegerValue == i.IntegerValue) &&
		(fin.Precision == i.Precision))
	return b
}
func (fin FloatIntegerNumber) LessThan(i FloatIntegerNumber) bool {
	//TODO Fix
	b := ((fin.IntegerValue == i.IntegerValue) &&
		(fin.Precision == i.Precision))
	return b
}
func (fin FloatIntegerNumber) GreaterThan(i FloatIntegerNumber) bool {
	//TODO fix
	b := ((fin.IntegerValue == i.IntegerValue) &&
		(fin.Precision == i.Precision))
	return b
}
func (fin FloatIntegerNumber) LessThanEqual(i FloatIntegerNumber) bool {
	//TODO Fix
	b := ((fin.IntegerValue == i.IntegerValue) &&
		(fin.Precision == i.Precision))
	return b
}
func (fin FloatIntegerNumber) GreaterThanEqual(i FloatIntegerNumber) bool {
	//TODO fix
	b := ((fin.IntegerValue == i.IntegerValue) &&
		(fin.Precision == i.Precision))
	return b
}
func (fin FloatIntegerNumber) ConvertFloatIntToFloat6Decimal() float64 {
	f64 := fin.ConvertFloatIntToFloatInputPlaces(6)
	return f64
}
func (fin FloatIntegerNumber) ConvertFloatIntToFloat64() float64 {
	f64 := fin.ConvertFloatIntToFloatInputPlaces(fin.Precision)
	return f64
}
func (fin FloatIntegerNumber) ConvertFloatIntToFloatInputPlaces(precision int64) float64 {
	numerator := float64(fin.IntegerValue)
	divisor := (float64)(math.Pow10(int(fin.Precision)))
	f64 := numerator / float64(divisor)
	return f64
}
func ConvertFloat64ToFloat64with6DecimalPlaces(f float64) float64 {

	s := fmt.Sprintf("%.7f", f)
	_, f64 := ConvertFloatStringToFloatIntegerNumber(s)
	return f64
}
func (fin FloatIntegerNumber) String() string {

	return fmt.Sprintf("%d (%d)", fin.IntegerValue, fin.Precision)
}
func (c CaseItem) String() string {
	t := fmt.Sprintf("Match: %s <> Result: %s", c.matchValue, c.resultValue)
	return t
}

type CaseStatement struct {
	expressionPencilResultValue PencilResult
	caseItems                   []CaseItem
	defaultCaseItem             CaseItem
}

func (p PencilResult) String() string {
	v := fmt.Sprintf("%v", p.PrValue)
	var t string
	switch p.Type {
	case PencilTypeString:
		t = "String"
	case PencilTypeInteger:
		t = "Integer"
	case PencilTypeFloat:
		t = "Float"
	case PencilTypeIntegerFloat:
		t = "InternalFloat (Integer)"
	case PencilTypeBoolean:
		t = "Boolean"
	case PencilTypeDateTime:
		t = "DateTime"
	case PencilTypeCaseItem:
		t = "CaseItem"
		v = p.PrValue.(CaseItem).String()
	case PencilTypeNameCalculator:
		t = "NameCalculator"
	case PencilTypeComma:
		t = "Comma"
	case PencilTypeLeftParen:
		t = "LeftParen"
	case PencilTypeOperation:
		t = "Operation"
	case PencilTypeDataAccessor:
		t = "DataAccessor"
	case PencilTypeDataAccessorElement:
		t = "DataAccessorElement"
	case PencilTypeNil:
		t = "Nil"

	default:
		t = "Unknown"
	}
	return t + " --> " + v
}
func determineBinaryOperationType(left PencilResult, right PencilResult) (BinaryElementType, string) {
	leftType := left.Type
	rightType := right.Type

	switch leftType {
	case PencilTypeInteger:
		switch rightType {
		case PencilTypeInteger:
			return LeftIntRightInt, "INT"
		case PencilTypeString, PencilTypeBoolean:
			return InvalidTypes, "INT"
		case PencilTypeFloat:
			return LeftIntRightFloat, "FLOAT"
		case PencilTypeIntegerFloat:
			return LeftIntRightIntFloat, "INT_FLOAT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeString:
		switch rightType {
		case PencilTypeString:
			return LeftStringRightString, "STRING"
		case PencilTypeInteger, PencilTypeFloat, PencilTypeIntegerFloat, PencilTypeBoolean:
			return InvalidTypes, "INT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeFloat:
		switch rightType {
		case PencilTypeFloat:
			return LeftFloatRightFloat, "FLOAT"
		case PencilTypeInteger:
			return LeftFloatRightInt, "FLOAT"
		case PencilTypeString, PencilTypeBoolean:
			return InvalidTypes, "INT"
		case PencilTypeIntegerFloat:
			return LeftFloatRightIntFloat, "FLOAT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeIntegerFloat:
		switch rightType {
		case PencilTypeFloat:
			return LeftFloatIntRightIntFloat, "INT"
		case PencilTypeInteger:
			return LeftFloatIntRightInt, "INT"
		case PencilTypeString, PencilTypeBoolean:
			return InvalidTypes, "INT"
		case PencilTypeIntegerFloat:
			return LeftFloatIntRightIntFloat, "INT_FLOAT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeBoolean:
		switch rightType {
		case PencilTypeBoolean:
			return LeftBooleanRightBoolean, "BOOLEAN"
		case PencilTypeInteger, PencilTypeFloat, PencilTypeIntegerFloat, PencilTypeString:
			return InvalidTypes, "INT"
		default:
			return InvalidTypes, "INT"
		}

	default:
		return InvalidTypes, "INT"
	}

}
