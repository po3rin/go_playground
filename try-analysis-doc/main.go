package main // import "try-analisis-doc"

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func isGoFile(fi os.FileInfo) bool {
	name := fi.Name()
	return !fi.IsDir() &&
		len(name) > 0 && name[0] != '.' && // ignore .files
		filepath.Ext(name) == ".go"
}

func main() {
	filter := isGoFile
	fset := token.NewFileSet()
	pkgs, _ := parser.ParseDir(fset, "./example", filter, parser.ParseComments)

	// for _, c := range f.Comments {
	// 	fmt.Printf("%s: %q\n", fset.Position(c.Pos()), c.Text())
	// }

	for _, pkg := range pkgs {
		importpath := "example/" + pkg.Name
		doc := doc.New(pkg, importpath, 0)

		// golden files always use / in filenames - canonicalize them
		lines := strings.Split(doc.Doc, "\n")
		fmt.Println(lines[0])
	}

	// fmt.Println(f.Doc.Text())

	// for _, l := range f.Doc.List {
	// 	fmt.Println(l.Text)
	// }

	// fmt.Println(f.Doc.List[0].Text)
}
