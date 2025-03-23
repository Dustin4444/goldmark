package goldmark_test

import (
	"encoding/json"
	"os"
	"testing"

	. "github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

type commonmarkSpecTestCase struct {
	Markdown  string `json:"markdown"`
	HTML      string `json:"html"`
	Example   int    `json:"example"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	Section   string `json:"section"`
}

func TestSpec(t *testing.T) {
	bs, err := os.ReadFile("_test/spec.json")
	if err != nil {
		panic(err)
	}
	var testCases []commonmarkSpecTestCase
	if err := json.Unmarshal(bs, &testCases); err != nil {
		panic(err)
	}
	cases := []testutil.MarkdownTestCase{}
	nos := testutil.ParseCliCaseArg()
	for _, c := range testCases {
		shouldAdd := len(nos) == 0
		if !shouldAdd {
			for _, no := range nos {
				if c.Example == no {
					shouldAdd = true
					break
				}
			}
		}

		if shouldAdd {
			cases = append(cases, testutil.MarkdownTestCase{
				No:       c.Example,
				Markdown: c.Markdown,
				Expected: c.HTML,
			})
		}
	}
	markdown := New(WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	))
	testutil.DoTestCases(markdown, cases, t)
}

func TestSpec_EdgeCases(t *testing.T) {
	edgeCases := []testutil.MarkdownTestCase{
		{
			No:       1,
			Markdown: "This is a [link](http://example.com) with a URL.",
			Expected: "<p>This is a <a href=\"http://example.com\">link</a> with a URL.</p>\n",
		},
		{
			No:       2,
			Markdown: "This is a [link](http://example.com) with a URL and a title.",
			Expected: "<p>This is a <a href=\"http://example.com\" title=\"Example\">link</a> with a URL and a title.</p>\n",
		},
		{
			No:       3,
			Markdown: "This is a [link](http://example.com) with a URL and an ID.",
			Expected: "<p>This is a <a href=\"http://example.com\" id=\"example\">link</a> with a URL and an ID.</p>\n",
		},
		{
			No:       4,
			Markdown: "This is a [link](http://example.com) with a URL, a title, and an ID.",
			Expected: "<p>This is a <a href=\"http://example.com\" title=\"Example\" id=\"example\">link</a> with a URL, a title, and an ID.</p>\n",
		},
		{
			No:       5,
			Markdown: "This is a [link](http://example.com) with a URL and a class.",
			Expected: "<p>This is a <a href=\"http://example.com\" class=\"example\">link</a> with a URL and a class.</p>\n",
		},
		{
			No:       6,
			Markdown: "This is a [link](http://example.com) with a URL, a title, and a class.",
			Expected: "<p>This is a <a href=\"http://example.com\" title=\"Example\" class=\"example\">link</a> with a URL, a title, and a class.</p>\n",
		},
		{
			No:       7,
			Markdown: "This is a [link](http://example.com) with a URL, an ID, and a class.",
			Expected: "<p>This is a <a href=\"http://example.com\" id=\"example\" class=\"example\">link</a> with a URL, an ID, and a class.</p>\n",
		},
		{
			No:       8,
			Markdown: "This is a [link](http://example.com) with a URL, a title, an ID, and a class.",
			Expected: "<p>This is a <a href=\"http://example.com\" title=\"Example\" id=\"example\" class=\"example\">link</a> with a URL, a title, an ID, and a class.</p>\n",
		},
	}

	markdown := New(WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	))
	testutil.DoTestCases(markdown, edgeCases, t)
}
