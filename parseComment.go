package goparser

import (
	"go/ast"
)

// CmtGrpToStr convert comment group to string
// read comment group to string
func CmtGrpToStr(cg *ast.CommentGroup) (str string) {

	// test if empty
	if cg == nil || cg.List == nil || len(cg.List) == 0 {
		return ""
	}

	// append all strings
	for _, cmt := range cg.List {
		if cmt != nil {
			str += CmtToStr(cmt)
		}
	}

	// TODO: remove comment signs "//", "/*", "*/" and etc

	return
}

// CmtToStr cmment to string
func CmtToStr(cmt *ast.Comment) (str string) {
	return cmt.Text
}
