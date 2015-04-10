package goparser

// TypeSpec representation of a general type spec
type TypeSpec struct {
	Name    string
	Type    string
	Fields  []FieldSpec
	Doc     string
	Comment string
}

// FieldSpec representation of a general struct field
type FieldSpec struct {
	Name    string
	Type    string
	Tag     string
	Doc     string
	Comment string
}
