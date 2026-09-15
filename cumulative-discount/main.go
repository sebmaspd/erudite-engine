package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	zen "github.com/gorules/zen-go/v2"
)

//go:embed cumulative-discount.json
var cumulativeDiscountJDM []byte

func main() {
	volume := flag.Float64("volume", 0, "Purchase Volume (units)")
	flag.Parse()

	engine := zen.NewEngine(zen.EngineConfig{})
	defer engine.Dispose()

	decision, err := engine.CreateDecision(cumulativeDiscountJDM)
	if err != nil {
		log.Fatalf("failed to load decision: %v", err)
	}
	defer decision.Dispose()

	response, err := decision.Evaluate(map[string]any{
		"purchaseVolume": *volume,
	})
	if err != nil {
		log.Fatalf("failed to evaluate decision: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(response.Result, &result); err != nil {
		log.Fatalf("failed to parse result: %v", err)
	}

	fmt.Printf("Unit-Price: $%.2f\n", result["unitPrice"])
	fmt.Printf("Discount: %.0f%%\n", result["discount"].(float64)*100)
	fmt.Printf("Net-Unit-Price: $%.2f\n", result["netUnitPrice"])
	fmt.Printf("Volume-Discounted-Price: $%.2f\n", result["volumeDiscountedPrice"])
}
