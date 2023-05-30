# gocyclo

#### A Cyclomatic complexity & maintainability index library for go language

#### This repository forked from <https://github.com/fzipp/gocyclo> and is enhanced with additional features

[![PkgGoDev](https://pkg.go.dev/badge/github.com/fzipp/gocyclo)](https://pkg.go.dev/github.com/fzipp/gocyclo)
![Build Status](https://github.com/fzipp/gocyclo/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/fzipp/gocyclo)](https://goreportcard.com/report/github.com/fzipp/gocyclo)

Gocyclo calculates [cyclomatic complexities](https://en.wikipedia.org/wiki/Cyclomatic_complexity) & [maintaninability index](https://learn.microsoft.com/en-us/visualstudio/code-quality/code-metrics-maintainability-index-range-and-meaning?view=vs-2022) of functions in Go source code.

**Cyclomatic complexity** is a software metric used to measure the complexity of a program's source code. It provides a quantitative measure of the number of independent paths through a program's control flow. In other words, it measures how many different possible paths the program can take based on different conditions and decisions.

A higher cyclomatic complexity value indicates that a program has more complex control flow and is likely to be more difficult to understand, test, and maintain. This metric helps software developers identify code that may be more prone to errors, as complex code paths can increase the likelihood of bugs.

A function with a higher cyclomatic complexity requires more test cases to
cover all possible paths and is potentially harder to understand. The
complexity can be reduced by applying common refactoring techniques that lead to smaller functions.

**The maintainability index** is a software metric used to assess the maintainability of a software system's source code. It provides a numerical rating that indicates how easy or difficult it is to maintain and modify the codebase. The maintainability index takes into account various factors such as code complexity, size, and code structure.

The maintainability index is typically calculated using a formula that incorporates metrics such as cyclomatic complexity, lines of code, and code duplication. The formula varies depending on the tool or methodology used, but the resulting index is usually represented as a score between 0 and 100, with higher scores indicating better maintainability.

A higher maintainability index implies that the codebase is easier to understand, modify, and enhance. It suggests that the code has clear structure, low complexity, and is well-documented. On the other hand, a lower maintainability index indicates that the code may be more convoluted, harder to comprehend, and prone to errors.

```
Maintainability Index = MAX(0, (171 - 5.2 * ln(Halstead Volume) - 0.23 * (Cyclomatic Complexity) - 16.2 * ln(LOC)) * 100 / 171)
```

## Installation

To install the `gocyclo` command, run

``` markdown
go install github.com/dwarakauttarkar/gocyclo/cmd/gocyclo@latest
```

Place the resulting binary in one of your PATH directories if
`$GOPATH/bin` isn't already in your PATH.

## Usage

``` markdown
Calculate cyclomatic complexities of Go functions.
Usage:
    gocyclo [flags] <Go file or directory> ...

Flags:
    -over N             show functions with complexity > N only and
                        return exit code 1 if the set is non-empty
    
    -top N              show the top N most complex functions only
    
    -avg, -avg-short    show the average complexity over all functions;
    
    -format             define the output format. Default is json. For csv 
                        the -file needs to be specified. if file is not specified 
                        then the output file is saved in /tmp/gocyclo-<datetime>.csv.
                        the short option prints the value without a label.
                        supported formats: tabular, json, csv(file output)
    
    -file               define the output file location. Default is /tmp/gocyclo-<datetime>.csv                          
    
    -ignore REGEX       exclude files matching the given regular expression

The output fields for each line are:
<complexity> <package> <function> <file:line:column>
```

## Examples

```
gocyclo .
gocyclo main.go
gocyclo -top 10 src/
gocyclo -avg .
gocyclo -top 20 -ignore "_test|Godeps|vendor/" .
gocyclo -over 3 -avg gocyclo/
gocyclo -format <json/tabular/csv> graph.go
gocyclo -format csv -file /<file_path>/code_analysis.csv graph.go
```

##### Tabular Output

`gocyclo -format tabular graph.go`
```
PackageName  FunctionName       CyclomaticComplexity  MaintainabilityIndex
-----------  ------------       --------------------  --------------------
graph        (*Graph).Dijkstra  8                     41
graph        (Item).More        1                     74
graph        (Item).Idx         1                     78

```
##### JSON Format

`gocyclo -format json graph.go`
```

[
  {
    "PkgName": "graph",
    "FuncName": "(*Graph).Dijkstra",
    "CyclomaticComplexity": 8,
    "MaintainabilityIndex": 41,
    "Pos": {
      "Filename": "graph.go",
      "Offset": 372,
      "Line": 24,
      "Column": 1
    }
  },
  {
    "PkgName": "graph",
    "FuncName": "(Item).More",
    "CyclomaticComplexity": 1,
    "MaintainabilityIndex": 74,
    "Pos": {
      "Filename": "graph.go",
      "Offset": 228,
      "Line": 16,
      "Column": 1
    }
  },
  {
    "PkgName": "graph",
    "FuncName": "(Item).Idx",
    "CyclomaticComplexity": 1,
    "MaintainabilityIndex": 78,
    "Pos": {
      "Filename": "graph.go",
      "Offset": 328,
      "Line": 20,
      "Column": 1
    }
  }
]

```

### Ignoring individual functions

Individual functions can be ignored with a `gocyclo:ignore` directive:

```
//gocyclo:ignore
func f1() {
 // ...
}
    
//gocyclo:ignore
var f2 = func() {
 // ...
}
```

## License

This project is free and open source software licensed under the
[BSD 3-Clause License](LICENSE).
