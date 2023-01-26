package matrix

import (
	"fmt"
	"hw1/ins_tester"
	"testing"
)

var m0 [][]int = [][]int{}

var m1 [][]int = [][]int{
	{3, 6, 4, 2, 1, 5},
	{1, 2, 5, 4, 6, 3},
	{5, 4, 2, 1, 3, 6},
	{6, 3, 1, 5, 2, 4},
	{4, 1, 6, 3, 5, 2},
	{2, 5, 3, 6, 4, 1},
}

var m1t [][]int = [][]int{
	{3, 1, 5, 6, 4, 2},
	{6, 2, 4, 3, 1, 5},
	{4, 5, 2, 1, 6, 3},
	{2, 4, 1, 5, 3, 6},
	{1, 6, 3, 2, 5, 4},
	{5, 3, 6, 4, 2, 1},
}

var m2 [][]int = [][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 5, 1, 1},
	{1, 2, 0, 4, 1},
	{1, 1, 3, 1, 1},
	{1, 1, 1, 1, 1},
}

var m2t [][]int = [][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 2, 1, 1},
	{1, 5, 0, 3, 1},
	{1, 1, 4, 1, 1},
	{1, 1, 1, 1, 1},
}

var m3 [][]int = [][]int{
	{1, 5, 0, 7, 1},
	{4, 1, 6, 1, 13},
	{0, 3, 1, 12, 0},
	{2, 1, 9, 1, 11},
	{1, 8, 0, 10, 1},
}

var m3t [][]int = [][]int{
	{1, 4, 0, 2, 1},
	{5, 1, 3, 1, 8},
	{0, 6, 1, 9, 0},
	{7, 1, 12, 1, 10},
	{1, 13, 0, 11, 1},
}

var m4 [][]int = [][]int{
	{0, 2, 9, 0},
	{3, 1, 1, 8},
	{4, 1, 1, 7},
	{0, 5, 6, 0},
}

var m4t [][]int = [][]int{
	{0, 3, 4, 0},
	{2, 1, 1, 5},
	{9, 1, 1, 6},
	{0, 8, 7, 0},
}

var m5 [][]int = [][]int{
	{0, 1, 2, 3},
	{1, 2, 3, 0},
	{2, 3, 0, 1},
	{2, 4, 6, 0},
	{3, 0, 1, 2},
}

var m5t [][]int = [][]int{
	{0, 1, 2, 2, 3},
	{1, 2, 3, 4, 0},
	{2, 3, 0, 6, 1},
	{3, 0, 1, 0, 2},
}

var m6 [][]int = [][]int{
	{0, 1, 2, 3},
	{1, 2, 3, 0},
	{2, 3, 0, 1},
	{2, 4, 7, 0},
	{3, 0, 1, 2},
}

var m6t [][]int = [][]int{
	{0, 1, 2, 2, 3},
	{1, 2, 3, 4, 0},
	{2, 3, 0, 7, 1},
	{3, 0, 1, 0, 2},
}

var m7 [][]int = [][]int{
	{1, 2},
	{3, 4},
}

var m7t [][]int = [][]int{
	{1, 3},
	{2, 4},
}

var m8 [][]int = [][]int{
	{1, 2, 3, 4},
}

var m8t [][]int = [][]int{
	{1},
	{2},
	{3},
	{4},
}

func matrix_equal(mat1, mat2 [][]int) bool {
	if mat1 == nil || mat2 == nil || len(mat1) != len(mat2) {
		return false
	}
	for index, i1 := range mat1 {
		i2 := mat2[index]
		if i1 == nil || i2 == nil || len(i1) != len(i2) {
			return false
		}
		for index, j1 := range i1 {
			j2 := i2[index]
			if j1 != j2 {
				return false
			}
		}
	}
	return true
}

// Tests for AreAdjacent

type LN_Test struct {
	id   string
	lst  []int
	a, b int
	out  bool
}

var ln_tests = []LN_Test{
	{"ECS140A_ID{AreAdjacent|0x00000001}", nil, 1, 1, false},
	{"ECS140A_ID{AreAdjacent|0x00000002}", []int{}, 1, 1, false},
	{"ECS140A_ID{AreAdjacent|0x00000003}", []int{1}, 1, 1, false},
	{"ECS140A_ID{AreAdjacent|0x00000004}", []int{1, 1}, 1, 1, true},
	{"ECS140A_ID{AreAdjacent|0x00000005}", []int{1, 2, 3}, 2, 1, true},
	{"ECS140A_ID{AreAdjacent|0x00000006}", []int{1, 2, 3}, 1, 2, true},
	{"ECS140A_ID{AreAdjacent|0x00000007}", []int{1, 2, 3}, 3, 2, true},
	{"ECS140A_ID{AreAdjacent|0x00000008}", []int{1, 2, 1}, 1, 1, false},
	// TODO: mutate tests
}

