package parser

import (
	"fmt"

	"github.com/sqldef/sqldef/parser"
)

type Mode parser.ParserMode

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(sql string, mode Mode) (*DDL, error) {
	stmt, err := parser.ParseDDL(sql, parser.ParserMode(mode))
	if err != nil {
		return nil, err
	}
	ddl, ok := stmt.(*parser.DDL)
	if !ok {
		return nil, fmt.Errorf("only supported ddl: %s", stmt)
	}
	return mapDDLPtr(ddl), nil
}
