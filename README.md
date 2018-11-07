# Progress Spinner for Go

[![Build Status](https://travis-ci.org/domdavis/gospin.svg?branch=master)](https://travis-ci.org/domdavis/gospin)
[![Coverage Status](https://coveralls.io/repos/github/domdavis/gospin/badge.svg?branch=master)](https://coveralls.io/github/domdavis/gospin?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/4fb6bb7263b9ef2da58b/maintainability)](https://codeclimate.com/github/domdavis/gospin/maintainability)
[![](https://godoc.org/github.com/domdavis/gospin?status.svg)](http://godoc.org/github.com/domdavis/gospin)

`gospin` provides a console _spinner_ to indicate activity with a running 
program. By default the spinner is animated on every tick, but a porcelain mode
for log output is also provided. 

## Installation

```
go get github.com/domdavis/gospin
```

## Basic Usage

```go
package main

import (
    "fmt"
    "github.com/domdavis/gospin"
    "time"
)

func main() {
	fmt.Print("Working ")
	s := gospin.New()
	
	for i := 0; i < 50; i++ {
		s.Advance()
	}
	
	s.Done()
	fmt.Println("... done")
}
```

## Advanced Usage

Spinners can be constructed by providing a set of frames, specifying the width
of the frame if non-standard characters are used, and the output destination.

```
s := gospin.New("☱", "☲", "☴")
s.Width(1)
s.Writer(os.Stderr)

// Start drawing the spinner
s.Advance()

// ...do work, calling Advance() as the work progresses.

// Finally remove the spinner.
s.Done()
```

See the examples for a full list of predefined spinners.
