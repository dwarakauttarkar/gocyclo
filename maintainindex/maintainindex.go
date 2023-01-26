package maintainindex

import (
	"go/ast"
	"go/token"
	"math"

	"github.com/dwarakauttarkar/gocyclo/cyclomaticindex"
	"github.com/dwarakauttarkar/gocyclo/halstvol"
	"github.com/dwarakauttarkar/gocyclo/loc"
)

func Calculate(fn ast.Node, fs *token.FileSet) int {
	v := MaintainIndexVisitor{
		MaintIdx: 0,
		fileSet:  fs,
	}
	ast.Walk(&v, fn)
	return v.MaintIdx
}

type MaintainIndexVisitor struct {
	MaintIdx             int
	HalstVol             float64
	Loc                  int
	CyclomaticComplexity int
	fileSet              *token.FileSet
}

func (m *MaintainIndexVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl:
		m.HalstVol = halstvol.Calculate(n)
		m.CyclomaticComplexity = cyclomaticindex.Calculate(n)
		m.Loc = loc.Calculate(m.fileSet, n)
		m.MaintIdx = m.calc()
	}
	return m
}

// Calc https://docs.microsoft.com/ja-jp/archive/blogs/codeanalysis/maintainability-index-range-and-meaning
func (v *MaintainIndexVisitor) calc() int {
	origVal := 171.0 - 5.2*math.Log(v.HalstVol) - 0.23*float64(v.CyclomaticComplexity) - 16.2*math.Log(float64(v.Loc))
	normVal := int(math.Max(0.0, origVal*100.0/171.0))
	return normVal
}
