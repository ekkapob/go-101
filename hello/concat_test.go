package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var output string

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		output = "hello" + "world"
	}
	checkError(b, output)
}

func BenchmarkConcatWithFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		output = fmt.Sprint("hello", "world")
	}
	checkError(b, output)
}
func BenchmarkConcatWithStringsJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		output = strings.Join([]string{"hello", "world"}, "")
	}
	checkError(b, output)
}

func BenchmarkConcatWithByteBuffer(b *testing.B) {
	var o bytes.Buffer
	for n := 0; n < b.N; n++ {
		o = bytes.Buffer{}
		o.WriteString("hello")
		o.WriteString("world")
	}
	checkError(b, o.String())
}

func BenchmarkConcatWithStringBuilder(b *testing.B) {
	var s strings.Builder
	for n := 0; n < b.N; n++ {
		s = strings.Builder{}
		s.WriteString("hello")
		s.WriteString("world")
	}
	checkError(b, s.String())
}

func checkError(b *testing.B, output string) {
	if output != "helloworld" {
		b.Error("wrong output")
	}
}
