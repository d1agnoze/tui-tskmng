package parser

import (
	_ "embed"
	"os"
	"testing"
)

var s string

func read(t *testing.T) {
	t.Helper()
	raw, err := os.ReadFile("../../test/demo.tmg")
	if err != nil {
		t.Fatalf("Error while reading file: %v", err)
	}
	s = string(raw)
}

func TestParser(t *testing.T) {
	read(t)
	// Test code here
	t.Run("Parser should run without error", func(t *testing.T) {
		parser, err := New()
		if err != nil {
			t.Fatalf("Error while creating parser: %v", err)
		}

		out, err := parser.Parse(s)
		if err != nil {
			t.Fatalf("Error when parsing file: %v", err)
		}

		t.Logf("Output: \n%s", out)
	})
}
