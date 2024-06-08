package main

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
	sqldefparser "github.com/sqldef/sqldef/parser"
	"github.com/wreulicke/sqldefvet/internal/linter"
	"github.com/wreulicke/sqldefvet/internal/parser"
	"github.com/wreulicke/sqldefvet/internal/reporter"
	"github.com/wreulicke/sqldefvet/linter/rules"
)

func New() *cobra.Command {
	var mode string
	cmd := &cobra.Command{
		Use:   "sqldefvet",
		Short: "sqldefvet is a linter for SQL schema definitions",
		Long:  `sqldefvet is a linter for SQL schema definitions`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			return root(w, args[0], mode)
		},
	}
	cmd.Flags().StringVarP(&mode, "mode", "m", "", "parser mode")
	return cmd
}

func toParserMode(mode string) (parser.Mode, error) {
	switch mode {
	case "mysql":
		return parser.Mode(sqldefparser.ParserModeMysql), nil
	case "postgres":
		return parser.Mode(sqldefparser.ParserModePostgres), nil
	case "sqlite3":
		return parser.Mode(sqldefparser.ParserModeSQLite3), nil
	case "mssql":
		return parser.Mode(sqldefparser.ParserModeMssql), nil
	default:
		return parser.Mode(sqldefparser.ParserModeMysql), errors.New("invalid mode")
	}
}

func root(w io.Writer, schema string, mode string) error {
	parserMode, err := toParserMode(mode)
	if err != nil {
		return err
	}
	f, err := os.Open(schema)
	if err != nil {
		return err
	}
	defer f.Close()
	bs, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	ddl, err := parser.New().Parse(string(bs), parserMode)
	if err != nil {
		return err
	}

	fs, err := linter.New().Run(ddl, []rules.Rule{})
	if err != nil {
		return err
	}
	return reporter.New(w).Report(fs)
}
