package term

import (
	"fmt"
	"testing"
	"time"
)

// Invalid terms
type InvalidTermTest struct {
	id     string
	handle int
	input  string
}

var invalid_term_tests = []InvalidTermTest{
	{"ECS140A_ID{INVALID|001}", 1, "func)"},
	{"ECS140A_ID{INVALID|002}", 2, "alpha()"},
	{"ECS140A_ID{INVALID|003}", 3, "bravo(("},
	{"ECS140A_ID{INVALID|004}", 4, "charlie(11)gamma"},
	{"ECS140A_ID{INVALID|005}", 5, ",delta(910)"},
	{"ECS140A_ID{INVALID|006}", 6, "echo(78),"},
	{"ECS140A_ID{INVALID|007}", 7, "foxtort(X"},
	{"ECS140A_ID{INVALID|008}", 8, "(XRAY, 56)"},
	{"ECS140A_ID{INVALID|009}", 9, "Golf, 4)"},
	{"ECS140A_ID{INVALID|010}", 10, ", 3)"},
	{"ECS140A_ID{INVALID|011}", 11, "E(C)"},
	{"ECS140A_ID{INVALID|012}", 12, "123(S)"},
	{"ECS140A_ID{INVALID|013}", 13, "abc x*"},
	{"ECS140A_ID{INVALID|014}", 14, "def(*"},
	{"ECS140A_ID{INVALID|015}", 15, "ghi(D *"},
	{"ECS140A_ID{INVALID|016}", 16, "jkl(C, *"},
	{"ECS140A_ID{INVALID|017}", 17, "banana(,"},
	{"ECS140A_ID{INVALID|018}", 18, "cherimoya pumpkin"},
	{"ECS140A_ID{INVALID|019}", 19, "kiwifruit 17"},
	{"ECS140A_ID{INVALID|020}", 20, "lemon(14 ("},
	{"ECS140A_ID{INVALID|021}", 21, "mango(33 longan"},
	{"ECS140A_ID{INVALID|022}", 22, "cherry(28 183"},
	{"ECS140A_ID{INVALID|023}", 23, "melon(plum) squash"},
	{"ECS140A_ID{INVALID|024}", 24, "equal(x(A, B, C), x(1, jkl, B)"},
	{"ECS140A_ID{INVALID|025}", 25, "equal(x(A, B, C), x(1, jkl, B)), hellow"},
}

func (test InvalidTermTest) GetId() string {
	return test.id
}

func (test InvalidTermTest) GetHandle() int {
	return test.handle
}

func (test InvalidTermTest) RunTest() (bool, error, string) {
	p := NewParser()
	if _, err := p.Parse(test.input); err == nil {
		msg := fmt.Sprintf(
			"Parser did not return an error for an invalid input, \"%s\".",
			test.input)
		return false, nil, msg
	}
	return true, nil, ""
}

type ATest interface {
	GetId() string
	RunTest() (bool, error, string)
}

type Result struct {
	value    int
	feedback string
}

func RunSomeTests(tests []ATest, t *testing.T) {

	for _, test := range tests {
		out := make(chan Result)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					panic_string := fmt.Sprintf("%v", r)
					if panic_string == "TODO: implement this!" {
						out <- Result{0, "FAIL: Presence of original panic TODO"}
					} else {
						out <- Result{0, "FAIL: " + test.GetId() + ": " + panic_string}
					}
				}
			}()

			res, err, msg := test.RunTest()

			if err != nil {
				out <- Result{0, "ERROR: " + test.GetId() + " got unexpected error."}
			} else if res {
				out <- Result{1, ""}
			} else {
				out <- Result{0, "FAIL: " + test.GetId() + ": " + msg}
			}
		}()

		go func() {
			time.Sleep(1 * time.Second)
			out <- Result{0, "FAIL: " + test.GetId() + ": Correctness test timed out"}
		}()

		result := <-out
		if result.value == 0 {
			fmt.Println(result.feedback)
		}
	}

}

func TestParserInvalidTerms_Ins(t *testing.T) {
	ins_tests := make([]ATest, len(invalid_term_tests))

	for i, v := range invalid_term_tests {
		ins_tests[i] = v
	}

	RunSomeTests(ins_tests, t)
}

type termTestGeneratorFunction func() (string, *Term)

type TermTest struct {
	id        string
	handle    int
	generator termTestGeneratorFunction
}

func (test TermTest) GetId() string {
	return test.id
}

func (test TermTest) GetHandle() int {
	return test.handle
}

