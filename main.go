package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

type Header struct {
	Title      string   `key:"title"      dc:"true"`
	Date       string   `key:"date"       dc:"true"`
	Draft      bool     `key:"draft"      dc:"false"`
	Toc        bool     `key:"toc"        dc:"false"`
	Categories []string `key:"categories" dc:"false"`
	Tags       []string `key:"tags"       dc:"false"`
	Author     string   `key:"author"     dc:"true"`
}

var (
	now = time.Now()
)

func main() {
	file, err := os.Create(now.Format("200601020304")+".md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	WriteHeader(file)
}

func WriteHeader(writer io.Writer) {
	var header Header
	header.Date = now.Format(time.RFC3339)
	rt, rv := reflect.TypeOf(header), reflect.ValueOf(header)

	fmt.Fprint(writer, "+++\n")
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get("key")
		dc := field.Tag.Get("dc")
		if dc == "true" {
			fmt.Fprintf(writer, "%v = \"%v\"\n", key, rv.Field(i).Interface())
		} else {
			fmt.Fprintf(writer, "%v = %v\n", key, rv.Field(i).Interface())
		}
	}
	fmt.Fprint(writer, "+++\n\n")
	fmt.Fprint(writer, "<!--more-->")
}
