package tagchecker_test

import (
	tagchecker "tagchecker"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, tagchecker.Analyzer, "a")
}