func (test TermTest) RunTest() (bool, error, string) {
	input, expected := test.generator()
	p := NewParser()
	actual, err := p.Parse(input)
	if err != nil {
		msg := fmt.Sprintf(
			"Parser returned an unexpected error while parsing a valid input, \"%s\". Error msg: %s", input, err)
		return false, nil, msg
	}
	if areIsomorphic, err := checkIsomorphic(expected, actual); !areIsomorphic {
		msg := fmt.Sprintf("Parser returned an incorrect output for a valid input, \"%s\". %s", input, err)
		return false, nil, msg
	}
	return true, nil, ""
}

// Basics
func basicTestEmptyString() (string, *Term) {
	return "", nil
}
func basicTestAtom() (string, *Term) {
	return "camelCase", &Term{Typ: TermAtom, Literal: "camelCase"}
}
func basicTestNumber() (string, *Term) {
	return "815", &Term{Typ: TermNumber, Literal: "815"}
}
func basicTestVariable() (string, *Term) {
	return "VAR", &Term{Typ: TermVariable, Literal: "VAR"}
}
func basicTestVariableUnderscore() (string, *Term) {
	return "_VAR_", &Term{Typ: TermVariable, Literal: "_VAR_"}
}

var basicTermTests = []TermTest{
	{"ECS140A_ID{BASIC|001}", 1, basicTestEmptyString},
	{"ECS140A_ID{BASIC|002}", 2, basicTestAtom},
	{"ECS140A_ID{BASIC|003}", 3, basicTestNumber},
	{"ECS140A_ID{BASIC|004}", 4, basicTestVariable},
	{"ECS140A_ID{BASIC|005}", 5, basicTestVariableUnderscore},
}

func TestParserBasic_Ins(t *testing.T) {
	insTests := make([]ATest, len(basicTermTests))

	for i, v := range basicTermTests {
		insTests[i] = v
	}

	RunSomeTests(insTests, t)
}

// Not shared
func notSharedTestSimpleCompound() (string, *Term) {
	f := &Term{Typ: TermAtom, Literal: "dNSoMv"}
	arg := &Term{Typ: TermVariable, Literal: "FNeAbC"}
	return "dNSoMv(FNeAbC)", &Term{Typ: TermCompound, Functor: f, Args: []*Term{arg}}
}
func notSharedTestCompoundWithMultipleArgs() (string, *Term) {
	f := &Term{Typ: TermAtom, Literal: "gTQ"}
	arg1 := &Term{Typ: TermAtom, Literal: "cdC"}
	arg2 := &Term{Typ: TermNumber, Literal: "670"}
	arg3 := &Term{Typ: TermVariable, Literal: "UVC"}
	return "gTQ(cdC, 670, UVC)", &Term{Typ: TermCompound, Functor: f, Args: []*Term{arg1, arg2, arg3}}
}
func notSharedTestCompoundWithCompoundArg() (string, *Term) {
	foo := &Term{Typ: TermAtom, Literal: "oee"}
	bar := &Term{Typ: TermAtom, Literal: "yRr"}
	barArg := &Term{Typ: TermNumber, Literal: "983"}
	fooArg := &Term{Typ: TermCompound, Functor: bar, Args: []*Term{barArg}}
	return "oee( yRr( 983 ))", &Term{Typ: TermCompound, Functor: foo, Args: []*Term{fooArg}}
}
func notSharedTestCompoundsWithCompoundArgDeep() (string, *Term) {
	f1 := &Term{Typ: TermAtom, Literal: "eRF"}
	f2 := &Term{Typ: TermAtom, Literal: "xEI"}
	f3 := &Term{Typ: TermAtom, Literal: "qfo"}
	f4 := &Term{Typ: TermAtom, Literal: "zoT"}
	f5 := &Term{Typ: TermAtom, Literal: "iTH"}
	f6 := &Term{Typ: TermAtom, Literal: "ocz"}
	f7 := &Term{Typ: TermAtom, Literal: "aTA"}
	f8 := &Term{Typ: TermAtom, Literal: "oLD"}
	f9 := &Term{Typ: TermAtom, Literal: "lSX"}
	f10 := &Term{Typ: TermAtom, Literal: "eiq"}
	arg := &Term{Typ: TermNumber, Literal: "983"}
	return "eRF(xEI(qfo(zoT(iTH(ocz(aTA(oLD(lSX(eiq(983))))))))))", &Term{
		Typ: TermCompound, Functor: f1, Args: []*Term{&Term{
			Typ: TermCompound, Functor: f2, Args: []*Term{&Term{
				Typ: TermCompound, Functor: f3, Args: []*Term{&Term{
					Typ: TermCompound, Functor: f4, Args: []*Term{&Term{
						Typ: TermCompound, Functor: f5, Args: []*Term{&Term{
							Typ: TermCompound, Functor: f6, Args: []*Term{&Term{
								Typ: TermCompound, Functor: f7, Args: []*Term{&Term{
									Typ: TermCompound, Functor: f8, Args: []*Term{&Term{
										Typ: TermCompound, Functor: f9, Args: []*Term{&Term{
											Typ: TermCompound, Functor: f10, Args: []*Term{arg}}}}}}}}}}}}}}}}}}}}
}
func notSharedTestCompoundWithMultipleCompoundArgs() (string, *Term) {
	f := &Term{Typ: TermAtom, Literal: "gTQ"}
	g1 := &Term{Typ: TermAtom, Literal: "oMv"}
	g2 := &Term{Typ: TermAtom, Literal: "eiq"}
	g3 := &Term{Typ: TermAtom, Literal: "nAh"}
	arg1 := &Term{Typ: TermAtom, Literal: "cdC"}
	arg2 := &Term{Typ: TermNumber, Literal: "670"}
	arg3 := &Term{Typ: TermVariable, Literal: "UVC"}
	return "gTQ(oMv(cdC), eiq(670), nAh(UVC))", &Term{Typ: TermCompound, Functor: f, Args: []*Term{
		&Term{Typ: TermCompound, Functor: g1, Args: []*Term{arg1}},
		&Term{Typ: TermCompound, Functor: g2, Args: []*Term{arg2}},
		&Term{Typ: TermCompound, Functor: g3, Args: []*Term{arg3}}}}
}
func notSharedTestCompound6() (string, *Term) {
	return "a(b(X,c(Y,3,e)),S,d(W),9)", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "a"},
		Args: []*Term{
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "b"},
				Args: []*Term{
					&Term{Typ: TermVariable, Literal: "X"},
					&Term{
						Typ:     TermCompound,
						Functor: &Term{Typ: TermAtom, Literal: "c"},
						Args: []*Term{
							&Term{Typ: TermVariable, Literal: "Y"},
							&Term{Typ: TermNumber, Literal: "3"},
							&Term{Typ: TermAtom, Literal: "e"},
						},
					},
				},
			},
			&Term{Typ: TermVariable, Literal: "S"},
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "d"},
				Args: []*Term{
					&Term{Typ: TermVariable, Literal: "W"},
				},
			},
			&Term{Typ: TermNumber, Literal: "9"},
		},
	}
}

