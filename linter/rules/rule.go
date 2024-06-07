package rules

import (
	"github.com/wreulicke/sqldefvet/internal/parser"
	"github.com/wreulicke/sqldefvet/linter/report"
)

type Rule interface {
	Apply(*parser.DDL) ([]*report.Failure, error)
}
