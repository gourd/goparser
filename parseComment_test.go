package goparser

import (
	"go/ast"
	"testing"
)

func testCmtAgainstExpt(t *testing.T, text, expected string) {
	// comment to test
	cmt := ast.Comment{
		Text: text,
	}

	// expectation and result
	result := CmtToStr(&cmt)

	// check result
	if expected != result {
		t.Errorf("Result doesn't match expectation")
		t.Errorf("Expected:\n%s", expected)
		t.Errorf("Get:\n%s", result)
	}
}

// test comment style 1 ("//")
func TestCmtToStr1(t *testing.T) {
	testCmtAgainstExpt(t,
		"//\n"+
			" // testing line 1\n"+
			"//testing line 2\n"+
			"// testing line 3",
		"testing line 1\n"+
			"testing line 2\n"+
			"testing line 3")
}

// test comment style 2 ("/* * */")
func TestCmtToStr2(t *testing.T) {
	testCmtAgainstExpt(t,
		"/**\n"+
			" * testing line 1\n"+
			"* testing line 2\n"+
			" * testing line 3\n"+
			" */",
		" testing line 1\n"+
			" testing line 2\n"+
			" testing line 3")
}

// test comment style 3 ("/*  */")
func TestCmtToStr3(t *testing.T) {
	testCmtAgainstExpt(t,
		"/*\n"+
			" testing line 1\n"+
			" testing line 2\n"+
			"  testing line 3\n"+
			"*/",
		" testing line 1\n"+
			" testing line 2\n"+
			"  testing line 3")
}
