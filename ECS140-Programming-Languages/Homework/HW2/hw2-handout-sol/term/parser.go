package term

import (
	"errors"
	// BEGIN SOLUTION
	"strconv"
	// END SOLUTION
	// HANDOUT: // "strconv"
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

// NewParser creates a struct of a type that satisfies the Parser Interface

func NewParser() Parser {
	// HANDOUT: panic("TODO: implement NewParser")
	// BEGIN_SOLUTION
	return &ParserImpl{
		lex:         nil,
		peekTok:     nil,
		terms:       make(map[string]*Term),
		termID:      make(map[*Term]int),
		termCounter: 0,
	}

	// END_SOLUTION
}

// BEGIN_SOLUTION

// ParserImpl for terms
type ParserImpl struct {
	// Lexer, initialized at each call to Parse.
	lex *lexer
	// Look ahead token, initialized at each call to Parse.
	peekTok *Token
	// Map from string representing a Term to a Term.
	terms map[string]*Term
	// Map from Term to its ID.
	termID map[*Term]int
	// Counter
	termCounter int
}

// nextToken gets the next token either by reading peekTok or
// from the lexer
func (p *ParserImpl) nextToken() (*Token, error) {
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}
	return p.lex.next()
}

// backToken puts back tok.
func (p *ParserImpl) backToken(tok *Token) {
	p.peekTok = tok
}

// Parse a Term
func (p *ParserImpl) Parse(input string) (*Term, error) {
	p.lex = newLexer(input)
	p.peekTok = nil

	// If the input is an empty string
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}

	if tok.typ == tokenEOF {
		return nil, nil
	}

	p.backToken(tok)
	term, err := p.parseNextTerm()
	//term, err := p.termNT() // Table-driven parser
	if err != nil {
		return nil, ErrParser
	}
	// Error if we have not consumed all of the input
	if tok, err := p.nextToken(); err != nil || tok.typ != tokenEOF {
		return nil, ErrParser
	}
	return term, nil
}

// parseNextTerm parses a prefix of the string (via the lexer) into a Term
// returns an error
func (p *ParserImpl) parseNextTerm() (*Term, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, err
	}
	switch tok.typ {
	case tokenEOF:
		return nil, nil
	case tokenNumber:
		return p.mkSimpleTerm(TermNumber, tok.literal), nil
	case tokenVariable:
		return p.mkSimpleTerm(TermVariable, tok.literal), nil
	case tokenAtom:
		a := p.mkSimpleTerm(TermAtom, tok.literal)
		nxt, err := p.nextToken()
		if err != nil {
			return nil, err
		}
		if nxt.typ != tokenLpar {
			// Atom is not the functor for a compound term.
			p.backToken(nxt)
			return a, nil
		}
		// Atom might be the functor of a compound term
		arg, err := p.parseNextTerm()
		if err != nil {
			return nil, err
		}
		// Args of a compound term contains at least one Term.
		args := []*Term{arg}
		nxt, err = p.nextToken()
		if err != nil {
			return nil, err
		}
		// Parse the rest of the arguments, if any
		for ; nxt.typ == tokenComma; nxt, err = p.nextToken() {
			arg, err = p.parseNextTerm()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}
		if nxt.typ != tokenRpar {
			return nil, ErrParser
		}
		return p.mkCompoundTerm(a, args), nil
	default:
		return nil, ErrParser
	}
}

// Table-driven parser.

// -- LL(1) equivalent grammar
// <term>    ::= ATOM <pars> | NUM | VAR
// <pars>    ::= LPAR <args> RPAR | \epsilon
// <args>    ::= <term> <otherargs>
// <otherargs> ::= COMMA <args> | \epsilon

// FIRST(<term>)       = {ATOM,NUM,VAR}
// FIRST(<pars>)       = {LPAR, \epsilon}
// FIRST(<args>)       = {ATOM,NUM,VAR}
// FIRST(<otherargs>)  = {COMMA, \epsilon}

// FOLLOW(<term>)      = {$,COMMA,RPAR}
// FOLLOW(<pars>)      = {$,COMMA,RPAR}
// FOLLOW(<args>)      = {RPAR}
// FOLLOW(<otherargs>) = {RPAR}

/* Parsing Table
                | ATOM                          | NUM / VAR                    | LPAR                    | RPAR                        | COMMA                        | $                    |
<term>          | <term> -> ATOM <pars>         | <term> -> NUM / VAR          | ERROR                   | ERROR                       | ERROR                        | ERROR                |
<pars>          | ERROR                         | ERROR                        | <pars> -> LPAR <args>   | <pars> -> \epsilon          | <pars> -> \epsilon           | <pars> -> \epsilon   |
<args>          | <args> -> <term> <otherargs>  | <args> -> <term> <otherargs> | ERROR                   | ERROR                       | ERROR                        | ERROR                |
<otherargs>     | ERROR                         | ERROR                        | ERROR                   | <otherargs> -> COMMA <args> | <otherargs> -> COMMA <args>  | ERROR                |
*/

