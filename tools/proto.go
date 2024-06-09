package main

import (
	"bytes"
	"reflect"
	"strconv"

	"github.com/sqldef/sqldef/parser"
)

type protoGenerator struct {
	*bytes.Buffer
	nopVisitor
	visitor Visitor
}

func newProtoGenerator(buf *bytes.Buffer) *protoGenerator {
	g := &protoGenerator{
		Buffer: buf,
	}
	g.visitor = newBaseVisitor(g)
	return g
}

func (g *protoGenerator) Ignore(typ reflect.Type) bool {
	if typ.PkgPath() == "github.com/sqldef/sqldef/parser" && typ.Name() == "ColIdent" {
		return true
	}
	if typ.PkgPath() == "github.com/sqldef/sqldef/parser" && typ.Name() == "TableIdent" {
		return true
	}
	return false
	// do nothing
}

func (g *protoGenerator) VisitStruct(typ reflect.Type) {
	writeMessage(g.Buffer, typ)
}

func (g *protoGenerator) VisitInterface(typ reflect.Type) {
	// TODO: implement
}

func (g *protoGenerator) VisitField(fieldName string, typ reflect.Type) {
	g.visitor.Visit(typ)
}

func (g *protoGenerator) Generate() {
	g.WriteString(`syntax = "proto3";
package sqldefvet;`)
	g.WriteString("\n")

	writeColIdent(g.Buffer)
	writeTableIdent(g.Buffer)

	typ := reflect.TypeOf(parser.DDL{})
	g.visitor.Visit(typ)
}

func writeProto(buf *bytes.Buffer) {
	g := newProtoGenerator(buf)
	g.Generate()
}

func writeMessage(g *bytes.Buffer, typ reflect.Type) {
	typName := typ.Name()
	g.WriteString("\nmessage ")
	g.WriteString(typName)
	g.WriteString(" {\n")
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		if fieldName[0] > 'A' && 'Z' < fieldName[0] {
			continue
		}
		kind := field.Type.Kind()
		if kind == reflect.Map {
			continue
		}
		g.WriteString("\t")
		if kind == reflect.Slice {
			g.WriteString("repeated ")
		}
		fieldTyp := getElementType(field.Type)
		g.WriteString(fieldTyp.Name())
		g.WriteString(" ")
		g.WriteString(toSnakeCase(field.Name))
		g.WriteString(" = ")
		g.WriteString(strconv.Itoa(i + 1))
		g.WriteString(";\n")
	}
	g.WriteString("}\n")
}

func getElementType(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Ptr {
		return typ.Elem()
	} else if typ.Kind() == reflect.Slice {
		return getElementType(typ.Elem())
	}
	return typ
}

func writeColIdent(buf *bytes.Buffer) {
	buf.WriteString("\nmessage ColIdent {\n")
	buf.WriteString("\tstring val = 1;\n")
	buf.WriteString("}\n")
}

func writeTableIdent(buf *bytes.Buffer) {
	buf.WriteString("\nmessage TableIdent {\n")
	buf.WriteString("\tstring val = 1;\n")
	buf.WriteString("}\n")
}
