package one

import (
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/token"
)

func init() {
	appendCommand(&cobra.Command{Use: "p:ArrAndMap", Short: "", Run: one})
}

func one(cmd *cobra.Command, args []string) {
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
}
