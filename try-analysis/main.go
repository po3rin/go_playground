package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// parseExpr()
	parseFile()
	// inspect()
}

func parseExpr() {
	expr, _ := parser.ParseExpr("A + 1")
	ast.Print(nil, expr)
}

func parseFile() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	for _, d := range f.Imports {
		ast.Print(fset, d.Name)
		fmt.Printf("position of first character belonging to the node: %v", d.Pos())
		fmt.Printf("position of first character immediately after the node: %v", d.End())
	}
}

func inspect() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	// fmt.Println(f.Decls[1].(*ast.FuncDecl).Name)

	// ast.Inspect(f, func(n ast.Node) bool {
	// 	if v, ok := n.(*ast.BasicLit); ok {
	// 		fmt.Println(v.Value)
	// 	}
	// 	return true
	// })

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.Ident); ok {
			if v.Obj != nil {
				if v.Obj.Kind == 5 {
					fmt.Println(v.Name)
				}
			}
		}
		return true
	})
}
