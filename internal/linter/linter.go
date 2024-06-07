package linter

import (
	"github.com/wreulicke/sqldefvet/internal/parser"
	"github.com/wreulicke/sqldefvet/linter/report"
	"github.com/wreulicke/sqldefvet/linter/rules"
)

type Linter struct {
}

func New() *Linter {
	return &Linter{}
}

func (l *Linter) Run(ddl *parser.DDL, rules []rules.Rule) ([]*report.Failure, error) {
	failures := []*report.Failure{}
	for _, rule := range rules {
		rs, err := rule.Apply(ddl)
		if err != nil {
			return nil, err
		}
		failures = append(failures, rs...)
	}
	return failures, nil
}
