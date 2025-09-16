/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2025 Leslie Dancsecs

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

package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/szlog"
)

// Valid template extensions.
const (
	GoTemplate = ".gtm.go"
	MdTemplate = ".gtm.md"
)

func isValidTemplate(path string) error {
	name := filepath.Base(path)

	valid := (strings.HasSuffix(name, GoTemplate) ||
		strings.HasSuffix(name, MdTemplate)) &&
		len(name) > len(GoTemplate)+1 &&
		strings.HasPrefix(name, ".")

	if valid {
		return nil
	}

	return fmt.Errorf(
		"%w: '%s': expected - (%s or %s)",
		ErrInvalidArgument,
		path,
		GoTemplate,
		MdTemplate,
	)
}

func appendFile(fName string) {
	if alreadyIncluded[fName] {
		szlog.Say2f("Excluding redundant template: '%s'\n", fName)
	} else {
		alreadyIncluded[fName] = true

		if strings.HasSuffix(fName, GoTemplate) {
			goFiles = append(goFiles, filepath.Clean(fName))
		} else {
			mdFiles = append(mdFiles, filepath.Clean(fName))
		}

		szlog.Say1("File to process: '", fName, "'\n")
	}
}

func addDir(dir string, recursive bool) error {
	files, err := os.ReadDir(dir)
	for i, mi := 0, len(files); i < mi && err == nil; i++ {
		filePath := filepath.Join(dir, files[i].Name())

		if files[i].IsDir() {
			if recursive {
				err = addDir(filePath, recursive)
			}
		} else {
			err = isValidTemplate(filePath)
			if err == nil {
				appendFile(filePath)
			} else {
				// Ignore non matching files.
				err = nil
			}
		}
	}

	return err //nolint:wrapCheck // Wrapped by caller.
}

// add includes the template into either the goTemplates or the mdTemplates.
func add(item string, recursive bool) error {
	var (
		stat os.FileInfo
		err  error
	)

	stat, err = os.Stat(item)

	if err == nil {
		if stat.IsDir() {
			err = addDir(item, recursive)
		} else {
			err = isValidTemplate(item)
			if err == nil {
				appendFile(item)
			}
		}
	}

	if err == nil {
		return nil
	}

	return fmt.Errorf("%w: %w", ErrUnknownObject, err)
}
