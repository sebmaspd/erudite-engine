package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	zen "github.com/gorules/zen-go/v2"
)

//go:embed haze-hourly.json
var hazeHourlyJDM []byte

func main() {
	pm25 := flag.Float64("pm25", 0, "1-hr PM2.5 concentration (µg/m3)")
	flag.Parse()

	engine := zen.NewEngine(zen.EngineConfig{})
	defer engine.Dispose()

	decision, err := engine.CreateDecision(hazeHourlyJDM)
	if err != nil {
		log.Fatalf("failed to load decision: %v", err)
	}
	defer decision.Dispose()

	response, err := decision.Evaluate(map[string]any{
		"pm25": *pm25,
	})
	if err != nil {
		log.Fatalf("failed to evaluate decision: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(response.Result, &result); err != nil {
		log.Fatalf("failed to parse result: %v", err)
	}

	fmt.Printf("hazeBand = %v\n", result["hazeBand"])
}
