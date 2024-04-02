package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-PencilCalculator/pencilCalculator"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/hlog"
)

func (s *ServerEnvironment) calculatorView(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	//var request CalculatorRequest

	expression := "1+2"
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("GET::CalculatorView Request: " + expression)
	//update response
	pr := evaluateAndRespond(s, expression, w, r)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("GET::CalculatorView Response: " + pr.String())

}
func evaluateAndRespond(s *ServerEnvironment, expression string, w http.ResponseWriter, r *http.Request) pencilCalculator.PencilResult {
	var response CalculatorResponse
	ctx := context.Background()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("../api/calculator.html")
	if err != nil {
		msg := fmt.Sprintf("%s", err)
		hlog.FromRequest(r).Fatal().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("Error: HTML Calculator Page Loading: " + msg)
	}
	pencilResult := evaluate(s, ctx, true, expression, r)
	response = CalculatorResponse{
		Type:       fmt.Sprintf("%v", pencilResult.Type),
		Result:     fmt.Sprintf("%v", pencilResult.PrValue),
		Expression: expression,
	}
	t.Execute(w, response)

	return pencilResult
}

func (s *ServerEnvironment) webEvaluate(w http.ResponseWriter, r *http.Request) {

	//declare response variable
	//var request CalculatorRequest
	//expression := "2+3"

	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s=%s\n", key, value)
	}
	expression := r.FormValue("expression")
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("POST::webEvaluate --> " + expression)
	pr := evaluateAndRespond(s, expression, w, r)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("Exit POST::webEvaluate --> " + pr.String())
}
func evaluate(s *ServerEnvironment, ctx context.Context, isLocal bool, expression string, r *http.Request) pencilCalculator.PencilResult {
	var pencilResult pencilCalculator.PencilResult
	if s.LocalEvaluate {
		pencilResult = evaluateLocal(ctx, expression, r)
	} else {
		pencilResult = evaluateRestApi(s, ctx, expression, r)

	}
	return pencilResult
}
func evaluateLocal(ctx context.Context, expression string, r *http.Request) pencilCalculator.PencilResult {
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("Enter POST::webEvaluate <> evaluateLocal --> " + expression)
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
	statementMap := make(map[string]*pencilCalculator.Statement)
	ctx = context.WithValue(ctx, pencilCalculator.StatementMapContextKey{}, statementMap)
	pencilResult := pencilCalculator.Evaluate(ctx, nil, expression)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("Exit POST::webEvaluate <> evaluateLocal --> " + pencilResult.String())
	return pencilResult

}

func evaluateRestApi(s *ServerEnvironment, ctx context.Context, expression string, r *http.Request) pencilCalculator.PencilResult {
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("Enter POST::webEvaluate <> evaluateRestApi --> " + expression)
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"expression": expression,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	url := s.ApiServerHostName + " api/evaluate"
	resp, err := http.Post(url, "application/json", responseBody)
	//Handle Error
	if err != nil {
		em := fmt.Sprintf("An Error Occurred after Post call: %v", err)
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("Enter POST::webEvaluate <> evaluateRestApi --> " + em)
	}
	defer resp.Body.Close()
	//Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		em := fmt.Sprintf("An Error Occurred after ReadAll:%v", err)
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("Enter POST::webEvaluate <> evaluateRestApi --> " + em)
	}
	var cr CalculatorResponse
	json.Unmarshal(body, &cr)

	if err != nil {
		em := fmt.Sprintf("An Error Occurred after Unmarshall:%v", err)
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("Enter POST::webEvaluate <> evaluateRestApi --> " + em)
	}
	pr := pencilCalculator.PencilResult{

		Type:    pencilCalculator.PencilTypeInteger,
		PrValue: fmt.Sprintf("%v", cr.Result),
	}
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("Exit POST::webEvaluate <> evaluateRestApi --> " + expression)
	return pr
}
