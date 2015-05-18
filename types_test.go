package goparser

import (
	"testing"
)

func testTypeSpec() *TypeSpec {
	return &TypeSpec{
		Fields: []FieldSpec{
			FieldSpec{
				Name: "Id",
				Tag:  `foo:"bar,omitempty"`,
			},
			FieldSpec{
				Name: "Title",
				Tag:  `foo:"boo,-"`,
			},
			FieldSpec{
				Name: "Body",
				Tag:  `hello:"world" foo:"-,lucky,omitempty"`,
			},
			FieldSpec{
				Name: "Author",
				Tag:  `foo:"boo"`,
			},
		},
	}
}

func TestTypeSpec_FieldsTagged(t *testing.T) {
	ts := testTypeSpec()
	expect := []string{"Title", "Author"}
	i := 0

	// check result order
	for f := range ts.FieldsTagged("foo", "boo") {
		if i < len(expect) && f.Name != expect[i] {
			t.Errorf("Failed to get correct field. Expecting \"%s\" but get \"%s\"",
				expect[i], f.Name)
		}
		i++
	}

	// check result quantity
	if i != len(expect) {
		t.Errorf("Incorrect number of fields. Expect %d but get %d", len(expect), i)
	}
}

func TestTypeSpec_FieldsTaggedExtra(t *testing.T) {
	ts := testTypeSpec()
	expect := []string{"Id", "Body"}
	i := 0

	// check result order
	for f := range ts.FieldsTaggedExtra("foo", "omitempty") {
		if i < len(expect) && f.Name != expect[i] {
			t.Errorf("Failed to get correct field. Expecting \"%s\" but get \"%s\"",
				expect[i], f.Name)
		}
		i++
	}

	// check result quantity
	if i != len(expect) {
		t.Errorf("Incorrect number of fields. Expect %d but get %d", len(expect), i)
	}
}
