/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023, 2024 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"go/doc"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/szlog"
	"golang.org/x/tools/go/packages"
)

const pkgLabel = "package"

type packageInfo struct {
	fSet      *token.FileSet
	docPkg    *doc.Package
	functions map[string]*doc.Func
	constants map[string]*doc.Value
	types     map[string]*doc.Type
}

//nolint:goCheckNoGlobals // Ok.
var packageCache = make(map[string]*packageInfo)

func clearPackageCache() {
	keys := make([]string, 0, len(packageCache))

	for k := range packageCache {
		keys = append(keys, k)
	}

	for _, k := range keys {
		delete(packageCache, k)
	}
}

func (pi *packageInfo) findFunc(name string) *doc.Func {
	if pi.functions == nil {
		addFunc := func(n string, f *doc.Func) {
			pi.functions[n] = f
		}
		pi.functions = make(map[string]*doc.Func)
		plainFunctions := append([]*doc.Func(nil), pi.docPkg.Funcs...)

		for _, t := range pi.docPkg.Types {
			plainFunctions = append(plainFunctions, t.Funcs...)

			for _, f := range t.Methods {
				addFunc(t.Name+"."+f.Name, f)
			}
		}

		for _, f := range plainFunctions {
			addFunc(f.Name, f)
		}
	}

	return pi.functions[name]
}

func (pi *packageInfo) findConst(name string) *doc.Value {
	if pi.constants == nil {
		addConst := func(n string, c *doc.Value) {
			pi.constants[n] = c
		}
		pi.constants = make(map[string]*doc.Value, len(pi.docPkg.Consts))

		for _, c := range pi.docPkg.Consts {
			for _, n := range c.Names {
				addConst(n, c)
			}
		}

		for _, t := range pi.docPkg.Types {
			for _, c := range t.Consts {
				for _, n := range c.Names {
					addConst(n, c)
				}
			}
		}
	}

	return pi.constants[name]
}

func (pi *packageInfo) findType(name string) *doc.Type {
	if pi.types == nil {
		addType := func(n string, t *doc.Type) {
			pi.types[n] = t
		}
		pi.types = make(map[string]*doc.Type, len(pi.docPkg.Types))

		for _, t := range pi.docPkg.Types {
			addType(t.Name, t)
		}
	}

	return pi.types[name]
}

// getInfoFunc looks up the documentation for a function.
func (pi *packageInfo) getInfoFunc(docFunc *doc.Func) (*docInfo, error) {
	var dInfo *docInfo

	dStart := pi.fSet.PositionFor(docFunc.Decl.Pos(), true)
	dEnd := pi.fSet.PositionFor(docFunc.Decl.Body.Lbrace, true)
	fEnd := pi.fSet.PositionFor(docFunc.Decl.End(), true)
	decl, body, err := pi.snipFile(
		dStart.Filename, dStart.Offset, dEnd.Offset, fEnd.Offset,
	)

	if err == nil {
		dInfo = &docInfo{
			header: decl,
			body:   body,
			doc:    strings.Split(strings.TrimSpace(docFunc.Doc), "\n"),
		}
	}

	return dInfo, err
}

// getInfoConst looks up the documentation for a function.
func (pi *packageInfo) getInfoConst(docConst *doc.Value) (*docInfo, error) {
	var dInfo *docInfo

	dStart := pi.fSet.PositionFor(docConst.Decl.Pos(), true)
	fEnd := pi.fSet.PositionFor(docConst.Decl.End(), true)
	decl, body, err := pi.snipFile(
		dStart.Filename, dStart.Offset, -1, fEnd.Offset,
	)

	if err == nil {
		dInfo = &docInfo{
			header: decl,
			body:   body,
			doc:    strings.Split(strings.TrimSpace(docConst.Doc), "\n"),
		}
	}

	return dInfo, err
}

// getInfoType looks up the documentation for a function.
func (pi *packageInfo) getInfoType(docType *doc.Type) (*docInfo, error) {
	var dInfo *docInfo

	dStart := pi.fSet.PositionFor(docType.Decl.Pos(), true)
	dEnd := pi.fSet.PositionFor(docType.Decl.Lparen, true)
	fEnd := pi.fSet.PositionFor(docType.Decl.End(), true)
	decl, body, err := pi.snipFile(
		dStart.Filename, dStart.Offset, dEnd.Offset, fEnd.Offset,
	)

	if err == nil {
		dInfo = &docInfo{
			header: decl,
			body:   body,
			doc:    strings.Split(strings.TrimSpace(docType.Doc), "\n"),
		}
	}

	return dInfo, err
}