func (test LN_Test) RunTest() (bool, error, string) {
	actual := AreAdjacent(test.lst, test.a, test.b)
	res := actual == test.out
	var msg string
	if !res {
		msg = fmt.Sprintf("AreAdjacent(%v, %d, %d)= %t ; Expected: %t",
			test.lst, test.a, test.b, actual, test.out)
	}
	return res, nil, msg
}

func (test LN_Test) GetId() string {
	return test.id
}

func TestAreAdjacent_Ins(t *testing.T) {
	ins_tests := make([]ins_tester.ATest, len(ln_tests))
	for i, v := range ln_tests {
		ins_tests[i] = v
	}
	ins_tester.RunSomeTests(ins_tests, t)
}

// Tests for Transpose

type MT_Test struct {
	id         string
	mat1, mat2 [][]int
}

var mt_tests = []MT_Test{
	{"ECS140A_ID{Transpose|0x00000001}", m1, m1t},
	{"ECS140A_ID{Transpose|0x00000002}", m2, m2t},
	{"ECS140A_ID{Transpose|0x00000003}", m3, m3t},
	{"ECS140A_ID{Transpose|0x00000004}", m4, m4t},
	{"ECS140A_ID{Transpose|0x00000005}", m5, m5t},
	{"ECS140A_ID{Transpose|0x00000006}", m6, m6t},
	{"ECS140A_ID{Transpose|0x00000007}", m7, m7t},
	{"ECS140A_ID{Transpose|0x00000008}", m8, m8t},
	// TODO: mutate tests
}

func (test MT_Test) RunTest() (bool, error, string) {
	actual := Transpose(test.mat1)
	res := matrix_equal(actual, test.mat2)
	var msg string
	if !res {
		msg = fmt.Sprintf("Transpose(%v)= %v ; Expected: %v",
			test.mat1, actual, test.mat2)
	}
	return res, nil, msg
}

func (test MT_Test) GetId() string {
	return test.id
}

func TestTranspose_Ins(t *testing.T) {
	ins_tests := make([]ins_tester.ATest, len(mt_tests))
	for i, v := range mt_tests {
		ins_tests[i] = v
	}
	ins_tester.RunSomeTests(ins_tests, t)
}

// Tests for AreNeighbors

type MN_Test struct {
	id   string
	mat  [][]int
	a, b int
	out  bool
}

var mn_tests = []MN_Test{
	{"ECS140A_ID{AreNeighbors|0x00000001}", m2, 0, 1, false},
	{"ECS140A_ID{AreNeighbors|0x00000002}", m2, 0, 2, true},
	{"ECS140A_ID{AreNeighbors|0x00000003}", m2, 0, 3, true},
	{"ECS140A_ID{AreNeighbors|0x00000004}", m2, 0, 4, true},
	{"ECS140A_ID{AreNeighbors|0x00000005}", m2, 0, 5, true},
	{"ECS140A_ID{AreNeighbors|0x00000006}", m2, 3, 2, false},
	{"ECS140A_ID{AreNeighbors|0x00000007}", m2, 4, 3, false},
	{"ECS140A_ID{AreNeighbors|0x00000008}", m2, 1, 1, true},
	{"ECS140A_ID{AreNeighbors|0x00000009}", m2, 4, 5, false},
	{"ECS140A_ID{AreNeighbors|0x0000000a}", m2, 3, 5, false},
	// TODO: mutate tests
	// TODO: add more tests on m3, m4, m5
}

func (test MN_Test) RunTest() (bool, error, string) {
	actual := AreNeighbors(test.mat, test.a, test.b)
	res := actual == test.out
	var msg string
	if !res {
		msg = fmt.Sprintf("AreNeighbors(%v, %v, %v) = %t ; Expected %t",
			test.mat, test.a, test.b, actual, test.out)
	}
	return res, nil, msg
}

func (test MN_Test) GetId() string {
	return test.id
}

func TestAreNeighbors_Ins(t *testing.T) {
	ins_tests := make([]ins_tester.ATest, len(mn_tests))
	for i, v := range mn_tests {
		ins_tests[i] = v
	}
	ins_tester.RunSomeTests(ins_tests, t)
}
