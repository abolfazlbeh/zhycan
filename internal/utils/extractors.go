package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// ExtractStructs - extract struct from source code
func ExtractStructs(sourceCode string) []*ast.StructType {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", sourceCode, parser.ParseComments)
	if err != nil {
		return nil
	}

	// Collect the struct types in this slice.
	var structTypes []*ast.StructType
	ast.Inspect(node, func(n ast.Node) bool {
		if n, ok := n.(*ast.StructType); ok {
			structTypes = append(structTypes, n)
		}
		return true
	})
	return structTypes
}
