package regexp

import (
	"fmt"
	"regexp"
	"strings"
)

type ParsingResult struct {
	str  string
	list []ParsingResult
}

// Parsers

type Parser interface {
	Regexp() string
	Parse(s string) ParsingResult
}

type literalParser struct {
	regexpString string
	regexp       *regexp.Regexp
}

func (t literalParser) Parse(s string) ParsingResult {
	if !t.regexp.MatchString(s) {
		panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
	}
	return ParsingResult{str: s}
}

func (t literalParser) Regexp() string {
	return t.regexpString
}

type regexpParser struct {
	regexpString string
	regexp       *regexp.Regexp
}

func (t regexpParser) Parse(s string) ParsingResult {
	if !t.regexp.MatchString(s) {
		panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
	}
	return ParsingResult{str: s}
}

func (t regexpParser) Regexp() string {
	return t.regexpString
}

type sequenceParser struct {
	regexpString string
	parts        []*regexp.Regexp
	children     []Parser
}

func (t sequenceParser) Parse(s string) ParsingResult {
	res := ParsingResult{list: make([]ParsingResult, 0)}
	for i, p := range t.parts {
		loc := p.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found for regexp '%v'", i, t.regexpString))
		}
		switch t.children[i].(type) {
		case literalParser:
			// nothing
		default:
			res.list = append(res.list, t.children[i].Parse(s[:loc[1]]))
		}
		s = s[loc[1]:]
	}
	return res
}

func (t sequenceParser) Regexp() string {
	return t.regexpString
}

type switchParser struct {
	regexpString string
	parts        []*regexp.Regexp
	children     []Parser
}

func (t switchParser) Parse(s string) ParsingResult {
	for i, p := range t.parts {
		if p.MatchString(s) {
			return t.children[i].Parse(s)
		}
	}
	panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
}

func (t switchParser) Regexp() string {
	return t.regexpString
}

type listParser struct {
	regexpString string
	part         *regexp.Regexp
	separator    *regexp.Regexp
	parser       Parser
}

func (t listParser) Parse(s string) ParsingResult {
	res := ParsingResult{list: make([]ParsingResult, 0)}
	for s != "" {
		loc := t.part.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found", len(res.list)))
		}
		res.list = append(res.list, t.parser.Parse(s[:loc[1]]))
		s = s[loc[1]:]
		if s != "" {
			loc := t.separator.FindStringIndex(s)
			if loc == nil || loc[0] != 0 {
				panic(fmt.Sprintf("%v separator not found", len(res.list)))
			}
			s = s[loc[1]:]
		}
	}
	return res
}

func (t listParser) Regexp() string {
	return t.regexpString
}

// Elements

type parsingElement interface {
	Complie() Parser
}

type Literal struct {
	unescaped string
}

func (t Literal) Complie() Parser {
	parser := literalParser{}
	parser.regexpString = regexp.QuoteMeta(t.unescaped)
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type Token struct {
	unescaped string
}

func (t Token) Complie() Parser {
	parser := regexpParser{}
	parser.regexpString = regexp.QuoteMeta(t.unescaped)
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type Regexp struct {
	escaped string
}

func (t Regexp) Complie() Parser {
	parser := regexpParser{}
	parser.regexpString = t.escaped
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type Number struct{}

func (t Number) Complie() Parser {
	return Regexp{`[-+.e\d]+`}.Complie()
}

type Word struct{}

func (t Word) Complie() Parser {
	return Regexp{`[\w]+`}.Complie()
}

type Sequence []parsingElement

func (t Sequence) Complie() Parser {
	parser := sequenceParser{}
	var b strings.Builder
	parser.parts = make([]*regexp.Regexp, len(t))
	parser.children = make([]Parser, len(t))
	for i, e := range t {
		parser.children[i] = e.Complie()
		cps := parser.children[i].Regexp()
		b.WriteString(cps)
		parser.parts[i] = regexp.MustCompile(cps)
	}
	parser.regexpString = b.String()
	return parser
}

type Switch []parsingElement

func (t Switch) Complie() Parser {
	parser := switchParser{}
	var b strings.Builder
	parser.parts = make([]*regexp.Regexp, len(t))
	parser.children = make([]Parser, len(t))
	for i, e := range t {
		parser.children[i] = e.Complie()
		cps := parser.children[i].Regexp()
		if i > 0 {
			b.WriteString(`|`)
		}
		b.WriteString(`(`)
		b.WriteString(cps)
		b.WriteString(`)`)
		parser.parts[i] = regexp.MustCompile(cps)
	}
	parser.regexpString = b.String()
	return parser
}

type List struct {
	element            parsingElement
	unescapedSeparator string
}

func (t List) Complie() Parser {
	parser := listParser{}
	parser.parser = t.element.Complie()
	cps := parser.parser.Regexp()
	parser.part = regexp.MustCompile(cps)
	sps := regexp.QuoteMeta(t.unescapedSeparator)
	parser.separator = regexp.MustCompile(sps)
	parser.regexpString = `(` + cps + `(` + sps + cps + `)*(` + sps + `)?)?`
	return parser
}
