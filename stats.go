// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"fmt"
	"go/token"
	"sort"
)

// Stat holds the cyclomatic complexity of a function, along with its package
// and and function name and its position in the source code.
type Stat struct {
	PkgName              string
	FuncName             string
	CyclomaticComplexity int
	MaintainabilityIndex int
	Pos                  token.Position
}

// String formats the cyclomatic complexity information of a function in
// the following format: "<complexity> <package> <function> <file:line:column>"
func (s Stat) String() string {
	return fmt.Sprintf("%d %s %s %s", s.CyclomaticComplexity, s.PkgName, s.FuncName, s.Pos)
}

// Stats hold the cyclomatic complexities of many functions.
type Stats []Stat

// AverageComplexity calculates the average cyclomatic complexity of the
// cyclomatic complexities in s.
func (s Stats) AverageComplexity() float64 {
	return float64(s.TotalCyclomaticComplexity()) / float64(len(s))
}

// TotalCyclomaticComplexity calculates the total sum of all cyclomatic
// complexities in s.
func (s Stats) TotalCyclomaticComplexity() uint64 {
	total := uint64(0)
	for _, stat := range s {
		total += uint64(stat.CyclomaticComplexity)
	}
	return total
}

// SortAndFilter sorts the cyclomatic complexities in s in descending order
// and returns a slice of s limited to the 'top' N entries with a cyclomatic
// complexity greater than 'over'. If 'top' is negative, i.e. -1, it does
// not limit the result. If 'over' is <= 0 it does not limit the result either,
// because a function has a base cyclomatic complexity of at least 1.
func (s Stats) SortAndFilter(top int) Stats {
	result := make(Stats, len(s))
	copy(result, s)
	sort.Stable(byComplexityDesc(result))
	for i := range result {
		if i == top {
			return result[:i]
		}
	}
	return result
}

type byComplexityDesc Stats

func (s byComplexityDesc) Len() int      { return len(s) }
func (s byComplexityDesc) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byComplexityDesc) Less(i, j int) bool {
	return s[j].MaintainabilityIndex >= s[i].MaintainabilityIndex
}
