package unify

import (
	"hw4/term"
	"reflect"
	"testing"
)

// Unification failure due to symbol clash
type SymbolClashTest struct {
	id             int
	input1, input2 string
}

var symbol_clash_tests = []SymbolClashTest{
	{1, "123", "fff(123, XXXXXX, YYYYYY)"},
	{2, "x", "bar(AAAAAA, 456, bbbbbb)"},
	{3, "fff(1, XXXXXX, 1234)", "fff(1, YYYYYY, aaaaaa)"},
	{4, "fff(AAAAAA, BBBBBB, CCCCCC, DDDDDD, EEEEEE, FFFFFF, GGGGGG)", "fff(AAAAAA, BBBBBB, CCCCCC, DDDDDD, EEEEEE, FFFFFF)"},
	{5, "123", "abc"},
	{6, "abc", "def"},
	{7,
		"foo(XXXXXX, p1(1, x, p2(YYYYYY)), p3(2, p4(p5(ZZZ), YYYYYY)))",
		"foo(aaaaaa, p1(1, x, p2(XXXXXX)), p3(2, p4(p5(YYYYYY), bbbbbb)))"},
	{8,
		"foo1(foo2(ZZZ, foo3(AAAAAA, 1, foo4(ZZZ, foo5(99, s1, s2), ggg(1)), s2), XXXXXX))",
		"foo1(foo2(ggg(AAAAAA), foo3(2, 1, foo4(XXXXXX, foo5(99, s1, s2), YYYYYY), s2), YYYYYY))"},
}

func TestUnifyErrorSymbolClashGrading(t *testing.T) {
	for idx, test := range symbol_clash_tests {
		unifier := NewUnifier()
		parser := term.NewParser()
		term1, err := parser.Parse(test.input1)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input1)
			continue
		}
		term2, err := parser.Parse(test.input2)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input2)
			continue
		}
		_, err = unifier.Unify(term1, term2)
		if err == nil {
			t.Errorf("\nin test %d when unifying %#v and %#v:\n\tdid not get error", idx, test.input1, test.input2)
		}
	}
}

// Unification failure due to a cycle (occurs check fails)
type CycleExistsTest struct {
	id             int
	input1, input2 string
}

var cycle_exists_tests = []CycleExistsTest{
	{1, "Var", "func(Var)"},
	{2, "bar(AAA, ggg(BBB), CCC)", "bar(BBB, CCC, AAA)"},
	{3, "aaa(123, ggg(456), b(XXX, 2, XXX), 0, ZZZ)", "XXX"},
	{
		4,
		"fff(AAA, BBB, CCC, DDD, EEE, FFF)",
		"fff(fff(BBB), DDD, AAA, FFF, CCC, EEE)",
	},
	{
		5,
		"fff(fff(AAA), ggg(BBB), hhh(CCC), ppp(DDD), qqq(EEE), rrr(FFF))",
		"fff(CCC, DDD, EEE, FFF, AAA, BBB)",
	},
	{
		6,
		"fff(ggg(FFF), DDD, ppp(DDD), qqq(AAA), FFF, rrr(BBB))",
		"fff(ggg(EEE), hhh(CCC), EEE, BBB, AAA, CCC)",
	},
	{
		7,
		"fff(BBB, CCC, DDD, EEE, FFF, AAA)",
		"fff(ggg(hhh(ppp(qqq(rrr(AAA))))), ggg(hhh(BBB)), ggg(hhh(CCC)), ggg(hhh(ppp(DDD))), ggg(hhh(ppp(qqq(EEE)))), ggg(FFF))",
	},
	{
		8,
		"fff(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(ggg(AAA))))))))))))))))))))))), hhh(BBB, CCC), fff(ggg(DDD, 1)), hhh(AAA, ggg(EEE)))",
		"fff(EEE, hhh(DDD, AAA), fff(ggg(CCC, 1)), hhh(AAA, ggg(BBB)))",
	},
}

func TestUnifyErrorExistsACycleGrading(t *testing.T) {
	for idx, test := range cycle_exists_tests {
		unifier := NewUnifier()
		parser := term.NewParser()
		term1, err := parser.Parse(test.input1)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input1)
			continue
		}
		term2, err := parser.Parse(test.input2)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input2)
			continue
		}
		_, err = unifier.Unify(term1, term2)
		if err == nil {
			t.Errorf("in test %d when unifying %#v and %#v: did not get error", idx, test.input1, test.input2)
		}
	}
}

