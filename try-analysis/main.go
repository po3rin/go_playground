package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
)

func main() {
	// parseExpr()
	// parseFile()
	// inspect()
	// scope()
	// printFile()
	write()
	// typesCheck()
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
	f, _ := parser.ParseFile(fset, "example/example.go", nil, parser.Mode(0))

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

func printFile() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example/example2.go", nil, parser.Mode(0))

	file, err := os.OpenFile("result/result.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(file, fset, f)
}

func write() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			v.Name = &ast.Ident{
				Name: "plus",
			}
		}
		return true
	})

	file, err := os.OpenFile("example/result.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(file, fset, f)
}

func typesCheck() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example/example2.go", nil, parser.Mode(0))
	if err != nil {
		log.Fatalln("Error:", err)
	}

	conf := types.Config{Importer: importer.Default()}

	pkg, err := conf.Check("example", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	scopes := map[*types.Scope]struct{}{}
	ast.Inspect(f, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			innerMost := pkg.Scope().Innermost(ident.Pos())
			s, _ := innerMost.LookupParent(ident.Name, ident.Pos())
			if s != nil {
				fmt.Println("==========")
				fmt.Println(s != pkg.Scope())
				fmt.Println(ident.Name)
				if s != pkg.Scope() && ident.Name == "helloWorld" {
					ident = &ast.Ident{
						Name: "heeeee",
					}
				}
				scopes[s] = struct{}{}
			}
		}
		return true
	})

	fmt.Println("====", len(scopes), "scopes ====")
	for s := range scopes {
		fmt.Println(s)
	}

	file, err := os.OpenFile("example/result2.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(file, fset, f)
}