func notSharedTestCompound7() (string, *Term) {
	return "foo(A,bar(B,h(C,D)),p(_ARG_))", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "foo"},
		Args: []*Term{
			&Term{Typ: TermVariable, Literal: "A"},
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "bar"},
				Args: []*Term{
					&Term{Typ: TermVariable, Literal: "B"},
					&Term{
						Typ:     TermCompound,
						Functor: &Term{Typ: TermAtom, Literal: "h"},
						Args: []*Term{
							&Term{Typ: TermVariable, Literal: "C"},
							&Term{Typ: TermVariable, Literal: "D"},
						},
					},
				},
			},
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "p"},
				Args: []*Term{
					&Term{Typ: TermVariable, Literal: "_ARG_"},
				},
			},
		},
	}
}


var notSharedTests = []TermTest{
	{"ECS140A_ID{NOTSHARED|001}", 1, notSharedTestSimpleCompound},
	{"ECS140A_ID{NOTSHARED|002}", 2, notSharedTestCompoundWithMultipleArgs},
	{"ECS140A_ID{NOTSHARED|003}", 3, notSharedTestCompoundWithCompoundArg},
	{"ECS140A_ID{NOTSHARED|004}", 4, notSharedTestCompoundsWithCompoundArgDeep},
	{"ECS140A_ID{NOTSHARED|005}", 5, notSharedTestCompoundWithMultipleCompoundArgs},
	{"ECS140A_ID{NOTSHARED|006}", 6, notSharedTestCompound6},
	{"ECS140A_ID{NOTSHARED|007}", 7, notSharedTestCompound7},
}

func TestParserTermNotShared_Ins(t *testing.T) {
	insTests := make([]ATest, len(notSharedTests))

	for i, v := range notSharedTests {
		insTests[i] = v
	}

	RunSomeTests(insTests, t)
}

// Sharing
func termWithSharingTest0() (string, *Term) {
	X := &Term{Typ: TermVariable, Literal: "X"}
	return "rel(X, X)", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "rel"},
		Args: []*Term{
			X,
			X,
		},
	}
}

