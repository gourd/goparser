package goparser

import (
	"go/ast"
)

// ChanDecls channel the slice of ast.Decl out for filtering
func ChanDecls(decls []ast.Decl) <-chan ast.Decl {
	cout := make(chan ast.Decl)
	go func(decls []ast.Decl, cout chan ast.Decl) {
		defer close(cout)
		for _, decl := range decls {
			cout <- decl
		}
	}(decls, cout)
	return cout
}
