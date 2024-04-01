package tests

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var variableNames []string

func BuildProgram(numberOfStatements int, programName string) string {
	variableNames = make([]string, 0)
	var sb strings.Builder
	var statement string
	for i := 1; i < numberOfStatements; i++ {
		statement = buildStatement(i)
		sb.WriteString(statement)
	}
	s := sb.String()
	b := []byte(s)
	fileName := fmt.Sprintf("./programs/%s.pp", programName)
	err := os.WriteFile(fileName, b, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return s
}
func buildStatement(variableNumber int) string {
	var expression string = "20"
	variable := fmt.Sprintf("W1:a%d", variableNumber)
	l := len(variableNames)
	j := getRandomInteger(l)
	var randomVariable string
	if j == 0 {
		randomVariable = "0"
	} else if j > 5 {
		randomVariable = variableNames[0]
	} else {
		randomVariable = variableNames[j]

	}
	i := getRandomInteger(4)
	switch i {
	case 0:
		expression = fmt.Sprintf("@Long (%d, %d)", 5, 5)
	case 1:
		expression = randomVariable
	case 2:
		ss := buildRecursiveStatements(1, variableNumber, variable, "")
		return ss
	case 3:
		expression = buildAddExpression(variableNumber)
	default:
		expression = fmt.Sprintf("%d", variableNumber)
	}
	statement := fmt.Sprintf("%s := %s;\n", variable, expression)
	variableNames = append(variableNames, variable)
	return statement
}
func buildAddExpression(c int) string {
	expression := "1"
	for i := 0; i < c; i++ {
		v := getRandomVariableName()
		expression = expression + " + " + v

	}

	return expression

}
func getRandomVariableName() string {
	l := len(variableNames)
	variable2 := "WE:a1"
	if l > 0 {
		j := getRandomInteger(l)
		variable2 = variableNames[j%10]
	}
	return variable2
}
func buildRecursiveStatements(level int, max int, variableName string, currentStatement string) string {
	var s string

	var t string
	if level == max {
		variable := fmt.Sprintf("%s%d", variableName, level)
		l := len(variableNames)
		variable2 := strconv.Itoa(max)
		if l > 0 {
			j := getRandomInteger(l)
			variable2 = variableNames[j]
		}
		t = fmt.Sprintf("%s := %s;\n", variable, variable2)
		//variableNames = append(variableNames, variable)
		return t
	}

	variable := fmt.Sprintf("%s%d", variableName, level)
	//variableNames = append(variableNames, variable)
	t = fmt.Sprintf("%s := %s%d;\n", variable, variableName, level+1)
	t2 := buildRecursiveStatements(level+1, max, variableName, currentStatement)
	s = currentStatement + t + t2
	return s
}
func getRandomInteger(max int) int {
	var i1 int
	rand.Seed(time.Now().UnixNano())
	if max == 0 {
		return 0
	}
	//i1 = int64(rand.Intn(max-min+1) + min
	i1 = rand.Intn(max)

	return i1
}

/*
func buildRandomProgram() string {
	variables := make([]string, 0)
	var sb strings.Builder
	for i := 0; i < 16; i++ {
		variable := fmt.Sprintf("WE:e%d", i)
		variables = append(variables, variable)
		sb.WriteString(buildSimpleElementalStatement(i))
	}
	for level := 0; level < 50; level++ {
		variable := fmt.Sprintf("W%d:a%d", level, level)
		elements := make([]string, 0)
		for i := 0; i < level; i++ {
			count := len(variables)
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(count)
			field := variables[n]
			elements = append(elements, field)
		}
		sb.WriteString(buildStatement(variable, elements))
		variables = append(variables, variable)
	}
	s := sb.String()
	b := []byte(s)
	err := os.WriteFile("./data/data.txt", b, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return s

}
func buildProgram1() string {
	var sb strings.Builder
	sb.WriteString("W1:e1 :=  20 ; ")
	sb.WriteString("W1:e2 :=  10 ; ")
	sb.WriteString("W1:e3 :=  @Long(5, 5) ; ")
	sb.WriteString("W1:c :=  W1:e1 + W1:e2 ; ")
	sb.WriteString("W1:d :=  W1:e1 - W1:e2  ; ")
	sb.WriteString("W1:e :=  W1:e1 * W1:e2  ; ")
	sb.WriteString("W1:f :=  W1:e1 -  W1:e2 ; ")
	sb.WriteString("W1:g :=  W1:e1 + W1:e2 + W1:e3 ; ")
	sb.WriteString("W1:h :=  W1:e1 - W1:e2 + W1:e3 ; ")
	sb.WriteString("W1:i :=  W1:e1 * W1:e2 + W1:e3 ; ")
	sb.WriteString("W1:j :=  W1:e1 -  W1:e2 + W1:e3 ; ")

	return sb.String()
}
func buildElementalStatement(number int) string {
	var sb strings.Builder
	s := fmt.Sprintf("WE:e%d", number)
	sb.WriteString(s)
	sb.WriteString(" := ")
	value := ""
	if number == 0 {
		value = fmt.Sprintf("@Long (%d, %d)", 5, 5)

	} else if number%2 == 1 {
		value = fmt.Sprintf("%d", number)
	} else {
		value = fmt.Sprintf("@Long (%d, %d)", number*2, number)

	}
	sb.WriteString(value)
	sb.WriteString(" ;\n")
	return sb.String()
}
func buildSimpleElementalStatement(number int) string {
	var sb strings.Builder
	s := fmt.Sprintf("WE:e%d", number)
	sb.WriteString(s)
	sb.WriteString(" := ")
	value := ""
	if number == 0 {
		value = fmt.Sprintf("%d + %d", 60, 10)

	} else if number%2 == 1 {
		value = fmt.Sprintf("%d", number)
	} else {
		value = fmt.Sprintf("%d * %d", number*2, number)

	}
	sb.WriteString(value)
	sb.WriteString(" ;\n")
	return sb.String()
}
func buildStatement(variable string, components []string) string {
	var sb strings.Builder
	sb.WriteString(variable)
	sb.WriteString(" := ")
	for _, element := range components {
		sb.WriteString(element + " + ")
	}
	sb.WriteString("WE:e0;\n")
	return sb.String()
}
*/
