package main

import (
	"testing"
	"bytes"
	"fmt"
)

func TestWriteHeader(t *testing.T) {
	buf := &bytes.Buffer{}
	WriteHeader(buf)
	fmt.Println(buf)
}
