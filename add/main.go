package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	zen "github.com/gorules/zen-go/v2"
)

//go:embed add.json
var addJDM []byte

func main() {
	a := flag.Float64("a", 0, "value of A")
	b := flag.Float64("b", 0, "value of B")
	flag.Parse()

	engine := zen.NewEngine(zen.EngineConfig{})
	defer engine.Dispose()

	decision, err := engine.CreateDecision(addJDM)
	if err != nil {
		log.Fatalf("failed to load decision: %v", err)
	}
	defer decision.Dispose()

	response, err := decision.Evaluate(map[string]any{
		"a": *a,
		"b": *b,
	})
	if err != nil {
		log.Fatalf("failed to evaluate decision: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(response.Result, &result); err != nil {
		log.Fatalf("failed to parse result: %v", err)
	}

	fmt.Printf("sum = %v\n", result["sum"])
}