// termNT parses the <term> non-terminal
// <term> ::= ATOM <pars> | NUM | VAR
// FIRST(<term>) = {ATOM, NUM,VAR}
// FIRST(<term>) = {$, COMMA, RPAR} (not used)

func (p *ParserImpl) termNT() (*Term, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}

	switch tok.typ {
	// <term> -> NUM
	case tokenNumber:
		return p.mkSimpleTerm(TermNumber, tok.literal), nil

	// <term> -> VAR
	case tokenVariable:
		return p.mkSimpleTerm(TermVariable, tok.literal), nil
	// <term> -> ATOM <pars>
	case tokenAtom:
		functor := p.mkSimpleTerm(TermAtom, tok.literal)
		args, err := p.parsNT()
		if err != nil {
			return nil, ErrParser
		}
		if args != nil {
			return p.mkCompoundTerm(functor, args), nil
		}
		return functor, nil
	default:
		return nil, ErrParser
	}
}

// parsNT parses the <pars> non-terminal.
// <pars> ::= LPAR <args> RPAR | \epsilon
// FIRST(<pars>) = {LPAR, \epsilon}
// FOLLOW(<pars>) = {$, COMMA, RPAR}
func (p *ParserImpl) parsNT() ([]*Term, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}

	switch tok.typ {
	// <pars> -> \epsilon
	case tokenEOF, tokenComma, tokenRpar:
		p.backToken(tok)
		return nil, nil

	// <pars> -> LPAR <args> RPAR
	case tokenLpar:
		args, err := p.argsNT()
		if err != nil {
			return nil, ErrParser
		}
		if tokRpar, err := p.nextToken(); err != nil || tokRpar.typ != tokenRpar {
			// Doing nothing here for 100% code coverage. Because
			// the next token must be 'RPAR', otherwise the last
			// 'p.otherargsNT()' will return 'nil, ErrParser'.
		}
		return args, nil

	default:
		return nil, ErrParser
	}
}

// argsNT parses the <args> non-terminal.
// <args> ::= <term> <otherargs>
// FIRST(<args>) = {ATOM, NUM, VAR} (not used)
// FOLLOw(<args>) = {RPAR}          (not used)
func (p *ParserImpl) argsNT() ([]*Term, error) {
	arg, err := p.termNT()
	if err != nil {
		return nil, ErrParser
	}

	args, err := p.otherargsNT()
	if err != nil {
		return nil, ErrParser
	}

	return append([]*Term{arg}, args...), nil
}

// otherargsNT() parses the <otherargs> non-terminal.
// <otherargs> ::= COMMA <args> | \epsilon
// FIRST(<otherargs>) = {COMMA, \epsilon}
// FOLLOW(<otherargs>) = {RPAR}
func (p *ParserImpl) otherargsNT() ([]*Term, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}

	switch tok.typ {
	// <otherargs> -> \epsilon
	case tokenRpar:
		p.backToken(tok)
		return nil, nil
	// <otherargs> -> COMMA <args>
	case tokenComma:
		return p.argsNT()

	default:
		return nil, ErrParser
	}
}

// Helper functions to make terms.

// mkSimpleTerm makes a simple term.
func (p *ParserImpl) mkSimpleTerm(typ TermType, lit string) *Term {
	key := lit // Use literal as the key for simple terms.
	term, ok := p.terms[key]
	if !ok {
		term = &Term{Typ: typ, Literal: lit}
		p.insertTerm(term, key)
	}

	return term
}

// mkCompoundTerm makes a compound term.
func (p *ParserImpl) mkCompoundTerm(functor *Term, args []*Term) *Term {
	key := strconv.Itoa(p.termID[functor])
	for _, arg := range args {
		key += ", " + strconv.Itoa(p.termID[arg])
	}
	term, ok := p.terms[key]
	if !ok {
		term = &Term{
			Typ:     TermCompound,
			Functor: functor,
			Args:    args,
		}
		p.insertTerm(term, key)
	}
	return term
}

// insertTerm inserts term with given key into the terms and termID maps.
func (p *ParserImpl) insertTerm(term *Term, key string) {
	p.terms[key] = term
	p.termID[term] = p.termCounter
	p.termCounter++
}

// END_SOLUTION
