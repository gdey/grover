package grover

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// FilterIgnoreTests is a filter function for parser.ParseDir
// that ignores all test files.
var FilterIgnoreTests = func(info os.FileInfo) bool {
	return !strings.HasSuffix(info.Name(), "_test.go")
}

// ParsePackagesFromDir parses all packages in a directory.
//
// It ignores all test files.
//
// If there are directories contained within dir, ParsePackagesFromDir
// attempts to traverse into those directories as well.
//
// If an error occurs whilst traversing the nested directories,
// ParsePackagesFromDir will return a map containing any correctly
// parsed packages and the error that occured.
func ParsePackagesFromDir(dir string, funcOption FuncOption) (map[string]*Package, error) {
	fset := token.NewFileSet()

	pkgs := make(map[string]*Package)

	return pkgs, filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() || strings.HasPrefix(path, filepath.Join(dir, "cmd")) {
			return nil
		}

		parsed, err := parser.ParseDir(fset, path, FilterIgnoreTests, parser.ParseComments)
		if err != nil {
			return err
		}

		for _, pkg := range parsed {
			p := ParsePackage(fset, pkg, funcOption)
			if _, exists := pkgs[pkg.Name]; exists {
				pkgs[pkg.Name].Funcs = append(pkgs[pkg.Name].Funcs, p.Funcs...)
			} else {
				pkgs[pkg.Name] = p
			}
		}

		return nil
	})
}
