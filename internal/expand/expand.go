/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023-2025 Leslie Dancsecs

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

package expand

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/format"
)

const (
	cmdSep     = "::"
	szCmdLabel = "gotomd"
)

const (
	szCmdPrefix = "<!--- " + szCmdLabel + cmdSep
	szDocPrefix = "doc::"
)

const (
	szAutoHeader1 = "*****  AUTO GENERATED:  DO NOT MODIFY  *****"
	szAutoHeader2 = "MODIFY TEMPLATE: "
	szAutoHeader3 = "See: 'https://github.com/dancsecs/gotomd'"
)

// func buildCommand(cmd string, args ...string) string {
// 	rootCmd := szCmdLabel + cmdSep + cmd

// 	if len(args) > 0 {
// 		rootCmd += cmdSep + strings.Join(args, cmdSep)
// 	}

// 	return format.Comment(rootCmd)
// }

//nolint:cyclop,funlen // Ok.
func processLines(lines []string, sentinel string) (string, error) {
	var (
		cmdIdx        int
		cmdStart      int
		lastLineBlank bool
		updatedFile   strings.Builder
		err           error
	)

	processLine := (sentinel == "")

	for i, mi := 0, len(lines); i < mi; i++ {
		line := strings.TrimRight(lines[i], " ")

		if !processLine {
			processLine = (line == sentinel)

			continue
		}

		if err == nil && line == "" && lastLineBlank {
			continue
		}

		lastLineBlank = line == ""

		cmdIdx, cmdStart, err = isCmd(line)
		if err == nil && cmdIdx >= 0 {
			i, err = expandCmd(&updatedFile, i, cmdIdx, cmdStart, lines)
			if err == nil {
				continue
			}
		}

		if err == nil && strings.HasPrefix(line, "```") {
			i, err = expandPreFormatted(&updatedFile, i, lines, line[3:])

			if err == nil {
				continue
			}
		}

		if err == nil {
			// Remove comment (keeping gopls from complaining.)
			const packageLabel = "package ////"

			if !format.IsForMarkdown() {
				if strings.HasPrefix(line, packageLabel) {
					line = "package " + line[len(packageLabel):]
				}
			}

			// Not a block. Just append the line.
			updatedFile.WriteString(line + "\n")
		} else {
			break
		}
	}

	if err == nil {
		return updatedFile.String(), nil
	}

	return "", err
}

func splitDir(rawPath string) (string, string, string, error) {
	dir, name := filepath.Split(rawPath)
	dir = filepath.Clean(dir)
	path := filepath.Join(dir, name)

	if !filepath.IsLocal(path) {
		return "", "", "", errs.ErrNotLocalDir
	}

	return dir, name, path, nil
}

func parse(fName, sentinel string) (string, error) {
	var (
		err       error
		fileBytes []byte
		res       string
	)

	fileBytes, err = os.ReadFile(fName) //nolint:gosec // Ok.

	if err == nil {
		lines := strings.Split(
			string(bytes.TrimRight(fileBytes, "\n")),
			"\n",
		)
		res, err = processLines(lines, sentinel)
	}

	if err == nil {
		return strings.TrimRight(res, "\n"), nil
	}

	return "", fmt.Errorf("%w: %w", errs.ErrParseError, err)
}
