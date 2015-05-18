package goparser

// TypeSpec representation of a general type spec
type TypeSpec struct {
	Name    string
	Type    string
	Fields  []FieldSpec
	Doc     string
	Comment string
}

// FieldsTagged return a channel that will yield all FieldSpec of field
// which is tagged a certain tag name and primary value
func (s *TypeSpec) FieldsTagged(name, value string) <-chan *FieldSpec {
	cout := make(chan *FieldSpec)
	go func() {
		defer close(cout)
		for _, f := range s.Fields {
			if f.Tagged(name, value) {
				out := &FieldSpec{}
				*out = f
				cout <- out
			}
		}
	}()
	return cout
}

// FieldsTagged return a channel that will yield all FieldSpec of field
// which is tagged a certain tag name and secondary value
func (s *TypeSpec) FieldsTaggedExtra(name, value string) <-chan *FieldSpec {
	cout := make(chan *FieldSpec)
	go func() {
		defer close(cout)
		for _, f := range s.Fields {
			if f.TaggedExtra(name, value) {
				out := &FieldSpec{}
				*out = f
				cout <- out
			}
		}
	}()
	return cout
}
