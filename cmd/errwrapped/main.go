package main

import (
	"github.com/bitkey-platform/errwrapped"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(errwrapped.Analyzer) }