func termWithSharingTest1() (string, *Term) {
	X := &Term{Typ: TermVariable, Literal: "X"}
	return "foo  ( X ,X, X)  ", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "foo"},
		Args: []*Term{
			X,
			X,
			X,
		},
	}
}

func termWithSharingTest2() (string, *Term) {
	X := &Term{Typ: TermVariable, Literal: "X"}
	return " foo( X, X ,f (X) )", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "foo"},
		Args: []*Term{
			X,
			X,
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "f"},
				Args: []*Term{
					X,
				}},
		},
	}
}

func termWithSharingTest3() (string, *Term) {
	f := &Term{Typ: TermAtom, Literal: "f"}
	X := &Term{Typ: TermVariable, Literal: "X"}
	fX := &Term{Typ: TermCompound, Functor: f, Args: []*Term{X}}
	return "foo ( X, X , X, f(X), f(f (X) ))", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "foo"},
		Args: []*Term{
			X,
			X,
			X,
			fX,
			&Term{
				Typ:     TermCompound,
				Functor: f,
				Args: []*Term{
					fX,
				}},
		},
	}
}

func termWithSharingTest4() (string, *Term) {
	fX := &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "f"},
		Args: []*Term{
			&Term{Typ: TermVariable, Literal: "X"},
		},
	}
	return "rel( f( X ) , f (X) )", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "rel"},
		Args: []*Term{
			fX,
			fX,
		},
	}
}

func termWithSharingTest5() (string, *Term) {
	X := &Term{Typ: TermVariable, Literal: "X"}
	return "append(X, 1, X)", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "append"},
		Args: []*Term{
			X,
			&Term{Typ: TermNumber, Literal: "1"},
			X,
		},
	}
}

// Only share variable
func termWithSharingTest6() (string, *Term) {
	X := &Term{Typ: TermVariable, Literal: "X"}
	return "append(X, 1, 2, X)", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "append"},
		Args: []*Term{
			X,
			&Term{Typ: TermNumber, Literal: "1"},
			&Term{Typ: TermNumber, Literal: "2"},
			X,
		},
	}
}

// Only share number
func termWithSharingTest7() (string, *Term) {
	num1 := &Term{Typ: TermNumber, Literal: "1"}
	return "member(1, list(1, 2, 3))", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "member"},
		Args: []*Term{
			num1,
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "list"},
				Args: []*Term{
					num1,
					&Term{Typ: TermNumber, Literal: "2"},
					&Term{Typ: TermNumber, Literal: "3"},
				},
			},
		},
	}
}

// Only share atoms
func termWithSharingTest8() (string, *Term) {
	a := &Term{Typ: TermAtom, Literal: "a"}
	b := &Term{Typ: TermAtom, Literal: "b"}
	c := &Term{Typ: TermAtom, Literal: "c"}
	d := &Term{Typ: TermAtom, Literal: "d"}
	return "intersection(set(b, c, d), list(a, b, c))", &Term{
		Typ:     TermCompound,
		Functor: &Term{Typ: TermAtom, Literal: "intersection"},
		Args: []*Term{
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "set"},
				Args:    []*Term{b, c, d},
			},
			&Term{
				Typ:     TermCompound,
				Functor: &Term{Typ: TermAtom, Literal: "list"},
				Args:    []*Term{a, b, c},
			},
		},
	}
}

// Shared function names
func termWithSharingTest9() (string, *Term) {
	list := &Term{Typ: TermAtom, Literal: "list"}
	a := &Term{Typ: TermAtom, Literal: "a"}
	b := &Term{Typ: TermAtom, Literal: "b"}
	c := &Term{Typ: TermAtom, Literal: "c"}
	d := &Term{Typ: TermAtom, Literal: "d"}
	return "list(list(b, c, d), list(a, b, c), list)", &Term{
		Typ:     TermCompound,
		Functor: list,
		Args: []*Term{
			&Term{
				Typ:     TermCompound,
				Functor: list,
				Args:    []*Term{b, c, d},
			},
			&Term{
				Typ:     TermCompound,
				Functor: list,
				Args:    []*Term{a, b, c},
			},
			list,
		},
	}
}

func termWithSharingTest10() (string, *Term) {
        gY := &Term{
                Typ:     TermCompound,
                Functor: &Term{Typ: TermAtom, Literal: "g"},
                Args: []*Term{
                       &Term{Typ: TermVariable, Literal: "Y"},
                },
        }
	d := &Term{Typ: TermAtom, Literal: "d"}
        return "try( g( Y ) , g (Y), d, g(Y) )", &Term{
                Typ:     TermCompound,
                Functor: &Term{Typ: TermAtom, Literal: "try"},
                Args: []*Term{
                        gY,
                        gY,
			d,
			gY,
                },
        }
}


