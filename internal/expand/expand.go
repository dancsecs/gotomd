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

func isCmd(line string) (int, int, error) {
	const sep = "::"

	var (
		cmdIdx = -1
		end    = 0
	)

	if strings.HasPrefix(line, szCmdPrefix) {
		cmd := line[len(szCmdPrefix):]
		end = strings.Index(cmd, sep)

		if end >= 0 {
			cmd = cmd[:end+2]
			cmdIdx = action.search(cmd)
		}

		if end < 0 || cmdIdx == -1 {
			return 0, 0, fmt.Errorf("%w: %q", errs.ErrUnknownCommand, line)
		}
	}

	return cmdIdx, len(szCmdPrefix) + end + len(sep), nil
}

func getBlock(
	i, cmdStart int,
	lines []string,
	terminator, cutSet, sep string,
) (int, string, error) {
	var (
		str          strings.Builder
		addSeparator bool
		line         string
		err          error
	)

	if i < len(lines) {
		line = strings.Trim(lines[i][cmdStart:], " ")
	} else {
		err = errs.ErrBlockNotTerminated
	}

	for err == nil {
		if strings.HasSuffix(line, terminator) {
			break
		}

		if addSeparator {
			str.WriteString(sep)
		} else {
			addSeparator = true
		}

		str.WriteString(line)

		i++
		if i < len(lines) {
			line = strings.Trim(lines[i], " ")
		} else {
			err = errs.ErrBlockNotTerminated

			continue
		}
	}

	if err == nil {
		line = strings.TrimRight(line, cutSet)
		if addSeparator && len(line) > 0 {
			str.WriteString(sep)
		}
	}

	if err == nil {
		str.WriteString(line)
	}

	return i,
		strings.TrimRight(str.String(), "\n"), err
}

func processCmd(
	file *strings.Builder,
	i,
	cmdIdx, cmdStart int,
	lines []string,
) (int, error) {
	var (
		cmd string
		res string
		err error
	)

	i, cmd, err = getBlock(i, cmdStart, lines, "-->", " ->", " ")

	if err == nil {
		res, err = action.run(cmdIdx, cmd)
	}

	if err == nil {
		file.WriteString(res + "\n")
	}

	return i, err
}

func processCodeBlock(
	file *strings.Builder,
	i int,
	lines []string,
	codeSyntaxName string,
) (int, error) {
	var (
		code string
		err  error
	)

	i, code, err = getBlock(i+1, 0, lines, "```", "`", "\n")
	if err == nil {
		file.WriteString(
			format.Inline(codeSyntaxName, code),
		)
	}

	return i, err
}

//nolint:cyclop // Ok.
func processLines(lines []string, sentinel string) (string, error) {
	var (
		cmdIdx      int
		cmdStart    int
		updatedFile strings.Builder
		err         error
	)

	processLine := (sentinel == "")

	for i, mi := 0, len(lines); i < mi; i++ {
		line := strings.TrimRight(lines[i], " ")

		if !processLine {
			processLine = (line == sentinel)

			continue
		}

		cmdIdx, cmdStart, err = isCmd(line)
		if err == nil && cmdIdx >= 0 {
			i, err = processCmd(&updatedFile, i, cmdIdx, cmdStart, lines)
			if err == nil {
				continue
			}
		}

		if err == nil && strings.HasPrefix(line, "```") {
			i, err = processCodeBlock(&updatedFile, i, lines, line[3:])

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
