package sexpr

import (
	"errors"
	// "math/big"
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

func NewParser() Parser {
	return &ParserImpl{}
}

func (p *ParserImpl) Parse(input string) (*SExpr, error) {

	if len(input) == 0 { //
		return nil, ErrParser
	}
	p.tknzer = newLexer(input)

	newExpr, err := p.sexprNT()
	if err != nil {
		return nil, ErrParser
	}

	if nextTok, err := p.nextToken(); err != nil || nextTok.typ != tokenEOF {
		return nil, ErrParser
	}

	return newExpr, nil
}

func (p *ParserImpl) sexprNT() (*SExpr, error) {
	tmpTkn, err := p.peekToken()

	var newExpr *SExpr
	newExpr = mkNil()
	switch tmpTkn.typ {

	case tokenNumber:
		newExpr = mkNumber(tmpTkn.num)

		p.nextToken()

	case tokenSymbol:

		newExpr = mkSymbol(tmpTkn.literal)

		p.nextToken()

	case tokenLpar: // <pars>
		newExpr, err = p.parsNT()
		if err != nil {
			return nil, ErrParser
		}

	case tokenRpar:
		newExpr = mkNil()
		return newExpr, nil
	case tokenQuote: // QUOTE <sexprs>
		p.nextToken()
		qExpr := mkSymbol("QUOTE")
		tmp, err := p.sexprNT()
		newExpr = mkConsCell(qExpr, mkConsCell(tmp, mkNil()))

		if err != nil {
			return nil, ErrParser
		}

	case tokenEOF:
		return nil, ErrParser
	}

	return newExpr, nil
}

func (p *ParserImpl) parsNT() (*SExpr, error) {
	p.nextToken()
	var newExprs *SExpr
	newExprs, err := p.pListNT()

	if err != nil {
		return nil, ErrParser
	}

	p.nextToken()

	return newExprs, nil
}

func (p *ParserImpl) pListNT() (*SExpr, error) {
	peek, err := p.peekToken()

	if peek.typ == tokenEOF {
		return nil, ErrParser
	} else if peek.typ == tokenRpar {
		return mkNil(), nil
	} else if peek.typ == tokenDot {
		p.nextToken()
		dotExpr, err := p.sexprNT()
		if err != nil {
			return nil, ErrParser
		}

		return dotExpr, nil
	}

	sexpr, err := p.sexprNT()
	if err != nil {
		return nil, ErrParser
	}

	plist, err := p.pListNT()
	if err != nil {
		return nil, ErrParser
	}

	return mkConsCell(sexpr, plist), nil
}
