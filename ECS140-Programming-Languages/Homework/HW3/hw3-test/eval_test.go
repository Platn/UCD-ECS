package sexpr

import (
	"testing"
)

func TestEvalInvalid(t *testing.T) {
	for _, test := range []struct {
		id    string
		input string
	}{
		{"ECS140A_ID{EvalInvalid|01}", "(42)"},
		{"ECS140A_ID{EvalInvalid|02}", "(QUOTE)"},
		{"ECS140A_ID{EvalInvalid|03}", "(QUOTE 1 2 3 4 5 6 7)"},
		{"ECS140A_ID{EvalInvalid|04}", "(QUOTE . 2)"},
		{"ECS140A_ID{EvalInvalid|05}", "(QUOTE . (2 . 'NIL))"},
		{"ECS140A_ID{EvalInvalid|06}", "(CAR)"},
		{"ECS140A_ID{EvalInvalid|07}", "(CAR '(1 2) '(1 2 3))"},
		{"ECS140A_ID{EvalInvalid|08}", "(CDR)"},
		{"ECS140A_ID{EvalInvalid|09}", "(CDR '(1 2) '(1 2 3))"},
		{"ECS140A_ID{EvalInvalid|10}", "(CONS)"},
		{"ECS140A_ID{EvalInvalid|11}", "(CONS 1 2 3 4 5 6 7)"},
		{"ECS140A_ID{EvalInvalid|12}", "(LENGTH)"},
		{"ECS140A_ID{EvalInvalid|13}", "(LENGTH 6 '(1 '(2)) '(3 4) 5 '(5))"},
		{"ECS140A_ID{EvalInvalid|14}", "(LENGTH 1 2 3 4 5 6 7 8 9)"},
		{"ECS140A_ID{EvalInvalid|15}", "(LENGTH (CONS 1 2))"},
		{"ECS140A_ID{EvalInvalid|16}", "(+ '(1 2 3))"},
		{"ECS140A_ID{EvalInvalid|17}", "(* '(1 2 3))"},
		{"ECS140A_ID{EvalInvalid|18}", "(ATOM)"},
		{"ECS140A_ID{EvalInvalid|19}", "(ATOM 1 2 3)"},
		{"ECS140A_ID{EvalInvalid|20}", "(ZEROP)"},
		{"ECS140A_ID{EvalInvalid|21}", "(ZEROP 1 2 3)"},
		{"ECS140A_ID{EvalInvalid|22}", "(ZEROP '())"},
		{"ECS140A_ID{EvalInvalid|23}", "(ZEROP '(1))"},
		{"ECS140A_ID{EvalInvalid|24}", "(LISTP)"},
		{"ECS140A_ID{EvalInvalid|25}", "(LISTP 1 2 3)"},
		{"ECS140A_ID{EvalInvalid|26}", "*"},
		{"ECS140A_ID{EvalInvalid|27}", "+"},
	} {
		p := NewParser()
		sexpr, err := p.Parse(test.input)
		if err != nil {
			t.Errorf("\nin test %s (\"%s\"):\nunexpected parse error", test.id, test.input)
			continue
		}
		_, err = sexpr.Eval()
		if err == nil {
			t.Errorf("\nin test %s (\"%s\"):\n\terror: should get an eval error", test.id, test.input)
		}
	}
}
func TestEvalQUOTE(t *testing.T) {
	for _, test := range []struct {
		id       string
		input    string
		expected string
	}{
		// QUOTE
		{"ECS140A_ID{EvalQUOTE|01}", "'42", "42"},
		{"ECS140A_ID{EvalQUOTE|02}", "''42", "(QUOTE . (42 . NIL))"},
		{"ECS140A_ID{EvalQUOTE|03}", "'(42)", "(42 . NIL)"},
		{"ECS140A_ID{EvalQUOTE|04}", "''(42)", "(QUOTE . ((42 . NIL) . NIL))"},
		{"ECS140A_ID{EvalQUOTE|05}", "(QUOTE (42))", "(42 . NIL)"},
		{"ECS140A_ID{EvalQUOTE|06}", "(QUOTE . (42))", "42"},
		{"ECS140A_ID{EvalQUOTE|07}", "(QUOTE . ('NIL . NIL))", "(QUOTE . (NIL . NIL))"},
		{"ECS140A_ID{EvalQUOTE|08}", "(QUOTE . (('5 . 6) . NIL))", "((QUOTE . (5 . NIL)) . 6)"},
		// Numbers
		{"ECS140A_ID{EvalNumber|01}", "43546", "43546"},
		{"ECS140A_ID{EvalNumber|02}", "-43546", "-43546"},
		{"ECS140A_ID{EvalNumber|03}", "-184467440737095516150000000000000000000000000000000", "-184467440737095516150000000000000000000000000000000"},
		// CAR
		{"ECS140A_ID{EvalCAR|01}", "(CAR NIL)", "NIL"},
		{"ECS140A_ID{EvalCAR|02}", "(CAR '((CONS 51 52) . 53))", "(CONS . (51 . (52 . NIL)))"},
		{"ECS140A_ID{EvalCAR|03}", "(CAR '((CAR '((CAR (CONS (LISTP (LENGTH '(51 (+ 52 53 54)))) 53)) . 52))))", "(CAR . ((QUOTE . (((CAR . ((CONS . ((LISTP . ((LENGTH . ((QUOTE . ((51 . ((+ . (52 . (53 . (54 . NIL)))) . NIL)) . NIL)) . NIL)) . NIL)) . (53 . NIL))) . NIL)) . 52) . NIL)) . NIL))"},
		// CDR
		{"ECS140A_ID{EvalCDR|01}", "(CDR NIL)", "NIL"},
		{"ECS140A_ID{EvalCDR|02}", "(CDR '(51 52))", "(52 . NIL)"},
		{"ECS140A_ID{EvalCDR|03}", "(CDR '(51 . 52))", "52"},
		{"ECS140A_ID{EvalCDR|04}", "(CDR '''(51 52))", "((QUOTE . ((51 . (52 . NIL)) . NIL)) . NIL)"},
		// CONS
		{"ECS140A_ID{EvalCONS|01}", "(CONS NIL NIL)", "(NIL . NIL)"},
		{"ECS140A_ID{EvalCONS|02}", "(CONS (CDR NIL) (CAR NIL))", "(NIL . NIL)"},
		{"ECS140A_ID{EvalCONS|03}", "(CONS 'QUOTE 51)", "(QUOTE . 51)"},
		// LENGTH
		{"ECS140A_ID{EvalLENGTH|01}", "(LENGTH '())", "0"},
		{"ECS140A_ID{EvalLENGTH|02}", "(LENGTH '(51 52 53 54 55 56 57))", "7"},
		{"ECS140A_ID{EvalLENGTH|03}", "(LENGTH (CONS (+ 1 (* +2 -3)) '(+ 4 5 (+ 6 7 (* 8 9)))))", "5"},
		// +
		{"ECS140A_ID{EvalSUM|01}", "(+)", "0"},
		{"ECS140A_ID{EvalSUM|02}", "(+ 51)", "51"},
		{"ECS140A_ID{EvalSUM|03}", "(+ 51 52)", "103"},
		{"ECS140A_ID{EvalSUM|04}", "(+ (+ 1 2) (+ 3 4) (+ 5 6))", "21"},
		{"ECS140A_ID{EvalSUM|05}", "(+ (+ (+ (+ (+ 1 2) (+ 3 4)) 5) (+ 6 (+ 7 -1)) (+ (+ -8 -2) (+ -9 10))))", "18"},
		{"ECS140A_ID{EvalSUM|06}", "(+ -18446744073709551615 -18446744073709551615)", "-36893488147419103230"},
		// *
		{"ECS140A_ID{EvalPRODUCT|01}", "(*)", "1"},
		{"ECS140A_ID{EvalPRODUCT|02}", "(* 0)", "0"},
		{"ECS140A_ID{EvalPRODUCT|03}", "(* 51 52)", "2652"},
		{"ECS140A_ID{EvalPRODUCT|04}", "(* (* 1 -2) (* 3 -4) (* -5 6))", "-720"},
		{"ECS140A_ID{EvalPRODUCT|05}", "(* (* (* (* (* 1 2) (* 3 4)) 5) (* 6 (* 7 -1)) (* (* -8 -2) (* -9 10))))", "7257600"},
		{"ECS140A_ID{EvalPRODUCT|06}", "(* -18446744073709551615 18446744073709551615)", "-340282366920938463426481119284349108225"},
		// ATOM
		{"ECS140A_ID{EvalATOM|01}", "(ATOM ())", "T"},
		{"ECS140A_ID{EvalATOM|02}", "(ATOM 'some-atom)", "T"},
		{"ECS140A_ID{EvalATOM|03}", "(ATOM (+ 51 597))", "T"},
		{"ECS140A_ID{EvalATOM|04}", "(ATOM '51)", "T"},
		{"ECS140A_ID{EvalATOM|05}", "(ATOM ''51)", "NIL"},
		// LISTP
		{"ECS140A_ID{EvalLISTP|01}", "(LISTP NIL)", "T"},
		{"ECS140A_ID{EvalLISTP|02}", "(LISTP '(NIL))", "T"},
		{"ECS140A_ID{EvalLISTP|03}", "(LISTP '(51 . 52))", "T"},
		{"ECS140A_ID{EvalLISTP|04}", "(LISTP '51)", "NIL"},
		{"ECS140A_ID{EvalLISTP|05}", "(LISTP ''51)", "T"},
		{"ECS140A_ID{EvalLISTP|06}", "(LISTP '(+ (+ (+ (+ (+ 1 2) (+ 3 4)) 5) (+ 6 (+ 7 -1)) (+ (+ -8 -2) (+ -9 10)))))", "T"},
		// ZEROP
		{"ECS140A_ID{EvalZEROP|01}", "(ZEROP 0)", "T"},
		{"ECS140A_ID{EvalZEROP|02}", "(ZEROP (LENGTH 'NIL))", "T"},
		{"ECS140A_ID{EvalZEROP|03}", "(ZEROP (+ 1953 -9863 (* 9657)))", "NIL"},
		// Mixed
		{"ECS140A_ID{EvalMixed|01}", "(CONS 51 ''52)", "(51 . (QUOTE . (52 . NIL)))"},
		{"ECS140A_ID{EvalMixed|02}", "(ATOM (CDR '(51 . 52)))", "T"},
		{"ECS140A_ID{EvalMixed|03}", "(ATOM (CDR '(51 52)))", "NIL"},
		{"ECS140A_ID{EvalMixed|04}", "(LISTP (CONS 51 52))", "T"},
		{"ECS140A_ID{EvalMixed|05}", "(ZEROP (+ (LENGTH '(7815215)) -1))", "T"},
		{"ECS140A_ID{EvalMixed|06}", "(CDR (CONS 53 (CONS NIL (+ 10000000 -22222 (LENGTH '(+ 52 53 54))))))", "(NIL . 9977782)"},
		{"ECS140A_ID{EvalMixed|07}", "(CAR (CONS '(51 52 53) '(CDR 52)))", "(51 . (52 . (53 . NIL)))"},
		{"ECS140A_ID{EvalMixed|08}", "(CAR (CONS (CONS (LISTP (LENGTH '(51 (+ 52 53 54)))) (CAR '(+ 53))) NIL))", "(NIL . +)"},
		{"ECS140A_ID{EvalMixed|09}", "(CONS (+ 5 (CAR '(32 . 33)) (CAR (CDR '(142222 9876)))) (* (LENGTH '(LENGTH 51 52 53)) -12786415))", "(9913 . -51145660)"},
	} {
		p := NewParser()
		sexpr, err := p.Parse(test.input)
		if err != nil {
			t.Errorf("\nin test %s (\"%s\"):\nunexpected parse error", test.id, test.input)
			continue
		}
		actual, err := sexpr.Eval()
		if err != nil {
			t.Errorf("\nin test %s (\"%s\"):\nunexpected eval error", test.id, test.input)
		} else if actual.SExprString() != test.expected {
			t.Errorf("\nin test %s (\"%s\"):\nerror:\tgot\t\t\t\"%s\"\n\t\texpected\t\"%s\"",
				test.id, test.input, actual.SExprString(), test.expected)
		}
	}
}
