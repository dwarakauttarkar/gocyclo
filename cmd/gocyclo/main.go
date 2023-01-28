// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gocyclo calculates the cyclomatic complexities of functions and
// methods in Go source code.
//
// Usage:
//
//	gocyclo [<flag> ...] <Go file or directory> ...
//
// Flags:
//
//	-over N               show functions with complexity > N only and
//	                      return exit code 1 if the output is non-empty
//	-top N                show the top N most complex functions only
//	-avg, -avg-short      show the average complexity;
//	                      the short option prints the value without a label
//	-ignore REGEX         exclude files matching the given regular expression
//
// The output fields for each line are:
// <complexity> <package> <function> <file:line:column>
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/dwarakauttarkar/gocyclo"
	"github.com/dwarakauttarkar/gocyclo/output"
)

const usageDoc = `Calculate cyclomatic complexities of Go functions.
Usage:
    gocyclo [flags] <Go file or directory> ...

Flags:
    -over N               show functions with complexity > N only and
                          return exit code 1 if the set is non-empty
    -top N                show the top N most complex functions only
    -avg, -avg-short      show the average complexity over all functions;
                          the short option prints the value without a label
    -ignore REGEX         exclude files matching the given regular expression

The output fields for each line are:
<complexity> <package> <function> <file:line:column>
`

func main() {
	top := flag.Int("top", -1, "show the top N most complex functions only")
	avg := flag.Bool("avg", false, "show the average complexity")
	ignore := flag.String("ignore", "", "exclude files matching the given regular expression")

	log.SetFlags(0)
	log.SetPrefix("gocyclo: ")
	flag.Usage = usage
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		usage()
		return
	}

	allStats, err := gocyclo.Analyze(paths, regex(*ignore))
	if err != nil {
		log.Fatal("error while analyzing the file", err)
		return
	}
	// If the -avg flag is set, we only want to print the average
	if *avg {
		output.PrintAverageJSON(allStats)
		return
	}

	shownStats := allStats.SortAndFilter(*top)

	output.PrintStatsJSON(shownStats)
	return
}

func regex(expr string) *regexp.Regexp {
	if expr == "" {
		return nil
	}
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}
	return re
}

func usage() {
	_, _ = fmt.Fprint(os.Stderr, usageDoc)
	os.Exit(2)
}
