package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/dwarakauttarkar/gocyclo"
	"github.com/dwarakauttarkar/gocyclo/entities"
)

type StatOutput struct {
	Stats        gocyclo.Stats
	formatFlag   string
	fileLocation *string
	sortOutput   *bool
	printTop     *int
}

// NewOutput creates a new Output object.
// DEFAULT: If sortOutput is nil, it will be set to true.
// DEFAULT: If printTop is nil, it will be set to 1000.
func NewStatOutput(stats gocyclo.Stats, formatFlag string, fileLocation *string, sortOutput *bool, printTop *int) *StatOutput {
	if sortOutput == nil {
		sortOutput = to.BoolPtr(true)
	}
	if printTop == nil {
		printTop = to.IntPtr(1000)
	}
	return &StatOutput{
		Stats:        stats,
		formatFlag:   formatFlag,
		fileLocation: fileLocation,
		sortOutput:   sortOutput,
		printTop:     printTop,
	}
}

// PrintStats prints the given stats in the given format.
// by default it prints in JSON format.
func (o *StatOutput) PrintStats() error {
	if *o.sortOutput {
		o.Stats = o.Stats.SortAndFilter(*o.printTop)
	}

	switch o.formatFlag {
	case "json":
		return o.printStatsJSONPretty()
	case "csv":
		return o.printStatsCSV()
	case "tabular":
		return o.printStatsTabular()
	default:
		return o.printStatsJSON()
	}
}

// PrintStatsJSON prints the given stats in JSON format.
func (o *StatOutput) printStatsJSON() error {
	jsonBytes, err := json.Marshal(o.Stats)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
	return nil
}

// PrintAverageJSON prints the average complexity of the given stats in JSON format.
func (o *StatOutput) PrintAverageJSON() {
	avgMap := map[string]string{
		"average": fmt.Sprintf("%.3g", o.Stats.AverageComplexity()),
	}
	jsonBytes, _ := json.Marshal(avgMap)
	fmt.Println(string(jsonBytes))
}

// PrintStatsJSONPretty prints the given stats in JSON format with indentation.
func (o *StatOutput) printStatsJSONPretty() error {
	jsonBytes, err := json.MarshalIndent(o.Stats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
	return nil
}

// printStatsCSV prints the given stats in CSV format to a file.
func (o *StatOutput) printStatsCSV() error {
	fileLocation := fmt.Sprintf("/tmp/gocyclo-%s.csv", time.Now().String())
	if o.fileLocation != nil && *o.fileLocation != "" {
		fileLocation = *o.fileLocation
	}
	if !strings.HasSuffix(fileLocation, ".csv") {
		return fmt.Errorf("file location must end with <file_name>.csv")
	}
	file, err := os.Create(fileLocation)
	if err != nil {
		fmt.Println("Error while creating csv file: ", err, " file location: ", fileLocation)
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{
		entities.PackageNameLabel,
		entities.FunctionNameLabel,
		entities.CyclomaticComplexityLabel,
		entities.MaintainabilityIndexLabel,
	})
	for _, stat := range o.Stats {
		writer.Write([]string{stat.PkgName, stat.FuncName, fmt.Sprintf("%d", stat.CyclomaticComplexity), fmt.Sprintf("%d", stat.MaintainabilityIndex)})
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error while writing to csv file: ", err)
		return err
	}
	fmt.Println("CSV file written to: ", fileLocation)
	return nil
}

func (o *StatOutput) printStatsTabular() error {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	tabWriter.Write([]byte(fmt.Sprintf("%s\t%s\t%s\t%s\n", entities.PackageNameLabel, entities.FunctionNameLabel, entities.CyclomaticComplexityLabel, entities.MaintainabilityIndexLabel)))
	tabWriter.Write([]byte("-----------\t------------\t--------------------\t--------------------\n"))
	for _, stat := range o.Stats {
		tabWriter.Write([]byte(fmt.Sprintf("%s\t%s\t%d\t%d\n", stat.PkgName, stat.FuncName, stat.CyclomaticComplexity, stat.MaintainabilityIndex)))
	}
	tabWriter.Flush()
	return nil
}