// Success
type SuccessTest struct {
	id             int
	input1, input2 string
}

var success_tests = []SuccessTest{
	{1, "fff(XXX, 74876156754246, YYY)", "fff(74876156754246, XXX, YYY)"},
	{2, "fff(XXX, YYY)", "fff(ggg(YYY), 2)"},
	{3, "fff(XXX, ggg(XXX), ggg(hhh(XXX)))", "fff(AAA, ggg(BBB), ggg(hhh(CCC)))"},
	{4, "fff(hhh(YYY), YYY)", "fff(XXX, ggg(bar))"},
	{5, "fff(XXX, ggg(aaa))", "fff(ggg(YYY), ggg(YYY))"},
	{6, "fff(XXX, YYY, ZZZ, fff(AAA, fff(fff(BBB), fff(fff(CCC)))))", "fff(2, fff(ZZZ), fff(XXX), fff(BBB, fff(fff(CCC), fff(fff(DDD)))))"},
	{
		7,
		"fff(AAA, BBB, fff(CCC), fff(ggg(DDD, EEE)), fff(ggg(hhh(FFF), fff(AAA, BBB, CCC, DDD, EEE, FFF))), ggg(1))",
		"fff(XXX, AAA, fff(BBB), fff(ggg(CCC, DDD)), fff(ggg(hhh(EEE), fff(XXX, XXX, XXX, XXX, XXX, XXX))), ggg(FFF))",
	},
	{
		8,
		"bar(fff(AAA, BBB, CCC, DDD, EEE, 2))",
		"bar(fff(CCC, AAA, 1, FFF, DDD, EEE))",
	},
}

func TestUnifySuccessGrading(t *testing.T) {
	for idx, test := range success_tests {
		unifier := NewUnifier()
		parser := term.NewParser()
		term1, err := parser.Parse(test.input1)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input1)
			continue
		}
		term2, err := parser.Parse(test.input2)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input2)
			continue
		}
		_, err = unifier.Unify(term1, term2)
		if err != nil {
			t.Errorf("\nin test %d when unifying %#v and %#v:\n\tgot unexpected error", idx, test.input1, test.input2)
			continue
		}
	}
}

// Helper type
type UnifyResultAsStr map[string]string

// ToStringMap converts `UnifyResult` to the string form
func (result UnifyResult) ToStringMap() UnifyResultAsStr {
	stringMap := UnifyResultAsStr{}
	for k, v := range result {
		stringMap[k.String()] = v.String()
	}
	return stringMap
}

// Unique results
type UniqueTest struct {
	id               int
	input1, input2   string
	expectedAsStrMap UnifyResultAsStr
}

var unique_tests = []UniqueTest{
	{1, "2418974654234", "VARRR", UnifyResultAsStr{"VARRR": "2418974654234"}},
	{2, "Var", "foo(aaa)", UnifyResultAsStr{"Var": "foo(aaa)"}},
	{3, "fff(ggg(hhh(ppp(qqq(rrr(VARXXX))))))", "fff(ggg(hhh(ppp(qqq(rrr(x))))))", UnifyResultAsStr{"VARXXX": "x"}},
	{4, "fff(XXX, 1)", "fff(2, YYY)", UnifyResultAsStr{"XXX": "2", "YYY": "1"}},
	{5,
		"foo(fff(XXX, 1), ggg(AAA))", "foo(fff(2, YYY), BBB)",
		UnifyResultAsStr{"XXX": "2", "YYY": "1", "BBB": "ggg(AAA)"}},
	{6, "fff(hhh(ZZZ), YYY)", "fff(XXX, ggg(1))", UnifyResultAsStr{"XXX": "hhh(ZZZ)", "YYY": "ggg(1)"}},
	{7, "aaa(b(c, 2, XXX), 0, ZZZ)", "AAA", UnifyResultAsStr{"AAA": "aaa(b(c, 2, XXX), 0, ZZZ)"}},
	{8,
		"fff(AAA, BBB, CCC, DDD, EEE, fff(XXX, ggg(1)))",
		"fff(1, 2, 3, 4, 5, fff(2, YYY))",
		UnifyResultAsStr{
			"AAA": "1", "BBB": "2", "CCC": "3", "DDD": "4", "EEE": "5", "XXX": "2", "YYY": "ggg(1)",
		},
	},
}