var share_term_tests = []TermTest{
	{"ECS140A_ID{SHARED|001}", 1, termWithSharingTest0},
	{"ECS140A_ID{SHARED|002}", 2, termWithSharingTest1},
	{"ECS140A_ID{SHARED|003}", 3, termWithSharingTest2},
	{"ECS140A_ID{SHARED|004}", 4, termWithSharingTest3},
	{"ECS140A_ID{SHARED|005}", 5, termWithSharingTest4},
	{"ECS140A_ID{SHARED|006}", 6, termWithSharingTest5},
	{"ECS140A_ID{SHARED|007}", 7, termWithSharingTest6},
	{"ECS140A_ID{SHARED|008}", 8, termWithSharingTest7},
	{"ECS140A_ID{SHARED|009}", 9, termWithSharingTest8},
	{"ECS140A_ID{SHARED|010}", 10, termWithSharingTest9},
	{"ECS140A_ID{SHARED|011}", 11, termWithSharingTest10},
}

func TestParserTermShared_Ins(t *testing.T) {
	ins_tests := make([]ATest, len(share_term_tests))

	for i, v := range share_term_tests {
		ins_tests[i] = v
	}

	RunSomeTests(ins_tests, t)
}

func checkIsomorphic(expected, actual *Term) (bool, error) {
	matchTerms := make(map[*Term]*Term)
	return checkTermIsomorphic(expected, actual, matchTerms)
}

func checkTermIsomorphic(expected, actual *Term, matchTerms map[*Term]*Term) (bool, error) {
	if x, ok := matchTerms[expected]; ok {
		if x == actual {
			return true, nil
		}
		return false, fmt.Errorf(
			"\nerror:\n|\tthe subterm:\n|\t\t\"%s\" (%#v(%p))\n|\tin the expected term matches more than one terms:\n|\t\t\"%s\" (%#v(%p))\n|\t\t\"%s\" (%#v(%p))\n|\tin the actual term",
			expected, expected, expected,
			x, x, x,
			actual, actual, actual)
	}
	if expected != actual {
		if (expected == nil || actual == nil) ||
			(expected.Typ != actual.Typ) ||
			(expected.Literal != actual.Literal) {
			return false, fmt.Errorf(
				"\nerror:\n|\texpected\n|\t\t\"%s\" (%#v(%p))\n|\tgot\n|\t\t\"%s\" (%#v(%p))",
				expected, expected, expected,
				actual, actual, actual)
		}
		if areIsomorphic, err := checkTermIsomorphic(expected.Functor, actual.Functor, matchTerms); !areIsomorphic {
			return false, fmt.Errorf(
				"\ncontext:\n|\tin the functor of\n|\t\t\"%s\" (%#v(%p))\n|\tand\n|\t\t\"%s\" (%#v(%p))%s",
				expected, expected, expected,
				actual, actual, actual,
				err)
		}
		if areIsomorphic, err := checkTermSliceIsomorphic(expected.Args, actual.Args, matchTerms); !areIsomorphic {
			return false, fmt.Errorf(
				"\ncontext:\n|\tin the arguments of\n|\t\t\"%s\" (%#v(%p))\n|\tand\n|\t\t\"%s\" (%#v(%p))%s",
				expected, expected, expected,
				actual, actual, actual,
				err)
		}
	}
	matchTerms[expected] = actual
	return true, nil
}

func checkTermSliceIsomorphic(expectedSlice, actualSlice []*Term, matchTerms map[*Term]*Term) (bool, error) {
	if (expectedSlice == nil && actualSlice != nil) ||
		(expectedSlice != nil && actualSlice == nil) ||
		(len(expectedSlice) != len(actualSlice)) {
		return false, fmt.Errorf(
			"\nerror:\n|\texpected:\n|\t\t\"(%s)\" (%#v(%p))\n|\tgot:\n|\t\t\"(%s)\" (%#v(%p))",
			TermSliceToString(expectedSlice), expectedSlice, expectedSlice,
			TermSliceToString(actualSlice), actualSlice, actualSlice)
	}
	for idx := range expectedSlice {
		if areIsomorphic, err := checkTermIsomorphic(expectedSlice[idx], actualSlice[idx], matchTerms); !areIsomorphic {
			return false, fmt.Errorf("\n|\tin the %d-th argument:%s", idx+1, err)
		}
	}
	return true, nil
}
