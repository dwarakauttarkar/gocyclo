package output

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dwarakauttarkar/gocyclo"
)

// PrintStatsJSON prints the given stats in JSON format.
func PrintStatsJSON(s gocyclo.Stats) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}

// PrintAverageJSON prints the average complexity of the given stats in JSON format.
func PrintAverageJSON(s gocyclo.Stats) {
	avgMap := map[string]string{
		"average": fmt.Sprintf("%.3g", s.AverageComplexity()),
	}
	jsonBytes, _ := json.Marshal(avgMap)
	fmt.Println(string(jsonBytes))
}
