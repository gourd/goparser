package goparser

import (
	"go/ast"
)

// FilterGenDecl filter only GetDecl out of ast.Decl
func FilterGenDecl(cin <-chan ast.Decl) <-chan *ast.GenDecl {
	cout := make(chan *ast.GenDecl)
	go func(cin <-chan ast.Decl, cout chan *ast.GenDecl) {
		defer close(cout)
		for decl := range cin {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				cout <- genDecl
			}
		}
	}(cin, cout)
	return cout
}

// FilterTypeSpec filter only TypeSpec from GenDecl
func FilterTypeSpec(cin <-chan *ast.GenDecl) <-chan *ast.TypeSpec {
	cout := make(chan *ast.TypeSpec)
	go func(cin <-chan *ast.GenDecl, cout chan *ast.TypeSpec) {
		defer close(cout)
		for decl := range cin {
			// Note: may further break up this into
			//       another pipeline function
			for _, spec := range decl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if typeSpec.Doc == nil {
						// Don't know why the typeSpec doesn't carry
						// the doc of the parent declaration.
						// Pass it through.
						typeSpec.Doc = decl.Doc
					}
					cout <- typeSpec
				}
			}
		}
	}(cin, cout)
	return cout
}

// FilterStructType filter only StructType from TypeSpec
func FilterStructType(cin <-chan *ast.TypeSpec) <-chan *ast.StructType {
	cout := make(chan *ast.StructType)
	go func(cin <-chan *ast.TypeSpec, cout chan *ast.StructType) {
		defer close(cout)
		for spec := range cin {
			if structType, ok := spec.Type.(*ast.StructType); ok {
				cout <- structType
			}
		}
	}(cin, cout)
	return cout
}
