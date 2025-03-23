package parser

import (
	"errors"
	"fmt"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type autoLinkParser struct {
	allowedTypes map[ast.AutoLinkType]bool
}

var defaultAutoLinkParser = &autoLinkParser{
	allowedTypes: map[ast.AutoLinkType]bool{
		ast.AutoLinkEmail: true,
		ast.AutoLinkURL:   true,
	},
}

// NewAutoLinkParser returns a new InlineParser that parses autolinks
// surrounded by '<' and '>' .
func NewAutoLinkParser() InlineParser {
	return defaultAutoLinkParser
}

func (s *autoLinkParser) Trigger() []byte {
	return []byte{'<'}
}

func (s *autoLinkParser) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	line, segment := block.PeekLine()
	stop, typ, err := s.findAutoLink(line[1:])
	if err != nil {
		fmt.Printf("Error finding autolink: %v\n", err)
		return nil
	}
	if stop >= len(line) || line[stop] != '>' {
		fmt.Println("Error: Autolink not properly closed with '>'")
		return nil
	}
	value := ast.NewTextSegment(text.NewSegment(segment.Start+1, segment.Start+stop))
	block.Advance(stop + 1)
	return ast.NewAutoLink(typ, value)
}

func (s *autoLinkParser) findAutoLink(line []byte) (int, ast.AutoLinkType, error) {
	if s.allowedTypes[ast.AutoLinkEmail] {
		if stop := util.FindEmailIndex(line); stop >= 0 {
			return stop + 1, ast.AutoLinkEmail, nil
		}
	}
	if s.allowedTypes[ast.AutoLinkURL] {
		if stop := util.FindURLIndex(line); stop >= 0 {
			return stop + 1, ast.AutoLinkURL, nil
		}
	}
	return -1, 0, errors.New("no valid autolink found")
}

// AddAllowedType adds a new allowed autolink type to the parser.
func (s *autoLinkParser) AddAllowedType(typ ast.AutoLinkType) {
	s.allowedTypes[typ] = true
}

// RemoveAllowedType removes an allowed autolink type from the parser.
func (s *autoLinkParser) RemoveAllowedType(typ ast.AutoLinkType) {
	delete(s.allowedTypes, typ)
}
