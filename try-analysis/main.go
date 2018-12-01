package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var src = `package p

import x "pkg"

func f() {
    if x := x.f(); x != nil {
        x(func(x int) int { return x + 1 })
    }
}
`

func main() {
	// parseExpr()
	// parseFile()
	inspect()
	// scope()
}

func parseExpr() {
	expr, _ := parser.ParseExpr("A + 1")
	ast.Print(nil, expr)
}

func parseFile() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
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

	// ast.Inspect(f, func(n ast.Node) bool {
	// 	if v, ok := n.(*ast.Ident); ok {
	// 		if v.Obj != nil {
	// 			if v.Obj.Kind == 5 {
	// 				fmt.Println(v.Name)
	// 			}
	// 		}
	// 	}
	// 	return true
	// })

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			fmt.Println(v.Name)
			fmt.Println(fset.Position(v.Pos()))
			// fmt.Println(v.Pos())
		}
		return true
	})
}

func scope() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example.go", src, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Name == "x" {
			var decl interface{}
			if ident != nil && ident.Obj != nil {
				decl = ident.Obj.Decl
			}
			var kind ast.ObjKind
			if ident.Obj != nil {
				kind = ident.Obj.Kind
			}
			fmt.Printf("%-17sobj=%-12p  kind=%s decl=%T\n", fset.Position(ident.Pos()), ident.Obj, kind, decl)
		}
		return true
	})
}
