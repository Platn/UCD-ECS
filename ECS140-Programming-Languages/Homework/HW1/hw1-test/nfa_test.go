package nfa

import (
	"fmt"
	"hw1/ins_tester"
	"testing"
)

// 0 --a-->1--b-->2--c-->3--d-->6
// \---a-->4--b-->5--------c---/

func dagTransitions(st state, sym rune) []state {

	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{1, 4},
		},
		1: map[rune][]state{
			'b': []state{2},
		},
		2: map[rune][]state{
			'c': []state{3},
		},
		3: map[rune][]state{
			'd': []state{6},
		},
		4: map[rune][]state{
			'b': []state{5},
		},
		5: map[rune][]state{
			'c': []state{6},
		},
	}[st][sym]
}

// 1--x-->2
// 1--x-->3
// 1--y-->3
// 2--y-->1
func expTransitions(st state, sym rune) []state {

	return map[state]map[rune][]state{
		1: map[rune][]state{
			'x': []state{2, 3},
			'y': []state{3},
		},
		2: map[rune][]state{
			'y': []state{1},
		},
	}[st][sym]
}

type Test struct {
	id           string
	nfa          string
	start, final state
	input        []rune
	expected     bool
}

var tests = []Test{
	{"ECS140A_ID{0x00000001}", "dagTransitions", 0, 6, []rune{'a', 'b', 'c'}, true},
	{"ECS140A_ID{0x00000002}", "dagTransitions", 0, 6, []rune{'a', 'b', 'c', 'd'}, true},
	{"ECS140A_ID{0x00000003}", "dagTransitions", 0, 6, []rune{'a', 'b', 'c', 'c'}, false},
	{"ECS140A_ID{0x00000004}", "dagTransitions", 0, 6, []rune{'a', 'b'}, false},

	{"ECS140A_ID{0x00000005}", "expTransitions", 1, 1, nil, true},
	{"ECS140A_ID{0x00000006}", "expTransitions", 1, 2, []rune{'x', 'y', 'x', 'y', 'x', 'y', 'x', 'y', 'x', 'y', 'x'}, true},
	{"ECS140A_ID{0x00000007}", "expTransitions", 1, 3, []rune{'x', 'y', 'x', 'y', 'x', 'y', 'x', 'y', 'x', 'y', 'x'}, true},
	{"ECS140A_ID{0x00000008}", "expTransitions", 1, 3, []rune{'x', 'y', 'x', 'y', 'x', 'y', 'x', 'y', 'x', 'y', 'x', 'x'}, false},
}

func (test Test) RunTest() (bool, error, string) {

	nfas := map[string]TransitionFunction{
		"dagTransitions": dagTransitions,
		"expTransitions": expTransitions,
	}

	actual := Reachable(nfas[test.nfa], test.start, test.final, test.input)
	res := actual == test.expected

	var msg string
	if !res {
		msg = fmt.Sprintf("Reachable (%s, %d, %d, %v)= %t ; Expected: %t",
			test.nfa, test.start, test.final, string(test.input), actual, test.expected)
	}

	return res, nil, msg
}

func (test Test) GetId() string {
	return test.id
}

func TestReachable_Ins(t *testing.T) {

	ins_tests := make([]ins_tester.ATest, len(tests))

	for i, v := range tests {
		ins_tests[i] = v
	}

	ins_tester.RunSomeTests(ins_tests, t)
}
