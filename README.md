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
type Person struct {
	Name string
	Parents Parents
}

type Parents struct {
	Mom string
	Dad string
}

func main() {
	p := Person{
		Name:   "Millhouse",
		Parents: Parents{
			Mom: "Cherry",
			Dad: "Dennis",
		},
	}

	fmt.Println(p) //{Millhouse {Cherry Dennis}}
	updated := R(p, "Parents.Dad", "Larry")
	fmt.Println(updated) //{Millhouse {Cherry Larry}}
}

```

Again, feel free to look in `r_test.go` to see more examples