// GetInfo looks up the documentation information for a declaration.
func (pi *packageInfo) getInfo(name string) (*docInfo, error) {
	szlog.Infof("getInfo(%q)\n", name)

	if name == pkgLabel {
		// Return Package information.
		return &docInfo{
			header: []string{pkgLabel + " " + pi.docPkg.Name},
			body:   []string{pkgLabel + " " + pi.docPkg.Name},
			doc: strings.Split(
				strings.TrimRight(pi.docPkg.Doc, "\n\t "),
				"\n",
			),
		}, nil
	}

	if f := pi.findFunc(name); f != nil {
		// Process function
		return pi.getInfoFunc(f)
	}

	if c := pi.findConst(name); c != nil {
		// Process Constant
		return pi.getInfoConst(c)
	}

	if t := pi.findType(name); t != nil {
		// Process Type
		return pi.getInfoType(t)
	}

	return nil, fmt.Errorf("%w: %s", ErrUnknownObject, name)
}

func leadingTabsToSpaces(lines []string) []string {
	const fourSpaces = "    "

	for i, line := range lines {
		newPrefix := ""

		for j, mj := 0, len(line); j < mj; j++ {
			if line[j] == '\t' {
				newPrefix += fourSpaces
			} else {
				lines[i] = newPrefix + line[j:]

				break
			}
		}
	}

	return lines
}

func (pi *packageInfo) snipFile(
	fPath string, fPos, bPos, endPos int,
) ([]string, []string, error) {
	var (
		decl []string
		body []string
		err  error
	)

	d, err := os.ReadFile(fPath) //nolint:gosec // Ok.

	if err == nil {
		res := string(d)

		switch {
		case bPos < 0:
			decl = nil
		case bPos == 0:
			decl = leadingTabsToSpaces(strings.Split(res[fPos:endPos], "\n"))
		default:
			decl = leadingTabsToSpaces(strings.Split(
				res[fPos:bPos-1],
				"\n",
			))
		}

		body = leadingTabsToSpaces(strings.Split(res[fPos:endPos], "\n"))
	}

	return decl, body, err //nolint:wrapcheck // Caller will wrap error.
}

func createPackageInfo(dir string) (*packageInfo, error) {
	var (
		docPkg        *doc.Package
		packagesToDoc []*packages.Package
		err           error
	)

	szlog.Info("Loading Package info for: ", dir)

	cfg := new(packages.Config)
	cfg.Mode = packages.NeedName |
		packages.NeedFiles |
		packages.NeedCompiledGoFiles |
		packages.NeedSyntax |
		packages.NeedTypes

	cfg.Fset = token.NewFileSet()
	cfg.Tests = false // Exclude test packages

	packagesToDoc, err = packages.Load(cfg, dir)

	if err == nil && len(packagesToDoc[0].Errors) > 0 {
		err = ErrInvalidPackage
	}

	if err == nil {
		// Use doc.NewFromFiles with the FileSet and parsed AST files.
		docPkg, err = doc.NewFromFiles(
			packagesToDoc[0].Fset,
			packagesToDoc[0].Syntax,
			packagesToDoc[0].Name,
			doc.PreserveAST|doc.AllDecls|doc.AllMethods,
		)
	}

	if err == nil {
		return &packageInfo{
			fSet:      packagesToDoc[0].Fset,
			docPkg:    docPkg,
			functions: nil,
			constants: nil,
			types:     nil,
		}, nil
	}

	return nil, err //nolint:wrapcheck // Caller will wrap error.
}

func getInfo(dir, name string) (*docInfo, error) {
	var (
		pkgInfo *packageInfo
		dInfo   *docInfo
		ok      bool
		err     error
	)

	cwd, err := os.Getwd()
	if err == nil {
		pDir := filepath.Join(cwd, dir)
		pkgInfo, ok = packageCache[pDir]

		if !ok {
			pkgInfo, err = createPackageInfo(dir)
			if err == nil {
				packageCache[pDir] = pkgInfo
			}
		}
	}

	if err == nil {
		dInfo, err = pkgInfo.getInfo(name)
	}

	if err == nil {
		return dInfo, nil
	}

	return nil, err
}
