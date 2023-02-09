package main

import (
	"github.com/nrnrk/errwrapped"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(errwrapped.Analyzer) }
