package goparser

import (
	"go/ast"
)

// DeclsToTypes filter and transform decls to types
func DeclsToTypes(decls []ast.Decl) <-chan *TypeSpec {
	cout := make(chan *TypeSpec)
	go func(decls []ast.Decl, cout chan *TypeSpec) {
		defer close(cout)
		for ts := range FilterTypeSpec(FilterGenDecl(ChanDecls(decls))) {
			pts := ParseTypeSpec(ts)
			cout <- &pts
		}
	}(decls, cout)
	return cout

}
