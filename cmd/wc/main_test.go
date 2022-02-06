package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		split    bufio.SplitFunc
		expected int
	}{
		{"a b 中", bufio.ScanWords, 3},
		{"a\tb\t中", bufio.ScanBytes, 7},
		{"polarisxu studygolang", bufio.ScanWords, 2},
		{"polarisxu Go 语言 中文网", bufio.ScanWords, 4},
		{"polarisxu Go 语言 中文网", bufio.ScanRunes, 19},
	}

	var b = new(bytes.Buffer)
	for _, tt := range tests {
		b.Reset()
		b.WriteString(tt.input)
		if got := count(b, tt.split); got != tt.expected {
			t.Errorf("count() = %v, want %v", got, tt.expected)
		}
	}
}
