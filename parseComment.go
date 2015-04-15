package goparser

import (
	"go/ast"
	"regexp"
	"strings"
)

var cmtStyles []cmtDef

func init() {
	cmtStyles = []cmtDef{
		cmtDef{
			Regexp: regexp.MustCompile("^[ ]*\\/\\/"),
			Filter: func(in string) (out string) {
				// break down into lines
				// remove " // " from each line and combine string
				r := regexp.MustCompile("(?m)^[ \t]*\\/\\/[ \t]*")
				return strings.Trim(r.ReplaceAllString(in, ""), " \t\n\r")
			},
		},
		cmtDef{
			Regexp: regexp.MustCompile("^[ ]*/\\*\\*"),
			Filter: func(in string) (out string) {
				// break down into lines
				// remove " // " from each line and combine string
				r1 := regexp.MustCompile("(?s)^/[\\*]+[\n\r]+(.+?)[\\*]+/$")
				r2 := regexp.MustCompile("[\n\r]+")
				r3 := regexp.MustCompile("^[ \t]*\\*")
				lines := r2.Split(r1.ReplaceAllString(in, "$1"), -1)
				for _, l := range lines {
					out += r3.ReplaceAllString(l, "") + "\n"
				}
				return strings.TrimRight(out, " \t\n\r")
			},
		},
		cmtDef{
			Regexp: regexp.MustCompile("^[ ]*/\\*"),
			Filter: func(in string) (out string) {
				// break down into lines
				// remove " // " from each line and combine string
				r1 := regexp.MustCompile("(?s)^/[\\*]+[\n\r]+(.+?)[\\*]+/$")
				r2 := regexp.MustCompile("[\n\r]+")
				lines := r2.Split(r1.ReplaceAllString(in, "$1"), -1)
				for _, l := range lines {
					out += l + "\n"
				}
				return strings.TrimRight(out, " \t\n\r")
			},
		},
	}
}

// cmtDef represents comment style definition
type cmtDef struct {
	Regexp *regexp.Regexp
	Filter func(in string) (out string)
}

func (d *cmtDef) Match(str string) bool {
	return d.Regexp.Match([]byte(str))
}

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

	return
}

// CmtToStr cmment to string
func CmtToStr(cmt *ast.Comment) string {
	// if the comment matches the comment style
	for _, cs := range cmtStyles {
		if cs.Match(cmt.Text) {
			return cs.Filter(cmt.Text)
		}
	}
	//panic("Comment doesn't match any comment style we have")
	return cmt.Text
}
