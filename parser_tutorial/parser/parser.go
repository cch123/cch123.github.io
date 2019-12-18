package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var src = `
package main

func main() {
	println("hello world")
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.Mode(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	ast.Print(fset, f)
}
