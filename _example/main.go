package main

import (
	"github.com/gourd/goparser"
	"go/parser"
	"go/token"
	"log"
	"fmt"
)

// read TypeSpec of given type names from a file
func readTypeFile(inputPath string, tns []string) (pkg string, specs []*goparser.TypeSpec, err error) {
	fset := token.NewFileSet()

	// inputPath can only be filename
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		return
	}

	// read package name
	if f.Name == nil {
		err = fmt.Errorf("Unknown package name")
	} else {
		pkg = f.Name.Name
	}

	// read types name and details
	for pts := range goparser.DeclsToTypes(f.Decls) {
		if stringInSlice(pts.Name, tns) {
			specs = append(specs, pts)
		}
	}

	// see if all types needed are found
	if len(tns) != len(specs) {
		// TODO: improve this error message. Be specific on missing type.
		err = fmt.Errorf("Not all types can be found.")
	}

	return
}

// test if a string exists in a string slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	// read the Post type from types file
	pkg, specs, err := readTypeFile("./data/types.go", []string{"Post"})
	if err != nil {
		log.Fatalf("Error reading type file: %s", err.Error())
	} else if len(specs) != 1 {
		log.Fatalf("Type \"Post\" not found")
	}
	spec := specs[0]
	log.Printf("pkg:   %s", pkg)
	log.Printf("spec:  %#v", spec)

	// find Post's field tag which db name is "id"
	for f := range spec.FieldsTagged("db", "id") {
		log.Printf("ID of Post: %#v", f)
	}

	// find Post's field tag which db name is "id"
	for f := range spec.FieldsTaggedExtra("foo", "omitempty") {
		log.Printf("Post fields that foo is marked omitempty: %#v", f)
	}

}