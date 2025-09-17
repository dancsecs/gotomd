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

package update

import (
	"errors"
	"fmt"
	"os"

	"github.com/dancsecs/szlog"
)

func writeFile(fPath string, data string, perm os.FileMode) error {
	var file *os.File

	//nolint:gosec // Ok.
	file, err := os.OpenFile(fPath,
		os.O_TRUNC|os.O_WRONLY|os.O_CREATE,
		perm,
	)
	if err == nil {
		defer szlog.Close("updating : "+fPath, file) // Just in case.
		_, err = file.WriteString(data + "\n")
	}

	if err == nil {
		err = file.Close()
	}

	return err //nolint:wrapcheck // Wrapped by caller.
}

func fileExists(fPath string) (bool, error) {
	stat, err := os.Stat(fPath)

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	if err == nil && stat.IsDir() {
		return false, ErrInvalidFileType
	}

	return true, nil
}

// File creates or replaces an existing file with the provided data and
// file permissions if and only if it has changed.  If changed and force is not
// true then a message asking for confirmation is presented giving an
// opportunity to review the changes.
//
//nolint:cyclop  // Ok.
func File(
	fPath string, force bool, data string, perm os.FileMode,
) (Result, error) {
	var (
		oldData       []byte
		okToOverwrite bool
		exists        bool
		err           error
	)

	exists, err = fileExists(fPath)

	if err == nil && !exists {
		err = writeFile(fPath, data, perm)
		if err == nil {
			return Created, nil
		}
	}

	if err == nil {
		oldData, err = os.ReadFile(fPath) //nolint:gosec // Ok.
	}

	if err == nil && string(oldData) == data {
		szlog.Say1("No change: ", fPath, "\n")

		return Unchanged, nil
	}

	okToOverwrite = force
	if err == nil && !force {
		okToOverwrite, err = confirm(fPath, string(oldData), data)
	}

	if err == nil && okToOverwrite {
		err = writeFile(fPath, data, perm)

		if err == nil {
			return Updated, nil
		}
	}

	if err == nil {
		return Cancelled, nil
	}

	return Failed, fmt.Errorf("%w: '%s': %w", ErrFileUpdate, fPath, err)
}
