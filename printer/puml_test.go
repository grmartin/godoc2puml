package printer_test

import (
	"bytes"
	"testing"

	"github.com/grmartin/godoc2puml/ast"
	"github.com/grmartin/godoc2puml/parser"
	"github.com/grmartin/godoc2puml/printer"
)

func TestFprintPlantUML(t *testing.T) {
	buf := &bytes.Buffer{}
	scope := ast.NewScope()
	scope.Packages["pkg1"] = &ast.Package{
		Name:    "pkg1",
		Classes: []*ast.Class{},
	}
	printer.FprintPlantUML(buf, scope, []string{})
}

func TestFprintPlantUMLStdLibs(t *testing.T) {
	for _, name := range []string{"io", "net", "net/http"} {
		scope := ast.NewScope()
		pkg, err := parser.ParsePackage(name)
		if err != nil {
			t.Fatal(err)
		}
		scope.Packages[pkg.Name] = pkg
		buf := &bytes.Buffer{}
		printer.FprintPlantUML(buf, scope, []string{})
		t.Log(buf)
	}

}
