package branch

import (
	"fmt"
	"testing"
	"time"
)

type Test struct {
	id       string
	name     string
	branches uint
}

var test_code = `
    package main

    import (
        "fmt"
        "eval"
    )

    func _sequential_if() {
        if true {
            fmt.Println("is true")
        } else {
            fmt.Println("isn't true")
        }

        if true {
            fmt.Println("is true")
        } else {
            fmt.Println("isn't true")
        }
    }

    func _single_for() {
        for i := 0; i < 10; i += 1 {
            fmt.Println("is loop")
        }

        for i := 0; i < 10; i += 1 {
            fmt.Println("is loop")
        }
    }

    func _single_range() {
        for i := range []int{1, 1, 2, 3, 5, 8} {
            fmt.Println("is fibonacci")
        }

        for i := range []int{1, 1, 2, 3, 5, 8} {
            fmt.Println("is fibonacci")
        }
    }

    func _no_branches() {
        return 42
    }

    func _single_switch() {
        switch 5 {
        case 0:
            // pass
        case 5:
            fmt.Println("It's five!")
        default:
            fmt.Println("It isn't five...")
        }

        switch 5 {
        case 0:
            // pass
        case 5:
            fmt.Println("It's five!")
        default:
            fmt.Println("It isn't five...")
        }
    }

    func _single_typeswitch() {
        var x interface{} = "test"
        switch x.(type) {
        case uint:
            // pass
        case string:
            fmt.Println("It's a string!")
        default:
            fmt.Println("It's not a string...")
        }
        switch x.(type) {
        case uint:
            // pass
        case string:
            fmt.Println("It's a string!")
        default:
            fmt.Println("It's not a string...")
        }
    }

    func _nested_if(x int) float64 {
        var result float64
        if x < 0 {
            if x > -5 {
                result := -0.5
            } else {
                result := -1
            }
        } else if x > 0 {
            result := 1
        } else {
            result := 0
        }
        if x < 0 {
            if x > -5 {
                result := -0.5
            } else {
                result := -1
            }
        } else if x > 0 {
            result := 1
        } else {
            result := 0
        }

        return result
    }

    func _nested_if_3(x int) float64 {
        var result float64
        if x < 0 {
            if x > -5 {
                if x > -3{
                    result := -0.3
                }else{
                    result := -0.5
                }

            } else {
                result := -1
            }
        } else if x > 0 {
            result := 1
        } else {
            result := 0
        }
        if x < 0 {
            if x > -5 {
                result := -0.5
            } else {
                result := -1
            }
        } else if x > 0 {
            result := 1
        } else {
            result := 0
        }

        return result
    }

    func _nested_for_if() {
        for i := 0; i < 10; i += 1 {
            if i > 5 {
                fmt.Println("is filter")
            }
        }
        for i := 0; i < 10; i += 1 {
            if i > 5 {
                fmt.Println("is filter")
            }
        }

    }

    func _nested_for_if_if() {
        for i := 0; i < 10; i += 1 {
            if i > 5 {
                if i < 8{
                    fmt.Println("is filter")
                }else{
                    fmt.Println("aaa")
                }

            }
        }
        for i := 0; i < 10; i += 1 {
            if i > 5 {
                fmt.Println("is filter")
            }
        }

    }

    func _nested_if_for() {
        if true{
            for i := 0; i < 10; i += 1 {
                    fmt.Println("is filter")
            }
        }else{
            for i := 0; i < 10; i += 1 {
                if i > 5 {
                    fmt.Println("is filter")
                }
            }
        }

    }

    func _nested_switch_if(x int) {
        switch x > 5 {
        case true:
            if x > 10 {
                fmt.Println("is really big")
            }
        default:
            fmt.Println("is default")
        }
        switch x > 5 {
        case true:
            if x > 10 {
                fmt.Println("is really big")
            }
        default:
            fmt.Println("is default")
        }
    }

    func _nested_if_switch_if(x int) {

        if x > 0{
            switch x > 5 {
            case true:
                if x > 10 {
                    fmt.Println("is really big")
                }
            default:
                fmt.Println("is default")
            }
            switch x > 5 {
            case true:
                if x > 10 {
                    fmt.Println("is really big")
                }
            default:
                fmt.Println("is default")
            }
        }
    }

    func _nested_for_switch(x int){

        for i:= 0; i < 10; i++{
            switch x{
            case 0:
                fmt.Println("Zero")
            default:
                fmt.Println("Nonzero")
            }
        }
    }

    func _nested_switch_for(x int){

        switch x > 0{
        case true:
            for i:= 0; i < x; i++{
                fmt.Println("Zero")
            }
        default:
            //pass
        }

    }

    func _mixed_switch_no_default_for_if() {
        switch 5 {
        case 0:
            // pass
        case 5:
            fmt.Println("It's five!")
        }

        for i := 0; i < 10; i += 1 {
            if i > 5 {
                fmt.Println("is filter")
            }
        }
    }

    func _single_typeswitch_no_default() {
        var x interface{} = "test"
        switch x.(type) {
        case uint:
            // pass
        case string:
            fmt.Println("It's a string!")
        }
    }

    func _nested_if_no_else(x int) float64 {
        var result float64 = 0
        if x < 0 {
            if x > -5 {
                result := -0.5
            } else {
                result := -1
            }
        } else if x > 0 {
            result := 1
        }
        return result
    }
    `
var tests = []Test{
	{"ECS140A_ID{0x00000001}", "_sequential_if", 2},
	{"ECS140A_ID{0x00000002}", "_single_for", 2},
	{"ECS140A_ID{0x00000003}", "_single_range", 2},
	{"ECS140A_ID{0x00000004}", "_no_branches", 0},
	{"ECS140A_ID{0x00000005}", "_single_switch", 2},
	{"ECS140A_ID{0x00000006}", "_single_typeswitch", 2},
	{"ECS140A_ID{0x00000007}", "_nested_if", 6},
	{"ECS140A_ID{0x00000008}", "_nested_for_if", 4},
	{"ECS140A_ID{0x00000009}", "_nested_switch_if", 4},
	{"ECS140A_ID{0x0000000a}", "_mixed_switch_no_default_for_if", 3},
	{"ECS140A_ID{0x0000000b}", "_single_typeswitch_no_default", 1},
	{"ECS140A_ID{0x0000000c}", "_nested_if_no_else", 3},
	{"ECS140A_ID{0x0000000d}", "_nested_for_switch", 2},
	{"ECS140A_ID{0x0000000e}", "_nested_switch_for", 2},
	{"ECS140A_ID{0x0000000f}", "_nested_if_switch_if", 5},
	{"ECS140A_ID{0x00000010}", "_nested_for_if_if", 5},
	{"ECS140A_ID{0x00000011}", "_nested_if_for", 4},
	{"ECS140A_ID{0x00000012}", "_nested_if_3", 7},
}

func (test Test) RunTest() (bool, error, string) {

	branch_factors := ComputeBranchFactors(test_code)
	res := branch_factors[test.name] == test.branches

	var msg string
	if !res {
		msg = fmt.Sprintf("ComputeBranchFactors(%v) = %d, want %d",
			test.name, branch_factors[test.name], test.branches)
	}

	return res, nil, msg

}

func (test Test) GetId() string {
	return test.id
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

func TestComputeBranchFactors_Ins(t *testing.T) {

	ins_tests := make([]ATest, len(tests))

	for i, v := range tests {
		ins_tests[i] = v
	}

	RunSomeTests(ins_tests, t)
}
