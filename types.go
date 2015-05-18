package goparser

// TypeSpec representation of a general type spec
type TypeSpec struct {
	Name    string
	Type    string
	Fields  []FieldSpec
	Doc     string
	Comment string
}
