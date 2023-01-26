package sexpr

import (
	"errors"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

// <sexpr>       ::= <atom> | <pars> | QUOTE <sexpr>
// <atom>        ::= NUMBER | SYMBOL
// <pars>        ::= LPAR <dotted_list> RPAR | LPAR <proper_list> RPAR
// <dotted_list> ::= <proper_list> <sexpr> DOT <sexpr>
// <proper_list> ::= <sexpr> <proper_list> | \epsilon
type Parser interface {
	Parse(string) (*SExpr, error)
}

type ParserImpl struct {
	tknzer *lexer
	pkTkn  *token
}

func (p *ParserImpl) nextToken() (*token, error) {
	if tok := p.pkTkn; tok != nil {
		p.pkTkn = nil
		return tok, nil
	}

	tok, _ := p.tknzer.next() // This is the same as nextToken

	return tok, nil
}

// Used to bring back during peek
func (p *ParserImpl) backToken(tok *token) {
	p.pkTkn = tok
}

func (p *ParserImpl) peekToken() (*token, error) {
	tok, _ := p.nextToken()

	p.backToken(tok)

	return tok, nil
}

func (p ParserImpl) Parse(input string) (*SExpr, error) {

	return nil, ErrParser // Placeholder
}

func NewParser() Parser {
	return &ParserImpl{}
}
