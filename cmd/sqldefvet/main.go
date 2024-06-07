package main

import (
	"io"
	"log"
	"os"

	sqldefparser "github.com/sqldef/sqldef/parser"
	"github.com/wreulicke/sqldefvet/internal/linter"
	"github.com/wreulicke/sqldefvet/internal/parser"
	"github.com/wreulicke/sqldefvet/internal/reporter"
	"github.com/wreulicke/sqldefvet/linter/rules"
)

func mainInternal() error {
	schema := os.Args[1]

	f, err := os.Open(schema)
	if err != nil {
		return err
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	ddl, err := parser.New().Parse(string(bs), parser.Mode(sqldefparser.ParserModePostgres))
	if err != nil {
		return err
	}

	fs, err := linter.New().Run(ddl, []rules.Rule{})
	if err != nil {
		return err
	}
	return reporter.New(os.Stdout).Report(fs)
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatal(err)
	}
}
