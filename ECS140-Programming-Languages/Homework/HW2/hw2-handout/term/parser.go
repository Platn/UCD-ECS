package term

import (
	"errors"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <start>    ::= <term> | \epsilon
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

// We will have the string passed in let's go look at the Term type
type ParserImpl struct {
	tknzer *lexer
	pkTkn  *Token // Uppercase instead of lower case
	mapper map[string]*Term
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &ParserImpl{}
}

// Create a new map and pass it back
func NewMap() map[string]*Term {
	m := make(map[string]*Term)
	return m
}

// Go to the next token
func (p *ParserImpl) nextToken() (*Token, error) {
	if tok := p.pkTkn; tok != nil {
		p.pkTkn = nil
		return tok, nil
	}

	tok, _ := p.tknzer.next() // This is the same as nextToken

	return tok, nil
}

// Used to bring back during peek
func (p *ParserImpl) backToken(tok *Token) {
	p.pkTkn = tok
}

func (p *ParserImpl) peekToken() (*Token, error) {
	tok, _ := p.nextToken()

	p.backToken(tok)

	return tok, nil
}

// Implementation of Parse Method
// Parse the non-terminal <start>
// <start> ::= <term> | \epsilon
func (p *ParserImpl) Parse(input string) (*Term, error) {
	if len(input) == 0 {
		var term *Term
		return term, nil
	} // Return a blank term if empty string
	p.tknzer = newLexer(input)
	p.mapper = NewMap()

	term, err := p.termNT()

	if err != nil {
		return nil, ErrParser
	}

	if nextTok, err := p.nextToken(); err != nil || nextTok.typ != tokenEOF {
		return nil, ErrParser // This might only be for first/follow
	} // This means that if we still have tokens left, its a dud
	return term, nil
}

// Parse the non-terminal <term>
// <term> ::= ATOM | NUM | VAR | <compound>
func (p *ParserImpl) termNT() (*Term, error) {
	// <term> -> ATOM | NUM | VAR | <compound>
	// If there are 4 choices, then we need to peek the token
	// tok, err := p.nextToken() // Next token only used when we consume
	tok, _ := p.peekToken()

	switch tok.typ {

	case tokenAtom:
		var fnctTerm *Term // Functor will be used to store the rest of the args

		if retTerm, ok := p.mapper[tok.literal]; ok { // If ATOM exists, pass to functor
			fnctTerm, _ = p.functorNT(retTerm)
		} else {
			var atom = &Term{Typ: TermAtom, Literal: tok.literal, Functor: nil, Args: nil}
			p.mapper[atom.String()] = atom
			fnctTerm, _ = p.functorNT(atom) // Functor is created in advance
		}

		p.nextToken()

		peek, _ := p.peekToken() // Peek holds lPar or invalid

		if peek.typ == tokenEOF || peek.typ == tokenComma || peek.typ == tokenRpar { // ATOM
			if retTerm, ok := p.mapper[tok.literal]; ok { // If it already exists
				return retTerm, nil
			}
		} else if fnctTerm != nil && peek.typ == tokenLpar { // Compound, check left
			_, err := p.compoundNT(fnctTerm)
			if err != nil {
				return nil, ErrParser
			}

			if retTerm, ok := p.mapper[fnctTerm.String()]; ok { // If it already exists
				return retTerm, nil
			} else {

				p.mapper[fnctTerm.String()] = fnctTerm
				return fnctTerm, nil
			}
		}

	case tokenNumber:
		p.nextToken() // Consume the number
		var mapTerm *Term = &Term{Typ: TermNumber, Literal: tok.literal, Functor: nil, Args: nil}
		p.mapper[mapTerm.String()] = mapTerm
		return mapTerm, nil
	case tokenVariable:
		p.nextToken()
		peek, _ := p.peekToken()

		if peek.typ == tokenEOF || peek.typ == tokenRpar || peek.typ == tokenComma {
			if retTerm, ok := p.mapper[tok.literal]; ok { // If it already exists
				return retTerm, nil
			} else {
				var mapTerm *Term = &Term{Typ: TermVariable, Literal: tok.literal, Functor: nil, Args: nil}
				p.mapper[mapTerm.String()] = mapTerm
				return mapTerm, nil
			}
		}

	default:
		return nil, ErrParser
	}

	return nil, ErrParser
}

// <compound> ::= <functor> LPAR <args> RPAR
func (p *ParserImpl) compoundNT(fnctTerm *Term) (*Term, error) {

	p.nextToken()

	arg, err := p.argsNT(fnctTerm)
	if err != nil {
		return nil, ErrParser
	}

	p.nextToken()
	return arg, nil
}

// <functor> ::= ATOM
func (p *ParserImpl) functorNT(atom *Term) (*Term, error) {
	var fnctr *Term

	fnctr = &Term{Typ: TermCompound, Literal: "", Functor: atom, Args: nil}
	return fnctr, nil
}

// <args> ::= <term> | <term> COMMA <args>
func (p *ParserImpl) argsNT(fnctTerm *Term) (*Term, error) { // We need to check ahead here and see if a comma exist
	term, err := p.termNT() // Every term consumes once it enters

	if err != nil {
		return nil, ErrParser
	}

	peek, err := p.peekToken()
	if err != nil || (peek.typ != tokenComma && peek.typ != tokenRpar) {
		return nil, ErrParser
	}

	fnctTerm.Args = append(fnctTerm.Args, term)

	if peek.typ == tokenComma {
		p.nextToken()
		_, err = p.argsNT(fnctTerm)

	}

	return fnctTerm, nil
}
