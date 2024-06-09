package main

import (
	"bytes"
	"reflect"
	"unicode"
)

type Visitor interface {
	Ignore(typ reflect.Type) bool
	Visit(typ reflect.Type)
	VisitStruct(typ reflect.Type)
	VisitInterface(typ reflect.Type)
	VisitTypeAlias(typ reflect.Type)
	VisitSlice(typ reflect.Type)
	VisitPtr(typ reflect.Type)
	VisitField(name string, typ reflect.Type)
}

var _ Visitor = (*baseVisitor)(nil)

type baseVisitor struct {
	Visitor
	processed map[string]bool
}

func newBaseVisitor(v Visitor) *baseVisitor {
	return &baseVisitor{
		Visitor:   v,
		processed: map[string]bool{},
	}
}

func (v *baseVisitor) Visit(typ reflect.Type) {
	if typ == nil {
		return
	}
	if v.Ignore(typ) {
		return
	}
	typName := typeName(typ)
	if v.processed[typName] {
		return
	}
	v.processed[typName] = true
	switch typ.Kind() {
	case reflect.Struct:
		v.VisitStruct(typ)
	case reflect.Interface:
		v.VisitInterface(typ)
	case reflect.Slice:
		v.VisitSlice(typ)
	case reflect.Ptr:
		v.VisitPtr(typ)
	case reflect.Map:
		// do nothing
	default:
		if typ.PkgPath() != "" {
			v.VisitTypeAlias(typ)
		}
	}
}

func (v *baseVisitor) VisitStruct(typ reflect.Type) {
	v.Visitor.VisitStruct(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		v.VisitField(field.Name, field.Type)
	}
}

func (v *baseVisitor) VisitInterface(typ reflect.Type) {
	v.Visitor.VisitInterface(typ)
}

func (v *baseVisitor) VisitTypeAlias(typ reflect.Type) {
	v.Visitor.VisitTypeAlias(typ)
}

func (v *baseVisitor) VisitSlice(typ reflect.Type) {
	v.Visitor.VisitSlice(typ)
	v.Visit(typ.Elem())
}

func (v *baseVisitor) VisitPtr(typ reflect.Type) {
	v.Visitor.VisitPtr(typ)
	v.Visit(typ.Elem())
}

func (v *baseVisitor) VisitField(name string, typ reflect.Type) {
	if name[0] > 'A' && 'Z' < name[0] { // if only exported field
		return
	}
	v.Visitor.VisitField(name, typ)
}

var _ Visitor = (*nopVisitor)(nil)

type nopVisitor struct{}

func (*nopVisitor) Ignore(typ reflect.Type) bool {
	return false
}
func (*nopVisitor) Visit(typ reflect.Type)                   {}
func (*nopVisitor) VisitStruct(typ reflect.Type)             {}
func (*nopVisitor) VisitInterface(typ reflect.Type)          {}
func (*nopVisitor) VisitTypeAlias(typ reflect.Type)          {}
func (*nopVisitor) VisitSlice(typ reflect.Type)              {}
func (*nopVisitor) VisitPtr(typ reflect.Type)                {}
func (*nopVisitor) VisitField(name string, typ reflect.Type) {}

func typeName(typ reflect.Type) string {
	if typ.Kind() == reflect.Ptr {
		return "*" + typeName(typ.Elem())
	} else if typ.Kind() == reflect.Slice {
		return "[]" + typeName(typ.Elem())
	} else {
		return typ.Name()
	}
}

func toSnakeCase(s string) string {
	var b bytes.Buffer

	for i, r := range s {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}
