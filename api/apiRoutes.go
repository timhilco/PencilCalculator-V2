package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-PencilCalculator/pencilCalculator"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

type CalculatorRequest struct {
	Expression string `json:"expression"`
}
type CalculatorResponse struct {
	Type       string `json:"type"`
	Result     string `json:"result"`
	ErrorCode  int    `json:"errorCode"`
	Error      string `json:"error"`
	Expression string `json:"expression"`
}

func (s *ServerEnvironment) apiEvaluate_POST(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var request CalculatorRequest
	var response CalculatorResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("POST::evaluate --> JSON unmarshaling failed")

	}
	rs := fmt.Sprintf("%v", request)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("POST::evaluate request--> JSON:" + rs)
	response = CalculatorResponse{
		Type:      "0",
		Result:    "aResult",
		ErrorCode: 0,
		Error:     "",
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
	pencilResult := pencilCalculator.Evaluate(ctx, nil, request.Expression)
	response.Type = fmt.Sprintf("%d", pencilResult.Type)
	response.Result = fmt.Sprintf("%v", pencilResult.PrValue)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	rs = fmt.Sprintf("%v", response)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("POST::evaluate --> JSON:" + rs)

}
