package main

import (
	"github.com/base-go/baseql/tools/listEntryNonNullablePatcher/listEntryNonNullablePatcher"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(listEntryNonNullablePatcher.Analyzer)
}
