package one

import (
	"go/ast"
	"go/parser"
	"go/token"
	"thh/base"
)

var commandList = make(map[string]base.Console)

func GetAllConsoles() map[string]base.Console {
	return commandList
}

func addConsole(signature string, description string, handle func()) {
	c := base.Console{Signature: signature, Description: description, Handle: handle}
	commandList[c.Signature] = c
}

func init() {
	addConsole("one", "", func() {
		src := `
package main
func main() {
    println("Hello, World!")
}
`
		// Create the AST by parsing src.
		fSet := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fSet, "", src, 0)
		if err != nil {
			panic(err)
		}

		// Print the AST.
		_ = ast.Print(fSet, f)
	})
}
