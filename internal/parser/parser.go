package parser

import (
	pt "github.com/alecthomas/participle/v2"
)

type Parser interface {
	Parse(string) (*ast, error)
}

type parser struct{ _p *pt.Parser[ast] }

func New() (Parser, error) {
	p, err := pt.Build[ast](pt.Unquote("String"), pt.Union[Value](String{}))
	return &parser{p}, err
}

func (p *parser) Parse(input string) (*ast, error) {
	return p._p.ParseString("", input)
}
