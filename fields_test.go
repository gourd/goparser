package goparser

import (
	"testing"
)

func TestFieldSpec_Tagged_Positive(t *testing.T) {
	// field exists with the value
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" foo:"bar,omitempty" hello:"world,-"`,
	}
	if !s.Tagged("foo", "bar") {
		t.Error("FieldSpec.Tagged couldn't be identify the tagged field as expected")
	}
}

func TestFieldSpec_Tagged_Negative1(t *testing.T) {
	// field doesn't exist
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" hello:"world,-"`,
	}
	if s.Tagged("foo", "bar") {
		t.Error("FieldSpec.Tagged falsely identified the tagged field")
	}
}

func TestFieldSpec_Tagged_Negative2(t *testing.T) {
	// field exists but value doesn't exist
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" foo:"notmatch,hi" hello:"world,-"`,
	}
	if s.Tagged("foo", "bar") {
		t.Error("FieldSpec.Tagged falsely identified the tagged field")
	}
}

func TestFieldSpec_TaggedExtra_Positive(t *testing.T) {
	// field exists with the value
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" foo:"bar,omitempty" hello:"world,-"`,
	}
	if !s.TaggedExtra("foo", "omitempty") {
		t.Error("FieldSpec.TaggedExtra couldn't be identify the tagged field as expected")
	}
}

func TestFieldSpec_TaggedExtra_Negative1(t *testing.T) {
	// field doesn't exist
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" hello:"world,-"`,
	}
	if s.TaggedExtra("foo", "omitempty") {
		t.Error("FieldSpec.TaggedExtra falsely identified the tagged field")
	}
}

func TestFieldSpec_TaggedExtra_Negative2(t *testing.T) {
	// field exists but value doesn't exist
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" foo:"notmatch,hi" hello:"world,-"`,
	}
	if s.TaggedExtra("foo", "omitempty") {
		t.Error("FieldSpec.TaggedExtra falsely identified the tagged field")
	}
}

func TestFieldSpec_TaggedExtra_Negative3(t *testing.T) {
	// field exists and value exists, but not secondary
	s := &FieldSpec{
		Tag: `db:"dummy_field,omitempty" foo:"notmatch,hi" hello:"world,-"`,
	}
	if s.TaggedExtra("foo", "notmatch") {
		t.Error("FieldSpec.TaggedExtra falsely identified the tagged field")
	}
}
