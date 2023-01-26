package sexpr

import (
	"testing"
)

func TestParserInvalidGrading(t *testing.T) {
	for _, test := range []struct {
		id    string
		input string
	}{
		{"ECS140A_ID{ParserInvalid|01}", ""},
		{"ECS140A_ID{ParserInvalid|02}", "("},
		{"ECS140A_ID{ParserInvalid|03}", "'"},
		{"ECS140A_ID{ParserInvalid|04}", ")"},
		{"ECS140A_ID{ParserInvalid|05}", "x)"},
		{"ECS140A_ID{ParserInvalid|06}", "( ) ( () ) ()"},
		{"ECS140A_ID{ParserInvalid|07}", "((a( . (()) . () . ())"},
		{"ECS140A_ID{ParserInvalid|08}", "((1 ."},
		{"ECS140A_ID{ParserInvalid|09}", "(123"},
		{"ECS140A_ID{ParserInvalid|10}", ".123"},
	} {
		_, err := NewParser().Parse(test.input)
		if err == nil {
			t.Errorf("\nin test %s\nshould error", test.id)
		}
	}
}

func TestParseValidGrading(t *testing.T) {
	for _, test := range []struct {
		id       string
		input    string
		expected string
	}{
		// basic
		{"ECS140A_ID{ParserBasic|01}", "()", "NIL"},
		{"ECS140A_ID{ParserBasic|02}", "abc", "ABC"},
		{"ECS140A_ID{ParserBasic|03}", "0", "0"},
		{"ECS140A_ID{ParserBasic|04}", "123", "123"},
		{"ECS140A_ID{ParserBasic|05}", "-0123", "-123"},
		// proper list
		{"ECS140A_ID{ParserProperList|01}", "(())", "(NIL . NIL)"},
		{"ECS140A_ID{ParserProperList|02}", "(a)", "(A . NIL)"},
		{"ECS140A_ID{ParserProperList|03}", "((a))", "((A . NIL) . NIL)"},
		{"ECS140A_ID{ParserProperList|04}", "( 1 2 3 4 5 6 7 a b c  d  e)", "(1 . (2 . (3 . (4 . (5 . (6 . (7 . (A . (B . (C . (D . (E . NIL))))))))))))"},
		{"ECS140A_ID{ParserProperList|05}", "(  (1 2 (3 4) 5) ((6) 7 (8 (9 (A (B (C (D (E (F (G(H(I(J(K(L))))))))))))) 10) 11) 12)", "((1 . (2 . ((3 . (4 . NIL)) . (5 . NIL)))) . (((6 . NIL) . (7 . ((8 . ((9 . ((A . ((B . ((C . ((D . ((E . ((F . ((G . ((H . ((I . ((J . ((K . ((L . NIL) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . NIL)) . (10 . NIL))) . (11 . NIL)))) . (12 . NIL)))"},
		{"ECS140A_ID{ParserProperList|06}", "(a b (c d))", "(A . (B . ((C . (D . NIL)) . NIL)))"},
		{"ECS140A_ID{ParserProperList|07}", "(a (()) () 1 (((()))) a)", "(A . ((NIL . NIL) . (NIL . (1 . ((((NIL . NIL) . NIL) . NIL) . (A . NIL))))))"},
		// dotted list
		{"ECS140A_ID{ParserDottedList|01}", "(1 (b . c))", "(1 . ((B . C) . NIL))"},
		{"ECS140A_ID{ParserDottedList|02}", "(1 . (b . (c . d)))", "(1 . (B . (C . D)))"},
		{"ECS140A_ID{ParserDottedList|03}", "(1 . ((b . c) . d))", "(1 . ((B . C) . D))"},
		{"ECS140A_ID{ParserDottedList|04}", "(1 (b c) d . e)", "(1 . ((B . (C . NIL)) . (D . E)))"},
		{"ECS140A_ID{ParserDottedList|05}", "(1 b (c d) . e)", "(1 . (B . ((C . (D . NIL)) . E)))"},
		{"ECS140A_ID{ParserDottedList|06}", "(1 b c . (d e))", "(1 . (B . (C . (D . (E . NIL)))))"},
		// quote
		{"ECS140A_ID{ParserQUOTE|01}", "'(a b c)", "(QUOTE . ((A . (B . (C . NIL))) . NIL))"}, // TODO
		{"ECS140A_ID{ParserQUOTE|02}", "'(1 . 2)", "(QUOTE . ((1 . 2) . NIL))"},
		{"ECS140A_ID{ParserQUOTE|03}", "(quote . (1 . 2))", "(QUOTE . (1 . 2))"},
		{"ECS140A_ID{ParserQUOTE|04}", "'42", "(QUOTE . (42 . NIL))"},
		{"ECS140A_ID{ParserQUOTE|05}", "'(42)", "(QUOTE . ((42 . NIL) . NIL))"},
		{"ECS140A_ID{ParserQUOTE|06}", "''42", "(QUOTE . ((QUOTE . (42 . NIL)) . NIL))"},
		{"ECS140A_ID{ParserQUOTE|07}", "''(42)", "(QUOTE . ((QUOTE . ((42 . NIL) . NIL)) . NIL))"},
		// mix
		{"ECS140A_ID{ParserMixed|01}", "(' 1 '' 2 a 3 'b '  c)", "((QUOTE . (1 . NIL)) . ((QUOTE . ((QUOTE . (2 . NIL)) . NIL)) . (A . (3 . ((QUOTE . (B . NIL)) . ((QUOTE . (C . NIL)) . NIL))))))"},
		{"ECS140A_ID{ParserMixed|02}", "(1 b c '2 3 4 ('( 5 6 . '7) 8 . 9). (d . e))", "(1 . (B . (C . ((QUOTE . (2 . NIL)) . (3 . (4 . (((QUOTE . ((5 . (6 . (QUOTE . (7 . NIL)))) . NIL)) . (8 . 9)) . (D . E))))))))"},
		{"ECS140A_ID{ParserMixed|03}", "(1 . ( ( '() ) . ( ( () . ()) . 1)))", "(1 . (((QUOTE . (NIL . NIL)) . NIL) . ((NIL . NIL) . 1)))"},
	} {
		actual, _ := NewParser().Parse(test.input)
		if actual.SExprString() != test.expected {
			t.Errorf("\nerror: in test %s (\"%s\"):\n\texpected: %s\n\tgot      %s",
				test.id, test.input, test.expected, actual.SExprString())
		}
	}
}
