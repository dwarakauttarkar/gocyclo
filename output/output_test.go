package output

import (
	"testing"

	"github.com/dwarakauttarkar/gocyclo"
)

func TestPrintStatsTabular(t *testing.T) {
	var stat gocyclo.Stats
	stat = append(stat, gocyclo.Stat{PkgName: "test", FuncName: "test", CyclomaticComplexity: 2, MaintainabilityIndex: 6})
	stat = append(stat, gocyclo.Stat{PkgName: "test2", FuncName: "test2", CyclomaticComplexity: 4, MaintainabilityIndex: 7})
	stat = append(stat, gocyclo.Stat{PkgName: "test3", FuncName: "test3", CyclomaticComplexity: 5, MaintainabilityIndex: 8})
	o := NewStatOutput(stat, "tabular", nil, nil, nil)
	o.PrintStats()
}

func TestPrintStatsJSON(t *testing.T) {
	var stat gocyclo.Stats
	stat = append(stat, gocyclo.Stat{PkgName: "test", FuncName: "test", CyclomaticComplexity: 2, MaintainabilityIndex: 6})
	stat = append(stat, gocyclo.Stat{PkgName: "test2", FuncName: "test2", CyclomaticComplexity: 4, MaintainabilityIndex: 7})
	stat = append(stat, gocyclo.Stat{PkgName: "test3", FuncName: "test3", CyclomaticComplexity: 5, MaintainabilityIndex: 8})
	o := NewStatOutput(stat, "json", nil, nil, nil)
	o.PrintStats()
}

func TestPrintStatsJSONPretty(t *testing.T) {
	var stat gocyclo.Stats
	stat = append(stat, gocyclo.Stat{PkgName: "test", FuncName: "test", CyclomaticComplexity: 2, MaintainabilityIndex: 6})
	stat = append(stat, gocyclo.Stat{PkgName: "test2", FuncName: "test2", CyclomaticComplexity: 4, MaintainabilityIndex: 7})
	stat = append(stat, gocyclo.Stat{PkgName: "test3", FuncName: "test3", CyclomaticComplexity: 5, MaintainabilityIndex: 8})
	o := NewStatOutput(stat, "jsonpretty", nil, nil, nil)
	o.PrintStats()
}
