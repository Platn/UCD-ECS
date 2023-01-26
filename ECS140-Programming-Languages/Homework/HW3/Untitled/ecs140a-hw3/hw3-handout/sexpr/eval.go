package sexpr

import (
	"errors"
	"math/big"
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrEval = errors.New("eval error")

// references: https://courses.cs.northwestern.edu/325/readings/interpreter.php#read

// helpers
func (expr *SExpr) isNilIncludingLiteral() bool {
	return expr.isNil() || (expr.isSymbol() && expr.atom.literal == "NIL")
}

func (expr *SExpr) GetLength() (int64, error) {
	if expr.isNil() {
		return 0, nil
	}
	if expr.isConsCell() {
		if expr.cdr == nil || expr.cdr.isNilIncludingLiteral() {
			return 1, nil
		} else {
			// we can continue to recurse on cdr
			recLen, recErr := expr.cdr.GetLength()
			if recErr != nil {
				return 0, recErr
			}
			return 1 + recLen, nil
		}
	}
	return 0, ErrEval
}

func (expr *SExpr) CalcSum() (*big.Int, error) {
	// assume cons cell
	eval, err := expr.car.Eval()
	if err != nil {
		return nil, ErrEval
	}
	if eval.isNumber() {
		if expr.cdr.isNilIncludingLiteral() {
			return eval.atom.num, nil
		} else {
			recSum, recErr := expr.cdr.CalcSum()
			if recErr != nil {
				return nil, ErrEval
			}
			return big.NewInt(0).Add(eval.atom.num, recSum), nil
		}
	}
	return nil, ErrEval
}

func (expr *SExpr) CalcProduct() (*big.Int, error) {
	// assume cons cell
	eval, err := expr.car.Eval()
	if err != nil {
		return nil, ErrEval
	}
	if eval.isNumber() {
		if expr.cdr.isNilIncludingLiteral() {
			return eval.atom.num, nil
		} else {
			recSum, recErr := expr.cdr.CalcProduct()
			if recErr != nil {
				return nil, ErrEval
			}
			return big.NewInt(0).Mul(eval.atom.num, recSum), nil
		}
	}
	return nil, ErrEval
}

func (expr *SExpr) Eval() (*SExpr, error) {
	//fmt.Println("EVAL: ", expr.SExprString())
	if expr.isNilIncludingLiteral() {
		return mkNil(), nil
	} else if expr.isAtom() {
		// atom
		if expr.isNumber() {
			return expr, nil
		}
		return nil, ErrEval
	} else {
		// list
		argLen, _ := expr.cdr.GetLength()
		//fmt.Println("arg len: ", argLen, " called on ", expr.cdr)
		if expr.car.isSymbol() {
			// the first element of the list is a symbol
			listSymbol := expr.car.atom.literal
			//fmt.Println("list symbol: ", listSymbol)
			switch listSymbol {
			case "QUOTE", "'":
				if argLen != 1 {
					return nil, ErrEval
				}
				// if !expr.cdr.car.isNil() && expr.cdr.car.isConsCell() && expr.cdr.car.car.isAtom() && expr.cdr.car.cdr.isAtom() {
				// 	return nil, ErrEval
				// }
				return expr.cdr.car, nil
			case "CAR":
				if argLen != 1 {
					return nil, ErrEval
				}
				// TODO I probably shouldnt need this
				// if expr.cdr.isConsCell() && expr.cdr.car.isNilIncludingLiteral() && expr.cdr.cdr.isNilIncludingLiteral() {
				// 	return mkNil(), nil
				// }
				cdrEval, cdrEvalErr := expr.cdr.car.Eval()
				if cdrEvalErr != nil {
					return nil, ErrEval
				}
				if !cdrEval.isNil() && cdrEval.isConsCell() {
					return cdrEval.car, nil
				}
				return mkNil(), nil
			case "CDR":
				if argLen != 1 {
					return nil, ErrEval
				}
				// TODO I probably shouldnt need this
				// if expr.cdr.isConsCell() && expr.cdr.car.isNilIncludingLiteral() && expr.cdr.cdr.isNilIncludingLiteral() {
				// 	return mkNil(), nil
				// }
				cdrEval, cdrEvalErr := expr.cdr.car.Eval()
				if cdrEvalErr != nil {
					return nil, ErrEval
				}
				if !cdrEval.isNil() && cdrEval.isConsCell() {
					return cdrEval.cdr, nil
				}
				return mkNil(), nil
			case "CONS":
				if argLen != 2 {
					return nil, ErrEval
				}
				cons1Eval, cons1Err := expr.cdr.car.Eval()
				if cons1Err != nil {
					return nil, ErrEval
				}
				cons2Eval, cons2Err := expr.cdr.cdr.car.Eval()
				if cons2Err != nil {
					return nil, ErrEval
				}
				return mkConsCell(cons1Eval, cons2Eval), nil
			case "LENGTH":
				if argLen != 1 {
					return nil, ErrEval
				}
				cdrEval, cdrEvalErr := expr.cdr.car.Eval()
				if cdrEvalErr != nil {
					return nil, ErrEval
				}
				if !cdrEval.isConsCell() {
					return nil, ErrEval
				}
				// if !cdrEval.isNil() && cdrEval.isConsCell() && cdrEval.car.isAtomExcludingLiteralNil() && cdrEval.cdr.isAtomExcludingLiteralNil() {
				// 	return nil, ErrEval
				// }
				lenVal, lenErr := cdrEval.GetLength()
				if lenErr != nil {
					return nil, ErrEval
				}
				return mkNumber(big.NewInt(lenVal)), nil
			case "ATOM":
				if argLen != 1 {
					return nil, ErrEval
				}
				cdrEval, cdrEvalErr := expr.cdr.car.Eval()
				if cdrEvalErr != nil {
					return nil, ErrEval
				}
				if cdrEval.isAtom() {
					return mkSymbol("T"), nil
				}
				return mkNil(), nil
			case "LISTP":
				if argLen != 1 {
					return nil, ErrEval
				}
				// TODO I probably shouldnt need this
				// if expr.cdr.isConsCell() && expr.cdr.car.isNilIncludingLiteral() && expr.cdr.cdr.isNilIncludingLiteral() {
				// 	return mkSymbol("T"), nil
				// }
				cdrEval, cdrEvalErr := expr.cdr.car.Eval()
				if cdrEvalErr != nil {
					return nil, ErrEval
				}
				if cdrEval.isConsCell() {
					return mkSymbol("T"), nil
				}
				return mkNil(), nil
			case "ZEROP":
				if argLen != 1 {
					return nil, ErrEval
				}
				cdrEval, cdrEvalErr := expr.cdr.car.Eval()
				if cdrEvalErr != nil {
					return nil, ErrEval
				}
				if cdrEval.isConsCell() {
					return nil, ErrEval
				}
				if cdrEval.isNumber() && cdrEval.atom.num.Cmp(big.NewInt(0)) == 0 {
					return mkSymbol("T"), nil
				}
				return mkNil(), nil
			case "+":
				if argLen == 0 {
					return mkNumber(big.NewInt(0)), nil
				}
				sumVal, sumErr := expr.cdr.CalcSum()
				if sumErr != nil {
					return nil, ErrEval
				}
				return mkNumber(sumVal), nil
			case "*":
				if argLen == 0 {
					return mkNumber(big.NewInt(1)), nil
				}
				sumVal, sumErr := expr.cdr.CalcProduct()
				if sumErr != nil {
					return nil, ErrEval
				}
				return mkNumber(sumVal), nil
			default:
				return nil, ErrEval
			}
		} else {
			// the list starts with something invalid
			return nil, ErrEval
		}
	}
}
