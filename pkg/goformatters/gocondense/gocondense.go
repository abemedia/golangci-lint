package gocondense

import (
	"errors"

	"github.com/abemedia/gocondense"

	"github.com/golangci/golangci-lint/v2/pkg/config"
)

const Name = "gocondense"

type Formatter struct {
	formatter *gocondense.Formatter
}

func New(settings *config.GoCondenseSettings) (*Formatter, error) {
	var cfg gocondense.Config

	if settings != nil {
		if settings.MaxLen < 0 || settings.TabWidth < 0 {
			return nil, errors.New("max-len and tab-width must not be negative")
		}

		cfg = gocondense.Config{
			MaxLen:   settings.MaxLen,
			TabWidth: settings.TabWidth,
		}
	}

	return &Formatter{formatter: gocondense.New(cfg)}, nil
}

func (*Formatter) Name() string {
	return Name
}

func (f *Formatter) Format(_ string, src []byte) ([]byte, error) {
	return f.formatter.Source(src)
}
