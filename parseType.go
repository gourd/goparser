package goparser

import (
	"go/ast"
)

// ParseTypeSpec parse a type spec into TypeSpec
func ParseTypeSpec(typeSpec *ast.TypeSpec) (spec TypeSpec) {

	// name of the type
	if typeSpec.Name != nil {
		spec.Name = typeSpec.Name.Name
	}

	// read comment and docs
	spec.Doc = CmtGrpToStr(typeSpec.Doc)
	spec.Comment = CmtGrpToStr(typeSpec.Comment)

	// if this is a struct
	if structType, ok := typeSpec.Type.(*ast.StructType); ok {
		spec.Type = "struct"

		// see FieldList, Field, Tag in package "os/ast"
		for _, f := range structType.Fields.List {
			fspec := FieldSpec{}
			if len(f.Names) > 0 {
				fspec.Name = f.Names[0].Name
			}
			if f.Tag != nil {
				// remove the "`" from beginning and end
				fspec.Tag = f.Tag.Value[1 : len(f.Tag.Value)-1]
			}
			if f.Type != nil {
				if id, ok := f.Type.(*ast.Ident); ok {
					fspec.Type = id.Name
				}
			}
			fspec.Doc = CmtGrpToStr(f.Doc)
			fspec.Comment = CmtGrpToStr(f.Comment)
			spec.Fields = append(spec.Fields, fspec)
		}
	}

	return
}
