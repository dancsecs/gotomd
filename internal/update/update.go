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
	"strings"

	"github.com/dancsecs/szlog"
)

//nolint:goCheckNoGlobals // Ok.
var upToDate bool

// ResetUpToDate sets the upToData flag to true.
func ResetUpToDate() {
	upToDate = true
}

// IsUpToDate returns the cumulative result of all processing.  Any write
// causes upToDate to be set to false.
func IsUpToDate() bool {
	return upToDate
}

func writeFile(fPath string, data string, perm os.FileMode) error {
	var file *os.File

	//nolint:gosec // Ok.
	file, err := os.OpenFile(fPath,
		os.O_TRUNC|os.O_WRONLY|os.O_CREATE,
		perm,
	)
	if err == nil {
		defer szlog.Close("updating : "+fPath, file) // Just in case.
		_, err = file.WriteString(data)
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
//nolint:cyclop,funlen  // Ok.
func File(
	fPath string, force, checkUpToDate bool, data string, perm os.FileMode,
) (Result, error) {
	var (
		oldData       string
		exists        bool
		err           error
		okToOverwrite = force
	)

	// Insure a data ends in a single linefeed.
	data = strings.TrimRight(data, "\n") + "\n"

	exists, err = fileExists(fPath)

	if err == nil && !exists {
		upToDate = false

		if checkUpToDate {
			szlog.Say1("Would have created: ", fPath, "\n")

			return Cancelled, nil
		}

		err = writeFile(fPath, data, perm)
		if err == nil {
			return Created, nil
		}
	}

	if err == nil {
		var rawOldData []byte

		rawOldData, err = os.ReadFile(fPath) //nolint:gosec // Ok.
		if err == nil {
			oldData = strings.TrimRight(string(rawOldData), "\n") + "\n"
		}
	}

	if err == nil && oldData == data {
		szlog.Say1("No change: ", fPath, "\n")

		return Unchanged, nil
	}

	upToDate = false

	if checkUpToDate {
		szlog.Say1("Would have updated: ", fPath, "\n")

		return Cancelled, nil
	}

	if err == nil && !okToOverwrite {
		okToOverwrite, err = confirm(fPath, oldData, data)
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
