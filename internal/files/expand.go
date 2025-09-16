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
	"strings"
)

// Expand scans the list replacing entries like "{dir}/..." with a
// recursive list of {dir} and all of its subdirectories.
func Expand(list []string) error {
	const recursiveSuffix = "/..."

	var err error

	Reset()

	for i, mi := 0, len(list); i < mi && err == nil; i++ {
		entry := list[i]
		if entry == "" {
			entry = "."
		}

		recursive := strings.HasSuffix(entry, recursiveSuffix)
		if recursive {
			entry = entry[:len(entry)-len(recursiveSuffix)]
		}

		err = add(entry, recursive)
	}

	if err == nil {
		return nil
	}

	return fmt.Errorf("%w: %w", ErrInvalidTemplate, err)
}
