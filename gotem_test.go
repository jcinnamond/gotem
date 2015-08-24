package main

import (
	"bytes"
	"fmt"
	"testing"
)

var test_partials = map[string]string{
	"header": "<html><body>",
	"footer": "</body></html>",
}

func TestLoadPartialsOK(t *testing.T) {
	partials, err := loadPartials("test/partials")

	if err != nil {
		t.Errorf("Unexpected error reading partials, %v", err)
		return
	}

	if len(partials) != len(test_partials) {
		t.Errorf("Expected %d partials, got %d", len(test_partials), len(partials))
		return
	}

	for name, expected := range test_partials {
		content, ok := partials[name]
		if !ok {
			t.Errorf("Missing partial %v", name)
		}

		if content != expected {
			t.Errorf("Expected %v to contain %v, got %v", name, expected, content)
		}
	}
}

func TestLoadPartialsBadPath(t *testing.T) {
	_, err := loadPartials("bad/path")

	if err == nil {
		t.Errorf("Expected error for bad path, got none")
	}
}

func TestCompileTemplate(t *testing.T) {
	src := "{{include \"header\"}}<h1>Hello</h1>{{include \"footer\"}}"
	in := bytes.NewBufferString(src)
	out := new(bytes.Buffer)

	err := compileTemplate(in, out, test_partials)
	if err != nil {
		t.Errorf("Unexpected error compiling template: %v", err)
		return
	}

	expected := fmt.Sprintf("%s<h1>Hello</h1>%s", test_partials["header"], test_partials["footer"])
	if out.String() != expected {
		t.Errorf("Wrong output.\n%v\n\nExpected\n%v", out.String(), expected)
	}
}
