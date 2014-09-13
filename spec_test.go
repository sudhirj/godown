package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	for i < len(spec.Examples) {
		example := spec.Examples[i]
		fmt.Println(example.Name)
		if example.Markdown != example.HTML {
			t.Error("===== " + example.Name + " failed. =====")
			t.Error("===== MARKDOWN =====")
			t.Error(example.Markdown)
			t.Error("===== EXPECTED HTML =====")
			t.Error(example.HTML)
			t.Error("===== GOT HTML =====")
			t.Error(example.HTML)
			t.Error(pretty.Diff(example.Markdown, example.HTML))
			t.Error("\n\n")
		}
		i = i + 1
	}

}
