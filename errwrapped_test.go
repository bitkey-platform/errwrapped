package errwrapped_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bitkey-platform/errwrapped"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, errwrapped.Analyzer, "a")
}
