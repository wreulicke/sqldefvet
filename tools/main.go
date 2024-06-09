package main

import (
	"bytes"
	"log"
	"os"
)

func mainInternal() error {
	outProto := &bytes.Buffer{}
	writeProto(outProto)

	outMapper := &bytes.Buffer{}
	writeMapper(outMapper)

	f, err := os.OpenFile("_proto/ddl.proto", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(outProto.Bytes()); err != nil {
		return err
	}

	f, err = os.OpenFile("internal/parser/mapper.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(outMapper.Bytes()); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatal(err)
	}
}
