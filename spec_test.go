package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

type Spec struct {
	Examples []Example
}
type Example struct {
	Name     string
	Markdown string
	HTML     string
}

func TestSpec(t *testing.T) {
	bs, err := ioutil.ReadFile("spec.json")
	if err != nil {
		t.Error("Unable to open the spec.")
	}
	var spec Spec
	err = json.Unmarshal(bs, &spec)
	if err != nil {
		t.Error("JSON format was wrong", err)
	}

	i := 0
	failed := 0
	for i < len(spec.Examples) {
		example := spec.Examples[i]
		renderedHTML := RenderHTML(Parse(example.Markdown))
		if renderedHTML != example.HTML {
			failed = failed + 1
			t.Error("===== " + example.Name + " failed. =====")
			t.Error("===== MARKDOWN =====")
			t.Error(strings.Replace(example.Markdown, "\t", "→", -1))
			t.Error("===== EXPECTED HTML =====")
			t.Error(example.HTML)
			t.Error("===== GOT HTML =====")
			t.Error(renderedHTML)
			t.Error(pretty.Diff(example.HTML, renderedHTML))
			t.Error("\n\n")
		}
		i = i + 1
		if failed > 3 {
			i = len(spec.Examples)
		}
	}
	passed := len(spec.Examples) - failed
	t.Log("Passed: ", passed, "/", len(spec.Examples))

}
