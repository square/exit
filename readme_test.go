package exit

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"regexp"
	"testing"
)

//go:embed README.md
var readme string

//go:embed exit.go
var src string

func TestExitCodesMatchReadme(t *testing.T) {
	var (
		f    *ast.File
		conf types.Config
		err  error
	)

	re := regexp.MustCompile("\\| (\\d+) \\| `(\\w+)` \\| .* \\|\n")
	expectedConstants := make(map[string]string)
	for _, submatch := range re.FindAllStringSubmatch(readme, -1) {
		expectedConstants[submatch[2]] = submatch[1]
	}

	fset := token.NewFileSet()
	if f, err = parser.ParseFile(fset, "", src, 0); err != nil {
		t.Error(err)
	}

	// Type-check the package.
	// We create an empty map for each kind of input
	// we're interested in, and Check populates them.
	info := types.Info{Defs: make(map[*ast.Ident]types.Object)}
	if _, err = conf.Check("", fset, []*ast.File{f}, &info); err != nil {
		t.Error(err)
	}

	actualConstants := make(map[string]string)
	for id, obj := range info.Defs {
		if c, ok := obj.(*types.Const); ok {
			actualConstants[id.Name] = c.Val().String()
		}
	}

	// Assertions

	for name, expectedValue := range expectedConstants {
		if value, ok := actualConstants[name]; !ok {
			t.Errorf("exit.go does not define the exit code %q (%s)", name, expectedValue)
		} else if value != expectedValue {
			t.Errorf("exit.go maps %q to %s, README.md expects it to be %s", name, value, expectedValue)
		}
	}
	for name := range actualConstants {
		if _, ok := expectedConstants[name]; !ok {
			t.Errorf("exit.go defines an undocumented exit code, %q", name)
		}
	}
}
