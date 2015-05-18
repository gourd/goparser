package goparser

import (
	"go/parser"
	"go/token"
	"testing"
)

// Test example
func TestExample(t *testing.T) {

	fset := token.NewFileSet()

	// filepath can only be filename
	filePath := "./_example/data/types.go"
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		t.Errorf(err.Error())
	}

	// read package name
	if f.Name == nil {
		t.Errorf("Unknown package name")
	} else if f.Name.Name != "example" {
		t.Errorf("Failed to parse packge name. "+
			"Expected \"example\" but get\"%s\"", f.Name.Name)
	}

	// read types name and details
	numt := 0
	expected := []string{"Post", "Comment", "User"}
	var found []TypeSpec
	for ts := range FilterTypeSpec(FilterGenDecl(ChanDecls(f.Decls))) {
		pts := ParseTypeSpec(ts)
		if pts.Name != expected[numt] {
			t.Errorf("Unexptected definition \"%s\"", pts.Name)
		}
		found = append(found, pts)
		numt++
	}

	// see if all types needed are found
	if len(expected) != len(found) {
		t.Errorf("Types found are not as expected")
		t.Errorf("Expected: %#v", expected)
		t.Errorf("Found:    %#v", found)
	}
}

// Test example
func TestDeclsToTypes(t *testing.T) {

	fset := token.NewFileSet()

	// filepath can only be filename
	filePath := "./_example/data/types.go"
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		t.Errorf(err.Error())
	}

	// read package name
	if f.Name == nil {
		t.Errorf("Unknown package name")
	} else if f.Name.Name != "example" {
		t.Errorf("Failed to parse packge name. "+
			"Expected \"example\" but get\"%s\"", f.Name.Name)
	}

	// read types name and details
	numt := 0
	expected := []string{"Post", "Comment", "User"}
	var found []*TypeSpec
	for pts := range DeclsToTypes(f.Decls) {
		if pts.Name != expected[numt] {
			t.Errorf("Unexptected definition \"%s\"", pts.Name)
		}
		found = append(found, pts)
		numt++
	}

	// see if all types needed are found
	if len(expected) != len(found) {
		t.Errorf("Types found are not as expected")
		t.Errorf("Expected: %#v", expected)
		t.Errorf("Found:    %#v", found)
	}
}
