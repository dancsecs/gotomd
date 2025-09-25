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

package godoc

import (
	"strings"

	"github.com/dancsecs/gotomd/internal/cmds"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/gotomd/internal/update"
)

// GetDoc returns the go documentation requested.
func GetDoc(cmd string) (string, error) {
	var (
		dInfo *gopkg.DocInfo
		res   string
	)

	dir, action, err := cmds.ParseCmds(cmd)
	for i, mi := 0, len(dir); i < mi && err == nil; i++ {
		dInfo, err = gopkg.Info(dir[i], action[i])
		if err == nil {
			if res != "" {
				res += "\n\n"
			}

			res += update.MarkGoCode(dInfo.Declaration()) + "\n\n" +
				dInfo.Comment()
		}
	}

	if err == nil {
		return res, nil
	}

	return "", err //nolint:wrapcheck // Ok.
}

// GetDocDecl returns go information for requested objects as presented in
// the source..
func GetDocDecl(cmd string) (string, error) {
	var (
		dInfo *gopkg.DocInfo
		res   string
	)

	dir, action, err := cmds.ParseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = gopkg.Info(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n"
				}

				res += strings.Join(dInfo.Header(), "\n")
			}
		}
	}

	if err == nil {
		return update.MarkGoCode(res), nil
	}

	return "", err //nolint:wrapcheck // Ok.
}

// GetDocDeclSingle returns the go declaration for requested objects on a
// single line.
func GetDocDeclSingle(cmd string) (string, error) {
	var (
		dInfo *gopkg.DocInfo
		res   string
	)

	dir, action, err := cmds.ParseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = gopkg.Info(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n"
				}

				res += dInfo.OneLine()
			}
		}
	}

	if err == nil {
		return update.MarkGoCode(res), nil
	}

	return "", err //nolint:wrapcheck // Ok.
}

// GetDocDeclNatural returns the go declaration for requested objects exactly
// as defined in the source code.
func GetDocDeclNatural(cmd string) (string, error) {
	var (
		dInfo *gopkg.DocInfo
		res   string
	)

	dir, action, err := cmds.ParseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = gopkg.Info(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n\n"
				}

				res += dInfo.NaturalComment() + "\n"
				res += dInfo.OneLine()
			}
		}
	}

	if err == nil {
		return update.MarkGoCode(res), nil
	}

	return "", err //nolint:wrapcheck // Ok.
}

// GetDocDeclConstantBlock returns the go declaration for a typed constant
// block.
func GetDocDeclConstantBlock(cmd string) (string, error) {
	var (
		dInfo *gopkg.DocInfo
		res   string
	)

	dir, action, err := cmds.ParseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = gopkg.Info(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n\n"
				}

				res += dInfo.ConstantBlock() + "\n"
			}
		}
	}

	if err == nil {
		return update.MarkGoCode(res), nil
	}

	return "", err //nolint:wrapcheck // Ok.
}
