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

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/dwarakauttarkar/gocyclo"
	"github.com/dwarakauttarkar/gocyclo/output"
)

const usageDoc = `Calculate cyclomatic complexities of Go functions.
Usage:
    gocyclo [flags] [values] <Go file or directory> ...

Flags:
    -over N             show functions with complexity > N only and
                        return exit code 1 if the set is non-empty
    -top N              show the top N most complex functions only
    -avg, -avg-short    show the average complexity over all functions;
                        the short option prints the value without a label
    -ignore REGEX       exclude files matching the given regular expression

	-format   			define the output format. Default is json. For csv 
						the -file needs to be specified. if file is not specified 
						then the output file is saved in /tmp/gocyclo-<datetime>.csv.
	-file				define the output file location. Default is /tmp/gocyclo-<datetime>.csv
`

func main() {
	top := flag.Int("top", -1, "show the top N most complex functions only")
	avg := flag.Bool("avg", false, "show the average complexity")
	ignore := flag.String("ignore", "", "exclude files matching the given regular expression")
	formatFlag := flag.String("format", "", "define the output format, valid values are: json (default), csv, tabular")
	file := flag.String("file", "", "define the output file location, valid values are: /path/to/file")

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
		out := output.NewStatOutput(allStats, "", nil, nil, nil)
		out.PrintAverageJSON()
		return
	}

	shownStats := allStats.SortAndFilter(*top)
	out := output.NewStatOutput(shownStats, *formatFlag, file, to.BoolPtr(true), to.IntPtr(10000))
	// if output flag is not set or set to json, print the json output
	out.PrintStats()

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
