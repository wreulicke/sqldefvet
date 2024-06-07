package reporter

import (
	"fmt"
	"io"

	"github.com/wreulicke/sqldefvet/linter/report"
)

type Reporter struct {
	w io.Writer
}

func New(w io.Writer) *Reporter {
	return &Reporter{w}
}

func (r *Reporter) Report(failures []*report.Failure) error {
	for _, f := range failures {
		_, err := fmt.Fprintln(r.w, f.Message)
		if err != nil {
			return err
		}
	}
	return nil
}
