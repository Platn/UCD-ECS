package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func branchCount(fn *ast.FuncDecl) uint {
	// TODO: Write the branchCount function,
	// count the number of branching statements in function fn
	count := 0
	// fmt.Printf("%s", funcs.Recv.NumFields)
	ast.Inspect(fn, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.SwitchStmt, *ast.RangeStmt, *ast.TypeSwitchStmt:
			count += 1
		}

		// If we return true, we keep recursing under this AST node.
		// If we return false, we won't visit anything under this AST node.
		return true
	})

	return uint(count)
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
