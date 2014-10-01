package main

import "strings"
import "container/list"

func wrap(tag, content string) string {
	if tag != "" {
		return "<" + tag + ">" + content + "</" + tag + ">"
	}
	return content
}

type block struct {
	openingTag string
	closingTag string
	kind       string
	children   []*block
	attrs      map[string]string
	content    string
}

func (b *block) renderHTML() string {
	if b.content != "" {
		return b.openingTag + b.content + b.closingTag
	}
	renderedChildren := make([]string, len(b.children))
	for i := 0; i < len(b.children); i++ {
		renderedChildren = append(renderedChildren, b.children[i].renderHTML())
	}
	return b.openingTag + strings.Join(renderedChildren, "") + b.closingTag
}

// Parse markdown string into custom tree
func Parse(markdown string) *block {
	document := &block{}
	stack := new(list.List)
	stack.PushBack(document)
	lines := strings.Split(markdown, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Index(line, "\t") == 0 || strings.Index(line, "    ") == 0 {
			codeBlock := &block{openingTag: "<pre><code>", closingTag: "\n</code></pre>", kind: "code"}
			document.children = append(document.children, codeBlock)
			codeBlock.content = strings.Replace(line, "\t", "    ", -1)
		} else if line == "***" || line == "---" || line == "___" {
			hrBlock := &block{openingTag: "<hr />\n"}
			document.children = append(document.children, hrBlock)
		} else {
			pBlock := &block{openingTag: "<p>", closingTag: "</p>", kind: "p"}
			document.children = append(document.children, pBlock)
			pBlock.content = line
		}

	}
	return document
}

// RenderHTML based on the given document
func RenderHTML(document *block) string {
	return strings.Trim(document.renderHTML(), "\n")
}
