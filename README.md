[![Build Status](https://travis-ci.org/brehv/r.svg?branch=master)](https://travis-ci.org/brehv/r)

## Description

r is a tool that allows you to replace a single field in a struct n-levels deep with a new value.

The idea is to use r in testing scenarios where a function updates a single field, and you don't want to pollute your code with giant struct declarations; like in config tests.

## Installation

r is available using the standard `go get` command.

Install by running:

    go get github.com/brehv/r

Run tests by running:

    go test github.com/brehv/r

## Usage
```go
r.R(subj, "Field.Name.Whatever", "NewValue") // Replaces "Field.Name.Whatever" with "NewValue" and returns a copy of subj 
r.G(subj, "Field.Name.Whatever") // Returns "Field.Name.Whatever" of subj
```

See [example_test.go](./example_test.go) or [r_test.go](./r_test.go) to see more examples