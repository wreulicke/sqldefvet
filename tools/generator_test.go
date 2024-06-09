package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/sqldef/sqldef/parser"
)

type ValTuple parser.Exprs

func TestForDebug(t *testing.T) {
	buf := &bytes.Buffer{}
	g := newMapperGenerator(buf)
	g.VisitSlice(reflect.TypeOf(parser.ValTuple{}))
	t.Error("err")
}

type test int

func TestReflect(t *testing.T) {
	typ := reflect.TypeOf(test(0))
	t.Errorf("kind: %v", typ.Kind())
	t.Errorf("pkg path: %v", typ.PkgPath())
	t.Errorf("pkg path: %v", typ.Name())
}
