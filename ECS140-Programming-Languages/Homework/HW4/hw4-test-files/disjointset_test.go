package disjointset

import (
	"math/rand"
	"testing"
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
			t.Errorf("In ECS140A_ID{0x00000001}, expected %d and %d to be in the same set", i, j)
			return
		}
	}
}

type UnionSetStepGrading struct {
	a, b int
}
type FindSetCheckGrading struct {
	a, b     int
	expected bool
}

var disjointSetGradingTests = []struct {
	id            string
	UnionSetStepGradings []UnionSetStepGrading
	FindSetCheckGradings []FindSetCheckGrading
}{
	{
		id: "ECS140A_ID{0x00000002}",
		UnionSetStepGradings: []UnionSetStepGrading{},
		FindSetCheckGradings: []FindSetCheckGrading{
			FindSetCheckGrading{1, 1, true},
			FindSetCheckGrading{3, 3, true},
			FindSetCheckGrading{2, 3, false},
			FindSetCheckGrading{6, 7, false},
		},
	},
	{
		id: "ECS140A_ID{0x00000003}",
		UnionSetStepGradings: []UnionSetStepGrading{
			UnionSetStepGrading{1, 2},
			UnionSetStepGrading{1, 3},
			UnionSetStepGrading{1, 4},
		},
		FindSetCheckGradings: []FindSetCheckGrading{
			FindSetCheckGrading{2, 1, true},
			FindSetCheckGrading{2, 3, true},
			FindSetCheckGrading{2, 5, false},
			FindSetCheckGrading{6, 7, false},
		},
	},
	{
		id: "ECS140A_ID{0x00000004}",
		UnionSetStepGradings: []UnionSetStepGrading{
			UnionSetStepGrading{1, 2},
			UnionSetStepGrading{2, 3},
			UnionSetStepGrading{3, 4},
		},
		FindSetCheckGradings: []FindSetCheckGrading{
			FindSetCheckGrading{2, 4, true},
			FindSetCheckGrading{3, 1, true},
			FindSetCheckGrading{2, 8, false},
			FindSetCheckGrading{3, 6, false},
		},
	},
	{
		id: "ECS140A_ID{0x00000005}",
		UnionSetStepGradings: []UnionSetStepGrading{
			UnionSetStepGrading{1, 2},
			UnionSetStepGrading{2, 3},
			UnionSetStepGrading{3, 1},
			UnionSetStepGrading{4, 5},
			UnionSetStepGrading{5, 6},
			UnionSetStepGrading{6, 4},
		},
		FindSetCheckGradings: []FindSetCheckGrading{
			FindSetCheckGrading{1, 3, true},
			FindSetCheckGrading{1, 6, false},
			FindSetCheckGrading{5, 6, true},
			FindSetCheckGrading{3, 5, false},
		},
	},
	{
		id: "ECS140A_ID{0x00000006}",
		UnionSetStepGradings: []UnionSetStepGrading{
			UnionSetStepGrading{1, 2},
			UnionSetStepGrading{2, 3},
			UnionSetStepGrading{3, 4},
			UnionSetStepGrading{5, 6},
			UnionSetStepGrading{6, 7},
			UnionSetStepGrading{7, 8},
			UnionSetStepGrading{8, 9},
		},
		FindSetCheckGradings: []FindSetCheckGrading{
			FindSetCheckGrading{1, 5, false},
			FindSetCheckGrading{1, 4, true},
			FindSetCheckGrading{5, 8, true},
		},
	},
}

func TestDisjointSetGrading(t *testing.T) {
	for _, test := range disjointSetGradingTests {
		func() {
			defer func() {
				if recover() != nil {
					t.Errorf("DisjointSet panicked in test %s", test.id)
				}
			}()
			s := NewDisjointSet()
			for _, UnionSetStepGrading := range test.UnionSetStepGradings {
				s.UnionSet(UnionSetStepGrading.a, UnionSetStepGrading.b)
			}
			for _, FindSetCheckGrading := range test.FindSetCheckGradings {
				if actual := s.FindSet(FindSetCheckGrading.a) == s.FindSet(FindSetCheckGrading.b); actual != FindSetCheckGrading.expected {
					t.Errorf("In test %s, FindSet(%d) == FindSet(%d) gives %t, expected %t",
						test.id, FindSetCheckGrading.a, FindSetCheckGrading.b, actual, FindSetCheckGrading.expected)
				}
			}
		}()
	}
}
