package disjointset

import (
	"fmt"
	"hw1/ins_tester"
	"math/rand"
	"testing"
	"regexp"
)

func TestDisjointSetOddEven_Ins(t *testing.T) {
	rand.Seed(1)
	s := NewDisjointSet()
	const N = 1000 * 1000
	// Union all even numbers
	for i := 2; i < N; i += 2 {
		s.UnionSet(i, i-2)
	}
	// Union all odd numbers
	for i := 3; i < N; i += 2 {
		s.UnionSet(i, i-2)
	}
	// Perform N random checks
	for i := 1; i < N; i++ {
		j := rand.Intn(i)
		sameMod := i%2 == j%2
		sameSet := s.FindSet(i) == s.FindSet(j)
		if sameMod != sameSet {
			msg := fmt.Sprintf("In test 0x00000001 (OddEven), expected %d and %d to be in the same set", i, j)
			ins_tester.WriteSecret("ECS140A_ID{0x00000001}", 0, msg)
			return
		}
	}
	ins_tester.WriteSecret("ECS140A_ID{0x00000001}", 1, "")
}

type UnionSetStep struct {
	a, b int
}
type FindSetCheck struct {
	a, b     int
	expected bool
}

var disjointSetTests = map[string]struct {
	unionSetSteps []UnionSetStep
	findSetChecks []FindSetCheck
}{
	"ECS140A_ID{0x00000002}": {
		unionSetSteps: []UnionSetStep{},
		findSetChecks: []FindSetCheck{
			FindSetCheck{1, 1, true},
			FindSetCheck{3, 3, true},
			FindSetCheck{2, 3, false},
			FindSetCheck{6, 7, false},
		},
	},
	"ECS140A_ID{0x00000003}": {
		unionSetSteps: []UnionSetStep{
			UnionSetStep{1, 2},
			UnionSetStep{1, 3},
			UnionSetStep{1, 4},
		},
		findSetChecks: []FindSetCheck{
			FindSetCheck{2, 1, true},
			FindSetCheck{2, 3, true},
			FindSetCheck{2, 5, false},
			FindSetCheck{6, 7, false},
		},
	},
	"ECS140A_ID{0x00000004}":{
		unionSetSteps: []UnionSetStep{
			UnionSetStep{-800001, -800002},
			UnionSetStep{-800002, -800003},
			UnionSetStep{-800003, -800004},
		},
		findSetChecks: []FindSetCheck{
			FindSetCheck{-800002, -800004, true},
			FindSetCheck{-800003, -800001, true},
			FindSetCheck{-800002, -800008, false},
			FindSetCheck{-800003, -800006, false},
		},
	},
	"ECS140A_ID{0x00000005}":{
		unionSetSteps: []UnionSetStep{
			UnionSetStep{100001, 100002},
			UnionSetStep{100002, 100003},
			UnionSetStep{100003, 100001},
			UnionSetStep{100004, 100005},
			UnionSetStep{100005, 100006},
			UnionSetStep{100006, 100004},
		},
		findSetChecks: []FindSetCheck{
			FindSetCheck{100001, 100003, true},
			FindSetCheck{100001, 100006, false},
			FindSetCheck{100005, 100006, true},
			FindSetCheck{100003, 100005, false},
		},
	},
	"ECS140A_ID{0x00000006}":{
		unionSetSteps: []UnionSetStep{
			UnionSetStep{2, 4},
			UnionSetStep{4, 6},
			UnionSetStep{6, 8},
			UnionSetStep{10, 12},
			UnionSetStep{12, 14},
			UnionSetStep{14, 16},
			UnionSetStep{16, 18},
		},
		findSetChecks: []FindSetCheck{
			FindSetCheck{2, 10, false},
			FindSetCheck{2, 8, true},
			FindSetCheck{10, 16, true},
		},
	},
        "ECS140A_ID{0x00000007}":{
                unionSetSteps: []UnionSetStep{
                        UnionSetStep{0, -4},
                        UnionSetStep{-4, 500000},
                        UnionSetStep{-999, 0},
                        UnionSetStep{10, -2},
                        UnionSetStep{10, -100},
                        UnionSetStep{10, 20},
                        UnionSetStep{10, 6000},
                },
                findSetChecks: []FindSetCheck{
                        FindSetCheck{0, 10, false},
                        FindSetCheck{-999, -4, true},
                        FindSetCheck{10, -100, true},
                },
        },
        "ECS140A_ID{0x00000008}":{
                unionSetSteps: []UnionSetStep{
                        UnionSetStep{0, -1},
                        UnionSetStep{0, -2},
                        UnionSetStep{0, -3},
                        UnionSetStep{0, 8},
                        UnionSetStep{0, 9},
                        UnionSetStep{0, 106},
                        UnionSetStep{0, 107},
                },
                findSetChecks: []FindSetCheck{
                        FindSetCheck{0, 108, false},
                        FindSetCheck{-2, 107, true},
                        FindSetCheck{8, 9, true},
                },
        },
        "ECS140A_ID{0x00000009}":{
                unionSetSteps: []UnionSetStep{
                        UnionSetStep{2, 4},
                        UnionSetStep{6, 8},
                        UnionSetStep{10, 12},
                        UnionSetStep{2000, 2002},
                        UnionSetStep{2018, 2020},
                        UnionSetStep{2014, 2016},
                        UnionSetStep{2022, 2024},
                },
                findSetChecks: []FindSetCheck{
                        FindSetCheck{2018, 2014, false},
                        FindSetCheck{0, 0, true},
                        FindSetCheck{12, 10, true},
                },
        },
}

func TestDisjointSet_Ins(t *testing.T) {

	testID := *ins_tester.Test_ID
	re := regexp.MustCompile(`[^\{]*\{([^\}]*)\}`)
	testNo := re.FindStringSubmatch(testID)[1]
	test := disjointSetTests[testID]
	s := NewDisjointSet()
	for _, unionSetStep := range test.unionSetSteps {
		s.UnionSet(unionSetStep.a, unionSetStep.b)
	}
	for _, findSetCheck := range test.findSetChecks {
		if actual := s.FindSet(findSetCheck.a) == s.FindSet(findSetCheck.b); actual != findSetCheck.expected {
			msg := fmt.Sprintf("In test %s, FindSet(%d) == FindSet(%d) gives %t, expected %t",
				testNo, findSetCheck.a, findSetCheck.b, actual, findSetCheck.expected)
			ins_tester.WriteSecret(testID, 0, msg)
			return
		}
	}
	ins_tester.WriteSecret(testID, 1, "")
}
