package regexp

import (
	"fmt"
	"regexp"
	"strings"
)

type ParsingResult struct {
	S string
	L []ParsingResult
}

// Parsers

type Parser interface {
	getRegexpString() string
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
	return ParsingResult{S: s}
}

func (t literalParser) getRegexpString() string {
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
	return ParsingResult{S: s}
}

func (t regexpParser) getRegexpString() string {
	return t.regexpString
}

type sequenceParser struct {
	regexpString string
	parts        []*regexp.Regexp
	children     []Parser
}

func (t sequenceParser) Parse(s string) ParsingResult {
	res := ParsingResult{L: make([]ParsingResult, 0)}
	for i, p := range t.parts {
		loc := p.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found for regexp '%v'", i, t.regexpString))
		}
		switch t.children[i].(type) {
		case literalParser:
			// nothing
		default:
			res.L = append(res.L, t.children[i].Parse(s[:loc[1]]))
		}
		s = s[loc[1]:]
	}
	return res
}

func (t sequenceParser) getRegexpString() string {
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

func (t switchParser) getRegexpString() string {
	return t.regexpString
}

type listParser struct {
	regexpString string
	part         *regexp.Regexp
	separator    *regexp.Regexp
	parser       Parser
}

func (t listParser) Parse(s string) ParsingResult {
	res := ParsingResult{L: make([]ParsingResult, 0)}
	for s != "" {
		loc := t.part.FindStringIndex(s)
		if loc == nil || loc[0] != 0 {
			panic(fmt.Sprintf("%v part not found", len(res.L)))
		}
		res.L = append(res.L, t.parser.Parse(s[:loc[1]]))
		s = s[loc[1]:]
		if s != "" {
			loc := t.separator.FindStringIndex(s)
			if loc == nil || loc[0] != 0 {
				panic(fmt.Sprintf("%v separator not found", len(res.L)))
			}
			s = s[loc[1]:]
		}
	}
	return res
}

func (t listParser) getRegexpString() string {
	return t.regexpString
}

// Elements

type ParsingElement interface {
	Complie() Parser
}

func Literal(unescaped string) ParsingElement {
	return literalElement{unescaped}
}

type literalElement struct {
	unescaped string
}

func (t literalElement) Complie() Parser {
	parser := literalParser{}
	parser.regexpString = regexp.QuoteMeta(t.unescaped)
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

func Token(unescaped string) ParsingElement {
	return tokenElement{unescaped}
}

type tokenElement struct {
	unescaped string
}

func (t tokenElement) Complie() Parser {
	parser := regexpParser{}
	parser.regexpString = regexp.QuoteMeta(t.unescaped)
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

func Regexp(escaped string) ParsingElement {
	return regexpElement{escaped}
}

type regexpElement struct {
	escaped string
}

func (t regexpElement) Complie() Parser {
	parser := regexpParser{}
	parser.regexpString = t.escaped
	parser.regexp = regexp.MustCompile(parser.regexpString)
	return parser
}

func Number() ParsingElement {
	return numberElement{}
}

type numberElement struct{}

func (t numberElement) Complie() Parser {
	return regexpElement{`[-+.e\d]+`}.Complie()
}

func Word() ParsingElement {
	return wordElement{}
}

type wordElement struct{}

func (t wordElement) Complie() Parser {
	return regexpElement{`[\w]+`}.Complie()
}

func Sequence(c ...ParsingElement) ParsingElement {
	return sequenceElement(c)
}

type sequenceElement []ParsingElement

func (t sequenceElement) Complie() Parser {
	parser := sequenceParser{}
	var b strings.Builder
	parser.parts = make([]*regexp.Regexp, len(t))
	parser.children = make([]Parser, len(t))
	for i, e := range t {
		parser.children[i] = e.Complie()
		cps := parser.children[i].getRegexpString()
		b.WriteString(cps)
		parser.parts[i] = regexp.MustCompile(cps)
	}
	parser.regexpString = b.String()
	return parser
}

func Switch(c ...ParsingElement) ParsingElement {
	return switchElement(c)
}

type switchElement []ParsingElement

func (t switchElement) Complie() Parser {
	parser := switchParser{}
	var b strings.Builder
	parser.parts = make([]*regexp.Regexp, len(t))
	parser.children = make([]Parser, len(t))
	for i, e := range t {
		parser.children[i] = e.Complie()
		cps := parser.children[i].getRegexpString()
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

func List(element ParsingElement, separator ParsingElement) ParsingElement {
	return listElement{element, separator}
}

type listElement struct {
	element   ParsingElement
	separator ParsingElement
}

func (t listElement) Complie() Parser {
	parser := listParser{}
	parser.parser = t.element.Complie()
	cps := parser.parser.getRegexpString()
	parser.part = regexp.MustCompile(cps)
	sps := t.separator.Complie().getRegexpString()
	parser.separator = regexp.MustCompile(sps)
	parser.regexpString = `(` + cps + `(` + sps + cps + `)*(` + sps + `)?)?`
	return parser
}