func TestUnifyUniqueGrading(t *testing.T) {
	for idx, test := range unique_tests {
		unifier := NewUnifier()
		parser := term.NewParser()
		term1, err := parser.Parse(test.input1)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input1)
			continue
		}
		term2, err := parser.Parse(test.input2)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input2)
			continue
		}
		actual, err := unifier.Unify(term1, term2)
		if err != nil {
			t.Errorf("\nin test %d when unifying %#v and %#v:\n\tgot unexpected error", idx, test.input1, test.input2)
			continue
		}
		if actualAsStrMap := actual.ToStringMap(); !reflect.DeepEqual(test.expectedAsStrMap, actualAsStrMap) {
			t.Errorf("\nin test %d when unifying %#v and %#v:\n\texpected: %#v\n\tgot     : %#v",
				idx, test.input1, test.input2, test.expectedAsStrMap, actualAsStrMap)
		}
	}
}

// Test case where the MGU is not unique but can be a small number of possibilities
type NotUniqueTest struct {
	id              int
	input1, input2  string
	possibleResults []UnifyResultAsStr
}

var not_unique_tests = []NotUniqueTest{
	{
		1,
		"fff(4156, bar(ggg(XXX)), baz(ggg(ppp(YYY))))",
		"fff(4156, bar(ggg(YYY)), baz(ggg(ppp(XXX))))",
		[]UnifyResultAsStr{
			{"XXX": "YYY"},
			{"YYY": "XXX"},
		},
	},
	{
		2,
		"fff(AAA, fff(hhh(ZZZ), YYY), ggg(aaa))",
		"fff(ggg(BBB), fff(XXX, ggg(1)), ggg(BBB))",
		[]UnifyResultAsStr{
			{"AAA": "ggg(BBB)", "BBB": "aaa", "XXX": "hhh(ZZZ)", "YYY": "ggg(1)"},
			{"AAA": "ggg(aaa)", "BBB": "aaa", "XXX": "hhh(ZZZ)", "YYY": "ggg(1)"},
		},
	},
	{
		3,
		"fff(AAA, CCC, DDD, EEE, fff(XXX, ggg(1)))",
		"fff(BBB, 3, 4, 5, fff(2, YYY))",
		[]UnifyResultAsStr{
			{"AAA": "BBB", "CCC": "3", "DDD": "4", "EEE": "5", "XXX": "2", "YYY": "ggg(1)"},
			{"BBB": "AAA", "CCC": "3", "DDD": "4", "EEE": "5", "XXX": "2", "YYY": "ggg(1)"},
		},
	},
	{
		4,
		"foo(fff(XXX, 1), ggg(AAA), Hwz, Hoq)",
		"foo(fff(2, YYY), BBB, iif, Hwz)",
		[]UnifyResultAsStr{
			{"XXX": "2", "YYY": "1", "BBB": "ggg(AAA)", "Hwz": "iif", "Hoq": "Hwz"},
			{"XXX": "2", "YYY": "1", "BBB": "ggg(AAA)", "Hwz": "iif", "Hoq": "iif"},
		},
	},
}

// Test case where the MGU is not unique but can be a small number of possibilities
func TestUnifyNotUniqueGrading(t *testing.T) {
	for idx, test := range not_unique_tests {
		unifier := NewUnifier()
		parser := term.NewParser()
		term1, err := parser.Parse(test.input1)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input1)
			continue
		}
		term2, err := parser.Parse(test.input2)
		if err != nil {
			t.Errorf("\nin test %d when parsing %#v:\n\tgot unexpected error", idx, test.input2)
			continue
		}
		result, err := unifier.Unify(term1, term2)
		if err != nil {
			t.Errorf("\nin test %d when unifying %#v and %#v:\n\tgot unexpected error", idx, test.input1, test.input2)
			continue
		}

		resultAsStr := result.ToStringMap()
		resultIsCorrect := false
		for _, possibleResult := range test.possibleResults {
			if reflect.DeepEqual(resultAsStr, possibleResult) {
				resultIsCorrect = true
				break
			}
		}

		if !resultIsCorrect {
			t.Errorf("\nin test %d when unifying %#v and %#v:\n\tpossible results are: %#v\n\tgot: %#v",
				idx, test.input1, test.input2, test.possibleResults, resultAsStr)
		}
	}
}
