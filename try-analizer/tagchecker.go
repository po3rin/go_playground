package tagchecker // import "tagchecker"

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is provided by go/analysis.
var Analyzer = &analysis.Analyzer{
	Name:             "tagchecker",
	Doc:              doc,
	Run:              run,
	RunDespiteErrors: true,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const doc = "tagchecker checks whether an appropriate tag is set in the field of the structure"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		styp := pass.TypesInfo.Types[n.(*ast.StructType)].Type.(*types.Struct)
		var seen map[[2]string]token.Pos
		for i := 0; i < styp.NumFields(); i++ {
			field := styp.Field(i)
			tag := styp.Tag(i)
			checkCanonicalFieldTag(pass, field, tag, &seen)
		}
	})
	return nil, nil
}

func checkCanonicalFieldTag(pass *analysis.Pass, field *types.Var, tag string, seen *map[[2]string]token.Pos) {
	fmt.Println(pass)
	fmt.Println(tag)
	fmt.Println(field)
	fmt.Println(seen)
}
