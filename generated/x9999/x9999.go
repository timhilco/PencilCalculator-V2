package x9999

import (
	"context"

	pc "go-PencilCalculator/pencilCalculator"
)

func RunProgram(ctx context.Context, inputJson []byte) (map[string]pc.PencilResult, error) {
	results := make(map[string]pc.PencilResult)
	// Statements start here
	prValue_W1a1 := W1_a1(results)
	results["W1:a1"] = pc.PencilResult{
		Type:    pc.PencilTypeInteger,
		PrValue: prValue_W1a1,
	}
	prValue_W1a2 := W1_a2(results)
	results["W1:a2"] = pc.PencilResult{
		Type:    pc.PencilTypeInteger,
		PrValue: prValue_W1a2,
	}

	prValue_W1a3 := W1_a3(results)
	results["W1:a3"] = pc.PencilResult{
		Type:    pc.PencilTypeInteger,
		PrValue: prValue_W1a3,
	}
	// Statements end here
	return results, nil

}
func W1_a1(results map[string]pc.PencilResult) int64 {
	r := results["W1:a1"]
	if r.PrValue != nil {
		return r.PrValue.(int64)
	}
	var v int64
	v = 100
	return v
}
func W1_a2(results map[string]pc.PencilResult) int64 {
	r := results["W1:a2"]
	if r.PrValue != nil {
		return r.PrValue.(int64)
	}
	var x int64
	x = 200
	return x
}
func W1_a3(results map[string]pc.PencilResult) int64 {
	r := results["W1:a3"]
	if r.PrValue != nil {
		return r.PrValue.(int64)
	}
	x := results["W1:a1"].PrValue.(int64)
	y := results["W1:a2"].PrValue.(int64)
	z := x + y
	return z
}
