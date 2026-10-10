package gocondense

import (
	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/gocondense"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.GoCondenseSettings) *goanalysis.Linter {
	formatter, err := gocondense.New(settings)
	if err != nil {
		internal.LinterLogger.Fatalf("%s: create analyzer: %v", gocondense.Name, err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(gocondense.Name),
				"Checks if code is formatted, and condenses multi-line constructs onto single lines where they fit",
				formatter,
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
