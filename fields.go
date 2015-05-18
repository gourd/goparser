package goparser

import (
	"regexp"
	"strings"
)

var tagFieldRe *regexp.Regexp

func init() {
	tagFieldRe = regexp.MustCompile("(\\w+):\"([^\"]*)\"")
}

// TagField representats structure of a tag field
type TagField struct {
	Name   string
	Value  string
	Extras []string
}

// FieldSpec representation of a general struct field
type FieldSpec struct {
	Name    string
	Type    string
	Tag     string
	Doc     string
	Comment string
}

// Tags read the Tag text into a map of string
func (s *FieldSpec) TagFields() <-chan *TagField {
	cout := make(chan *TagField)
	go func(cout chan *TagField) {
		defer close(cout)

		tagStr := strings.Trim(s.Tag, "\n\r\t ")
		fields := strings.Split(tagStr, " ")
		for _, f := range fields {
			f = strings.Trim(f, "\n\r\t ")
			if tagFieldRe.MatchString(f) {
				ms := tagFieldRe.FindStringSubmatch(f)
				tf := &TagField{
					Name: ms[1],
				}

				// parse value and extras, if any
				vals := strings.Split(ms[2], ",")
				if len(vals) > 0 {
					tf.Value = vals[0]
					tf.Extras = vals[1:]
				}

				// yield field
				cout <- tf
			}
		}
	}(cout)
	return cout
}

// Tagged determine if the field is tagged with value
// as the primary value (the 1st in comma separated list)
func (s *FieldSpec) Tagged(name, value string) bool {
	for f := range s.TagFields() {
		if f.Name == name {
			if f.Value == value {
				return true
			}
			return false
		}
	}
	return false
}

// TaggedExtra determine if the field is tagged with
// given secondary values (not 1st in comma separated list)
func (s *FieldSpec) TaggedExtra(name, value string) bool {
	for f := range s.TagFields() {
		if f.Name == name {
			for _, v := range f.Extras {
				if v == value {
					return true
				}
			}
			return false
		}
	}
	return false
}
