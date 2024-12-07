package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ParsingResult struct {
	stringValue string
	intValue    int
	values      []ParsingResult
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
	return ParsingResult{stringValue: s}
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
	return ParsingResult{stringValue: s}
}

func (t regexpParser) Regexp() string {
	return t.regexpString
}

type intParser struct{}

func (t intParser) Parse(s string) ParsingResult {
	return ParsingResult{intValue: Must(strconv.Atoi(s))}
}

func (t intParser) Regexp() string {
	return `\-?[\d]+`
}

type sequenceParser struct {
	regexpString string
	parts        []*regexp.Regexp
	children     []Parser
}

func (t sequenceParser) Parse(s string) ParsingResult {
	res := ParsingResult{values: make([]ParsingResult, len(t.parts))}
	for i, p := range t.parts {
		loc := p.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found for regexp '%v'", i, t.regexpString))
		}
		res.values[i] = t.children[i].Parse(s[:loc[1]])
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
	res := ParsingResult{values: make([]ParsingResult, 0)}
	for s != "" {
		loc := t.part.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found", len(res.values)))
		}
		res.values = append(res.values, t.parser.Parse(s[:loc[1]]))
		s = s[loc[1]:]
		if s != "" {
			loc := t.separator.FindStringIndex(s)
			if loc == nil || loc[0] != 0 {
				panic(fmt.Sprintf("%v separator not found", len(res.values)))
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

type Regexp struct {
	escaped string
}

func (t Regexp) Complie() Parser {
	parser := regexpParser{}
	parser.regexpString = t.escaped
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type Int struct{}

func (t Int) Complie() Parser {
	return intParser{}
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
