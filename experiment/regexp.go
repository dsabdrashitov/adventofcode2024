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

type iLiteralParser struct {
	regexpString string
	regexp       *regexp.Regexp
}

func (t iLiteralParser) Parse(s string) ParsingResult {
	if !t.regexp.MatchString(s) {
		panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
	}
	return ParsingResult{stringValue: s}
}

func (t iLiteralParser) Regexp() string {
	return t.regexpString
}

type iIntParser struct{}

func (t iIntParser) Parse(s string) ParsingResult {
	return ParsingResult{intValue: Must(strconv.Atoi(s))}
}

func (t iIntParser) Regexp() string {
	return `[\d]+`
}

type iVariableParser struct {
	name   string
	parser Parser
}

func (t iVariableParser) Parse(s string) ParsingResult {
	return ParsingResult{stringValue: t.name, values: []ParsingResult{t.parser.Parse(s)}}
}

func (t iVariableParser) Regexp() string {
	return t.parser.Regexp()
}

type iSequenceParser struct {
	regexpString string
	parts        []*regexp.Regexp
	children     []Parser
}

func (t iSequenceParser) Parse(s string) ParsingResult {
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

func (t iSequenceParser) Regexp() string {
	return t.regexpString
}

type iAlternativeParser struct {
	regexpString string
	parts        []*regexp.Regexp
	children     []Parser
}

func (t iAlternativeParser) Parse(s string) ParsingResult {
	for i, p := range t.parts {
		if p.MatchString(s) {
			return t.children[i].Parse(s)
		}
	}
	panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
}

func (t iAlternativeParser) Regexp() string {
	return t.regexpString
}

type iMultipleParser struct {
	regexpString string
	part         *regexp.Regexp
	parser       Parser
}

func (t iMultipleParser) Parse(s string) ParsingResult {
	res := ParsingResult{values: make([]ParsingResult, 0)}
	for s != "" {
		loc := t.part.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found", len(res.values)))
		}
		res.values = append(res.values, t.parser.Parse(s[:loc[1]]))
		s = s[loc[1]:]
	}
	return res
}

func (t iMultipleParser) Regexp() string {
	return t.regexpString
}

// Elements

type iParsingElement interface {
	Complie(rules map[string]Parser) Parser
}

type Literal struct {
	escaped string
}

func (t Literal) Complie(rules map[string]Parser) Parser {
	parser := iLiteralParser{}
	parser.regexpString = t.escaped
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type IntNumber struct{}

func (t IntNumber) Complie(rules map[string]Parser) Parser {
	return iIntParser{}
}

type Variable struct {
	name string
}

func (t Variable) Complie(rules map[string]Parser) Parser {
	parser := iVariableParser{}
	parser.name = t.name
	parser.parser = rules[t.name]
	return parser
}

type Sequence []iParsingElement

func (t Sequence) Complie(rules map[string]Parser) Parser {
	parser := iSequenceParser{}
	var b strings.Builder
	parser.parts = make([]*regexp.Regexp, len(t))
	parser.children = make([]Parser, len(t))
	for i, e := range t {
		parser.children[i] = e.Complie(rules)
		cps := parser.children[i].Regexp()
		b.WriteString(cps)
		parser.parts[i] = regexp.MustCompile(cps)
	}
	parser.regexpString = b.String()
	return parser
}

type Alternative []iParsingElement

func (t Alternative) Complie(rules map[string]Parser) Parser {
	parser := iAlternativeParser{}
	var b strings.Builder
	parser.parts = make([]*regexp.Regexp, len(t))
	parser.children = make([]Parser, len(t))
	for i, e := range t {
		parser.children[i] = e.Complie(rules)
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

type Multiple struct {
	element  iParsingElement
	reSuffix string
}

func (t Multiple) Complie(rules map[string]Parser) Parser {
	parser := iMultipleParser{}
	parser.parser = t.element.Complie(rules)
	cps := parser.parser.Regexp()
	parser.part = regexp.MustCompile(cps)
	parser.regexpString = `(` + cps + `)` + t.reSuffix
	return parser
}
