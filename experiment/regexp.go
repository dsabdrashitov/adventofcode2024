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

type iOptionalParser struct {
	regexpString string
	regexp       *regexp.Regexp
}

func (t iOptionalParser) Parse(s string) ParsingResult {
	if !t.regexp.MatchString(s) {
		panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
	}
	return ParsingResult{stringValue: s}
}

func (t iOptionalParser) Regexp() string {
	return t.regexpString
}

type iRegexpParser struct {
	regexpString string
	regexp       *regexp.Regexp
}

func (t iRegexpParser) Parse(s string) ParsingResult {
	if !t.regexp.MatchString(s) {
		panic(fmt.Sprintf("mismatch regexp '%v'", t.regexpString))
	}
	return ParsingResult{stringValue: s}
}

func (t iRegexpParser) Regexp() string {
	return t.regexpString
}

type iIntParser struct{}

func (t iIntParser) Parse(s string) ParsingResult {
	return ParsingResult{intValue: Must(strconv.Atoi(s))}
}

func (t iIntParser) Regexp() string {
	return `\-?[\d]+`
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

type iListParser struct {
	regexpString string
	part         *regexp.Regexp
	separator    *regexp.Regexp
	parser       Parser
}

func (t iListParser) Parse(s string) ParsingResult {
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

func (t iListParser) Regexp() string {
	return t.regexpString
}

// Elements

type iParsingElement interface {
	Complie() Parser
}

type Literal struct {
	unescaped string
}

func (t Literal) Complie() Parser {
	parser := iLiteralParser{}
	parser.regexpString = regexp.QuoteMeta(t.unescaped)
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type Optional struct {
	unescaped string
}

func (t Optional) Complie() Parser {
	parser := iOptionalParser{}
	parser.regexpString = `(` + regexp.QuoteMeta(t.unescaped) + `)?`
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type Regexp struct {
	escaped string
}

func (t Regexp) Complie() Parser {
	parser := iRegexpParser{}
	parser.regexpString = t.escaped
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

type IntNumber struct{}

func (t IntNumber) Complie() Parser {
	return iIntParser{}
}

type Sequence []iParsingElement

func (t Sequence) Complie() Parser {
	parser := iSequenceParser{}
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

type Alternative []iParsingElement

func (t Alternative) Complie() Parser {
	parser := iAlternativeParser{}
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

type Multiple struct {
	element  iParsingElement
	reSuffix string
}

func (t Multiple) Complie() Parser {
	parser := iMultipleParser{}
	parser.parser = t.element.Complie()
	cps := parser.parser.Regexp()
	parser.part = regexp.MustCompile(cps)
	parser.regexpString = `(` + cps + `)` + t.reSuffix
	return parser
}

type List struct {
	element            iParsingElement
	unescapedSeparator string
}

func (t List) Complie() Parser {
	parser := iListParser{}
	parser.parser = t.element.Complie()
	cps := parser.parser.Regexp()
	parser.part = regexp.MustCompile(cps)
	sps := regexp.QuoteMeta(t.unescapedSeparator)
	parser.separator = regexp.MustCompile(sps)
	parser.regexpString = `(` + cps + `(` + sps + cps + `)*(` + sps + `)?)?`
	return parser
}
